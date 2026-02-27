// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auction

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// NFTAuctionMetaData contains all meta data concerning the NFTAuction contract.
var NFTAuctionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctionData\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"highestBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bidAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"cancelAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAuction\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeRecipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextAuctionId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nftToken2AuctionId\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"platformFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeRecipient\",\"inputs\":[{\"name\":\"_feeRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPlatformFee\",\"inputs\":[{\"name\":\"_feeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AuctionCanceled\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionCreated\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionEnded\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"winner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidPlaced\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AuctionAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionAlreadyStarted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionEndederror\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionNotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionNotEnded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionNotStarted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BidTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BidsAlreadyPlaced\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeePercent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotNFTOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSeller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RefundFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// NFTAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTAuctionMetaData.ABI instead.
var NFTAuctionABI = NFTAuctionMetaData.ABI

// NFTAuction is an auto generated Go binding around an Ethereum contract.
type NFTAuction struct {
	NFTAuctionCaller     // Read-only binding to the contract
	NFTAuctionTransactor // Write-only binding to the contract
	NFTAuctionFilterer   // Log filterer for contract events
}

// NFTAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTAuctionSession struct {
	Contract     *NFTAuction       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTAuctionCallerSession struct {
	Contract *NFTAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NFTAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTAuctionTransactorSession struct {
	Contract     *NFTAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NFTAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTAuctionRaw struct {
	Contract *NFTAuction // Generic contract binding to access the raw methods on
}

// NFTAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTAuctionCallerRaw struct {
	Contract *NFTAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// NFTAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTAuctionTransactorRaw struct {
	Contract *NFTAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTAuction creates a new instance of NFTAuction, bound to a specific deployed contract.
func NewNFTAuction(address common.Address, backend bind.ContractBackend) (*NFTAuction, error) {
	contract, err := bindNFTAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTAuction{NFTAuctionCaller: NFTAuctionCaller{contract: contract}, NFTAuctionTransactor: NFTAuctionTransactor{contract: contract}, NFTAuctionFilterer: NFTAuctionFilterer{contract: contract}}, nil
}

// NewNFTAuctionCaller creates a new read-only instance of NFTAuction, bound to a specific deployed contract.
func NewNFTAuctionCaller(address common.Address, caller bind.ContractCaller) (*NFTAuctionCaller, error) {
	contract, err := bindNFTAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionCaller{contract: contract}, nil
}

// NewNFTAuctionTransactor creates a new write-only instance of NFTAuction, bound to a specific deployed contract.
func NewNFTAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTAuctionTransactor, error) {
	contract, err := bindNFTAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionTransactor{contract: contract}, nil
}

// NewNFTAuctionFilterer creates a new log filterer instance of NFTAuction, bound to a specific deployed contract.
func NewNFTAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTAuctionFilterer, error) {
	contract, err := bindNFTAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionFilterer{contract: contract}, nil
}

// bindNFTAuction binds a generic wrapper to an already deployed contract.
func bindNFTAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NFTAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTAuction *NFTAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTAuction.Contract.NFTAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTAuction *NFTAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuction.Contract.NFTAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTAuction *NFTAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTAuction.Contract.NFTAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTAuction *NFTAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTAuction *NFTAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTAuction *NFTAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTAuction.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NFTAuction *NFTAuctionCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NFTAuction *NFTAuctionSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NFTAuction.Contract.UPGRADEINTERFACEVERSION(&_NFTAuction.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NFTAuction *NFTAuctionCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NFTAuction.Contract.UPGRADEINTERFACEVERSION(&_NFTAuction.CallOpts)
}

