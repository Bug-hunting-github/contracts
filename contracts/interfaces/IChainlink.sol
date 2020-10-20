pragma solidity ^0.6.0;

interface IChainlink {
  function latestAnswer() external view returns (int256);
  function latestTimestamp() external view returns (uint256);
  function getAnswer(uint256 roundId) external view returns (int256);
  function getTimestamp(uint256 roundId) external view returns (uint256);
}