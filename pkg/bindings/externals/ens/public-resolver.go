// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ens

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PublicResolverABI is the input ABI used to generate the binding from.
const PublicResolverABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"content\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setMultihash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"setContent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"multihash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"ensAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ContentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"MultihashChanged\",\"type\":\"event\"}]"

// PublicResolverBin is the compiled bytecode used for deploying new contracts.
const PublicResolverBin = `608060405234801561001057600080fd5b50604051602080611400833981016040525160008054600160a060020a03909216600160a060020a03199092169190911790556113ae806100526000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a781146100df57806310f13a8c146101155780632203ab56146101b357806329cd62ea1461024d5780632dff69411461026b5780633b3b57de1461029557806359d1d43c146102c9578063623195b01461039c578063691f3431146103fc5780637737221314610414578063aa4cb54714610472578063c3d014d6146104d0578063c8690233146104eb578063d5fa2b001461051c578063e89401a114610540575b600080fd5b3480156100eb57600080fd5b50610101600160e060020a031960043516610558565b604080519115158252519081900360200190f35b34801561012157600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101b195833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506106f99650505050505050565b005b3480156101bf57600080fd5b506101ce60043560243561091f565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156102115781810151838201526020016101f9565b50505050905090810190601f16801561023e5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34801561025957600080fd5b506101b1600435602435604435610a2b565b34801561027757600080fd5b50610283600435610b2b565b60408051918252519081900360200190f35b3480156102a157600080fd5b506102ad600435610b41565b60408051600160a060020a039092168252519081900360200190f35b3480156102d557600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610327958335953695604494919390910191908190840183828082843750949750610b5c9650505050505050565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610361578181015183820152602001610349565b50505050905090810190601f16801561038e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103a857600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101b1948235946024803595369594606494920191908190840183828082843750949750610c659650505050505050565b34801561040857600080fd5b50610327600435610d66565b34801561042057600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101b1958335953695604494919390910191908190840183828082843750949750610e0a9650505050505050565b34801561047e57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101b1958335953695604494919390910191908190840183828082843750949750610f609650505050505050565b3480156104dc57600080fd5b506101b1600435602435611076565b3480156104f757600080fd5b50610503600435611157565b6040805192835260208301919091528051918290030190f35b34801561052857600080fd5b506101b1600435600160a060020a0360243516611174565b34801561054c57600080fd5b50610327600435611278565b6000600160e060020a031982167f3b3b57de0000000000000000000000000000000000000000000000000000000014806105bb5750600160e060020a031982167fd8389dc500000000000000000000000000000000000000000000000000000000145b806105ef5750600160e060020a031982167f691f343100000000000000000000000000000000000000000000000000000000145b806106235750600160e060020a031982167f2203ab5600000000000000000000000000000000000000000000000000000000145b806106575750600160e060020a031982167fc869023300000000000000000000000000000000000000000000000000000000145b8061068b5750600160e060020a031982167f59d1d43c00000000000000000000000000000000000000000000000000000000145b806106bf5750600160e060020a031982167fe89401a100000000000000000000000000000000000000000000000000000000145b806106f35750600160e060020a031982167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b600080546040805160e060020a6302571be302815260048101879052905186933393600160a060020a0316926302571be39260248083019360209383900390910190829087803b15801561074c57600080fd5b505af1158015610760573d6000803e3d6000fd5b505050506040513d602081101561077657600080fd5b5051600160a060020a03161461078b57600080fd5b6000848152600160209081526040918290209151855185936005019287929182918401908083835b602083106107d25780518252601f1990920191602091820191016107b3565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101909320845161081395919491909101925090506112e7565b5083600019167fd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a75508485604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561087d578181015183820152602001610865565b50505050905090810190601f1680156108aa5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156108dd5781810151838201526020016108c5565b50505050905090810190601f16801561090a5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a250505050565b60008281526001602081905260409091206060905b838311610a1e578284161580159061096d5750600083815260068201602052604081205460026000196101006001841615020190911604115b15610a1357600083815260068201602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610a075780601f106109dc57610100808354040283529160200191610a07565b820191906000526020600020905b8154815290600101906020018083116109ea57829003601f168201915b50505050509150610a23565b600290920291610934565b600092505b509250929050565b600080546040805160e060020a6302571be302815260048101879052905186933393600160a060020a0316926302571be39260248083019360209383900390910190829087803b158015610a7e57600080fd5b505af1158015610a92573d6000803e3d6000fd5b505050506040513d6020811015610aa857600080fd5b5051600160a060020a031614610abd57600080fd5b604080518082018252848152602080820185815260008881526001835284902092516003840155516004909201919091558151858152908101849052815186927f1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46928290030190a250505050565b6000908152600160208190526040909120015490565b600090815260016020526040902054600160a060020a031690565b600082815260016020908152604091829020915183516060936005019285929182918401908083835b60208310610ba45780518252601f199092019160209182019101610b85565b518151600019602094850361010090810a820192831692199390931691909117909252949092019687526040805197889003820188208054601f6002600183161590980290950116959095049283018290048202880182019052818752929450925050830182828015610c585780601f10610c2d57610100808354040283529160200191610c58565b820191906000526020600020905b815481529060010190602001808311610c3b57829003601f168201915b5050505050905092915050565b600080546040805160e060020a6302571be302815260048101879052905186933393600160a060020a0316926302571be39260248083019360209383900390910190829087803b158015610cb857600080fd5b505af1158015610ccc573d6000803e3d6000fd5b505050506040513d6020811015610ce257600080fd5b5051600160a060020a031614610cf757600080fd5b6000198301831615610d0857600080fd5b600084815260016020908152604080832086845260060182529091208351610d32928501906112e7565b50604051839085907faa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe390600090a350505050565b6000818152600160208181526040928390206002908101805485516000199582161561010002959095011691909104601f81018390048302840183019094528383526060939091830182828015610dfe5780601f10610dd357610100808354040283529160200191610dfe565b820191906000526020600020905b815481529060010190602001808311610de157829003601f168201915b50505050509050919050565b600080546040805160e060020a6302571be302815260048101869052905185933393600160a060020a0316926302571be39260248083019360209383900390910190829087803b158015610e5d57600080fd5b505af1158015610e71573d6000803e3d6000fd5b505050506040513d6020811015610e8757600080fd5b5051600160a060020a031614610e9c57600080fd5b60008381526001602090815260409091208351610ec1926002909201918501906112e7565b50604080516020808252845181830152845186937fb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f79387939092839283019185019080838360005b83811015610f21578181015183820152602001610f09565b50505050905090810190601f168015610f4e5780820380516001836020036101000a031916815260200191505b509250505060405180910390a2505050565b600080546040805160e060020a6302571be302815260048101869052905185933393600160a060020a0316926302571be39260248083019360209383900390910190829087803b158015610fb357600080fd5b505af1158015610fc7573d6000803e3d6000fd5b505050506040513d6020811015610fdd57600080fd5b5051600160a060020a031614610ff257600080fd5b60008381526001602090815260409091208351611017926007909201918501906112e7565b50604080516020808252845181830152845186937fc0b0fc07269fc2749adada3221c095a1d2187b2d075b51c915857b520f3a502193879390928392830191850190808383600083811015610f21578181015183820152602001610f09565b600080546040805160e060020a6302571be302815260048101869052905185933393600160a060020a0316926302571be39260248083019360209383900390910190829087803b1580156110c957600080fd5b505af11580156110dd573d6000803e3d6000fd5b505050506040513d60208110156110f357600080fd5b5051600160a060020a03161461110857600080fd5b6000838152600160208181526040928390209091018490558151848152915185927f0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc92908290030190a2505050565b600090815260016020526040902060038101546004909101549091565b600080546040805160e060020a6302571be302815260048101869052905185933393600160a060020a0316926302571be39260248083019360209383900390910190829087803b1580156111c757600080fd5b505af11580156111db573d6000803e3d6000fd5b505050506040513d60208110156111f157600080fd5b5051600160a060020a03161461120657600080fd5b600083815260016020908152604091829020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386169081179091558251908152915185927f52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd292908290030190a2505050565b60008181526001602081815260409283902060070180548451600260001995831615610100029590950190911693909304601f81018390048302840183019094528383526060939091830182828015610dfe5780601f10610dd357610100808354040283529160200191610dfe565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061132857805160ff1916838001178555611355565b82800160010185558215611355579182015b8281111561135557825182559160200191906001019061133a565b50611361929150611365565b5090565b61137f91905b80821115611361576000815560010161136b565b905600a165627a7a723058205895912debbb21d592ea4ac446032db3d22e64f7eecfc902f21074990a50d6c20029`

