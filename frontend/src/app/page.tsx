'use client'

import { useEffect, useState } from 'react'
import { fetchStats, fetchAuctions } from '@/services/api'
// 👉 引入 usePublicClient 用于手动精准监听任何一笔交易
import { useAccount, useConnect, useDisconnect, useWriteContract, useReadContract, usePublicClient } from 'wagmi'
import { injected } from 'wagmi/connectors'
import { parseEther } from 'viem'

// 👉 从刚才建好的常量文件中引入配置
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

  // 👉 1. 新增：首页 Tabs 状态
  const [activeTab, setActiveTab] = useState<'active' | 'ended'>('active')

  // 👉 2. 新增：拍卖详情与出价历史状态
  const [showDetailsModal, setShowDetailsModal] = useState(false)
  const [bidHistory, setBidHistory] = useState<any[]>([])
  const [isLoadingBids, setIsLoadingBids] = useState(false)

  // 👉 新增：为每个核心动作设置独立的 Loading 状态，体验拉满
  const [isMinting, setIsMinting] = useState(false)
  const [isListing, setIsListing] = useState(false)
  const [isBidding, setIsBidding] = useState(false)

  const [processingId, setProcessingId] = useState<string | null>(null)

  // 👉 1. 新增：我的 NFT 画廊状态
  const [showNFTModal, setShowNFTModal] = useState(false)
  const [myNFTs, setMyNFTs] = useState<any[]>([])
  const [isLoadingNFTs, setIsLoadingNFTs] = useState(false)

  const { address, isConnected } = useAccount()
  const { connect } = useConnect()
  const { disconnect } = useDisconnect()

  // 👉 关键核心：引入公共客户端(查收据) 和 异步写入函数
  const publicClient = usePublicClient()
  const { writeContractAsync } = useWriteContract()

  // 实时读取造币厂当前发到了第几号，提取 refetch 方法以便铸造完立刻刷新
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

  const latestTokenId = nextTokenIdRaw ? Number(nextTokenIdRaw) - 1 : 0

  // 👉 替换原有的 useEffect
  useEffect(() => {
    setMounted(true)
    fetchStats().then(setStats)
    
    // 每次 activeTab 改变，就传给 fetchAuctions 去获取新数据
    fetchAuctions(activeTab).then(setAuctions)
    
    const timer = setInterval(() => {
      setMounted(prev => !prev); setMounted(true);
    }, 10000); 
    return () => clearInterval(timer);
  }, [activeTab]) // 👈 别忘了把 activeTab 加到依赖数组里


  // 👉 2. 新增：读取并展示我的 NFT 列表
  const handleOpenMyNFTs = async () => {
    if (!address) return alert("请先连接钱包！")
    setShowNFTModal(true)
    setIsLoadingNFTs(true)
    try {
      // 请确保这里的请求路径与你的后端/代理一致
      const res = await fetch(`http://localhost:8080/api/users/${address}/nfts`)
      const data = await res.json()
      // 假设后端返回 { "nfts": [...] }，请根据实际字段调整
      setMyNFTs(data.nfts || []) 
    } catch (err) {
      console.error("获取个人 NFT 失败", err)
    } finally {
      setIsLoadingNFTs(false)
    }
  }

  // 👉 新增：在画廊点击 NFT 后，直接跳转上架表单
  const handleSelectNFTForAuction = (nft: any) => {
    // 假设后端返回的字段是 contract_address 和 token_id，请根据实际情况替换
    setFormNftContract(nft.contractAddress || nft.contract_address || nft.contract) 
    setFormTokenId(nft.tokenId || nft.token_id)
    setShowNFTModal(false) // 关掉画廊
    setShowModal(true)     // 弹出上架向导（数据已自动填好！）
  }

  // ==========================================
  // 1. 丝滑铸造：带链上监听 + 自动回填表单
  // ==========================================
  const handleMintTestNFT = async () => {
    if (!address) return alert("请先连接钱包！")
    try {
      setIsMinting(true) // 开启 Loading 动画
      
      // 1. 发送铸造交易
      const hash = await writeContractAsync({
        address: MOCK_NFT_ADDRESS,
        abi: ERC721_ABI,
        functionName: 'mint',
        args: [address], 
      })

      // 2. 死死盯住这笔交易，直到它被打包进区块！
      await publicClient!.waitForTransactionReceipt({ hash })

      // 3. 打包成功后，立刻重新查询最新发到了几号
      const { data: newNext } = await refetchNextTokenId()
      const mintedId = Number(newNext) - 1

      // 4. 自动把刚刚铸造的 ID 填入上架表单中，彻底告别手动输入！
      setFormTokenId(mintedId.toString())
      alert(`🎉 铸造成功且已确认上链！\n\n你获得了 Token #${mintedId}\n系统已自动为你填入上架表单，去点击【+ 发布拍卖】吧！`)

    } catch (err) {
      console.error("铸造失败", err)
      alert("交易失败或被取消！")
    } finally {
      setIsMinting(false) // 结束 Loading
    }
  }

  // ==========================================
  // 2. 丝滑上架：一键完成 授权(Approve) + 上架(Create)
  // ==========================================
  const handleApproveAndList = async () => {
    if (!formTokenId || !formStartPrice) return alert("请完整填写信息")
    try {
      setIsListing(true) // 开启 Loading
      
      const startPriceWei = parseEther(formStartPrice)
      const durationSeconds = BigInt(Number(formDuration) * 24 * 60 * 60) 

      // 步骤 A：发送【授权】交易
      const hashApprove = await writeContractAsync({
        address: formNftContract as `0x${string}`,
        abi: ERC721_ABI,
        functionName: 'approve',
        args: [AUCTION_CONTRACT_ADDRESS, BigInt(formTokenId)],
      })
      // 盯盘：等待授权上链...
      await publicClient!.waitForTransactionReceipt({ hash: hashApprove })

      // 步骤 B：授权上链成功后，无缝发送【上架】交易
      const hashCreate = await writeContractAsync({
        address: AUCTION_CONTRACT_ADDRESS,
        abi: AUCTION_ABI,
        functionName: 'createAuction',
        args: [formNftContract as `0x${string}`, BigInt(formTokenId), startPriceWei, BigInt(0), durationSeconds],
      })
      // 盯盘：等待上架上链...
      await publicClient!.waitForTransactionReceipt({ hash: hashCreate })

      // 走到这里说明大功告成！
      setShowModal(false)
      
      // 给 Go 后端 1 秒钟把数据存进 MySQL，然后刷新页面
      setTimeout(() => {
        fetchStats().then(setStats)
        fetchAuctions().then(setAuctions)
      }, 1000)

    } catch (err) {
      console.error("上架流程失败", err)
      alert("流程中断！如果你拒绝了交易，请重新操作。")
    } finally {
      setIsListing(false)
    }
  }

  // ==========================================
  // 3. 丝滑出价：带链上监听
  // ==========================================
  const openBidModal = (item: any) => {
    if (!isConnected) return alert("请先连接钱包！")
    setSelectedItem(item)

    const isFirstBid = item.highest_bid === "0"
    const currentPriceEth = isFirstBid ? Number(item.start_price) / 1e18 : Number(item.highest_bid) / 1e18
    const minRequiredBid = isFirstBid ? currentPriceEth : currentPriceEth * 1.05

    setBidAmountInput(minRequiredBid.toFixed(4).replace(/\.?0+$/, '')) 
    setShowBidModal(true)
  }

