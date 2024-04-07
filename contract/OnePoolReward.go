// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// OnePoolRewardMetaData contains all meta data concerning the OnePoolReward contract.
var OnePoolRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"book_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lifetimeMonthes\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pricingIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributeReward\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accumulatedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeDonation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"book\",\"outputs\":[{\"internalType\":\"contractAddressBook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricingIndex\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"claimMineReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beforeLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardSectors\",\"type\":\"uint256\"}],\"name\":\"fillReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstValidChunk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastValidChunk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lifetimeInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextChunkDonation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refresh\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutHead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timeoutRecords\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"numPriceChunks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timeoutTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"donation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604081815260049182361015610031575b505050361561002057600080fd5b61002c346007546104e7565b600755005b600092833560e01c91826305a8da7214610457575081630dbbe0a21461040157816314bcec9f146103e257816318ca1b2b146103c357816322ee4cb8146103a457816328fedefe14610385578163514ccef01461036657816359e96700146101c85781636baff3d3146101a957816377bd602b1461018a57816380fad3251461016c578163b7375ddb14610131578163ccec92c4146100fa575063f8ac93e8146100db5780610012565b346100f657816003193601126100f6576100f36105db565b51f35b5080fd5b90503461012d578160031936011261012d576024356001600160a01b0381168103610129576100f39135610829565b8380fd5b8280fd5b5050346100f657816003193601126100f657602090517f00000000000000000000000000000000000000000000000000000000000000008152f35b90503461012d578260031936011261012d5760209250549051908152f35b5050346100f657816003193601126100f6576020906008549051908152f35b5050346100f657816003193601126100f6576020906007549051908152f35b90508160031936011261012d5781516380f5560560e01b81528135906001600160a01b0360208285817f000000000000000000000000000000000000000000000000000000000000000085165afa801561035c57869061031a575b61023092501633146107dd565b6102386105db565b610244602435826104e7565b9060191c9060191c9181831161025b575b50505051f35b6102719167ffffffffffffffff9283918561050a565b169161029d7f0000000000000000000000000000000000000000000000000000000000000000426104e7565b16600754908551936102ae85610517565b8452602084015284830152845490680100000000000000008210156103075750906102e18260016102e79401875561049a565b9061056b565b6003556102f86007546008546104e7565b60085581600755388080610255565b634e487b7160e01b865260419052602485fd5b50906020813d8211610354575b8161033460209383610549565b8101031261035057519080821682036103505761023091610223565b8580fd5b3d9150610327565b85513d88823e3d90fd5b5050346100f657816003193601126100f6576020906005549051908152f35b5050346100f657816003193601126100f6576020906001549051908152f35b5050346100f657816003193601126100f6576020906002549051908152f35b5050346100f657816003193601126100f6576020906003549051908152f35b5050346100f657816003193601126100f6576020906006549051908152f35b90503461012d57602036600319011261012d5735918054831015610454575061042b60609261049a565b5090600182549201549080519267ffffffffffffffff908181168552821c166020840152820152f35b80fd5b8490346100f657816003193601126100f6577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b9060009182548110156104d35782805260011b7f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019190565b634e487b7160e01b83526032600452602483fd5b919082018092116104f457565b634e487b7160e01b600052601160045260246000fd5b919082039182116104f457565b6060810190811067ffffffffffffffff82111761053357604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff82111761053357604052565b91906105c5578051825460208301516fffffffffffffffffffffffffffffffff1990911667ffffffffffffffff9290921691909117604091821b6fffffffffffffffff000000000000000016178355015160019190910155565b634e487b7160e01b600052600060045260246000fd5b600080549081156106c8575b600180546105f48161049a565b505467ffffffffffffffff916040914290831c8416116106bd578261062c610684946106226102e19461049a565b5054851c166106ea565b8454906106388261049a565b50541661064860029182546104e7565b9055846106548261049a565b500154610664600891825461050a565b90558583519361067385610517565b81855281602086015284015261049a565b8054908082018092116106a9579080849255036105e75750505b6106a7426106ea565b565b634e487b7160e01b83526011600452602483fd5b50505050505061069e565b5050426006556106a7426106ea565b80600019048211811515166104f4570290565b600654808211156107d957610717610702828461050a565b6107116003546002549061050a565b906106d7565b600019919080151581840464020000000011166104f45760211b91828104600a11831515166104f457600a83028015159104670de0b6b3a764000011166104f45761076990610711600854918561050a565b7f00000000000000000000000000000000000000000000000000000000000000009182156107c3576301ea6e00678ac7230489e800006107bb946107b394049202601e1c046104e7565b6004546104e7565b600455600655565b634e487b7160e01b600052601260045260246000fd5b5050565b156107e457565b60405162461bcd60e51b815260206004820152601f60248201527f53656e64657220646f6573206e6f742068617665207065726d697373696f6e006044820152606490fd5b6040516399f4b25160e01b81526001600160a01b03906020816004817f000000000000000000000000000000000000000000000000000000000000000086165afa908115610912578290600092610935575b50610888911633146107dd565b6002548210610930576108996105db565b6108a86004546005549061050a565b90478211610928575b81159384156108c2575b5050505050565b1692818460009261091e575b600092839283928392f1156109125760207f83617a1b0f847971f005bd162dde513cfe93df96e6293c3bbb5fe9c40629dd4c91604051908152a338808080806108bb565b6040513d6000823e3d90fd5b6108fc92506108ce565b4791506108b1565b505050565b90506020813d821161096c575b8161094f60209383610549565b810103126100f6575190828216820361045457508161088861087b565b3d915061094256fea2646970667358221220debe4007830f804a6e58c5dcbbffe49842ab6146c111d0c6bc9d637ce2f9222f64736f6c63430008100033",
}

// OnePoolRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use OnePoolRewardMetaData.ABI instead.
var OnePoolRewardABI = OnePoolRewardMetaData.ABI

// OnePoolRewardBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OnePoolRewardMetaData.Bin instead.
var OnePoolRewardBin = OnePoolRewardMetaData.Bin

// DeployOnePoolReward deploys a new Ethereum contract, binding an instance of OnePoolReward to it.
func DeployOnePoolReward(auth *bind.TransactOpts, backend bind.ContractBackend, book_ common.Address, lifetimeMonthes *big.Int) (common.Address, *types.Transaction, *OnePoolReward, error) {
	parsed, err := OnePoolRewardMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnePoolRewardBin), backend, book_, lifetimeMonthes)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnePoolReward{OnePoolRewardCaller: OnePoolRewardCaller{contract: contract}, OnePoolRewardTransactor: OnePoolRewardTransactor{contract: contract}, OnePoolRewardFilterer: OnePoolRewardFilterer{contract: contract}}, nil
}

// OnePoolReward is an auto generated Go binding around an Ethereum contract.
type OnePoolReward struct {
	OnePoolRewardCaller     // Read-only binding to the contract
	OnePoolRewardTransactor // Write-only binding to the contract
	OnePoolRewardFilterer   // Log filterer for contract events
}

// OnePoolRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type OnePoolRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnePoolRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OnePoolRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnePoolRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OnePoolRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnePoolRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OnePoolRewardSession struct {
	Contract     *OnePoolReward    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OnePoolRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OnePoolRewardCallerSession struct {
	Contract *OnePoolRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OnePoolRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OnePoolRewardTransactorSession struct {
	Contract     *OnePoolRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OnePoolRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type OnePoolRewardRaw struct {
	Contract *OnePoolReward // Generic contract binding to access the raw methods on
}

// OnePoolRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OnePoolRewardCallerRaw struct {
	Contract *OnePoolRewardCaller // Generic read-only contract binding to access the raw methods on
}

// OnePoolRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OnePoolRewardTransactorRaw struct {
	Contract *OnePoolRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOnePoolReward creates a new instance of OnePoolReward, bound to a specific deployed contract.
func NewOnePoolReward(address common.Address, backend bind.ContractBackend) (*OnePoolReward, error) {
	contract, err := bindOnePoolReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnePoolReward{OnePoolRewardCaller: OnePoolRewardCaller{contract: contract}, OnePoolRewardTransactor: OnePoolRewardTransactor{contract: contract}, OnePoolRewardFilterer: OnePoolRewardFilterer{contract: contract}}, nil
}

// NewOnePoolRewardCaller creates a new read-only instance of OnePoolReward, bound to a specific deployed contract.
func NewOnePoolRewardCaller(address common.Address, caller bind.ContractCaller) (*OnePoolRewardCaller, error) {
	contract, err := bindOnePoolReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardCaller{contract: contract}, nil
}

// NewOnePoolRewardTransactor creates a new write-only instance of OnePoolReward, bound to a specific deployed contract.
func NewOnePoolRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*OnePoolRewardTransactor, error) {
	contract, err := bindOnePoolReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardTransactor{contract: contract}, nil
}

