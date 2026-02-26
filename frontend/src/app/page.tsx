'use client'

import { useEffect, useState } from 'react'
import { fetchStats, fetchAuctions } from '@/services/api'
import { useAccount, useConnect, useDisconnect, useWriteContract, useReadContract, usePublicClient } from 'wagmi'
import { injected } from 'wagmi/connectors'
import { parseEther } from 'viem'
import Link from 'next/link'
import toast from 'react-hot-toast' // 👉 引入丝滑弹窗
import AuctionCard from '@/components/AuctionCard'

import { 
  AUCTION_CONTRACT_ADDRESS, 
  MOCK_NFT_ADDRESS, 
  ERC721_ABI, 
  AUCTION_ABI 
} from '@/constants/contracts'

import { 
  CreateAuctionModal, 
  NFTGalleryModal, 
  AuctionDetailsModal 
} from '@/components/modals/AllModals'

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

  // 首页 Tabs 状态
  const [activeTab, setActiveTab] = useState<'active' | 'ended'>('active')

  // 拍卖详情与出价历史状态
  const [showDetailsModal, setShowDetailsModal] = useState(false)
  const [bidHistory, setBidHistory] = useState<any[]>([])
  const [isLoadingBids, setIsLoadingBids] = useState(false)

  // 独立 Loading 状态
  const [isMinting, setIsMinting] = useState(false)
  const [isListing, setIsListing] = useState(false)
  const [isBidding, setIsBidding] = useState(false)

  // 我的 NFT 画廊状态
  const [showNFTModal, setShowNFTModal] = useState(false)
  const [myNFTs, setMyNFTs] = useState<any[]>([])
  const [isLoadingNFTs, setIsLoadingNFTs] = useState(false)

  const { address, isConnected } = useAccount()
  const { connect } = useConnect()
  const { disconnect } = useDisconnect()

  const publicClient = usePublicClient()
  const { writeContractAsync } = useWriteContract()

  // 实时读取造币厂当前发到了第几号
  const { data: nextTokenIdRaw, refetch: refetchNextTokenId } = useReadContract({
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

  useEffect(() => {
    setMounted(true)
    fetchStats().then(setStats)
    
    // 每次 activeTab 改变，重新获取数据
    fetchAuctions(activeTab).then(setAuctions)
    
    const timer = setInterval(() => {
      fetchAuctions(activeTab).then(setAuctions);
    }, 10000); 
    return () => clearInterval(timer);
  }, [activeTab]) 

  // 读取并展示我的 NFT 列表
  const handleOpenMyNFTs = async () => {
    if (!address) return toast.error("请先连接钱包！") // 👉 替换 alert
    setShowNFTModal(true)
    setIsLoadingNFTs(true)
    try {
      const res = await fetch(`http://localhost:8080/api/users/${address}/nfts`)
      const data = await res.json()
      setMyNFTs(data.data || []) 
    } catch (err) {
      console.error("获取个人 NFT 失败", err)
      toast.error("获取资产失败，请检查网络")
    } finally {
      setIsLoadingNFTs(false)
    }
  }

// 画廊点击 NFT 后跳转上架表单 (已修复 [object Object] 问题)
  const handleSelectNFTForAuction = (nft: any) => {
    // 1. 精准提取合约地址：处理 Alchemy 的 { contract: { address: "..." } } 结构
    let address = MOCK_NFT_ADDRESS;
    if (nft.contract?.address) {
      address = nft.contract.address;
    } else if (typeof nft.contract === 'string') {
      address = nft.contract;
    } else if (nft.contractAddress) {
      address = nft.contractAddress;
    }

    // 2. 精准提取 Token ID：处理 Alchemy 的 { id: { tokenId: "..." } } 结构
    let tId = "";
    const rawTokenId = nft.id?.tokenId || nft.tokenId || "";
    if (rawTokenId) {
      // Alchemy 有时会返回十六进制的 tokenId (例如 "0x0c")，我们将它转为十进制，方便输入框显示
      tId = rawTokenId.toString().startsWith('0x') 
        ? BigInt(rawTokenId).toString() 
        : rawTokenId;
    }

    // 3. 回填表单并打开弹窗
    setFormNftContract(address);
    setFormTokenId(tId);
    setShowNFTModal(false); 
    setShowModal(true);    
  }
  // ==========================================
  // 1. 丝滑铸造 (接入 toast.promise)
  // ==========================================
  const handleMintTestNFT = async () => {
    if (!address) return toast.error("请先连接钱包！")
    try {
      setIsMinting(true) 
      const hash = await writeContractAsync({
        address: MOCK_NFT_ADDRESS,
        abi: ERC721_ABI,
        functionName: 'mint',
        args: [address], 
      })

      // 👉 核心：用 toast.promise 包装等待上链的过程
      await toast.promise(
        publicClient!.waitForTransactionReceipt({ hash }),
        {
          loading: '⛓️ 铸造上链中，请在钱包确认并等待区块打包...',
          success: '🎉 铸造成功且已确认上链！',
          error: '❌ 铸造失败，请重试',
        }
      )

      const { data: newNext } = await refetchNextTokenId()
      const mintedId = Number(newNext) - 1

      setFormTokenId(mintedId.toString())
      toast.success(`你获得了 Token #${mintedId}\n系统已自动为你填入表单，快去发布吧！`, { duration: 5000, icon: '🎁' })
    } catch (err) {
      console.error("铸造失败", err)
      toast.error("交易被取消或发起失败") // 👉 替换 alert
    } finally {
      setIsMinting(false) 
    }
  }

  // ==========================================
  // 2. 丝滑上架 (接入双重 toast.promise)
  // ==========================================
  const handleApproveAndList = async () => {
    if (!formTokenId || !formStartPrice) return toast.error("请完整填写信息")
    try {
      setIsListing(true) 
      const startPriceWei = parseEther(formStartPrice)
      const durationSeconds = BigInt(Number(formDuration) * 24 * 60 * 60) 

      // 步骤 A：授权交易
      const hashApprove = await writeContractAsync({
        address: formNftContract as `0x${string}`,
        abi: ERC721_ABI,
        functionName: 'approve',
        args: [AUCTION_CONTRACT_ADDRESS, BigInt(formTokenId)],
      })
      await toast.promise(
        publicClient!.waitForTransactionReceipt({ hash: hashApprove }),
        {
          loading: '🔓 正在授权 NFT，请耐心等待上链...',
          success: '✅ 授权成功！准备发起上架交易...',
          error: '❌ 授权失败',
        }
      )

      // 步骤 B：上架交易
      const hashCreate = await writeContractAsync({
        address: AUCTION_CONTRACT_ADDRESS,
        abi: AUCTION_ABI,
        functionName: 'createAuction',
        args: [formNftContract as `0x${string}`, BigInt(formTokenId), startPriceWei, BigInt(0), durationSeconds],
      })
      await toast.promise(
        publicClient!.waitForTransactionReceipt({ hash: hashCreate }),
        {
          loading: '📢 正在上架拍卖，等待区块打包...',
          success: '🎉 上架成功！你的 NFT 已进入拍卖大厅！',
          error: '❌ 上架失败',
        }
      )

      setShowModal(false)
      setTimeout(() => {
        fetchStats().then(setStats)
        fetchAuctions(activeTab).then(setAuctions)
      }, 1000)

    } catch (err) {
      console.error("上架流程失败", err)
      toast.error("流程中断！如果你拒绝了交易，请重新操作。")
    } finally {
      setIsListing(false)
    }
  }

  // ==========================================
  // 3. 丝滑出价 (接入 toast.promise)
  // ==========================================
  const openBidModal = (item: any) => {
    if (!isConnected) return toast.error("请先连接钱包！")
    setSelectedItem(item)

    const isFirstBid = item.highest_bid === "0"
    const currentPriceEth = isFirstBid ? Number(item.start_price) / 1e18 : Number(item.highest_bid) / 1e18
    const minRequiredBid = isFirstBid ? currentPriceEth : currentPriceEth * 1.05

    setBidAmountInput(minRequiredBid.toFixed(4).replace(/\.?0+$/, '')) 
    setShowBidModal(true)
  }

  const handleOpenDetails = async (item: any) => {
    setSelectedItem(item)
    setShowDetailsModal(true)
    setIsLoadingBids(true)
    try {
      const res = await fetch(`http://localhost:8080/api/auctions/${item.auction_id}/bids`)
      const json = await res.json()
      setBidHistory(json.data || []) 
    } catch (err) {
      console.error("获取出价历史失败", err)
      setBidHistory([])
    } finally {
      setIsLoadingBids(false)
    }
  }

  const submitBid = async () => {
    if (!selectedItem || !bidAmountInput) return
    const inputEth = Number(bidAmountInput)
    const isFirstBid = selectedItem.highest_bid === "0"
    const currentPriceEth = isFirstBid ? Number(selectedItem.start_price) / 1e18 : Number(selectedItem.highest_bid) / 1e18
    
    const minRequiredBid = isFirstBid ? currentPriceEth : currentPriceEth * 1.05
    if (isNaN(inputEth) || inputEth < minRequiredBid * 0.9999) {
      return toast.error(`出价无效！当前最低要求为: ${minRequiredBid.toFixed(4)} ETH`) // 👉 替换 alert
    }

    try {
      setIsBidding(true)
      const hash = await writeContractAsync({
        address: AUCTION_CONTRACT_ADDRESS,
        abi: AUCTION_ABI,
        functionName: 'bidAuction',
        args: [BigInt(selectedItem.auction_id)],
        value: parseEther(bidAmountInput),
      })
      
      await toast.promise(
        publicClient!.waitForTransactionReceipt({ hash }),
        {
          loading: '🔨 出价上链中，请耐心等待区块打包...',
          success: '🎉 出价成功！你目前是最高出价者！',
          error: '❌ 出价失败，请重试',
        }
      )
      
      setShowBidModal(false) 
      setTimeout(() => {
        fetchStats().then(setStats)
        fetchAuctions(activeTab).then(setAuctions)
      }, 1000)

    } catch(err) {
      console.error("出价失败", err)
      toast.error("交易已被取消或发生错误")
    } finally {
      setIsBidding(false)
    }
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
              onClick={handleMintTestNFT} disabled={isMinting}
              className="bg-purple-100 text-purple-600 px-4 py-2 rounded-xl font-bold hover:bg-purple-200 transition border border-purple-200 disabled:opacity-50"
            >
              {isMinting ? '⛓️ 处理中...' : '🎁 领测试 NFT'}
            </button>
            <button 
              onClick={() => setShowModal(true)}
              className="bg-green-500 text-white px-5 py-2 rounded-xl font-bold hover:bg-green-600 transition shadow-md shadow-green-200"
            >
              + 发布拍卖
            </button>
            <button 
              onClick={handleOpenMyNFTs}
              className="bg-indigo-50 text-indigo-600 border border-indigo-200 px-5 py-2 rounded-xl font-bold hover:bg-indigo-100 transition"
            >
              🖼️ 我的 NFT
            </button>
            
            <Link 
              href="/profile" 
              className="bg-gray-800 text-white px-5 py-2 rounded-xl font-bold hover:bg-gray-900 transition shadow-md"
            >
              👤 个人中心
            </Link>

            <span className="text-sm font-mono bg-blue-50 text-blue-600 px-3 py-2 rounded-xl border border-blue-100">
              {address?.slice(0, 6)}...{address?.slice(-4)}
            </span>
            <button onClick={() => { disconnect(); toast.success('钱包已断开'); }} className="text-sm text-gray-400 hover:text-red-500">断开</button>
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

      {/* 拍卖大厅 */}
      <div className="max-w-6xl mx-auto">
        <div className="flex justify-between items-end mb-8 border-b pb-4">
          <h2 className="text-3xl font-black">拍卖大厅 🏛️</h2>
          
          <div className="flex bg-gray-100 p-1 rounded-xl">
            <button 
              onClick={() => setActiveTab('active')}
              className={`px-6 py-2 rounded-lg font-bold text-sm transition-all ${activeTab === 'active' ? 'bg-white shadow-sm text-blue-600' : 'text-gray-500 hover:text-gray-700'}`}
            >
              🔥 进行中
            </button>
            <button 
              onClick={() => setActiveTab('ended')}
              className={`px-6 py-2 rounded-lg font-bold text-sm transition-all ${activeTab === 'ended' ? 'bg-white shadow-sm text-gray-800' : 'text-gray-500 hover:text-gray-700'}`}
            >
              🏁 已结束
            </button>
          </div>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
          {auctions.map((item: any) => (
            <AuctionCard 
              key={item.auction_id} 
              item={item} 
              currentUserAddress={address} 
              onRefresh={() => {
                fetchStats().then(setStats)
                fetchAuctions(activeTab).then(setAuctions)
              }}
              onOpenDetails={handleOpenDetails}
            />
          ))}
        </div>
      </div>

      {/* 精美的出价弹窗 (Bid Modal)  */}
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
                  disabled={isBidding}
                  className="w-full bg-slate-900 text-white py-4 rounded-xl font-bold shadow-xl shadow-slate-200 hover:bg-blue-600 hover:shadow-blue-200 transition-all duration-300 disabled:opacity-50 text-lg"
                >
                  {isBidding ? '⏳ 钱包确认中...' : '确认支付'}
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
      
      {/* 抽离的独立弹窗区域 */}
      <CreateAuctionModal 
        isOpen={showModal} onClose={() => setShowModal(false)}
        formNftContract={formNftContract} setFormNftContract={setFormNftContract}
        formTokenId={formTokenId} setFormTokenId={setFormTokenId}
        formStartPrice={formStartPrice} setFormStartPrice={setFormStartPrice}
        formDuration={formDuration} setFormDuration={setFormDuration}
        handleApproveAndList={handleApproveAndList} isListing={isListing}
      />

      <NFTGalleryModal 
        isOpen={showNFTModal} onClose={() => setShowNFTModal(false)}
        isLoadingNFTs={isLoadingNFTs} myNFTs={myNFTs}
        handleSelectNFTForAuction={handleSelectNFTForAuction}
      />

      <AuctionDetailsModal 
        isOpen={showDetailsModal} onClose={() => setShowDetailsModal(false)}
        selectedItem={selectedItem} isLoadingBids={isLoadingBids}
        bidHistory={bidHistory} openBidModal={openBidModal}
      />
    </main>
  )
}