// 👉 修正版：打开详情弹窗并拉取出价历史
  const handleOpenDetails = async (item: any) => {
    setSelectedItem(item)
    setShowDetailsModal(true)
    setIsLoadingBids(true)
    try {
      const res = await fetch(`http://localhost:8080/api/auctions/${item.auction_id}/bids`)
      const json = await res.json()
      // ✅ 必须取 json.data，如果后端返回 null 则用 [] 兜底
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
      return alert(`❌ 出价无效！\n\n当前最低要求为: ${minRequiredBid.toFixed(4)} ETH`)
    }

    try {
      setIsBidding(true)
      // 发送出价交易
      const hash = await writeContractAsync({
        address: AUCTION_CONTRACT_ADDRESS,
        abi: AUCTION_ABI,
        functionName: 'bidAuction',
        args: [BigInt(selectedItem.auction_id)],
        value: parseEther(bidAmountInput),
      })
      
      // 盯盘：等待出价被矿工打包...
      await publicClient!.waitForTransactionReceipt({ hash })
      
      setShowBidModal(false) 
      setTimeout(() => {
        fetchStats().then(setStats)
        fetchAuctions().then(setAuctions)
      }, 1000)

    } catch(err) {
      console.error("出价失败", err)
    } finally {
      setIsBidding(false)
    }
  }
  