// NewOnePoolRewardFilterer creates a new log filterer instance of OnePoolReward, bound to a specific deployed contract.
func NewOnePoolRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*OnePoolRewardFilterer, error) {
	contract, err := bindOnePoolReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardFilterer{contract: contract}, nil
}

// bindOnePoolReward binds a generic wrapper to an already deployed contract.
func bindOnePoolReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OnePoolRewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnePoolReward *OnePoolRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnePoolReward.Contract.OnePoolRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnePoolReward *OnePoolRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnePoolReward.Contract.OnePoolRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnePoolReward *OnePoolRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnePoolReward.Contract.OnePoolRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnePoolReward *OnePoolRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnePoolReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnePoolReward *OnePoolRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnePoolReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnePoolReward *OnePoolRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnePoolReward.Contract.contract.Transact(opts, method, params...)
}

// AccumulatedReward is a free data retrieval call binding the contract method 0x80fad325.
//
// Solidity: function accumulatedReward() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) AccumulatedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "accumulatedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedReward is a free data retrieval call binding the contract method 0x80fad325.
//
// Solidity: function accumulatedReward() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) AccumulatedReward() (*big.Int, error) {
	return _OnePoolReward.Contract.AccumulatedReward(&_OnePoolReward.CallOpts)
}

// AccumulatedReward is a free data retrieval call binding the contract method 0x80fad325.
//
// Solidity: function accumulatedReward() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) AccumulatedReward() (*big.Int, error) {
	return _OnePoolReward.Contract.AccumulatedReward(&_OnePoolReward.CallOpts)
}

// ActiveDonation is a free data retrieval call binding the contract method 0x77bd602b.
//
// Solidity: function activeDonation() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) ActiveDonation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "activeDonation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveDonation is a free data retrieval call binding the contract method 0x77bd602b.
//
// Solidity: function activeDonation() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) ActiveDonation() (*big.Int, error) {
	return _OnePoolReward.Contract.ActiveDonation(&_OnePoolReward.CallOpts)
}

// ActiveDonation is a free data retrieval call binding the contract method 0x77bd602b.
//
// Solidity: function activeDonation() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) ActiveDonation() (*big.Int, error) {
	return _OnePoolReward.Contract.ActiveDonation(&_OnePoolReward.CallOpts)
}

// Book is a free data retrieval call binding the contract method 0x05a8da72.
//
// Solidity: function book() view returns(address)
func (_OnePoolReward *OnePoolRewardCaller) Book(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "book")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Book is a free data retrieval call binding the contract method 0x05a8da72.
//
// Solidity: function book() view returns(address)
func (_OnePoolReward *OnePoolRewardSession) Book() (common.Address, error) {
	return _OnePoolReward.Contract.Book(&_OnePoolReward.CallOpts)
}

// Book is a free data retrieval call binding the contract method 0x05a8da72.
//
// Solidity: function book() view returns(address)
func (_OnePoolReward *OnePoolRewardCallerSession) Book() (common.Address, error) {
	return _OnePoolReward.Contract.Book(&_OnePoolReward.CallOpts)
}

// ClaimedReward is a free data retrieval call binding the contract method 0x514ccef0.
//
// Solidity: function claimedReward() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) ClaimedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "claimedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedReward is a free data retrieval call binding the contract method 0x514ccef0.
//
// Solidity: function claimedReward() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) ClaimedReward() (*big.Int, error) {
	return _OnePoolReward.Contract.ClaimedReward(&_OnePoolReward.CallOpts)
}

// ClaimedReward is a free data retrieval call binding the contract method 0x514ccef0.
//
// Solidity: function claimedReward() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) ClaimedReward() (*big.Int, error) {
	return _OnePoolReward.Contract.ClaimedReward(&_OnePoolReward.CallOpts)
}