// AuctionData is a free data retrieval call binding the contract method 0x55fc62d2.
//
// Solidity: function auctionData(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 duration, uint256 endTime, address highestBidder, uint256 highestBid, bool active)
func (_NFTAuction *NFTAuctionCaller) AuctionData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartPrice    *big.Int
	StartTime     *big.Int
	Duration      *big.Int
	EndTime       *big.Int
	HighestBidder common.Address
	HighestBid    *big.Int
	Active        bool
}, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "auctionData", arg0)

	outstruct := new(struct {
		Seller        common.Address
		NftContract   common.Address
		TokenId       *big.Int
		StartPrice    *big.Int
		StartTime     *big.Int
		Duration      *big.Int
		EndTime       *big.Int
		HighestBidder common.Address
		HighestBid    *big.Int
		Active        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.HighestBidder = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.HighestBid = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// AuctionData is a free data retrieval call binding the contract method 0x55fc62d2.
//
// Solidity: function auctionData(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 duration, uint256 endTime, address highestBidder, uint256 highestBid, bool active)
func (_NFTAuction *NFTAuctionSession) AuctionData(arg0 *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartPrice    *big.Int
	StartTime     *big.Int
	Duration      *big.Int
	EndTime       *big.Int
	HighestBidder common.Address
	HighestBid    *big.Int
	Active        bool
}, error) {
	return _NFTAuction.Contract.AuctionData(&_NFTAuction.CallOpts, arg0)
}

// AuctionData is a free data retrieval call binding the contract method 0x55fc62d2.
//
// Solidity: function auctionData(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 duration, uint256 endTime, address highestBidder, uint256 highestBid, bool active)
func (_NFTAuction *NFTAuctionCallerSession) AuctionData(arg0 *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartPrice    *big.Int
	StartTime     *big.Int
	Duration      *big.Int
	EndTime       *big.Int
	HighestBidder common.Address
	HighestBid    *big.Int
	Active        bool
}, error) {
	return _NFTAuction.Contract.AuctionData(&_NFTAuction.CallOpts, arg0)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_NFTAuction *NFTAuctionCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_NFTAuction *NFTAuctionSession) FeeRecipient() (common.Address, error) {
	return _NFTAuction.Contract.FeeRecipient(&_NFTAuction.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_NFTAuction *NFTAuctionCallerSession) FeeRecipient() (common.Address, error) {
	return _NFTAuction.Contract.FeeRecipient(&_NFTAuction.CallOpts)
}

// NextAuctionId is a free data retrieval call binding the contract method 0xfc528482.
//
// Solidity: function nextAuctionId() view returns(uint256)
func (_NFTAuction *NFTAuctionCaller) NextAuctionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "nextAuctionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextAuctionId is a free data retrieval call binding the contract method 0xfc528482.
//
// Solidity: function nextAuctionId() view returns(uint256)
func (_NFTAuction *NFTAuctionSession) NextAuctionId() (*big.Int, error) {
	return _NFTAuction.Contract.NextAuctionId(&_NFTAuction.CallOpts)
}

// NextAuctionId is a free data retrieval call binding the contract method 0xfc528482.
//
// Solidity: function nextAuctionId() view returns(uint256)
func (_NFTAuction *NFTAuctionCallerSession) NextAuctionId() (*big.Int, error) {
	return _NFTAuction.Contract.NextAuctionId(&_NFTAuction.CallOpts)
}

// NftToken2AuctionId is a free data retrieval call binding the contract method 0x980fda47.
//
// Solidity: function nftToken2AuctionId(address , uint256 ) view returns(uint256)
func (_NFTAuction *NFTAuctionCaller) NftToken2AuctionId(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "nftToken2AuctionId", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftToken2AuctionId is a free data retrieval call binding the contract method 0x980fda47.
//
// Solidity: function nftToken2AuctionId(address , uint256 ) view returns(uint256)
func (_NFTAuction *NFTAuctionSession) NftToken2AuctionId(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NFTAuction.Contract.NftToken2AuctionId(&_NFTAuction.CallOpts, arg0, arg1)
}

// NftToken2AuctionId is a free data retrieval call binding the contract method 0x980fda47.
//
// Solidity: function nftToken2AuctionId(address , uint256 ) view returns(uint256)
func (_NFTAuction *NFTAuctionCallerSession) NftToken2AuctionId(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NFTAuction.Contract.NftToken2AuctionId(&_NFTAuction.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTAuction *NFTAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTAuction *NFTAuctionSession) Owner() (common.Address, error) {
	return _NFTAuction.Contract.Owner(&_NFTAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTAuction *NFTAuctionCallerSession) Owner() (common.Address, error) {
	return _NFTAuction.Contract.Owner(&_NFTAuction.CallOpts)
}

// PlatformFeeBasisPoints is a free data retrieval call binding the contract method 0xc58bfb66.
//
// Solidity: function platformFeeBasisPoints() view returns(uint256)
func (_NFTAuction *NFTAuctionCaller) PlatformFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "platformFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFeeBasisPoints is a free data retrieval call binding the contract method 0xc58bfb66.
//
// Solidity: function platformFeeBasisPoints() view returns(uint256)
func (_NFTAuction *NFTAuctionSession) PlatformFeeBasisPoints() (*big.Int, error) {
	return _NFTAuction.Contract.PlatformFeeBasisPoints(&_NFTAuction.CallOpts)
}

// PlatformFeeBasisPoints is a free data retrieval call binding the contract method 0xc58bfb66.
//
// Solidity: function platformFeeBasisPoints() view returns(uint256)
func (_NFTAuction *NFTAuctionCallerSession) PlatformFeeBasisPoints() (*big.Int, error) {
	return _NFTAuction.Contract.PlatformFeeBasisPoints(&_NFTAuction.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NFTAuction *NFTAuctionCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFTAuction.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NFTAuction *NFTAuctionSession) ProxiableUUID() ([32]byte, error) {
	return _NFTAuction.Contract.ProxiableUUID(&_NFTAuction.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NFTAuction *NFTAuctionCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NFTAuction.Contract.ProxiableUUID(&_NFTAuction.CallOpts)
}

// BidAuction is a paid mutator transaction binding the contract method 0x64a7d7c7.
//
// Solidity: function bidAuction(uint256 auctionId) payable returns()
func (_NFTAuction *NFTAuctionTransactor) BidAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "bidAuction", auctionId)
}

// BidAuction is a paid mutator transaction binding the contract method 0x64a7d7c7.
//
// Solidity: function bidAuction(uint256 auctionId) payable returns()
func (_NFTAuction *NFTAuctionSession) BidAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuction.Contract.BidAuction(&_NFTAuction.TransactOpts, auctionId)
}

// BidAuction is a paid mutator transaction binding the contract method 0x64a7d7c7.
//
// Solidity: function bidAuction(uint256 auctionId) payable returns()
func (_NFTAuction *NFTAuctionTransactorSession) BidAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuction.Contract.BidAuction(&_NFTAuction.TransactOpts, auctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 auctionId) returns()
func (_NFTAuction *NFTAuctionTransactor) CancelAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "cancelAuction", auctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 auctionId) returns()
func (_NFTAuction *NFTAuctionSession) CancelAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuction.Contract.CancelAuction(&_NFTAuction.TransactOpts, auctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 auctionId) returns()
func (_NFTAuction *NFTAuctionTransactorSession) CancelAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuction.Contract.CancelAuction(&_NFTAuction.TransactOpts, auctionId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x961c9ae4.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 duration) returns()
func (_NFTAuction *NFTAuctionTransactor) CreateAuction(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, startPrice *big.Int, startTime *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "createAuction", nftContract, tokenId, startPrice, startTime, duration)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x961c9ae4.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 duration) returns()
func (_NFTAuction *NFTAuctionSession) CreateAuction(nftContract common.Address, tokenId *big.Int, startPrice *big.Int, startTime *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _NFTAuction.Contract.CreateAuction(&_NFTAuction.TransactOpts, nftContract, tokenId, startPrice, startTime, duration)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x961c9ae4.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 duration) returns()
func (_NFTAuction *NFTAuctionTransactorSession) CreateAuction(nftContract common.Address, tokenId *big.Int, startPrice *big.Int, startTime *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _NFTAuction.Contract.CreateAuction(&_NFTAuction.TransactOpts, nftContract, tokenId, startPrice, startTime, duration)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_NFTAuction *NFTAuctionTransactor) EndAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "endAuction", auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_NFTAuction *NFTAuctionSession) EndAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuction.Contract.EndAuction(&_NFTAuction.TransactOpts, auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_NFTAuction *NFTAuctionTransactorSession) EndAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuction.Contract.EndAuction(&_NFTAuction.TransactOpts, auctionId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_NFTAuction *NFTAuctionTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_NFTAuction *NFTAuctionSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _NFTAuction.Contract.Initialize(&_NFTAuction.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_NFTAuction *NFTAuctionTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _NFTAuction.Contract.Initialize(&_NFTAuction.TransactOpts, initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTAuction *NFTAuctionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTAuction *NFTAuctionSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTAuction.Contract.RenounceOwnership(&_NFTAuction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTAuction *NFTAuctionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTAuction.Contract.RenounceOwnership(&_NFTAuction.TransactOpts)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_NFTAuction *NFTAuctionTransactor) SetFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "setFeeRecipient", _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_NFTAuction *NFTAuctionSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _NFTAuction.Contract.SetFeeRecipient(&_NFTAuction.TransactOpts, _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_NFTAuction *NFTAuctionTransactorSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _NFTAuction.Contract.SetFeeRecipient(&_NFTAuction.TransactOpts, _feeRecipient)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _feeBasisPoints) returns()
func (_NFTAuction *NFTAuctionTransactor) SetPlatformFee(opts *bind.TransactOpts, _feeBasisPoints *big.Int) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "setPlatformFee", _feeBasisPoints)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _feeBasisPoints) returns()
func (_NFTAuction *NFTAuctionSession) SetPlatformFee(_feeBasisPoints *big.Int) (*types.Transaction, error) {
	return _NFTAuction.Contract.SetPlatformFee(&_NFTAuction.TransactOpts, _feeBasisPoints)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _feeBasisPoints) returns()
func (_NFTAuction *NFTAuctionTransactorSession) SetPlatformFee(_feeBasisPoints *big.Int) (*types.Transaction, error) {
	return _NFTAuction.Contract.SetPlatformFee(&_NFTAuction.TransactOpts, _feeBasisPoints)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTAuction *NFTAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTAuction *NFTAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTAuction.Contract.TransferOwnership(&_NFTAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTAuction *NFTAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTAuction.Contract.TransferOwnership(&_NFTAuction.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NFTAuction *NFTAuctionTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NFTAuction.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NFTAuction *NFTAuctionSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NFTAuction.Contract.UpgradeToAndCall(&_NFTAuction.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NFTAuction *NFTAuctionTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NFTAuction.Contract.UpgradeToAndCall(&_NFTAuction.TransactOpts, newImplementation, data)
}

// NFTAuctionAuctionCanceledIterator is returned from FilterAuctionCanceled and is used to iterate over the raw logs and unpacked data for AuctionCanceled events raised by the NFTAuction contract.
type NFTAuctionAuctionCanceledIterator struct {
	Event *NFTAuctionAuctionCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionAuctionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionAuctionCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionAuctionCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionAuctionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionAuctionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionAuctionCanceled represents a AuctionCanceled event raised by the NFTAuction contract.
type NFTAuctionAuctionCanceled struct {
	AuctionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCanceled is a free log retrieval operation binding the contract event 0x28601d865dccc9f113e15a7185c1b38c085d598c71250d3337916a428536d771.
//
// Solidity: event AuctionCanceled(uint256 indexed auctionId)
func (_NFTAuction *NFTAuctionFilterer) FilterAuctionCanceled(opts *bind.FilterOpts, auctionId []*big.Int) (*NFTAuctionAuctionCanceledIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _NFTAuction.contract.FilterLogs(opts, "AuctionCanceled", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionAuctionCanceledIterator{contract: _NFTAuction.contract, event: "AuctionCanceled", logs: logs, sub: sub}, nil
}

// WatchAuctionCanceled is a free log subscription operation binding the contract event 0x28601d865dccc9f113e15a7185c1b38c085d598c71250d3337916a428536d771.
//
// Solidity: event AuctionCanceled(uint256 indexed auctionId)
func (_NFTAuction *NFTAuctionFilterer) WatchAuctionCanceled(opts *bind.WatchOpts, sink chan<- *NFTAuctionAuctionCanceled, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _NFTAuction.contract.WatchLogs(opts, "AuctionCanceled", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionAuctionCanceled)
				if err := _NFTAuction.contract.UnpackLog(event, "AuctionCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionCanceled is a log parse operation binding the contract event 0x28601d865dccc9f113e15a7185c1b38c085d598c71250d3337916a428536d771.
//
// Solidity: event AuctionCanceled(uint256 indexed auctionId)
func (_NFTAuction *NFTAuctionFilterer) ParseAuctionCanceled(log types.Log) (*NFTAuctionAuctionCanceled, error) {
	event := new(NFTAuctionAuctionCanceled)
	if err := _NFTAuction.contract.UnpackLog(event, "AuctionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the NFTAuction contract.
type NFTAuctionAuctionCreatedIterator struct {
	Event *NFTAuctionAuctionCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionAuctionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionAuctionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionAuctionCreated represents a AuctionCreated event raised by the NFTAuction contract.
type NFTAuctionAuctionCreated struct {
	AuctionId   *big.Int
	Seller      common.Address
	NftContract common.Address
	TokenId     *big.Int
	StartPrice  *big.Int
	StartTime   *big.Int
	EndTime     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xcaf0ae751fb2b122e8718bf7c0d4b7584d1418a853a4d0cdaba45418d3da138b.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime)
func (_NFTAuction *NFTAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts, auctionId []*big.Int, seller []common.Address, nftContract []common.Address) (*NFTAuctionAuctionCreatedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _NFTAuction.contract.FilterLogs(opts, "AuctionCreated", auctionIdRule, sellerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionAuctionCreatedIterator{contract: _NFTAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xcaf0ae751fb2b122e8718bf7c0d4b7584d1418a853a4d0cdaba45418d3da138b.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime)
func (_NFTAuction *NFTAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *NFTAuctionAuctionCreated, auctionId []*big.Int, seller []common.Address, nftContract []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _NFTAuction.contract.WatchLogs(opts, "AuctionCreated", auctionIdRule, sellerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionAuctionCreated)
				if err := _NFTAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionCreated is a log parse operation binding the contract event 0xcaf0ae751fb2b122e8718bf7c0d4b7584d1418a853a4d0cdaba45418d3da138b.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime)
func (_NFTAuction *NFTAuctionFilterer) ParseAuctionCreated(log types.Log) (*NFTAuctionAuctionCreated, error) {
	event := new(NFTAuctionAuctionCreated)
	if err := _NFTAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the NFTAuction contract.
type NFTAuctionAuctionEndedIterator struct {
	Event *NFTAuctionAuctionEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionAuctionEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionAuctionEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionAuctionEnded represents a AuctionEnded event raised by the NFTAuction contract.
type NFTAuctionAuctionEnded struct {
	AuctionId *big.Int
	Winner    common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xd2aa34a4fdbbc6dff6a3e56f46e0f3ae2a31d7785ff3487aa5c95c642acea501.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address indexed winner, uint256 amount)
func (_NFTAuction *NFTAuctionFilterer) FilterAuctionEnded(opts *bind.FilterOpts, auctionId []*big.Int, winner []common.Address) (*NFTAuctionAuctionEndedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _NFTAuction.contract.FilterLogs(opts, "AuctionEnded", auctionIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionAuctionEndedIterator{contract: _NFTAuction.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xd2aa34a4fdbbc6dff6a3e56f46e0f3ae2a31d7785ff3487aa5c95c642acea501.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address indexed winner, uint256 amount)
func (_NFTAuction *NFTAuctionFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *NFTAuctionAuctionEnded, auctionId []*big.Int, winner []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _NFTAuction.contract.WatchLogs(opts, "AuctionEnded", auctionIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionAuctionEnded)
				if err := _NFTAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionEnded is a log parse operation binding the contract event 0xd2aa34a4fdbbc6dff6a3e56f46e0f3ae2a31d7785ff3487aa5c95c642acea501.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address indexed winner, uint256 amount)
func (_NFTAuction *NFTAuctionFilterer) ParseAuctionEnded(log types.Log) (*NFTAuctionAuctionEnded, error) {
	event := new(NFTAuctionAuctionEnded)
	if err := _NFTAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the NFTAuction contract.
type NFTAuctionBidPlacedIterator struct {
	Event *NFTAuctionBidPlaced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionBidPlaced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionBidPlaced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionBidPlaced represents a BidPlaced event raised by the NFTAuction contract.
type NFTAuctionBidPlaced struct {
	AuctionId *big.Int
	Bidder    common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x51db8e23b3f4479b162fd48823b8402895442b8f6cfd94f66239391881ec7b6f.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount, uint256 timestamp)
func (_NFTAuction *NFTAuctionFilterer) FilterBidPlaced(opts *bind.FilterOpts, auctionId []*big.Int, bidder []common.Address) (*NFTAuctionBidPlacedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _NFTAuction.contract.FilterLogs(opts, "BidPlaced", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionBidPlacedIterator{contract: _NFTAuction.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x51db8e23b3f4479b162fd48823b8402895442b8f6cfd94f66239391881ec7b6f.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount, uint256 timestamp)
func (_NFTAuction *NFTAuctionFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *NFTAuctionBidPlaced, auctionId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _NFTAuction.contract.WatchLogs(opts, "BidPlaced", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionBidPlaced)
				if err := _NFTAuction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBidPlaced is a log parse operation binding the contract event 0x51db8e23b3f4479b162fd48823b8402895442b8f6cfd94f66239391881ec7b6f.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount, uint256 timestamp)
func (_NFTAuction *NFTAuctionFilterer) ParseBidPlaced(log types.Log) (*NFTAuctionBidPlaced, error) {
	event := new(NFTAuctionBidPlaced)
	if err := _NFTAuction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NFTAuction contract.
type NFTAuctionInitializedIterator struct {
	Event *NFTAuctionInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionInitialized represents a Initialized event raised by the NFTAuction contract.
type NFTAuctionInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NFTAuction *NFTAuctionFilterer) FilterInitialized(opts *bind.FilterOpts) (*NFTAuctionInitializedIterator, error) {

	logs, sub, err := _NFTAuction.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NFTAuctionInitializedIterator{contract: _NFTAuction.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NFTAuction *NFTAuctionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NFTAuctionInitialized) (event.Subscription, error) {

	logs, sub, err := _NFTAuction.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionInitialized)
				if err := _NFTAuction.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NFTAuction *NFTAuctionFilterer) ParseInitialized(log types.Log) (*NFTAuctionInitialized, error) {
	event := new(NFTAuctionInitialized)
	if err := _NFTAuction.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NFTAuction contract.
type NFTAuctionOwnershipTransferredIterator struct {
	Event *NFTAuctionOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionOwnershipTransferred represents a OwnershipTransferred event raised by the NFTAuction contract.
type NFTAuctionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTAuction *NFTAuctionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NFTAuctionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTAuction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionOwnershipTransferredIterator{contract: _NFTAuction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTAuction *NFTAuctionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NFTAuctionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTAuction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionOwnershipTransferred)
				if err := _NFTAuction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTAuction *NFTAuctionFilterer) ParseOwnershipTransferred(log types.Log) (*NFTAuctionOwnershipTransferred, error) {
	event := new(NFTAuctionOwnershipTransferred)
	if err := _NFTAuction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NFTAuction contract.
type NFTAuctionUpgradedIterator struct {
	Event *NFTAuctionUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTAuctionUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTAuctionUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTAuctionUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionUpgraded represents a Upgraded event raised by the NFTAuction contract.
type NFTAuctionUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NFTAuction *NFTAuctionFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NFTAuctionUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NFTAuction.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionUpgradedIterator{contract: _NFTAuction.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NFTAuction *NFTAuctionFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NFTAuctionUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NFTAuction.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionUpgraded)
				if err := _NFTAuction.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NFTAuction *NFTAuctionFilterer) ParseUpgraded(log types.Log) (*NFTAuctionUpgraded, error) {
	event := new(NFTAuctionUpgraded)
	if err := _NFTAuction.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
