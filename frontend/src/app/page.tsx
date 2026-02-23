'use client'

import { useEffect, useState } from 'react'
import { fetchStats, fetchAuctions } from '@/services/api'
import { useAccount, useConnect, useDisconnect, useWriteContract, useReadContract, useWaitForTransactionReceipt } from 'wagmi'
import { injected } from 'wagmi/connectors'
import { parseEther } from 'viem'

// 合约地址配置
const AUCTION_CONTRACT_ADDRESS = '0xbf1304F2D77D110a32B150792Ee481212926763D'
const MOCK_NFT_ADDRESS = '0x4E9b354Bc82728dA2A37E1D281CA12401c032652' // 你的模拟造币厂地址

// 1. NFT 合约的 ABI (用于授权和铸造)
const ERC721_ABI = [
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
const AUCTION_ABI = [
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
  }
]

export default function Home() {
  const [stats, setStats] = useState<any>(null)
  const [auctions, setAuctions] = useState([])
  const [mounted, setMounted] = useState(false)
  
  // 弹窗与表单状态
  const [showModal, setShowModal] = useState(false)
  const [formNftContract, setFormNftContract] = useState(MOCK_NFT_ADDRESS) 
  const [formTokenId, setFormTokenId] = useState('')
  const [formStartPrice, setFormStartPrice] = useState('0.01')
  const [formDuration, setFormDuration] = useState('1') 

  // 出价弹窗的状态
  const [showBidModal, setShowBidModal] = useState(false)
  const [selectedItem, setSelectedItem] = useState<any>(null) 
  const [bidAmountInput, setBidAmountInput] = useState('')    

  const { address, isConnected } = useAccount()
  const { connect } = useConnect()
  const { disconnect } = useDisconnect()

  // 1. 获取发送交易的哈希值 (hash)
  const { data: hash, writeContract, isPending } = useWriteContract()

  // 2. 自动盯盘监听器
  const { isLoading: isConfirming, isSuccess: isConfirmed } = useWaitForTransactionReceipt({
    hash,
  })

  // 3. 监听到交易确认成功后，延迟拉取刷新数据
  useEffect(() => {
    if (isConfirmed) {
      setTimeout(() => {
        fetchStats().then(setStats)
        fetchAuctions().then(setAuctions)
      }, 1500) 
    }
  }, [isConfirmed])

  // 实时读取造币厂当前发到了第几号
  const { data: nextTokenIdRaw } = useReadContract({
    address: MOCK_NFT_ADDRESS as `0x${string}`,
    abi: [{
      "inputs": [],
      "name": "nextTokenId",
      "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
      "stateMutability": "view",
      "type": "function"
    }],
    functionName: 'nextTokenId',
  })

  const latestTokenId = nextTokenIdRaw ? Number(nextTokenIdRaw) - 1 : 0

  useEffect(() => {
    setMounted(true)
    fetchStats().then(setStats)
    fetchAuctions().then(setAuctions)
  }, [])

  // 铸造测试 NFT
  const handleMintTestNFT = () => {
    if (!address) return alert("请先连接钱包！")

    writeContract({
      address: MOCK_NFT_ADDRESS,
      abi: ERC721_ABI,
      functionName: 'mint',
      args: [address], 
    }, {
      onSuccess: (hash) => alert(`🎉 铸造请求已发送！\n交易哈希: ${hash}\n\n请等待十几秒区块确认。由于这是自动发号的新工厂，如果你是第一次领，你的 Token ID 就是 1！`),
      onError: (err) => {
        console.error("铸造失败", err)
        alert("铸造失败！")
      }
    })
  }

  // 打开出价弹窗
  const openBidModal = (item: any) => {
    if (!isConnected) return alert("请先连接钱包！")
    setSelectedItem(item)

    const isFirstBid = item.highest_bid === "0"
    const currentPriceEth = isFirstBid ? Number(item.start_price) / 1e18 : Number(item.highest_bid) / 1e18
    const minRequiredBid = isFirstBid ? currentPriceEth : currentPriceEth * 1.05

    setBidAmountInput(minRequiredBid.toFixed(4).replace(/\.?0+$/, '')) 
    setShowBidModal(true)
  }

  // 提交出价
  const submitBid = () => {
    if (!selectedItem || !bidAmountInput) return

    const inputEth = Number(bidAmountInput)
    const isFirstBid = selectedItem.highest_bid === "0"
    const currentPriceEth = isFirstBid ? Number(selectedItem.start_price) / 1e18 : Number(selectedItem.highest_bid) / 1e18
    
    const minRequiredBid = isFirstBid ? currentPriceEth : currentPriceEth * 1.05
    if (isNaN(inputEth) || inputEth < minRequiredBid * 0.9999) {
      return alert(`❌ 出价无效！\n\n规则：每次加价幅度不得低于当前价格的 5%。\n当前最低要求为: ${minRequiredBid.toFixed(4)} ETH`)
    }

    writeContract({
      address: AUCTION_CONTRACT_ADDRESS,
      abi: AUCTION_ABI,
      functionName: 'bidAuction',
      args: [BigInt(selectedItem.auction_id)],
      value: parseEther(bidAmountInput),
    }, {
      onSuccess: () => setShowBidModal(false) 
    })
  }
  
  // 步骤 1 - 授权 NFT
  const handleApprove = () => {
    if (!formTokenId) return alert("请填写 Token ID")
    writeContract({
      address: formNftContract as `0x${string}`,
      abi: ERC721_ABI,
      functionName: 'approve',
      args: [AUCTION_CONTRACT_ADDRESS, BigInt(formTokenId)],
    }, {
      onSuccess: () => alert("✅ 授权交易已发送！请等待狐狸钱包提示【交易确认】后，再点击下方的【步骤 2: 确认上架】。"),
      onError: (err) => console.error("授权失败", err)
    })
  }

  // 步骤 2 - 确认上架
  const handleCreateAuction = () => {
    if (!formTokenId || !formStartPrice) return alert("请完整填写信息")
    const startPriceWei = parseEther(formStartPrice)
    const durationSeconds = BigInt(Number(formDuration) * 24 * 60 * 60) 

    writeContract({
      address: AUCTION_CONTRACT_ADDRESS,
      abi: AUCTION_ABI,
      functionName: 'createAuction',
      args: [formNftContract as `0x${string}`, BigInt(formTokenId), startPriceWei, BigInt(0), durationSeconds],
    }, {
      onSuccess: () => setShowModal(false) 
    })
  }

  return (
    <main className="min-h-screen bg-gray-50 text-slate-900 p-8">
      {/* 头部：标题与钱包 */}
      <div className="max-w-6xl mx-auto flex justify-between items-center mb-12">
        <div>
          <h1 className="text-4xl font-black text-blue-600 tracking-tighter">NFT AUCTION</h1>
          <p className="text-gray-400 text-sm font-medium">去中心化拍卖平台</p>
        </div>

        {mounted && (isConnected ? (
          <div className="flex items-center gap-3">
            <button 
              onClick={handleMintTestNFT} disabled={isPending}
              className="bg-purple-100 text-purple-600 px-4 py-2 rounded-xl font-bold hover:bg-purple-200 transition border border-purple-200"
            >
              🎁 领测试 NFT
            </button>
            <button 
              onClick={() => setShowModal(true)}
              className="bg-green-500 text-white px-5 py-2 rounded-xl font-bold hover:bg-green-600 transition shadow-md shadow-green-200"
            >
              + 发布拍卖
            </button>
            <span className="text-sm font-mono bg-blue-50 text-blue-600 px-3 py-2 rounded-xl border border-blue-100">
              {address?.slice(0, 6)}...{address?.slice(-4)}
            </span>
            <button onClick={() => disconnect()} className="text-sm text-gray-400 hover:text-red-500">断开</button>
          </div>
        ) : (
          <button 
            onClick={() => connect({ connector: injected() })}
            className="bg-blue-600 text-white px-8 py-3 rounded-2xl font-bold shadow-lg shadow-blue-200 hover:scale-105 transition"
          >
            连接钱包
          </button>
        ))}
      </div>

      {/* 统计栏 */}
      <div className="max-w-6xl mx-auto grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
        <div className="bg-white p-6 rounded-3xl shadow-sm border border-gray-100">
          <p className="text-gray-400 text-xs font-bold uppercase mb-1">拍卖总数</p>
          <p className="text-3xl font-black">{stats?.total_auctions || 0}</p>
        </div>
        <div className="bg-white p-6 rounded-3xl shadow-sm border border-gray-100">
          <p className="text-gray-400 text-xs font-bold uppercase mb-1">累计出价</p>
          <p className="text-3xl font-black">{stats?.total_bids || 0}</p>
        </div>
        <div className="bg-white p-6 rounded-3xl shadow-sm border border-gray-100">
          <p className="text-gray-400 text-xs font-bold uppercase mb-1">当前网络</p>
          <p className="text-3xl font-black text-blue-500">Sepolia Testnet</p>
        </div>
      </div>

      {/* 拍卖卡片网格 */}
      <div className="max-w-6xl mx-auto">
        <h2 className="text-2xl font-black mb-8">热门拍卖 🔥</h2>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
          {auctions.map((item: any) => (
            <div key={item.auction_id} className="bg-white rounded-[2.5rem] p-4 shadow-xl shadow-gray-200/50 hover:-translate-y-2 transition-all duration-300 group">
              <div className="aspect-square bg-linear-to-tr from-blue-50 to-purple-50 rounded-4xl mb-4 flex items-center justify-center text-6xl group-hover:scale-95 transition-transform">
                🖼️
              </div>
              
              <div className="px-2">
                <div className="flex justify-between items-center mb-2">
                  <h3 className="text-xl font-bold">Token #{item.token_id}</h3>
                  <span className="text-[10px] bg-blue-600 text-white px-2 py-0.5 rounded-full font-bold">LIVE</span>
                </div>
                <p className="text-gray-400 text-xs mb-6 truncate" title={item.nft_contract}>合约: {item.nft_contract}</p>
                
                <div className="flex justify-between items-end bg-gray-50 p-4 rounded-2xl">
                  <div>
                    <p className="text-[10px] text-gray-400 font-bold uppercase">当前最高价</p>
                    <p className="text-xl font-black text-blue-600">
                      {item.highest_bid === "0" ? (Number(item.start_price) / 1e18).toFixed(3) : (Number(item.highest_bid) / 1e18).toFixed(3)} ETH
                    </p>
                  </div>
                  <button 
                    onClick={() => openBidModal(item)} 
                    disabled={isPending || isConfirming}
                    className="bg-slate-900 text-white px-5 py-2.5 rounded-xl text-sm font-bold hover:bg-blue-600 transition-colors disabled:opacity-50"
                  >
                    {isPending ? '钱包确认...' : isConfirming ? '链上打包 ⏳' : '参与竞拍'}
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>

      {/* ========================================================= */}
      {/* 上架拍卖的弹窗 (Modal) */}
      {/* ========================================================= */}
      {showModal && (
        <div className="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-3xl p-8 max-w-md w-full shadow-2xl">
            <div className="flex justify-between items-center mb-6">
              <h2 className="text-2xl font-black">上架 NFT</h2>
              <button onClick={() => setShowModal(false)} className="text-gray-400 hover:text-red-500 font-bold">✕</button>
            </div>
            
            <div className="space-y-4">
              <div>
                <label className="block text-xs font-bold text-gray-500 mb-1">NFT 合约地址 (已自动填入测试造币厂)</label>
                <input value={formNftContract} onChange={e => setFormNftContract(e.target.value)} className="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-2 text-sm focus:outline-blue-500" />
              </div>
              <div className="grid grid-cols-2 gap-4">
              <div>
                  <label className="block text-xs font-bold text-gray-500 mb-1">Token ID</label>
                  <input type="number" placeholder="例如: 1" value={formTokenId} onChange={e => setFormTokenId(e.target.value)} className="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-2 text-sm focus:outline-blue-500" />
                  
                  {latestTokenId > 0 && (
                    <p className="text-[10px] text-blue-500 font-bold mt-1">
                      💡 提示：最新铸造出的 ID 是 {latestTokenId} (如果你刚领完，填这个准没错！)
                    </p>
                  )}
                </div>
                <div>
                  <label className="block text-xs font-bold text-gray-500 mb-1">起拍价 (ETH)</label>
                  <input type="number" step="0.01" value={formStartPrice} onChange={e => setFormStartPrice(e.target.value)} className="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-2 text-sm focus:outline-blue-500" />
                </div>
              </div>
              <div>
                <label className="block text-xs font-bold text-gray-500 mb-1">拍卖时长 (天)</label>
                <input type="number" value={formDuration} onChange={e => setFormDuration(e.target.value)} className="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-2 text-sm focus:outline-blue-500" />
              </div>

              <div className="mt-8 space-y-3">
                <button 
                  onClick={handleApprove} disabled={isPending}
                  className="w-full bg-blue-50 text-blue-600 border border-blue-200 py-3 rounded-xl font-bold hover:bg-blue-100 transition disabled:opacity-50"
                >
                  步骤 1: 授权 NFT (Approve)
                </button>
                <button 
                  onClick={handleCreateAuction} disabled={isPending}
                  className="w-full bg-blue-600 text-white py-3 rounded-xl font-bold shadow-lg shadow-blue-200 hover:bg-blue-700 transition disabled:opacity-50"
                >
                  步骤 2: 确认上架 (Create)
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

      {/* ========================================================= */}
      {/* 👉 新增：精美的出价弹窗 (Bid Modal)  */}
      {/* ========================================================= */}
      {showBidModal && selectedItem && (
        <div className="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-3xl p-8 max-w-md w-full shadow-2xl">
            <div className="flex justify-between items-center mb-6">
              <h2 className="text-2xl font-black text-slate-900">参与竞拍</h2>
              <button onClick={() => setShowBidModal(false)} className="text-gray-400 hover:text-red-500 font-bold">✕</button>
            </div>
            
            <div className="bg-blue-50/50 p-4 rounded-2xl mb-6 border border-blue-100">
              <div className="flex justify-between text-sm mb-2">
                <span className="text-gray-500 font-medium">拍卖品:</span>
                <span className="font-bold text-slate-800">Token #{selectedItem.token_id}</span>
              </div>
              <div className="flex justify-between text-sm mb-2">
                <span className="text-gray-500 font-medium">当前最高价:</span>
                <span className="font-bold text-blue-600">
                  {selectedItem.highest_bid === "0" 
                    ? (Number(selectedItem.start_price) / 1e18).toFixed(4) 
                    : (Number(selectedItem.highest_bid) / 1e18).toFixed(4)} ETH
                </span>
              </div>
              <div className="flex justify-between text-sm">
                <span className="text-gray-500 font-medium">加价规则:</span>
                <span className="text-green-600 font-bold">至少加价 5%</span>
              </div>
            </div>

            <div className="space-y-4">
              <div>
                <label className="block text-xs font-bold text-gray-500 mb-2">你的出价 (ETH)</label>
                <div className="relative">
                  <input 
                    type="number" 
                    step="0.0001"
                    value={bidAmountInput} 
                    onChange={e => setBidAmountInput(e.target.value)} 
                    className="w-full bg-gray-50 border border-gray-200 rounded-xl pl-4 pr-16 py-4 text-xl font-black text-blue-600 focus:outline-blue-500 focus:bg-white transition-colors" 
                  />
                  <span className="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 font-bold">ETH</span>
                </div>
              </div>

              <div className="mt-8">
                <button 
                  onClick={submitBid} 
                  disabled={isPending}
                  className="w-full bg-slate-900 text-white py-4 rounded-xl font-bold shadow-xl shadow-slate-200 hover:bg-blue-600 hover:shadow-blue-200 transition-all duration-300 disabled:opacity-50 text-lg"
                >
                  {isPending ? '唤起钱包中...' : '确认支付'}
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
      
    </main>
  )
}