// FirstValidChunk is a free data retrieval call binding the contract method 0x22ee4cb8.
//
// Solidity: function firstValidChunk() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) FirstValidChunk(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "firstValidChunk")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstValidChunk is a free data retrieval call binding the contract method 0x22ee4cb8.
//
// Solidity: function firstValidChunk() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) FirstValidChunk() (*big.Int, error) {
	return _OnePoolReward.Contract.FirstValidChunk(&_OnePoolReward.CallOpts)
}

// FirstValidChunk is a free data retrieval call binding the contract method 0x22ee4cb8.
//
// Solidity: function firstValidChunk() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) FirstValidChunk() (*big.Int, error) {
	return _OnePoolReward.Contract.FirstValidChunk(&_OnePoolReward.CallOpts)
}

// LastUpdateTimestamp is a free data retrieval call binding the contract method 0x14bcec9f.
//
// Solidity: function lastUpdateTimestamp() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) LastUpdateTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "lastUpdateTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTimestamp is a free data retrieval call binding the contract method 0x14bcec9f.
//
// Solidity: function lastUpdateTimestamp() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) LastUpdateTimestamp() (*big.Int, error) {
	return _OnePoolReward.Contract.LastUpdateTimestamp(&_OnePoolReward.CallOpts)
}

// LastUpdateTimestamp is a free data retrieval call binding the contract method 0x14bcec9f.
//
// Solidity: function lastUpdateTimestamp() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) LastUpdateTimestamp() (*big.Int, error) {
	return _OnePoolReward.Contract.LastUpdateTimestamp(&_OnePoolReward.CallOpts)
}

// LastValidChunk is a free data retrieval call binding the contract method 0x18ca1b2b.
//
// Solidity: function lastValidChunk() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) LastValidChunk(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "lastValidChunk")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastValidChunk is a free data retrieval call binding the contract method 0x18ca1b2b.
//
// Solidity: function lastValidChunk() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) LastValidChunk() (*big.Int, error) {
	return _OnePoolReward.Contract.LastValidChunk(&_OnePoolReward.CallOpts)
}

// LastValidChunk is a free data retrieval call binding the contract method 0x18ca1b2b.
//
// Solidity: function lastValidChunk() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) LastValidChunk() (*big.Int, error) {
	return _OnePoolReward.Contract.LastValidChunk(&_OnePoolReward.CallOpts)
}

// LifetimeInSeconds is a free data retrieval call binding the contract method 0xb7375ddb.
//
// Solidity: function lifetimeInSeconds() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) LifetimeInSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "lifetimeInSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LifetimeInSeconds is a free data retrieval call binding the contract method 0xb7375ddb.
//
// Solidity: function lifetimeInSeconds() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) LifetimeInSeconds() (*big.Int, error) {
	return _OnePoolReward.Contract.LifetimeInSeconds(&_OnePoolReward.CallOpts)
}

// LifetimeInSeconds is a free data retrieval call binding the contract method 0xb7375ddb.
//
// Solidity: function lifetimeInSeconds() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) LifetimeInSeconds() (*big.Int, error) {
	return _OnePoolReward.Contract.LifetimeInSeconds(&_OnePoolReward.CallOpts)
}

// NextChunkDonation is a free data retrieval call binding the contract method 0x6baff3d3.
//
// Solidity: function nextChunkDonation() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) NextChunkDonation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "nextChunkDonation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextChunkDonation is a free data retrieval call binding the contract method 0x6baff3d3.
//
// Solidity: function nextChunkDonation() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) NextChunkDonation() (*big.Int, error) {
	return _OnePoolReward.Contract.NextChunkDonation(&_OnePoolReward.CallOpts)
}

// NextChunkDonation is a free data retrieval call binding the contract method 0x6baff3d3.
//
// Solidity: function nextChunkDonation() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) NextChunkDonation() (*big.Int, error) {
	return _OnePoolReward.Contract.NextChunkDonation(&_OnePoolReward.CallOpts)
}

// TimeoutHead is a free data retrieval call binding the contract method 0x28fedefe.
//
// Solidity: function timeoutHead() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCaller) TimeoutHead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "timeoutHead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeoutHead is a free data retrieval call binding the contract method 0x28fedefe.
//
// Solidity: function timeoutHead() view returns(uint256)
func (_OnePoolReward *OnePoolRewardSession) TimeoutHead() (*big.Int, error) {
	return _OnePoolReward.Contract.TimeoutHead(&_OnePoolReward.CallOpts)
}

