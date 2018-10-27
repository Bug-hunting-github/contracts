/**
 *  TokenWallet - The Consumer Contract Wallet
 *  Copyright (C) 2018 Token Group Ltd
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.

 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.

 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

pragma solidity ^0.4.25;

import "./internals/controllable.sol";
import "./internals/date.sol";
import "./internals/json.sol";
import "./externals/strings.sol";
import "./externals/SafeMath.sol";
import "./externals/oraclizeAPI_0.4.25.sol";
import "./externals/base64.sol";


/// @title Oracle converts ERC20 token amounts into equivalent ether amounts based on cryptocurrency exchange rates.
interface IOracle {
    function convert(address, uint) external view returns (uint);
}


/// @title Oracle provides asset exchange rates and conversion functionality.
contract Oracle is usingOraclize, Base64, Date, JSON, Controllable, IOracle {
    using strings for *;
    using SafeMath for uint256;

    event AddedToken(address _sender, address _token, string _symbol, uint _magnitude);
    event RemovedToken(address _sender, address _token);
    event UpdatedTokenRate(address _sender, address _token, uint _rate);

    event SetGasPrice(address _sender, uint _gasPrice);
    event Converted(address _sender, address _token, uint _amount, uint _ether);

    event RequestedUpdate(string _symbol);
    event FailedUpdateRequest(string _reason);

    event VerifiedProof(bytes _publicKey, string _result);
    event FailedProofVerification(bytes _publicKey, string _result, string _reason);

    event SetCryptoComparePublicKey(address _sender, bytes _publicKey);

    struct Token {
        string symbol;    // Token symbol
        uint magnitude;   // 10^decimals
        uint rate;        // Token exchange rate in wei
        uint lastUpdate;  // Time of the last rate update
        bool exists;      // Flags if the struct is empty or not
    }

    mapping(address => Token) public tokens;
    address[] private _tokenAddresses;

    bytes public APIPublicKey;
    mapping(bytes32 => address) private _queryToToken;

    // @dev Construct the oracle with multiple controllers, address resolver and custom gas price.
    // @dev _resolver is the oraclize address resolver contract address.
    // @param _ens is the address of the ENS.
    // @param _controllerName is the ENS name of the Controller.
    constructor(address _resolver, address _ens, bytes32 _controllerName) Controllable(_ens, _controllerName) public {
        APIPublicKey = hex"a0f4f688350018ad1b9785991c0bde5f704b005dc79972b114dbed4a615a983710bfc647ebe5a320daa28771dce6a2d104f5efa2e4a85ba3760b76d46f8571ca";
        OAR = OraclizeAddrResolverI(_resolver);
        oraclize_setCustomGasPrice(10000000000);
        oraclize_setProof(proofType_Native);
    }

    // @dev Updates the Crypto Compare public API key.
    function updateAPIPublicKey(bytes _publicKey) external onlyController {
        APIPublicKey = _publicKey;
        emit SetCryptoComparePublicKey(msg.sender, _publicKey);
    }

    // @dev Sets the gas price used by oraclize query.
    function setCustomGasPrice(uint _gasPrice) external onlyController {
        oraclize_setCustomGasPrice(_gasPrice);
        emit SetGasPrice(msg.sender, _gasPrice);
    }

    // @dev Convert ERC20 token amount to the corresponding ether amount (used by the wallet contract).
    // @param _token ERC20 token contract address.
    // @param _amount amount of token in base units.
    function convert(address _token, uint _amount) external view returns (uint) {
        // Store the token in memory to save map entry lookup gas.
        Token storage token = tokens[_token];
        // Require that the token exists and that its rate is not zero.
        require(token.exists && token.rate != 0, "token does not exist");
        // Safely convert the token amount to ether based on the exchange rate.
        return _amount.mul(token.rate).div(token.magnitude);
    }

    // @dev Add ERC20 tokens to the list of supported tokens.
    // @param _tokens ERC20 token contract addresses.
    // @param _symbols ERC20 token names.
    // @param _magnitude 10 to the power of number of decimal places used by each ERC20 token.
    // @param _updateDate date for the token updates. This will be compared to when oracle updates are received.
    function addTokens(address[] _tokens, bytes32[] _symbols, uint[] _magnitude, uint _updateDate) external onlyController {
        // Require that all parameters have the same length.
        require(_tokens.length == _symbols.length && _tokens.length == _magnitude.length, "parameter lengths do not match");
        // Add each token to the list of supported tokens.
        for (uint i = 0; i < _tokens.length; i++) {
            // Require that the token doesn't already exist.
            address token = _tokens[i];
            require(!tokens[token].exists, "token already exists");
            // Store the intermediate values.
            string memory symbol = _symbols[i].toSliceB32().toString();
            uint magnitude = _magnitude[i];
            // Add the token to the token list.
            tokens[token] = Token({
                symbol : symbol,
                magnitude : magnitude,
                rate : 0,
                exists : true,
                lastUpdate: _updateDate
            });
            // Add the token address to the address list.
            _tokenAddresses.push(token);
            // Emit token addition event.
            emit AddedToken(msg.sender, token, symbol, magnitude);
        }
    }

    // @dev Remove ERC20 tokens from the list of supported tokens.
    // @param _tokens ERC20 token contract addresses.
    function removeTokens(address[] _tokens) external onlyController {
        // Delete each token object from the list of supported tokens based on the addresses provided.
        for (uint i = 0; i < _tokens.length; i++) {
            //token must exist, reverts on duplicates as well
            require(tokens[_tokens[i]].exists, "token does not exist");
            // Store the token address.
            address token = _tokens[i];
            // Delete the token object.
            delete tokens[token];
            // Remove the token address from the address list.
            for (uint j = 0; j < _tokenAddresses.length.sub(1); j++) {
                if (_tokenAddresses[j] == token) {
                    _tokenAddresses[j] = _tokenAddresses[_tokenAddresses.length.sub(1)];
                    break;
                }
            }
            _tokenAddresses.length--;
            // Emit token removal event.
            emit RemovedToken(msg.sender, token);
        }
    }

    // @dev Update ERC20 token exchange rate manually.
    // @param _token ERC20 token contract address.
    // @param _rate ERC20 token exchange rate in wei.
    // @param _updateDate date for the token updates. This will be compared to when oracle updates are received.
    function updateTokenRate(address _token, uint _rate, uint _updateDate) external onlyController {
        // Require that the token exists.
        require(tokens[_token].exists, "token does not exist");
        // Update the token's rate.
        tokens[_token].rate = _rate;
        // Update the token's last update timestamp.
        tokens[_token].lastUpdate = _updateDate;
        // Emit the rate update event.
        emit UpdatedTokenRate(msg.sender, _token, _rate);
    }

    // @dev Update ERC20 token exchange rates for all supported tokens.
    function updateTokenRates() external payable onlyController {
        _updateTokenRates();
    }

    /// @dev Withdraw ether from the smart contract to the specified account.
    function withdraw(address _to, uint _amount) external onlyController {
        _to.transfer(_amount);
    }

    /// @dev Handle Oraclize query callback and verifiy the provided origin proof.
    /// @param _queryID Oraclize query ID.
    /// @param _result query result in JSON format.
    /// @param _proof origin proof from crypto compare.
    // solium-disable-next-line mixedcase
    function __callback(bytes32 _queryID, string _result, bytes _proof) public {
        // Require that the caller is the Oraclize contract.
        require(msg.sender == oraclize_cbAddress(), "sender is not oraclize");
        // Use the query ID to find the matching token address.
        address _token = _queryToToken[_queryID];
        require(_token != address(0), "queryID matches to address 0");
        // Get the corresponding token object.
        Token storage token = tokens[_token];

        bool valid;
        uint timestamp;
        (valid, timestamp) = verifyProof(_result, _proof, APIPublicKey, token.lastUpdate);

        // Require that the proof is valid.
        if (valid) {
            // Parse the JSON result to get the rate in wei.
            token.rate = parseInt(parseRate(_result, "ETH"), 18);
            // Set the update time of the token rate.
            token.lastUpdate = timestamp;
            // Remove query from the list.
            delete _queryToToken[_queryID];
            // Emit the rate update event.
            emit UpdatedTokenRate(msg.sender, _token, token.rate);
        } 
    }

    // @dev Re-usable helper function that performs the Oraclize Query.
    function _updateTokenRates() private {
        // Check if there are any existing tokens.
        if (_tokenAddresses.length == 0) {
            // Emit a query failure event.
            emit FailedUpdateRequest("no tokens");
        // Check if the contract has enough Ether to pay for the query.
        } else if (oraclize_getPrice("URL") * _tokenAddresses.length > address(this).balance) {
            // Emit a query failure event.
            emit FailedUpdateRequest("insufficient balance");
        } else {
            // Set up the cryptocompare API query strings.
            strings.slice memory apiPrefix = "https://min-api.cryptocompare.com/data/price?fsym=".toSlice();
            strings.slice memory apiSuffix = "&tsyms=ETH&sign=true".toSlice();
  
            uint gaslimit = 2000000;
            // Create a new oraclize query for each supported token.
            for (uint i = 0; i < _tokenAddresses.length; i++) {
                // Store the token symbol used in the query.
                strings.slice memory symbol = tokens[_tokenAddresses[i]].symbol.toSlice();
                // Create a new oraclize query from the component strings.
                bytes32 queryID = oraclize_query("URL", apiPrefix.concat(symbol).toSlice().concat(apiSuffix), gaslimit);
                // Store the query ID together with the associated token address.
                _queryToToken[queryID] = _tokenAddresses[i];
                // Emit the query success event.
                emit RequestedUpdate(symbol.toString());
            }
        }
    }

    // @dev Verify the origin proof returned by the cryptocompare API.
    // @param _result query result in JSON format.
    // @param _proof origin proof from cryptocompare.
    // @param _publicKey cryptocompare public key.
    // @param _lastUpdate timestamp of the last time the requested token was updated.
    function verifyProof(string _result, bytes _proof, bytes _publicKey, uint _lastUpdate) private returns (bool, uint) {
        // Extract the signature.
        uint signatureLength = uint(_proof[1]);
        bytes memory signature = new bytes(signatureLength);
        signature = copyBytes(_proof, 2, signatureLength, signature, 0);
        // Extract the headers.
        uint headersLength = uint(_proof[2 + signatureLength]) * 256 + uint(_proof[2 + signatureLength + 1]);
        bytes memory headers = new bytes(headersLength);
        headers = copyBytes(_proof, 4 + signatureLength, headersLength, headers, 0);
        // Check if the date is valid.
        bytes memory dateHeader = new bytes(30);
        dateHeader = copyBytes(headers, 5, 30, dateHeader, 0);

        bool dateValid;
        uint timestamp;
        (dateValid, timestamp) = verifyDate(string(dateHeader), _lastUpdate);

        // Check if the Date returned is valid or not
        if (!dateValid) {
            emit FailedProofVerification(_publicKey, _result, "date");
            return (false, 0);
        }

        // Check if the signed digest hash matches the result hash.
        bytes memory digest = new bytes(headersLength - 52);
        digest = copyBytes(headers, 52, headersLength - 52, digest, 0);
        if (keccak256(sha256(_result)) != keccak256(base64decode(digest))) {
            emit FailedProofVerification(_publicKey, _result, "hash");
            return (false, 0);
        }

        // Check if the signature is valid and if the signer addresses match.
        if (!verifySignature(headers, signature, _publicKey)) {
            emit FailedProofVerification(_publicKey, _result, "signature");
            return (false, 0);
        }

        emit VerifiedProof(_publicKey, _result);
        return (true, timestamp);
    }

    // @dev Verify the HTTP headers and the signature
    // @param _headers HTTP headers provided by the cryptocompare api
    // @param _signature signature provided by the cryptocompare api
    // @param _publicKey cryptocompare public key.
    function verifySignature(bytes _headers, bytes _signature, bytes _publicKey) private returns (bool) {
        address signer;
        bool signatureOK;
        
        // Checks if the signature is valid by hashing the headers
        (signatureOK, signer) = ecrecovery(sha256(_headers), _signature);
        return signatureOK && signer == address(keccak256(_publicKey));
    }

    // @dev Verify the signed HTTP date header.
    // @param _dateHeader extracted date string e.g. Wed, 12 Sep 2018 15:18:14 GMT.
    // @param _lastUpdate timestamp of the last time the requested token was updated.
    function verifyDate(string _dateHeader, uint _lastUpdate) private pure returns (bool, uint) {
        strings.slice memory date = _dateHeader.toSlice();
        strings.slice memory timeDelimiter = ":".toSlice();
        strings.slice memory dateDelimiter = " ".toSlice();
        // Split the date string.
        date.split(",".toSlice());
        date.split(dateDelimiter);
        // Get individual date components.
        uint day = parseInt(date.split(dateDelimiter).toString());
        uint month = monthToNumber(date.split(dateDelimiter).toString());
        uint year = parseInt(date.split(dateDelimiter).toString());
        uint hour = parseInt(date.split(timeDelimiter).toString());
        uint minute = parseInt(date.split(timeDelimiter).toString());
        uint second = parseInt(date.split(timeDelimiter).toString());

        if (day > 31 || day < 1) {
            return (false, 0);
        }

        if (month > 12 || month < 1) {
            return (false, 0);
        }

        if (year < 2018 || year > 3000) {
            return (false, 0);
        }
    
        if (hour >= 24) {
            return (false, 0);
        }

        if (minute >= 60) {
            return (false, 0);
        }

        if (second >= 60) {
            return (false, 0);
        }

        uint timestamp = year * (10 ** 10) + month * (10 ** 8) + day * (10 ** 6) + hour * (10 ** 4) + minute * (10 ** 2) + second;

        return (timestamp > _lastUpdate, timestamp);
    }
}
