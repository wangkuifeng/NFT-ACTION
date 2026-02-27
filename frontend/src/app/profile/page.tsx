// app/profile/page.tsx (或者 src/app/profile/page.tsx)
'use client'

import { useState, useEffect } from 'react'
import { useAccount, usePublicClient, useWriteContract } from 'wagmi'
import Link from 'next/link'
import AuctionCard from '@/components/AuctionCard'
import { AuctionDetailsModal } from '@/components/modals/AllModals'
import { parseEther } from 'viem'
import { AUCTION_CONTRACT_ADDRESS, AUCTION_ABI } from '@/constants/contracts'

export default function ProfilePage() {
  const { address, isConnected } = useAccount()
  const [mounted, setMounted] = useState(false)
  
  // Tabs 状态
  const [activeTab, setActiveTab] = useState<'CREATED' | 'PARTICIPATED'>('CREATED')
  const [subTab, setSubTab] = useState<'ACTIVE' | 'ENDED'>('ACTIVE')

  // 数据状态
  const [stats, setStats] = useState({ created_count: 0, participated_count: 0 })
  const [auctions, setAuctions] = useState([])
  const [isLoading, setIsLoading] = useState(false)

  // 弹窗状态 (复用主页逻辑)
  const [showDetailsModal, setShowDetailsModal] = useState(false)
  const [selectedItem, setSelectedItem] = useState<any>(null)
  const [bidHistory, setBidHistory] = useState<any[]>([])
  const [isLoadingBids, setIsLoadingBids] = useState(false)
  const [showBidModal, setShowBidModal] = useState(false)
  const [bidAmountInput, setBidAmountInput] = useState('')
  const [isBidding, setIsBidding] = useState(false)

  const publicClient = usePublicClient()
  const { writeContractAsync } = useWriteContract()

  // 获取统计数据
  const loadStats = () => {
    if (!address) return
    fetch(`http://localhost:8080/api/users/${address}/stats`)
      .then(res => res.json())
      .then(json => {
        if (json.success) setStats(json.data)
      })
  }

  // 获取列表数据
  const loadAuctions = () => {
    if (!address) return
    setIsLoading(true)
    let url = `http://localhost:8080/api/auctions?`
    if (activeTab === 'CREATED') {
      url += `seller=${address}&status=${subTab === 'ACTIVE' ? 'active' : 'ended'}`
    } else {
      url += `bidder=${address}`
    }
    fetch(url)
      .then(res => res.json())
      .then(json => {
        if(json.success) setAuctions(json.data)
      })
      .finally(() => setIsLoading(false))
  }

  useEffect(() => {
    setMounted(true)
    loadStats()
  }, [address])

  useEffect(() => {
    loadAuctions()
  }, [address, activeTab, subTab])

  // --- 弹窗与出价逻辑 (与主页保持一致) ---
  const handleOpenDetails = async (item: any) => {
    setSelectedItem(item)
    setShowDetailsModal(true)
    setIsLoadingBids(true)
    try {
      const res = await fetch(`http://localhost:8080/api/auctions/${item.auction_id}/bids`)
      const json = await res.json()
      setBidHistory(json.data || []) 
    } catch (err) {
      setBidHistory([])
    } finally {
      setIsLoadingBids(false)
    }
  }

  const openBidModal = (item: any) => {
    setSelectedItem(item)
    const isFirstBid = item.highest_bid === "0"
    const currentPriceEth = isFirstBid ? Number(item.start_price) / 1e18 : Number(item.highest_bid) / 1e18
    const minRequiredBid = isFirstBid ? currentPriceEth : currentPriceEth * 1.05
    setBidAmountInput(minRequiredBid.toFixed(4).replace(/\.?0+$/, '')) 
    setShowBidModal(true)
  }

  const submitBid = async () => {
    if (!selectedItem || !bidAmountInput) return
    try {
      setIsBidding(true)
      const hash = await writeContractAsync({
        address: AUCTION_CONTRACT_ADDRESS,
        abi: AUCTION_ABI,
        functionName: 'bidAuction',
        args: [BigInt(selectedItem.auction_id)],
        value: parseEther(bidAmountInput),
      })
      await publicClient!.waitForTransactionReceipt({ hash })
      setShowBidModal(false) 
      setTimeout(() => { loadStats(); loadAuctions(); }, 1000)
    } catch(err) {
      console.error("出价失败", err)
    } finally {
      setIsBidding(false)
    }
  }

  if (!mounted) return null
  if (!isConnected) {
    return (
      <div className="flex flex-col items-center justify-center min-h-screen bg-gray-50 text-slate-900">
        <h2 className="text-2xl font-bold mb-4">请先连接钱包</h2>
        <Link href="/" className="text-blue-600 hover:underline font-bold">🏠 返回首页大厅</Link>
      </div>
    )
  }

  return (
    <main className="min-h-screen bg-gray-50 text-slate-900 p-8">
      <div className="max-w-6xl mx-auto flex justify-between items-center mb-12">
        <h1 className="text-4xl font-black text-slate-900 tracking-tighter">个人中心 Profile</h1>
        <Link href="/" className="bg-white border border-gray-200 px-6 py-2 rounded-xl font-bold hover:bg-gray-50 transition shadow-sm">
          🏠 返回大厅
        </Link>
      </div>

      <div className="max-w-6xl mx-auto">
        {/* 数据统计卡片 */}
        <div className="bg-slate-900 rounded-3xl p-8 mb-8 text-white shadow-xl flex flex-col md:flex-row gap-8 justify-between items-center">
          <div>
            <p className="text-gray-400 font-bold mb-1">当前钱包</p>
            <p className="font-mono text-lg bg-white/10 px-4 py-2 rounded-xl border border-white/10 break-all">{address}</p>
          </div>
          <div className="flex gap-12">
            <div className="text-center">
              <p className="text-gray-400 text-sm font-bold uppercase mb-2">我创建的拍卖</p>
              <p className="text-5xl font-black">{stats?.created_count || 0}</p>
            </div>
            <div className="text-center">
              <p className="text-gray-400 text-sm font-bold uppercase mb-2">我参与的竞拍</p>
              <p className="text-5xl font-black text-blue-400">{stats?.participated_count || 0}</p>
            </div>
          </div>
        </div>

        {/* 一级 Tabs */}
        <div className="flex space-x-6 border-b-2 border-gray-200 mb-8 pb-4">
          <button 
            onClick={() => { setActiveTab('CREATED'); setSubTab('ACTIVE'); }}
            className={`text-xl font-black transition-colors ${activeTab === 'CREATED' ? 'text-blue-600' : 'text-gray-400 hover:text-gray-600'}`}
          >
            我发布的
          </button>
          <button 
            onClick={() => setActiveTab('PARTICIPATED')}
            className={`text-xl font-black transition-colors ${activeTab === 'PARTICIPATED' ? 'text-blue-600' : 'text-gray-400 hover:text-gray-600'}`}
          >
            我参与的
          </button>
        </div>

        {/* 二级 Tabs (仅在我发布的下显示) */}
        {activeTab === 'CREATED' && (
          <div className="flex gap-2 mb-8 bg-gray-100 w-fit p-1 rounded-xl">
            <button 
              onClick={() => setSubTab('ACTIVE')}
              className={`px-6 py-2 rounded-lg font-bold text-sm transition-all ${subTab === 'ACTIVE' ? 'bg-white shadow-sm text-blue-600' : 'text-gray-500 hover:text-gray-700'}`}
            >🔥 进行中</button>
            <button 
              onClick={() => setSubTab('ENDED')}
              className={`px-6 py-2 rounded-lg font-bold text-sm transition-all ${subTab === 'ENDED' ? 'bg-white shadow-sm text-gray-800' : 'text-gray-500 hover:text-gray-700'}`}
            >🏁 已结束 / 已取消</button>
          </div>
        )}

        {/* 卡片列表展示 */}
        {isLoading ? (
          <div className="text-center py-20 text-gray-400 font-bold animate-pulse">📡 数据同步中...</div>
        ) : auctions.length === 0 ? (
          <div className="text-center py-32 bg-white rounded-3xl border border-dashed border-gray-200 shadow-sm">
            <span className="text-4xl mb-4 block">🍃</span>
            <p className="text-gray-500 font-bold">这里空空如也，什么也没有</p>
          </div>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
            {auctions.map((item: any) => (
              <AuctionCard 
                key={item.auction_id} 
                item={item} 
                currentUserAddress={address} 
                onRefresh={loadAuctions}
                onOpenDetails={handleOpenDetails}
              />
            ))}
          </div>
        )}
      </div>

      {/* 详情与流水记录弹窗复用 */}
      <AuctionDetailsModal 
        isOpen={showDetailsModal} onClose={() => setShowDetailsModal(false)}
        selectedItem={selectedItem} isLoadingBids={isLoadingBids}
        bidHistory={bidHistory} openBidModal={openBidModal}
      />

      {/* 精简版竞拍弹窗 */}
      {showBidModal && selectedItem && (
        <div className="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-3xl p-8 max-w-md w-full shadow-2xl">
            <div className="flex justify-between items-center mb-6">
              <h2 className="text-2xl font-black">参与竞拍</h2>
              <button onClick={() => setShowBidModal(false)} className="text-gray-400 hover:text-red-500 font-bold">✕</button>
            </div>
            <div className="space-y-4">
              <label className="block text-xs font-bold text-gray-500 mb-2">你的出价 (ETH)</label>
              <div className="relative">
                <input 
                  type="number" step="0.0001" value={bidAmountInput} 
                  onChange={e => setBidAmountInput(e.target.value)} 
                  className="w-full bg-gray-50 border border-gray-200 rounded-xl pl-4 pr-16 py-4 text-xl font-black text-blue-600 focus:outline-blue-500" 
                />
                <span className="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 font-bold">ETH</span>
              </div>
              <button onClick={submitBid} disabled={isBidding} className="w-full bg-slate-900 text-white py-4 rounded-xl font-bold hover:bg-blue-600 transition disabled:opacity-50 mt-4">
                {isBidding ? '⏳ 链上打包中...' : '确认支付'}
              </button>
            </div>
          </div>
        </div>
      )}
    </main>
  )
}