// TimeoutHead is a free data retrieval call binding the contract method 0x28fedefe.
//
// Solidity: function timeoutHead() view returns(uint256)
func (_OnePoolReward *OnePoolRewardCallerSession) TimeoutHead() (*big.Int, error) {
	return _OnePoolReward.Contract.TimeoutHead(&_OnePoolReward.CallOpts)
}

// TimeoutRecords is a free data retrieval call binding the contract method 0x0dbbe0a2.
//
// Solidity: function timeoutRecords(uint256 ) view returns(uint64 numPriceChunks, uint64 timeoutTimestamp, uint256 donation)
func (_OnePoolReward *OnePoolRewardCaller) TimeoutRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NumPriceChunks   uint64
	TimeoutTimestamp uint64
	Donation         *big.Int
}, error) {
	var out []interface{}
	err := _OnePoolReward.contract.Call(opts, &out, "timeoutRecords", arg0)

	outstruct := new(struct {
		NumPriceChunks   uint64
		TimeoutTimestamp uint64
		Donation         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NumPriceChunks = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.TimeoutTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Donation = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TimeoutRecords is a free data retrieval call binding the contract method 0x0dbbe0a2.
//
// Solidity: function timeoutRecords(uint256 ) view returns(uint64 numPriceChunks, uint64 timeoutTimestamp, uint256 donation)
func (_OnePoolReward *OnePoolRewardSession) TimeoutRecords(arg0 *big.Int) (struct {
	NumPriceChunks   uint64
	TimeoutTimestamp uint64
	Donation         *big.Int
}, error) {
	return _OnePoolReward.Contract.TimeoutRecords(&_OnePoolReward.CallOpts, arg0)
}

// TimeoutRecords is a free data retrieval call binding the contract method 0x0dbbe0a2.
//
// Solidity: function timeoutRecords(uint256 ) view returns(uint64 numPriceChunks, uint64 timeoutTimestamp, uint256 donation)
func (_OnePoolReward *OnePoolRewardCallerSession) TimeoutRecords(arg0 *big.Int) (struct {
	NumPriceChunks   uint64
	TimeoutTimestamp uint64
	Donation         *big.Int
}, error) {
	return _OnePoolReward.Contract.TimeoutRecords(&_OnePoolReward.CallOpts, arg0)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xccec92c4.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary) returns()
func (_OnePoolReward *OnePoolRewardTransactor) ClaimMineReward(opts *bind.TransactOpts, pricingIndex *big.Int, beneficiary common.Address) (*types.Transaction, error) {
	return _OnePoolReward.contract.Transact(opts, "claimMineReward", pricingIndex, beneficiary)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xccec92c4.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary) returns()
func (_OnePoolReward *OnePoolRewardSession) ClaimMineReward(pricingIndex *big.Int, beneficiary common.Address) (*types.Transaction, error) {
	return _OnePoolReward.Contract.ClaimMineReward(&_OnePoolReward.TransactOpts, pricingIndex, beneficiary)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xccec92c4.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary) returns()
func (_OnePoolReward *OnePoolRewardTransactorSession) ClaimMineReward(pricingIndex *big.Int, beneficiary common.Address) (*types.Transaction, error) {
	return _OnePoolReward.Contract.ClaimMineReward(&_OnePoolReward.TransactOpts, pricingIndex, beneficiary)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 rewardSectors) payable returns()
func (_OnePoolReward *OnePoolRewardTransactor) FillReward(opts *bind.TransactOpts, beforeLength *big.Int, rewardSectors *big.Int) (*types.Transaction, error) {
	return _OnePoolReward.contract.Transact(opts, "fillReward", beforeLength, rewardSectors)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 rewardSectors) payable returns()
func (_OnePoolReward *OnePoolRewardSession) FillReward(beforeLength *big.Int, rewardSectors *big.Int) (*types.Transaction, error) {
	return _OnePoolReward.Contract.FillReward(&_OnePoolReward.TransactOpts, beforeLength, rewardSectors)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 rewardSectors) payable returns()
func (_OnePoolReward *OnePoolRewardTransactorSession) FillReward(beforeLength *big.Int, rewardSectors *big.Int) (*types.Transaction, error) {
	return _OnePoolReward.Contract.FillReward(&_OnePoolReward.TransactOpts, beforeLength, rewardSectors)
}

// Refresh is a paid mutator transaction binding the contract method 0xf8ac93e8.
//
// Solidity: function refresh() returns()
func (_OnePoolReward *OnePoolRewardTransactor) Refresh(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnePoolReward.contract.Transact(opts, "refresh")
}

// Refresh is a paid mutator transaction binding the contract method 0xf8ac93e8.
//
// Solidity: function refresh() returns()
func (_OnePoolReward *OnePoolRewardSession) Refresh() (*types.Transaction, error) {
	return _OnePoolReward.Contract.Refresh(&_OnePoolReward.TransactOpts)
}

// Refresh is a paid mutator transaction binding the contract method 0xf8ac93e8.
//
// Solidity: function refresh() returns()
func (_OnePoolReward *OnePoolRewardTransactorSession) Refresh() (*types.Transaction, error) {
	return _OnePoolReward.Contract.Refresh(&_OnePoolReward.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnePoolReward *OnePoolRewardTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnePoolReward.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnePoolReward *OnePoolRewardSession) Receive() (*types.Transaction, error) {
	return _OnePoolReward.Contract.Receive(&_OnePoolReward.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnePoolReward *OnePoolRewardTransactorSession) Receive() (*types.Transaction, error) {
	return _OnePoolReward.Contract.Receive(&_OnePoolReward.TransactOpts)
}

// OnePoolRewardDistributeRewardIterator is returned from FilterDistributeReward and is used to iterate over the raw logs and unpacked data for DistributeReward events raised by the OnePoolReward contract.
type OnePoolRewardDistributeRewardIterator struct {
	Event *OnePoolRewardDistributeReward // Event containing the contract specifics and raw log

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
func (it *OnePoolRewardDistributeRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnePoolRewardDistributeReward)
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
		it.Event = new(OnePoolRewardDistributeReward)
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
func (it *OnePoolRewardDistributeRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnePoolRewardDistributeRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnePoolRewardDistributeReward represents a DistributeReward event raised by the OnePoolReward contract.
type OnePoolRewardDistributeReward struct {
	PricingIndex *big.Int
	Beneficiary  common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDistributeReward is a free log retrieval operation binding the contract event 0x83617a1b0f847971f005bd162dde513cfe93df96e6293c3bbb5fe9c40629dd4c.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, uint256 amount)
func (_OnePoolReward *OnePoolRewardFilterer) FilterDistributeReward(opts *bind.FilterOpts, pricingIndex []*big.Int, beneficiary []common.Address) (*OnePoolRewardDistributeRewardIterator, error) {

	var pricingIndexRule []interface{}
	for _, pricingIndexItem := range pricingIndex {
		pricingIndexRule = append(pricingIndexRule, pricingIndexItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _OnePoolReward.contract.FilterLogs(opts, "DistributeReward", pricingIndexRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &OnePoolRewardDistributeRewardIterator{contract: _OnePoolReward.contract, event: "DistributeReward", logs: logs, sub: sub}, nil
}

// WatchDistributeReward is a free log subscription operation binding the contract event 0x83617a1b0f847971f005bd162dde513cfe93df96e6293c3bbb5fe9c40629dd4c.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, uint256 amount)
func (_OnePoolReward *OnePoolRewardFilterer) WatchDistributeReward(opts *bind.WatchOpts, sink chan<- *OnePoolRewardDistributeReward, pricingIndex []*big.Int, beneficiary []common.Address) (event.Subscription, error) {

	var pricingIndexRule []interface{}
	for _, pricingIndexItem := range pricingIndex {
		pricingIndexRule = append(pricingIndexRule, pricingIndexItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _OnePoolReward.contract.WatchLogs(opts, "DistributeReward", pricingIndexRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnePoolRewardDistributeReward)
				if err := _OnePoolReward.contract.UnpackLog(event, "DistributeReward", log); err != nil {
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

// ParseDistributeReward is a log parse operation binding the contract event 0x83617a1b0f847971f005bd162dde513cfe93df96e6293c3bbb5fe9c40629dd4c.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, uint256 amount)
func (_OnePoolReward *OnePoolRewardFilterer) ParseDistributeReward(log types.Log) (*OnePoolRewardDistributeReward, error) {
	event := new(OnePoolRewardDistributeReward)
	if err := _OnePoolReward.contract.UnpackLog(event, "DistributeReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
