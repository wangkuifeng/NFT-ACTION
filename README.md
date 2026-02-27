🏛️ NFT Auction DApp (去中心化 NFT 拍卖平台)
一个具备完整商业闭环的去中心化 NFT 拍卖平台。本项目采用事件驱动架构，通过 Go 后端实时监听链上状态，结合 Next.js 提供丝滑的 Web3 交互体验，并内置了基于 UUPS 代理模式的可升级智能合约与平台抽成经济模型。

✨ 核心功能与亮点 (Features)
🎨 前端体验 (Frontend - Next.js + React)
丝滑交互 (UX)：全面抛弃原生 Alert，深度集成 react-hot-toast，实现“钱包唤起 -> 链上打包中⏳ -> 交易确认🎉”的全链路精美弹窗反馈。

智能资金面板：为卖家提供实时的“交割资金预估”小票（动态计算最高出价、2% 平台手续费及预计净收入）。

全局网络守卫：内置 <NetworkChecker />，当用户连接非目标网络（如 Sepolia 测试网）时，全局拦截并提供一键切换网络功能。

TVL 与数据大盘：基于高精度计算，实时呈现全平台进行中拍卖的 ETH 锁仓总额 (TVL) 及核心统计数据。

个人中心 (Profile)：独立的资产仪表盘，支持一键切换“我发布的”与“我参与的”拍卖历史。

⚙️ 智能合约 (Smart Contracts - Solidity + Foundry)
UUPS 可升级架构：基于 OpenZeppelin 的 UUPS 代理模式，支持合约逻辑的无缝热升级（已成功完成包含状态变量新增的复杂逻辑升级）。

平台经济模型：内置基于基点 (Basis Points, BPS) 的高精度手续费引擎。在 endAuction 交割时，自动完成资金切分，将利润打入平台金库。

安全机制：

自动退款：买家抢拍出价时，智能合约使用底层的 call 方法自动将上一任最高出价者的资金原路退回。

状态锁与防重入：严谨的 Checks-Effects-Interactions 模式，避免重入攻击；拍卖状态（活跃/结束/流拍）多重校验。

📡 后端服务 (Backend - Go)
实时链上监听：通过 go-ethereum (abigen) 生成原生 Go 绑定代码，实时监听 AuctionCreated、BidPlaced、AuctionEnded 等合约事件。

高并发处理：平稳处理 WebSocket 节点断连或 API 嵌套数据，内置容错重试退出机制，确保链下数据库与链上状态的最终一致性。

多维度 API：提供高效的 RESTful API 接口，支持复杂维度的过滤查询（如卖家、竞拍者历史记录、NFT 画廊回填等）。

🛠️ 技术栈 (Tech Stack)
Smart Contracts: Solidity ^0.8.20, Foundry, OpenZeppelin (Upgradeable)

Frontend: Next.js (App/Pages Router), React, Tailwind CSS, Wagmi, Viem, react-hot-toast

Backend: Golang, GORM, go-ethereum (Geth)

Network: Sepolia Testnet (Ethereum)

Infrastructure: Alchemy / Infura RPC Nodes

📖 核心业务流程 (Workflow)
铸造与上架 (Mint & List)

用户可免费铸造 Mock ERC721 测试 NFT。

授权 (Approve) 并设置起拍价与倒计时，一键上架进入拍卖大厅。

参与竞拍 (Bid)

买家浏览大厅，发起竞拍（必须高于当前最高价 5%）。

智能合约自动将上一个买家锁定的 ETH 全额退还。

平台交割与分润 (Settlement & Fee)

倒计时结束后，任何人可触发交割 (endAuction)。

平台自动扣除 2% 服务费进入金库，剩余 98% 资金自动打入卖家钱包，NFT 转移至获胜买家。

卖家反悔 (Cancel)

在无人出价的前提下，卖家随时可以安全取消拍卖，原路退回 NFT。

📸 项目截图 (Screenshots)
(提示：你可以在 GitHub 上传刚才发给我的那两张极其漂亮的前端截图，并在这里替换为真实的图片链接)

1. 拍卖大厅与数据展板
展示了清爽的 UI、TVL 统计、进行中/已结束的选项卡。

2. 卖家专属净收入面板
展示了拍卖卡片中内置的平台手续费与预计收入计算器。

🚀 后续规划 (Roadmap)
[ ] 接入跨链支付，支持多币种（如 USDC/USDT）竞拍。

[ ] 引入 IPFS 去中心化存储，支持用户自由上传和铸造真实艺术品 NFT。

[ ] 优化后端 WebSocket 架构，向前端推送实时出价通知 (Server-Sent Events)。