package models

import (
	"gorm.io/gorm"
)

// AuctionRecord 对应 MySQL 中的 auctions 表
type AuctionRecord struct {
	gorm.Model
	AuctionID     string `gorm:"type:varchar(255);uniqueIndex;not null" json:"auction_id"`
	Seller        string `gorm:"type:varchar(42);index;not null" json:"seller"`
	NFTContract   string `gorm:"type:varchar(42);index;not null" json:"nft_contract"`
	TokenID       string `gorm:"type:varchar(255);not null" json:"token_id"`
	StartPrice    string `gorm:"type:varchar(255);not null" json:"start_price"`
	StartTime     uint64 `gorm:"not null" json:"start_time"`
	EndTime       uint64 `gorm:"not null" json:"end_time"`
	HighestBid    string `gorm:"type:varchar(255);default:'0'" json:"highest_bid"`
	HighestBidder string `gorm:"type:varchar(42)" json:"highest_bidder"` // 新增：当前最高出价人
	Winner        string `gorm:"type:varchar(42)" json:"winner"`
	Status        string `gorm:"type:varchar(20);default:'Active'" json:"status"`
}

// BidRecord 对应 MySQL 中的 bid_records 表，用于记录出价历史
type BidRecord struct {
	gorm.Model
	AuctionID string `gorm:"type:varchar(255);index;not null" json:"auction_id"`
	Bidder    string `gorm:"type:varchar(42);not null" json:"bidder"`
	Amount    string `gorm:"type:varchar(255);not null" json:"amount"`
	Timestamp uint64 `gorm:"not null" json:"timestamp"`
}
