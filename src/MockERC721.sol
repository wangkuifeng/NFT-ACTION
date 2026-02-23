// src/MockERC721.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import {ERC721} from  "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract MockERC721 is ERC721 {
    // 新增：自动递增的 Token ID 计数器
    uint256 public nextTokenId = 1;

    constructor() ERC721("Mock NFT", "MNFT") {}

    // 修复：现在只需要传入钱包地址，不需要手动传 ID 了
    function mint(address to) external { 
        uint256 currentId = nextTokenId; // 获取当前该发几号
        _mint(to, currentId);            // 铸造
        nextTokenId++;                   // 计数器 +1，留给下一个人
    }
}