// DeployPublicResolver deploys a new Ethereum contract, binding an instance of PublicResolver to it.
func DeployPublicResolver(auth *bind.TransactOpts, backend bind.ContractBackend, ensAddr common.Address) (common.Address, *types.Transaction, *PublicResolver, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicResolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PublicResolverBin), backend, ensAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicResolver{PublicResolverCaller: PublicResolverCaller{contract: contract}, PublicResolverTransactor: PublicResolverTransactor{contract: contract}, PublicResolverFilterer: PublicResolverFilterer{contract: contract}}, nil
}

// PublicResolver is an auto generated Go binding around an Ethereum contract.
type PublicResolver struct {
	PublicResolverCaller     // Read-only binding to the contract
	PublicResolverTransactor // Write-only binding to the contract
	PublicResolverFilterer   // Log filterer for contract events
}

// PublicResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicResolverSession struct {
	Contract     *PublicResolver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PublicResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicResolverCallerSession struct {
	Contract *PublicResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PublicResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicResolverTransactorSession struct {
	Contract     *PublicResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PublicResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicResolverRaw struct {
	Contract *PublicResolver // Generic contract binding to access the raw methods on
}

// PublicResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicResolverCallerRaw struct {
	Contract *PublicResolverCaller // Generic read-only contract binding to access the raw methods on
}

// PublicResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicResolverTransactorRaw struct {
	Contract *PublicResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicResolver creates a new instance of PublicResolver, bound to a specific deployed contract.
func NewPublicResolver(address common.Address, backend bind.ContractBackend) (*PublicResolver, error) {
	contract, err := bindPublicResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicResolver{PublicResolverCaller: PublicResolverCaller{contract: contract}, PublicResolverTransactor: PublicResolverTransactor{contract: contract}, PublicResolverFilterer: PublicResolverFilterer{contract: contract}}, nil
}

// NewPublicResolverCaller creates a new read-only instance of PublicResolver, bound to a specific deployed contract.
func NewPublicResolverCaller(address common.Address, caller bind.ContractCaller) (*PublicResolverCaller, error) {
	contract, err := bindPublicResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicResolverCaller{contract: contract}, nil
}

// NewPublicResolverTransactor creates a new write-only instance of PublicResolver, bound to a specific deployed contract.
func NewPublicResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicResolverTransactor, error) {
	contract, err := bindPublicResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicResolverTransactor{contract: contract}, nil
}

