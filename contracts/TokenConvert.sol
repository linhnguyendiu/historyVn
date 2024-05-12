// // SPDX-License-Identifier: MIT
// pragma solidity ^0.8.20;

// import "./LinkToken.sol";
// import "@openzeppelin/contracts/access/Ownable.sol";
// contract TokenConverter is Ownable {
//     LINKToken private token;

//     event TokensConverted(address indexed user, uint256 amount);

//     constructor(LINKToken _token) Ownable(msg.sender) { 
//         token = _token;
//     }

//     function convertTokensToEther(uint256 _amount) external {
//         require(_amount > 0, "Amount must be greater than 0");
//         require(token.balanceOf(msg.sender) >= _amount, "Insufficient balance");

//         // Transfer tokens from user to contract
//         token.transferFrom(msg.sender, address(this), _amount);

//         // Convert tokens to ether and send to user
//         payable(msg.sender).transfer(_amount);

//         emit TokensConverted(msg.sender, _amount);
//     }

//     function withdrawEther() external onlyOwner {
//         payable(owner()).transfer(address(this).balance);
//     }
// }