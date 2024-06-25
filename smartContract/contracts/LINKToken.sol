// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract LINKToken is ERC20 {
  address public owner;

  constructor(uint256 initialSupply) ERC20("LINK Token", "LINK") {
    owner = msg.sender;
    _mint(owner, initialSupply * 10 ** decimals());
  }

}