// NewPublicResolverFilterer creates a new log filterer instance of PublicResolver, bound to a specific deployed contract.
func NewPublicResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicResolverFilterer, error) {
	contract, err := bindPublicResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicResolverFilterer{contract: contract}, nil
}

// bindPublicResolver binds a generic wrapper to an already deployed contract.
func bindPublicResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicResolver *PublicResolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PublicResolver.Contract.PublicResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicResolver *PublicResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicResolver.Contract.PublicResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicResolver *PublicResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicResolver.Contract.PublicResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicResolver *PublicResolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PublicResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicResolver *PublicResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicResolver *PublicResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicResolver.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_PublicResolver *PublicResolverCaller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	ret := new(struct {
		ContentType *big.Int
		Data        []byte
	})
	out := ret
	err := _PublicResolver.contract.Call(opts, out, "ABI", node, contentTypes)
	return *ret, err
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_PublicResolver *PublicResolverSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _PublicResolver.Contract.ABI(&_PublicResolver.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_PublicResolver *PublicResolverCallerSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _PublicResolver.Contract.ABI(&_PublicResolver.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(address)
func (_PublicResolver *PublicResolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PublicResolver.contract.Call(opts, out, "addr", node)
	return *ret0, err
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(address)
func (_PublicResolver *PublicResolverSession) Addr(node [32]byte) (common.Address, error) {
	return _PublicResolver.Contract.Addr(&_PublicResolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(address)
func (_PublicResolver *PublicResolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _PublicResolver.Contract.Addr(&_PublicResolver.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(bytes32)
func (_PublicResolver *PublicResolverCaller) Content(opts *bind.CallOpts, node [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PublicResolver.contract.Call(opts, out, "content", node)
	return *ret0, err
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(bytes32)
func (_PublicResolver *PublicResolverSession) Content(node [32]byte) ([32]byte, error) {
	return _PublicResolver.Contract.Content(&_PublicResolver.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(bytes32)
func (_PublicResolver *PublicResolverCallerSession) Content(node [32]byte) ([32]byte, error) {
	return _PublicResolver.Contract.Content(&_PublicResolver.CallOpts, node)
}

// Multihash is a free data retrieval call binding the contract method 0xe89401a1.
//
// Solidity: function multihash(node bytes32) constant returns(bytes)
func (_PublicResolver *PublicResolverCaller) Multihash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _PublicResolver.contract.Call(opts, out, "multihash", node)
	return *ret0, err
}

// Multihash is a free data retrieval call binding the contract method 0xe89401a1.
//
// Solidity: function multihash(node bytes32) constant returns(bytes)
func (_PublicResolver *PublicResolverSession) Multihash(node [32]byte) ([]byte, error) {
	return _PublicResolver.Contract.Multihash(&_PublicResolver.CallOpts, node)
}

// Multihash is a free data retrieval call binding the contract method 0xe89401a1.
//
// Solidity: function multihash(node bytes32) constant returns(bytes)
func (_PublicResolver *PublicResolverCallerSession) Multihash(node [32]byte) ([]byte, error) {
	return _PublicResolver.Contract.Multihash(&_PublicResolver.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(string)
func (_PublicResolver *PublicResolverCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PublicResolver.contract.Call(opts, out, "name", node)
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(string)
func (_PublicResolver *PublicResolverSession) Name(node [32]byte) (string, error) {
	return _PublicResolver.Contract.Name(&_PublicResolver.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(string)
func (_PublicResolver *PublicResolverCallerSession) Name(node [32]byte) (string, error) {
	return _PublicResolver.Contract.Name(&_PublicResolver.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_PublicResolver *PublicResolverCaller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	ret := new(struct {
		X [32]byte
		Y [32]byte
	})
	out := ret
	err := _PublicResolver.contract.Call(opts, out, "pubkey", node)
	return *ret, err
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_PublicResolver *PublicResolverSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _PublicResolver.Contract.Pubkey(&_PublicResolver.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_PublicResolver *PublicResolverCallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _PublicResolver.Contract.Pubkey(&_PublicResolver.CallOpts, node)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceID bytes4) constant returns(bool)
func (_PublicResolver *PublicResolverCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PublicResolver.contract.Call(opts, out, "supportsInterface", interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceID bytes4) constant returns(bool)
func (_PublicResolver *PublicResolverSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _PublicResolver.Contract.SupportsInterface(&_PublicResolver.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceID bytes4) constant returns(bool)
func (_PublicResolver *PublicResolverCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _PublicResolver.Contract.SupportsInterface(&_PublicResolver.CallOpts, interfaceID)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(string)
func (_PublicResolver *PublicResolverCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PublicResolver.contract.Call(opts, out, "text", node, key)
	return *ret0, err
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(string)
func (_PublicResolver *PublicResolverSession) Text(node [32]byte, key string) (string, error) {
	return _PublicResolver.Contract.Text(&_PublicResolver.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(string)
func (_PublicResolver *PublicResolverCallerSession) Text(node [32]byte, key string) (string, error) {
	return _PublicResolver.Contract.Text(&_PublicResolver.CallOpts, node, key)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_PublicResolver *PublicResolverTransactor) SetABI(opts *bind.TransactOpts, node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _PublicResolver.contract.Transact(opts, "setABI", node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_PublicResolver *PublicResolverSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetABI(&_PublicResolver.TransactOpts, node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_PublicResolver *PublicResolverTransactorSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetABI(&_PublicResolver.TransactOpts, node, contentType, data)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_PublicResolver *PublicResolverTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _PublicResolver.contract.Transact(opts, "setAddr", node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_PublicResolver *PublicResolverSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetAddr(&_PublicResolver.TransactOpts, node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_PublicResolver *PublicResolverTransactorSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetAddr(&_PublicResolver.TransactOpts, node, addr)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_PublicResolver *PublicResolverTransactor) SetContent(opts *bind.TransactOpts, node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _PublicResolver.contract.Transact(opts, "setContent", node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_PublicResolver *PublicResolverSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetContent(&_PublicResolver.TransactOpts, node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_PublicResolver *PublicResolverTransactorSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetContent(&_PublicResolver.TransactOpts, node, hash)
}

// SetMultihash is a paid mutator transaction binding the contract method 0xaa4cb547.
//
// Solidity: function setMultihash(node bytes32, hash bytes) returns()
func (_PublicResolver *PublicResolverTransactor) SetMultihash(opts *bind.TransactOpts, node [32]byte, hash []byte) (*types.Transaction, error) {
	return _PublicResolver.contract.Transact(opts, "setMultihash", node, hash)
}

// SetMultihash is a paid mutator transaction binding the contract method 0xaa4cb547.
//
// Solidity: function setMultihash(node bytes32, hash bytes) returns()
func (_PublicResolver *PublicResolverSession) SetMultihash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetMultihash(&_PublicResolver.TransactOpts, node, hash)
}

// SetMultihash is a paid mutator transaction binding the contract method 0xaa4cb547.
//
// Solidity: function setMultihash(node bytes32, hash bytes) returns()
func (_PublicResolver *PublicResolverTransactorSession) SetMultihash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetMultihash(&_PublicResolver.TransactOpts, node, hash)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_PublicResolver *PublicResolverTransactor) SetName(opts *bind.TransactOpts, node [32]byte, name string) (*types.Transaction, error) {
	return _PublicResolver.contract.Transact(opts, "setName", node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_PublicResolver *PublicResolverSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetName(&_PublicResolver.TransactOpts, node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_PublicResolver *PublicResolverTransactorSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetName(&_PublicResolver.TransactOpts, node, name)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_PublicResolver *PublicResolverTransactor) SetPubkey(opts *bind.TransactOpts, node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _PublicResolver.contract.Transact(opts, "setPubkey", node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_PublicResolver *PublicResolverSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetPubkey(&_PublicResolver.TransactOpts, node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_PublicResolver *PublicResolverTransactorSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetPubkey(&_PublicResolver.TransactOpts, node, x, y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_PublicResolver *PublicResolverTransactor) SetText(opts *bind.TransactOpts, node [32]byte, key string, value string) (*types.Transaction, error) {
	return _PublicResolver.contract.Transact(opts, "setText", node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_PublicResolver *PublicResolverSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetText(&_PublicResolver.TransactOpts, node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_PublicResolver *PublicResolverTransactorSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetText(&_PublicResolver.TransactOpts, node, key, value)
}

// PublicResolverABIChangedIterator is returned from FilterABIChanged and is used to iterate over the raw logs and unpacked data for ABIChanged events raised by the PublicResolver contract.
type PublicResolverABIChangedIterator struct {
	Event *PublicResolverABIChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PublicResolverABIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicResolverABIChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PublicResolverABIChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PublicResolverABIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicResolverABIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicResolverABIChanged represents a ABIChanged event raised by the PublicResolver contract.
type PublicResolverABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterABIChanged is a free log retrieval operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: e ABIChanged(node indexed bytes32, contentType indexed uint256)
func (_PublicResolver *PublicResolverFilterer) FilterABIChanged(opts *bind.FilterOpts, node [][32]byte, contentType []*big.Int) (*PublicResolverABIChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _PublicResolver.contract.FilterLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return &PublicResolverABIChangedIterator{contract: _PublicResolver.contract, event: "ABIChanged", logs: logs, sub: sub}, nil
}

// WatchABIChanged is a free log subscription operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: e ABIChanged(node indexed bytes32, contentType indexed uint256)
func (_PublicResolver *PublicResolverFilterer) WatchABIChanged(opts *bind.WatchOpts, sink chan<- *PublicResolverABIChanged, node [][32]byte, contentType []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _PublicResolver.contract.WatchLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicResolverABIChanged)
				if err := _PublicResolver.contract.UnpackLog(event, "ABIChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PublicResolverAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the PublicResolver contract.
type PublicResolverAddrChangedIterator struct {
	Event *PublicResolverAddrChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PublicResolverAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicResolverAddrChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PublicResolverAddrChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PublicResolverAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicResolverAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicResolverAddrChanged represents a AddrChanged event raised by the PublicResolver contract.
type PublicResolverAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: e AddrChanged(node indexed bytes32, a address)
func (_PublicResolver *PublicResolverFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicResolverAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicResolverAddrChangedIterator{contract: _PublicResolver.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: e AddrChanged(node indexed bytes32, a address)
func (_PublicResolver *PublicResolverFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *PublicResolverAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicResolverAddrChanged)
				if err := _PublicResolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PublicResolverContentChangedIterator is returned from FilterContentChanged and is used to iterate over the raw logs and unpacked data for ContentChanged events raised by the PublicResolver contract.
type PublicResolverContentChangedIterator struct {
	Event *PublicResolverContentChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PublicResolverContentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicResolverContentChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PublicResolverContentChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PublicResolverContentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicResolverContentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicResolverContentChanged represents a ContentChanged event raised by the PublicResolver contract.
type PublicResolverContentChanged struct {
	Node [32]byte
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContentChanged is a free log retrieval operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: e ContentChanged(node indexed bytes32, hash bytes32)
func (_PublicResolver *PublicResolverFilterer) FilterContentChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicResolverContentChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.FilterLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicResolverContentChangedIterator{contract: _PublicResolver.contract, event: "ContentChanged", logs: logs, sub: sub}, nil
}

// WatchContentChanged is a free log subscription operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: e ContentChanged(node indexed bytes32, hash bytes32)
func (_PublicResolver *PublicResolverFilterer) WatchContentChanged(opts *bind.WatchOpts, sink chan<- *PublicResolverContentChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.WatchLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicResolverContentChanged)
				if err := _PublicResolver.contract.UnpackLog(event, "ContentChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PublicResolverMultihashChangedIterator is returned from FilterMultihashChanged and is used to iterate over the raw logs and unpacked data for MultihashChanged events raised by the PublicResolver contract.
type PublicResolverMultihashChangedIterator struct {
	Event *PublicResolverMultihashChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PublicResolverMultihashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicResolverMultihashChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PublicResolverMultihashChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PublicResolverMultihashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicResolverMultihashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicResolverMultihashChanged represents a MultihashChanged event raised by the PublicResolver contract.
type PublicResolverMultihashChanged struct {
	Node [32]byte
	Hash []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMultihashChanged is a free log retrieval operation binding the contract event 0xc0b0fc07269fc2749adada3221c095a1d2187b2d075b51c915857b520f3a5021.
//
// Solidity: e MultihashChanged(node indexed bytes32, hash bytes)
func (_PublicResolver *PublicResolverFilterer) FilterMultihashChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicResolverMultihashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.FilterLogs(opts, "MultihashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicResolverMultihashChangedIterator{contract: _PublicResolver.contract, event: "MultihashChanged", logs: logs, sub: sub}, nil
}

// WatchMultihashChanged is a free log subscription operation binding the contract event 0xc0b0fc07269fc2749adada3221c095a1d2187b2d075b51c915857b520f3a5021.
//
// Solidity: e MultihashChanged(node indexed bytes32, hash bytes)
func (_PublicResolver *PublicResolverFilterer) WatchMultihashChanged(opts *bind.WatchOpts, sink chan<- *PublicResolverMultihashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.WatchLogs(opts, "MultihashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicResolverMultihashChanged)
				if err := _PublicResolver.contract.UnpackLog(event, "MultihashChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PublicResolverNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the PublicResolver contract.
type PublicResolverNameChangedIterator struct {
	Event *PublicResolverNameChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PublicResolverNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicResolverNameChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PublicResolverNameChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PublicResolverNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicResolverNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicResolverNameChanged represents a NameChanged event raised by the PublicResolver contract.
type PublicResolverNameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: e NameChanged(node indexed bytes32, name string)
func (_PublicResolver *PublicResolverFilterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicResolverNameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicResolverNameChangedIterator{contract: _PublicResolver.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: e NameChanged(node indexed bytes32, name string)
func (_PublicResolver *PublicResolverFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *PublicResolverNameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicResolverNameChanged)
				if err := _PublicResolver.contract.UnpackLog(event, "NameChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PublicResolverPubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the PublicResolver contract.
type PublicResolverPubkeyChangedIterator struct {
	Event *PublicResolverPubkeyChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PublicResolverPubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicResolverPubkeyChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PublicResolverPubkeyChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PublicResolverPubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicResolverPubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicResolverPubkeyChanged represents a PubkeyChanged event raised by the PublicResolver contract.
type PublicResolverPubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: e PubkeyChanged(node indexed bytes32, x bytes32, y bytes32)
func (_PublicResolver *PublicResolverFilterer) FilterPubkeyChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicResolverPubkeyChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.FilterLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicResolverPubkeyChangedIterator{contract: _PublicResolver.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: e PubkeyChanged(node indexed bytes32, x bytes32, y bytes32)
func (_PublicResolver *PublicResolverFilterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *PublicResolverPubkeyChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.WatchLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicResolverPubkeyChanged)
				if err := _PublicResolver.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PublicResolverTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the PublicResolver contract.
type PublicResolverTextChangedIterator struct {
	Event *PublicResolverTextChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PublicResolverTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicResolverTextChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PublicResolverTextChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PublicResolverTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicResolverTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicResolverTextChanged represents a TextChanged event raised by the PublicResolver contract.
type PublicResolverTextChanged struct {
	Node       [32]byte
	IndexedKey string
	Key        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: e TextChanged(node indexed bytes32, indexedKey string, key string)
func (_PublicResolver *PublicResolverFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicResolverTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.FilterLogs(opts, "TextChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicResolverTextChangedIterator{contract: _PublicResolver.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: e TextChanged(node indexed bytes32, indexedKey string, key string)
func (_PublicResolver *PublicResolverFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *PublicResolverTextChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.WatchLogs(opts, "TextChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicResolverTextChanged)
				if err := _PublicResolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
