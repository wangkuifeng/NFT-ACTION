// src/constants/contracts.ts

// 合约地址配置
export const AUCTION_CONTRACT_ADDRESS = '0xbf1304F2D77D110a32B150792Ee481212926763D'
export const MOCK_NFT_ADDRESS = '0x4E9b354Bc82728dA2A37E1D281CA12401c032652' 

// 1. NFT 合约的 ABI
export const ERC721_ABI = [
  {
    "inputs": [
      { "internalType": "address", "name": "to", "type": "address" },
      { "internalType": "uint256", "name": "tokenId", "type": "uint256" }
    ],
    "name": "approve",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [{ "internalType": "address", "name": "to", "type": "address" }],
    "name": "mint",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]

// 2. 拍卖合约的 ABI
export const AUCTION_ABI = [
  {
    "inputs": [{ "internalType": "uint256", "name": "auctionId", "type": "uint256" }],
    "name": "bidAuction",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "nftContract", "type": "address" },
      { "internalType": "uint256", "name": "tokenId", "type": "uint256" },
      { "internalType": "uint256", "name": "startPrice", "type": "uint256" },
      { "internalType": "uint256", "name": "startTime", "type": "uint256" },
      { "internalType": "uint256", "name": "duration", "type": "uint256" }
    ],
    "name": "createAuction",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [{ "internalType": "uint256", "name": "auctionId", "type": "uint256" }],
    "name": "cancelAuction",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [{ "internalType": "uint256", "name": "auctionId", "type": "uint256" }],
    "name": "endAuction",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]