// 👉 [新增] 取消拍卖方法
  const handleCancelAuction = async (auctionId: string) => {
    try {
      setProcessingId(`${auctionId}-cancel`)
      const hash = await writeContractAsync({
        address: AUCTION_CONTRACT_ADDRESS,
        abi: AUCTION_ABI,
        functionName: 'cancelAuction',
        args: [BigInt(auctionId)],
      })
      await publicClient!.waitForTransactionReceipt({ hash })
      setTimeout(() => {
        fetchStats().then(setStats)
        fetchAuctions().then(setAuctions)
      }, 1500) // 稍微多等一下，让后端数据库更新完
    } catch(err) {
      console.error("取消失败", err)
    } finally {
      setProcessingId(null)
    }
  }

  // 👉 [新增] 结束拍卖方法
  const handleEndAuction = async (auctionId: string) => {
    try {
      setProcessingId(`${auctionId}-end`)
      const hash = await writeContractAsync({
        address: AUCTION_CONTRACT_ADDRESS,
        abi: AUCTION_ABI,
        functionName: 'endAuction',
        args: [BigInt(auctionId)],
      })
      await publicClient!.waitForTransactionReceipt({ hash })
      setTimeout(() => {
        fetchStats().then(setStats)
        fetchAuctions().then(setAuctions)
      }, 1500)
    } catch(err) {
      console.error("结束失败", err)
    } finally {
      setProcessingId(null)
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
              {isMinting ? '⛓️ 铸造上链中...' : '🎁 领测试 NFT'}
            </button>
            <button 
              onClick={() => setShowModal(true)}
              className="bg-green-500 text-white px-5 py-2 rounded-xl font-bold hover:bg-green-600 transition shadow-md shadow-green-200"
            >
              + 发布拍卖
            </button>
            {/* 👉 3. 新增：在上面那个按钮旁边加上这个： */}
            <button 
              onClick={handleOpenMyNFTs}
              className="bg-indigo-50 text-indigo-600 border border-indigo-200 px-5 py-2 rounded-xl font-bold hover:bg-indigo-100 transition"
            >
              🖼️ 我的 NFT
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
        <div className="flex justify-between items-end mb-8 border-b pb-4">
          <h2 className="text-3xl font-black">拍卖大厅 🏛️</h2>
          
          {/* Tabs 切换组件 */}
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
          {auctions.map((item: any) => {
            // 👉 [新增] 动态计算这三个核心状态
            const isSeller = address?.toLowerCase() === item.seller?.toLowerCase();
            const hasNoBids = item.highest_bid === "0";
            // 后端的 end_time 默认是秒级时间戳，乘以 1000 转为毫秒比对
            const isExpired = Date.now() > Number(item.end_time) * 1000;

            return (
              <div key={item.auction_id} className="bg-white rounded-[2.5rem] p-4 shadow-xl shadow-gray-200/50 hover:-translate-y-2 transition-all duration-300 group">
                {/* 1. 修改图片区域：过期加上蒙版和印章 */}
                <div className="aspect-square bg-linear-to-tr from-blue-50 to-purple-50 rounded-4xl mb-4 flex items-center justify-center text-6xl group-hover:scale-95 transition-transform relative">
                  🖼️
                  {isExpired && (
                    <div className="absolute inset-0 bg-white/60 backdrop-blur-[2px] rounded-4xl flex items-center justify-center">
                      <span className="bg-red-500 text-white px-4 py-2 rounded-xl font-bold rotate-[-10deg] shadow-lg">已结束</span>
                    </div>
                  )}
                </div>
                
                <div className="px-2">
                  <div className="flex justify-between items-center mb-2">
                    <h3 className="text-xl font-bold">Token #{item.token_id}</h3>
                    {/* 2. 修改 LIVE 标签：只有未过期才显示 */}
                    {!isExpired && <span className="text-[10px] bg-blue-600 text-white px-2 py-0.5 rounded-full font-bold">LIVE</span>}
                  </div>
                  <p className="text-gray-400 text-xs mb-6 truncate" title={item.nft_contract}>合约: {item.nft_contract}</p>
                  
                  <div className="flex justify-between items-end bg-gray-50 p-4 rounded-2xl">
                    {/* 👇 从这里开始替换 👇 */}
                  <div className="flex justify-between items-center bg-gray-50 p-4 rounded-2xl">
                    
                    {/* 1. 左侧价格区：增加 min-w-0 防止价格过长撑爆容器 */}
                    <div className="min-w-0 pr-2">
                      <p className="text-[10px] text-gray-400 font-bold uppercase">当前最高价</p>
                      <p className={`text-xl font-black truncate ${isExpired ? 'text-gray-500' : 'text-blue-600'}`}>
                        {item.highest_bid === "0" ? (Number(item.start_price) / 1e18).toFixed(4) : (Number(item.highest_bid) / 1e18).toFixed(4)} ETH
                      </p>
                    </div>
                    
                    {/* 2. 右侧按钮区 */}
                    <div className="flex gap-2 shrink-0">
                      
                      {/* 👇 [升级逻辑] 优先判断数据库的闭环状态：包含"已交割(Ended)"和"已取消(Cancelled)" */}
                      {['Ended', 'Cancelled', 'ended', 'cancelled'].includes(item.status || item.Status) ? (
                        <button 
                          disabled
                          className="bg-gray-100 text-gray-400 border border-gray-200 px-5 py-2.5 rounded-xl text-sm font-bold cursor-not-allowed whitespace-nowrap"
                        >
                          {/* 正则判断：如果是取消状态显示已取消，否则显示已完结 */}
                          {/cancel/i.test(item.status || item.Status || '') ? '已取消 🚫' : '已完结 🏁'}
                        </button>
                      ) : (
                        /* 如果还没完结，走之前的逻辑 */
                        <>
                          {isSeller && hasNoBids && !isExpired && (
                            <button 
                              onClick={() => handleCancelAuction(item.auction_id)} 
                              disabled={processingId === `${item.auction_id}-cancel`}
                              className="bg-red-50 text-red-600 border border-red-200 px-4 py-2.5 rounded-xl text-sm font-bold hover:bg-red-100 transition-colors disabled:opacity-50 whitespace-nowrap"
                            >
                              {processingId === `${item.auction_id}-cancel` ? '...' : '取消'}
                            </button>
                          )}

                          {isExpired ? (
                            <button 
                              onClick={() => handleEndAuction(item.auction_id)} 
                              disabled={processingId === `${item.auction_id}-end`}
                              className="bg-green-500 text-white px-5 py-2.5 rounded-xl text-sm font-bold hover:bg-green-600 transition-colors shadow-md disabled:opacity-50 whitespace-nowrap"
                            >
                              {processingId === `${item.auction_id}-end` ? '结算中...' : '结束并交割'}
                            </button>
                          ) : (
                            <button 
                              onClick={() => handleOpenDetails(item)} 
                              className="bg-slate-900 text-white px-5 py-2.5 rounded-xl text-sm font-bold hover:bg-blue-600 transition-colors whitespace-nowrap"
                            >
                            看详情
                            </button>
                          )}
                        </>
                      )}
                    </div>

                  </div>
                  {/* 👆 到这里结束替换 👆 */}
                  </div>
                </div>
              </div>
            )
          })}
        </div>
      </div>

      {/* ========================================================= */}
      {/* 精美的出价弹窗 (Bid Modal)  */}
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
                  disabled={isBidding}
                  className="w-full bg-slate-900 text-white py-4 rounded-xl font-bold shadow-xl shadow-slate-200 hover:bg-blue-600 hover:shadow-blue-200 transition-all duration-300 disabled:opacity-50 text-lg"
                >
                  {isBidding ? '⏳ 链上打包中...' : '确认支付'}
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
      
{/* ==================== 抽离的独立弹窗区域 ==================== */}
      
      {/* 极速上架弹窗 */}
      <CreateAuctionModal 
        isOpen={showModal} onClose={() => setShowModal(false)}
        formNftContract={formNftContract} setFormNftContract={setFormNftContract}
        formTokenId={formTokenId} setFormTokenId={setFormTokenId}
        formStartPrice={formStartPrice} setFormStartPrice={setFormStartPrice}
        formDuration={formDuration} setFormDuration={setFormDuration}
        handleApproveAndList={handleApproveAndList} isListing={isListing}
      />

      {/* 我的 NFT 资产库画廊 */}
      <NFTGalleryModal 
        isOpen={showNFTModal} onClose={() => setShowNFTModal(false)}
        isLoadingNFTs={isLoadingNFTs} myNFTs={myNFTs}
        handleSelectNFTForAuction={handleSelectNFTForAuction}
      />

      {/* 拍卖详情与流水记录 */}
      <AuctionDetailsModal 
        isOpen={showDetailsModal} onClose={() => setShowDetailsModal(false)}
        selectedItem={selectedItem} isLoadingBids={isLoadingBids}
        bidHistory={bidHistory} openBidModal={openBidModal}
      />
      {/* ======================================================= */}

    </main>
  )
}