// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// EduManageCourse is an auto generated low-level Go binding around an user-defined struct.
type EduManageCourse struct {
	Id         *big.Int
	Name       string
	Users      *big.Int
	Price      *big.Int
	Reward     *big.Int
	CourseType string
	HashCourse string
	CreateTime *big.Int
}

// EduManageGrade is an auto generated low-level Go binding around an user-defined struct.
type EduManageGrade struct {
	Mark            *big.Int
	HashResultExam  string
	IsSet           bool
	IsReceiveReward bool
	IsMintCert      bool
}

// EduManagePost is an auto generated low-level Go binding around an user-defined struct.
type EduManagePost struct {
	Id          *big.Int
	Owner       common.Address
	Point       *big.Int
	RejectCount *big.Int
	RewardCount *big.Int
}

// EduManageStudent is an auto generated low-level Go binding around an user-defined struct.
type EduManageStudent struct {
	Id        *big.Int
	StuAdd    common.Address
	StuName   string
	CoursesId []*big.Int
}

// EduManageMetaData contains all meta data concerning the EduManage contract.
var EduManageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_NFTAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cerType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"imageUri\",\"type\":\"string\"}],\"name\":\"CertificateMinted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CertNFT\",\"outputs\":[{\"internalType\":\"contractCertificateNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLINKToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"course_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"courseType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hashCourse\",\"type\":\"string\"}],\"name\":\"addCourse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"post_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner_add\",\"type\":\"address\"}],\"name\":\"addPost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stu_add\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stu_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"addStudent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowanceBuyCourse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stu_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"course_id\",\"type\":\"uint256\"}],\"name\":\"buyCourse\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stu_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"course_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image_uri\",\"type\":\"string\"}],\"name\":\"checkAndTransferRewardCourse\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"post_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"point\",\"type\":\"uint256\"}],\"name\":\"checkAndTransferRewardPost\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stu_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"course_id\",\"type\":\"uint256\"}],\"name\":\"checkEnrolledCourse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getCourse\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"users\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"courseType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hashCourse\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"}],\"internalType\":\"structEduManage.Course\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCourseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stu_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_course_id\",\"type\":\"uint256\"}],\"name\":\"getGrades\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"mark\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hashResultExam\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isReceiveReward\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMintCert\",\"type\":\"bool\"}],\"internalType\":\"structEduManage.Grade\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getPosts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"point\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejectCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardCount\",\"type\":\"uint256\"}],\"internalType\":\"structEduManage.Post\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getStudent\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stuAdd\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"stu_name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"courses_id\",\"type\":\"uint256[]\"}],\"internalType\":\"structEduManage.Student\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"grades\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mark\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hashResultExam\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isReceiveReward\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMintCert\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postPointToReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postTokenReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfTokens\",\"type\":\"uint256\"}],\"name\":\"rewardToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stu_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"course_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mark\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hash_result_exam\",\"type\":\"string\"}],\"name\":\"submitGrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EduManageABI is the input ABI used to generate the binding from.
// Deprecated: Use EduManageMetaData.ABI instead.
var EduManageABI = EduManageMetaData.ABI

// EduManage is an auto generated Go binding around an Ethereum contract.
type EduManage struct {
	EduManageCaller     // Read-only binding to the contract
	EduManageTransactor // Write-only binding to the contract
	EduManageFilterer   // Log filterer for contract events
}

// EduManageCaller is an auto generated read-only Go binding around an Ethereum contract.
type EduManageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EduManageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EduManageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EduManageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EduManageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EduManageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EduManageSession struct {
	Contract     *EduManage        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EduManageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EduManageCallerSession struct {
	Contract *EduManageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EduManageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EduManageTransactorSession struct {
	Contract     *EduManageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EduManageRaw is an auto generated low-level Go binding around an Ethereum contract.
type EduManageRaw struct {
	Contract *EduManage // Generic contract binding to access the raw methods on
}

// EduManageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EduManageCallerRaw struct {
	Contract *EduManageCaller // Generic read-only contract binding to access the raw methods on
}

// EduManageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EduManageTransactorRaw struct {
	Contract *EduManageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEduManage creates a new instance of EduManage, bound to a specific deployed contract.
func NewEduManage(address common.Address, backend bind.ContractBackend) (*EduManage, error) {
	contract, err := bindEduManage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EduManage{EduManageCaller: EduManageCaller{contract: contract}, EduManageTransactor: EduManageTransactor{contract: contract}, EduManageFilterer: EduManageFilterer{contract: contract}}, nil
}

// NewEduManageCaller creates a new read-only instance of EduManage, bound to a specific deployed contract.
func NewEduManageCaller(address common.Address, caller bind.ContractCaller) (*EduManageCaller, error) {
	contract, err := bindEduManage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EduManageCaller{contract: contract}, nil
}

// NewEduManageTransactor creates a new write-only instance of EduManage, bound to a specific deployed contract.
func NewEduManageTransactor(address common.Address, transactor bind.ContractTransactor) (*EduManageTransactor, error) {
	contract, err := bindEduManage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EduManageTransactor{contract: contract}, nil
}

// NewEduManageFilterer creates a new log filterer instance of EduManage, bound to a specific deployed contract.
func NewEduManageFilterer(address common.Address, filterer bind.ContractFilterer) (*EduManageFilterer, error) {
	contract, err := bindEduManage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EduManageFilterer{contract: contract}, nil
}

// bindEduManage binds a generic wrapper to an already deployed contract.
func bindEduManage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EduManageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EduManage *EduManageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EduManage.Contract.EduManageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EduManage *EduManageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EduManage.Contract.EduManageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EduManage *EduManageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EduManage.Contract.EduManageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EduManage *EduManageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EduManage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EduManage *EduManageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EduManage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EduManage *EduManageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EduManage.Contract.contract.Transact(opts, method, params...)
}

// CertNFT is a free data retrieval call binding the contract method 0x029a36fc.
//
// Solidity: function CertNFT() view returns(address)
func (_EduManage *EduManageCaller) CertNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "CertNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CertNFT is a free data retrieval call binding the contract method 0x029a36fc.
//
// Solidity: function CertNFT() view returns(address)
func (_EduManage *EduManageSession) CertNFT() (common.Address, error) {
	return _EduManage.Contract.CertNFT(&_EduManage.CallOpts)
}

// CertNFT is a free data retrieval call binding the contract method 0x029a36fc.
//
// Solidity: function CertNFT() view returns(address)
func (_EduManage *EduManageCallerSession) CertNFT() (common.Address, error) {
	return _EduManage.Contract.CertNFT(&_EduManage.CallOpts)
}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_EduManage *EduManageCaller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_EduManage *EduManageSession) LINK() (common.Address, error) {
	return _EduManage.Contract.LINK(&_EduManage.CallOpts)
}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_EduManage *EduManageCallerSession) LINK() (common.Address, error) {
	return _EduManage.Contract.LINK(&_EduManage.CallOpts)
}

// AllowanceBuyCourse is a free data retrieval call binding the contract method 0x9337c079.
//
// Solidity: function allowanceBuyCourse() view returns(uint256)
func (_EduManage *EduManageCaller) AllowanceBuyCourse(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "allowanceBuyCourse")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowanceBuyCourse is a free data retrieval call binding the contract method 0x9337c079.
//
// Solidity: function allowanceBuyCourse() view returns(uint256)
func (_EduManage *EduManageSession) AllowanceBuyCourse() (*big.Int, error) {
	return _EduManage.Contract.AllowanceBuyCourse(&_EduManage.CallOpts)
}

// AllowanceBuyCourse is a free data retrieval call binding the contract method 0x9337c079.
//
// Solidity: function allowanceBuyCourse() view returns(uint256)
func (_EduManage *EduManageCallerSession) AllowanceBuyCourse() (*big.Int, error) {
	return _EduManage.Contract.AllowanceBuyCourse(&_EduManage.CallOpts)
}

// CheckEnrolledCourse is a free data retrieval call binding the contract method 0x4887bef0.
//
// Solidity: function checkEnrolledCourse(uint256 stu_id, uint256 course_id) view returns(bool)
func (_EduManage *EduManageCaller) CheckEnrolledCourse(opts *bind.CallOpts, stu_id *big.Int, course_id *big.Int) (bool, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "checkEnrolledCourse", stu_id, course_id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckEnrolledCourse is a free data retrieval call binding the contract method 0x4887bef0.
//
// Solidity: function checkEnrolledCourse(uint256 stu_id, uint256 course_id) view returns(bool)
func (_EduManage *EduManageSession) CheckEnrolledCourse(stu_id *big.Int, course_id *big.Int) (bool, error) {
	return _EduManage.Contract.CheckEnrolledCourse(&_EduManage.CallOpts, stu_id, course_id)
}

// CheckEnrolledCourse is a free data retrieval call binding the contract method 0x4887bef0.
//
// Solidity: function checkEnrolledCourse(uint256 stu_id, uint256 course_id) view returns(bool)
func (_EduManage *EduManageCallerSession) CheckEnrolledCourse(stu_id *big.Int, course_id *big.Int) (bool, error) {
	return _EduManage.Contract.CheckEnrolledCourse(&_EduManage.CallOpts, stu_id, course_id)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_EduManage *EduManageCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_EduManage *EduManageSession) Decimals() (*big.Int, error) {
	return _EduManage.Contract.Decimals(&_EduManage.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_EduManage *EduManageCallerSession) Decimals() (*big.Int, error) {
	return _EduManage.Contract.Decimals(&_EduManage.CallOpts)
}

// GetCourse is a free data retrieval call binding the contract method 0x0b91e28d.
//
// Solidity: function getCourse(uint256 _id) view returns((uint256,string,uint256,uint256,uint256,string,string,uint256))
func (_EduManage *EduManageCaller) GetCourse(opts *bind.CallOpts, _id *big.Int) (EduManageCourse, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "getCourse", _id)

	if err != nil {
		return *new(EduManageCourse), err
	}

	out0 := *abi.ConvertType(out[0], new(EduManageCourse)).(*EduManageCourse)

	return out0, err

}

// GetCourse is a free data retrieval call binding the contract method 0x0b91e28d.
//
// Solidity: function getCourse(uint256 _id) view returns((uint256,string,uint256,uint256,uint256,string,string,uint256))
func (_EduManage *EduManageSession) GetCourse(_id *big.Int) (EduManageCourse, error) {
	return _EduManage.Contract.GetCourse(&_EduManage.CallOpts, _id)
}

// GetCourse is a free data retrieval call binding the contract method 0x0b91e28d.
//
// Solidity: function getCourse(uint256 _id) view returns((uint256,string,uint256,uint256,uint256,string,string,uint256))
func (_EduManage *EduManageCallerSession) GetCourse(_id *big.Int) (EduManageCourse, error) {
	return _EduManage.Contract.GetCourse(&_EduManage.CallOpts, _id)
}

// GetCourseCount is a free data retrieval call binding the contract method 0x96cfda06.
//
// Solidity: function getCourseCount() view returns(uint256)
func (_EduManage *EduManageCaller) GetCourseCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "getCourseCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCourseCount is a free data retrieval call binding the contract method 0x96cfda06.
//
// Solidity: function getCourseCount() view returns(uint256)
func (_EduManage *EduManageSession) GetCourseCount() (*big.Int, error) {
	return _EduManage.Contract.GetCourseCount(&_EduManage.CallOpts)
}

// GetCourseCount is a free data retrieval call binding the contract method 0x96cfda06.
//
// Solidity: function getCourseCount() view returns(uint256)
func (_EduManage *EduManageCallerSession) GetCourseCount() (*big.Int, error) {
	return _EduManage.Contract.GetCourseCount(&_EduManage.CallOpts)
}

// GetGrades is a free data retrieval call binding the contract method 0x3956c5d5.
//
// Solidity: function getGrades(uint256 _stu_id, uint256 _course_id) view returns((uint256,string,bool,bool,bool))
func (_EduManage *EduManageCaller) GetGrades(opts *bind.CallOpts, _stu_id *big.Int, _course_id *big.Int) (EduManageGrade, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "getGrades", _stu_id, _course_id)

	if err != nil {
		return *new(EduManageGrade), err
	}

	out0 := *abi.ConvertType(out[0], new(EduManageGrade)).(*EduManageGrade)

	return out0, err

}

// GetGrades is a free data retrieval call binding the contract method 0x3956c5d5.
//
// Solidity: function getGrades(uint256 _stu_id, uint256 _course_id) view returns((uint256,string,bool,bool,bool))
func (_EduManage *EduManageSession) GetGrades(_stu_id *big.Int, _course_id *big.Int) (EduManageGrade, error) {
	return _EduManage.Contract.GetGrades(&_EduManage.CallOpts, _stu_id, _course_id)
}

// GetGrades is a free data retrieval call binding the contract method 0x3956c5d5.
//
// Solidity: function getGrades(uint256 _stu_id, uint256 _course_id) view returns((uint256,string,bool,bool,bool))
func (_EduManage *EduManageCallerSession) GetGrades(_stu_id *big.Int, _course_id *big.Int) (EduManageGrade, error) {
	return _EduManage.Contract.GetGrades(&_EduManage.CallOpts, _stu_id, _course_id)
}

// GetPosts is a free data retrieval call binding the contract method 0x1d45cddc.
//
// Solidity: function getPosts(uint256 _id) view returns((uint256,address,uint256,uint256,uint256))
func (_EduManage *EduManageCaller) GetPosts(opts *bind.CallOpts, _id *big.Int) (EduManagePost, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "getPosts", _id)

	if err != nil {
		return *new(EduManagePost), err
	}

	out0 := *abi.ConvertType(out[0], new(EduManagePost)).(*EduManagePost)

	return out0, err

}

// GetPosts is a free data retrieval call binding the contract method 0x1d45cddc.
//
// Solidity: function getPosts(uint256 _id) view returns((uint256,address,uint256,uint256,uint256))
func (_EduManage *EduManageSession) GetPosts(_id *big.Int) (EduManagePost, error) {
	return _EduManage.Contract.GetPosts(&_EduManage.CallOpts, _id)
}

// GetPosts is a free data retrieval call binding the contract method 0x1d45cddc.
//
// Solidity: function getPosts(uint256 _id) view returns((uint256,address,uint256,uint256,uint256))
func (_EduManage *EduManageCallerSession) GetPosts(_id *big.Int) (EduManagePost, error) {
	return _EduManage.Contract.GetPosts(&_EduManage.CallOpts, _id)
}

// GetStudent is a free data retrieval call binding the contract method 0x642f9163.
//
// Solidity: function getStudent(uint256 _id) view returns((uint256,address,string,uint256[]))
func (_EduManage *EduManageCaller) GetStudent(opts *bind.CallOpts, _id *big.Int) (EduManageStudent, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "getStudent", _id)

	if err != nil {
		return *new(EduManageStudent), err
	}

	out0 := *abi.ConvertType(out[0], new(EduManageStudent)).(*EduManageStudent)

	return out0, err

}

// GetStudent is a free data retrieval call binding the contract method 0x642f9163.
//
// Solidity: function getStudent(uint256 _id) view returns((uint256,address,string,uint256[]))
func (_EduManage *EduManageSession) GetStudent(_id *big.Int) (EduManageStudent, error) {
	return _EduManage.Contract.GetStudent(&_EduManage.CallOpts, _id)
}

// GetStudent is a free data retrieval call binding the contract method 0x642f9163.
//
// Solidity: function getStudent(uint256 _id) view returns((uint256,address,string,uint256[]))
func (_EduManage *EduManageCallerSession) GetStudent(_id *big.Int) (EduManageStudent, error) {
	return _EduManage.Contract.GetStudent(&_EduManage.CallOpts, _id)
}

// Grades is a free data retrieval call binding the contract method 0x5bd6fa74.
//
// Solidity: function grades(uint256 , uint256 ) view returns(uint256 mark, string hashResultExam, bool isSet, bool isReceiveReward, bool isMintCert)
func (_EduManage *EduManageCaller) Grades(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Mark            *big.Int
	HashResultExam  string
	IsSet           bool
	IsReceiveReward bool
	IsMintCert      bool
}, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "grades", arg0, arg1)

	outstruct := new(struct {
		Mark            *big.Int
		HashResultExam  string
		IsSet           bool
		IsReceiveReward bool
		IsMintCert      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mark = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.HashResultExam = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.IsSet = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.IsReceiveReward = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsMintCert = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Grades is a free data retrieval call binding the contract method 0x5bd6fa74.
//
// Solidity: function grades(uint256 , uint256 ) view returns(uint256 mark, string hashResultExam, bool isSet, bool isReceiveReward, bool isMintCert)
func (_EduManage *EduManageSession) Grades(arg0 *big.Int, arg1 *big.Int) (struct {
	Mark            *big.Int
	HashResultExam  string
	IsSet           bool
	IsReceiveReward bool
	IsMintCert      bool
}, error) {
	return _EduManage.Contract.Grades(&_EduManage.CallOpts, arg0, arg1)
}

// Grades is a free data retrieval call binding the contract method 0x5bd6fa74.
//
// Solidity: function grades(uint256 , uint256 ) view returns(uint256 mark, string hashResultExam, bool isSet, bool isReceiveReward, bool isMintCert)
func (_EduManage *EduManageCallerSession) Grades(arg0 *big.Int, arg1 *big.Int) (struct {
	Mark            *big.Int
	HashResultExam  string
	IsSet           bool
	IsReceiveReward bool
	IsMintCert      bool
}, error) {
	return _EduManage.Contract.Grades(&_EduManage.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EduManage *EduManageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EduManage *EduManageSession) Owner() (common.Address, error) {
	return _EduManage.Contract.Owner(&_EduManage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EduManage *EduManageCallerSession) Owner() (common.Address, error) {
	return _EduManage.Contract.Owner(&_EduManage.CallOpts)
}

// PostPointToReward is a free data retrieval call binding the contract method 0xa351c287.
//
// Solidity: function postPointToReward() view returns(uint256)
func (_EduManage *EduManageCaller) PostPointToReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "postPointToReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostPointToReward is a free data retrieval call binding the contract method 0xa351c287.
//
// Solidity: function postPointToReward() view returns(uint256)
func (_EduManage *EduManageSession) PostPointToReward() (*big.Int, error) {
	return _EduManage.Contract.PostPointToReward(&_EduManage.CallOpts)
}

// PostPointToReward is a free data retrieval call binding the contract method 0xa351c287.
//
// Solidity: function postPointToReward() view returns(uint256)
func (_EduManage *EduManageCallerSession) PostPointToReward() (*big.Int, error) {
	return _EduManage.Contract.PostPointToReward(&_EduManage.CallOpts)
}

// PostTokenReward is a free data retrieval call binding the contract method 0xe8e07882.
//
// Solidity: function postTokenReward() view returns(uint256)
func (_EduManage *EduManageCaller) PostTokenReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EduManage.contract.Call(opts, &out, "postTokenReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostTokenReward is a free data retrieval call binding the contract method 0xe8e07882.
//
// Solidity: function postTokenReward() view returns(uint256)
func (_EduManage *EduManageSession) PostTokenReward() (*big.Int, error) {
	return _EduManage.Contract.PostTokenReward(&_EduManage.CallOpts)
}

// PostTokenReward is a free data retrieval call binding the contract method 0xe8e07882.
//
// Solidity: function postTokenReward() view returns(uint256)
func (_EduManage *EduManageCallerSession) PostTokenReward() (*big.Int, error) {
	return _EduManage.Contract.PostTokenReward(&_EduManage.CallOpts)
}

// AddCourse is a paid mutator transaction binding the contract method 0xd3bc9aca.
//
// Solidity: function addCourse(uint256 course_id, string name, uint256 price, uint256 reward, string courseType, string hashCourse) returns()
func (_EduManage *EduManageTransactor) AddCourse(opts *bind.TransactOpts, course_id *big.Int, name string, price *big.Int, reward *big.Int, courseType string, hashCourse string) (*types.Transaction, error) {
	return _EduManage.contract.Transact(opts, "addCourse", course_id, name, price, reward, courseType, hashCourse)
}

// AddCourse is a paid mutator transaction binding the contract method 0xd3bc9aca.
//
// Solidity: function addCourse(uint256 course_id, string name, uint256 price, uint256 reward, string courseType, string hashCourse) returns()
func (_EduManage *EduManageSession) AddCourse(course_id *big.Int, name string, price *big.Int, reward *big.Int, courseType string, hashCourse string) (*types.Transaction, error) {
	return _EduManage.Contract.AddCourse(&_EduManage.TransactOpts, course_id, name, price, reward, courseType, hashCourse)
}

// AddCourse is a paid mutator transaction binding the contract method 0xd3bc9aca.
//
// Solidity: function addCourse(uint256 course_id, string name, uint256 price, uint256 reward, string courseType, string hashCourse) returns()
func (_EduManage *EduManageTransactorSession) AddCourse(course_id *big.Int, name string, price *big.Int, reward *big.Int, courseType string, hashCourse string) (*types.Transaction, error) {
	return _EduManage.Contract.AddCourse(&_EduManage.TransactOpts, course_id, name, price, reward, courseType, hashCourse)
}

// AddPost is a paid mutator transaction binding the contract method 0x016a7522.
//
// Solidity: function addPost(uint256 post_id, address owner_add) returns()
func (_EduManage *EduManageTransactor) AddPost(opts *bind.TransactOpts, post_id *big.Int, owner_add common.Address) (*types.Transaction, error) {
	return _EduManage.contract.Transact(opts, "addPost", post_id, owner_add)
}

// AddPost is a paid mutator transaction binding the contract method 0x016a7522.
//
// Solidity: function addPost(uint256 post_id, address owner_add) returns()
func (_EduManage *EduManageSession) AddPost(post_id *big.Int, owner_add common.Address) (*types.Transaction, error) {
	return _EduManage.Contract.AddPost(&_EduManage.TransactOpts, post_id, owner_add)
}

// AddPost is a paid mutator transaction binding the contract method 0x016a7522.
//
// Solidity: function addPost(uint256 post_id, address owner_add) returns()
func (_EduManage *EduManageTransactorSession) AddPost(post_id *big.Int, owner_add common.Address) (*types.Transaction, error) {
	return _EduManage.Contract.AddPost(&_EduManage.TransactOpts, post_id, owner_add)
}

// AddStudent is a paid mutator transaction binding the contract method 0x60e5f3bb.
//
// Solidity: function addStudent(address stu_add, uint256 stu_id, string name) returns()
func (_EduManage *EduManageTransactor) AddStudent(opts *bind.TransactOpts, stu_add common.Address, stu_id *big.Int, name string) (*types.Transaction, error) {
	return _EduManage.contract.Transact(opts, "addStudent", stu_add, stu_id, name)
}

// AddStudent is a paid mutator transaction binding the contract method 0x60e5f3bb.
//
// Solidity: function addStudent(address stu_add, uint256 stu_id, string name) returns()
func (_EduManage *EduManageSession) AddStudent(stu_add common.Address, stu_id *big.Int, name string) (*types.Transaction, error) {
	return _EduManage.Contract.AddStudent(&_EduManage.TransactOpts, stu_add, stu_id, name)
}

// AddStudent is a paid mutator transaction binding the contract method 0x60e5f3bb.
//
// Solidity: function addStudent(address stu_add, uint256 stu_id, string name) returns()
func (_EduManage *EduManageTransactorSession) AddStudent(stu_add common.Address, stu_id *big.Int, name string) (*types.Transaction, error) {
	return _EduManage.Contract.AddStudent(&_EduManage.TransactOpts, stu_add, stu_id, name)
}

// BuyCourse is a paid mutator transaction binding the contract method 0x99dbff01.
//
// Solidity: function buyCourse(uint256 stu_id, uint256 course_id) payable returns()
func (_EduManage *EduManageTransactor) BuyCourse(opts *bind.TransactOpts, stu_id *big.Int, course_id *big.Int) (*types.Transaction, error) {
	return _EduManage.contract.Transact(opts, "buyCourse", stu_id, course_id)
}

// BuyCourse is a paid mutator transaction binding the contract method 0x99dbff01.
//
// Solidity: function buyCourse(uint256 stu_id, uint256 course_id) payable returns()
func (_EduManage *EduManageSession) BuyCourse(stu_id *big.Int, course_id *big.Int) (*types.Transaction, error) {
	return _EduManage.Contract.BuyCourse(&_EduManage.TransactOpts, stu_id, course_id)
}

// BuyCourse is a paid mutator transaction binding the contract method 0x99dbff01.
//
// Solidity: function buyCourse(uint256 stu_id, uint256 course_id) payable returns()
func (_EduManage *EduManageTransactorSession) BuyCourse(stu_id *big.Int, course_id *big.Int) (*types.Transaction, error) {
	return _EduManage.Contract.BuyCourse(&_EduManage.TransactOpts, stu_id, course_id)
}

// CheckAndTransferRewardCourse is a paid mutator transaction binding the contract method 0x48339a89.
//
// Solidity: function checkAndTransferRewardCourse(uint256 stu_id, uint256 course_id, uint256 token_id, string image_uri) payable returns()
func (_EduManage *EduManageTransactor) CheckAndTransferRewardCourse(opts *bind.TransactOpts, stu_id *big.Int, course_id *big.Int, token_id *big.Int, image_uri string) (*types.Transaction, error) {
	return _EduManage.contract.Transact(opts, "checkAndTransferRewardCourse", stu_id, course_id, token_id, image_uri)
}

// CheckAndTransferRewardCourse is a paid mutator transaction binding the contract method 0x48339a89.
//
// Solidity: function checkAndTransferRewardCourse(uint256 stu_id, uint256 course_id, uint256 token_id, string image_uri) payable returns()
func (_EduManage *EduManageSession) CheckAndTransferRewardCourse(stu_id *big.Int, course_id *big.Int, token_id *big.Int, image_uri string) (*types.Transaction, error) {
	return _EduManage.Contract.CheckAndTransferRewardCourse(&_EduManage.TransactOpts, stu_id, course_id, token_id, image_uri)
}

// CheckAndTransferRewardCourse is a paid mutator transaction binding the contract method 0x48339a89.
//
// Solidity: function checkAndTransferRewardCourse(uint256 stu_id, uint256 course_id, uint256 token_id, string image_uri) payable returns()
func (_EduManage *EduManageTransactorSession) CheckAndTransferRewardCourse(stu_id *big.Int, course_id *big.Int, token_id *big.Int, image_uri string) (*types.Transaction, error) {
	return _EduManage.Contract.CheckAndTransferRewardCourse(&_EduManage.TransactOpts, stu_id, course_id, token_id, image_uri)
}

// CheckAndTransferRewardPost is a paid mutator transaction binding the contract method 0xeda0a4dc.
//
// Solidity: function checkAndTransferRewardPost(uint256 post_id, uint256 point) payable returns()
func (_EduManage *EduManageTransactor) CheckAndTransferRewardPost(opts *bind.TransactOpts, post_id *big.Int, point *big.Int) (*types.Transaction, error) {
	return _EduManage.contract.Transact(opts, "checkAndTransferRewardPost", post_id, point)
}

// CheckAndTransferRewardPost is a paid mutator transaction binding the contract method 0xeda0a4dc.
//
// Solidity: function checkAndTransferRewardPost(uint256 post_id, uint256 point) payable returns()
func (_EduManage *EduManageSession) CheckAndTransferRewardPost(post_id *big.Int, point *big.Int) (*types.Transaction, error) {
	return _EduManage.Contract.CheckAndTransferRewardPost(&_EduManage.TransactOpts, post_id, point)
}

// CheckAndTransferRewardPost is a paid mutator transaction binding the contract method 0xeda0a4dc.
//
// Solidity: function checkAndTransferRewardPost(uint256 post_id, uint256 point) payable returns()
func (_EduManage *EduManageTransactorSession) CheckAndTransferRewardPost(post_id *big.Int, point *big.Int) (*types.Transaction, error) {
	return _EduManage.Contract.CheckAndTransferRewardPost(&_EduManage.TransactOpts, post_id, point)
}

// RewardToken is a paid mutator transaction binding the contract method 0xfb41138c.
//
// Solidity: function rewardToken(address recipient, uint256 _numberOfTokens) payable returns()
func (_EduManage *EduManageTransactor) RewardToken(opts *bind.TransactOpts, recipient common.Address, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _EduManage.contract.Transact(opts, "rewardToken", recipient, _numberOfTokens)
}

// RewardToken is a paid mutator transaction binding the contract method 0xfb41138c.
//
// Solidity: function rewardToken(address recipient, uint256 _numberOfTokens) payable returns()
func (_EduManage *EduManageSession) RewardToken(recipient common.Address, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _EduManage.Contract.RewardToken(&_EduManage.TransactOpts, recipient, _numberOfTokens)
}

// RewardToken is a paid mutator transaction binding the contract method 0xfb41138c.
//
// Solidity: function rewardToken(address recipient, uint256 _numberOfTokens) payable returns()
func (_EduManage *EduManageTransactorSession) RewardToken(recipient common.Address, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _EduManage.Contract.RewardToken(&_EduManage.TransactOpts, recipient, _numberOfTokens)
}

// SubmitGrade is a paid mutator transaction binding the contract method 0x40a7497d.
//
// Solidity: function submitGrade(uint256 stu_id, uint256 course_id, uint256 _mark, string hash_result_exam) returns(bool)
func (_EduManage *EduManageTransactor) SubmitGrade(opts *bind.TransactOpts, stu_id *big.Int, course_id *big.Int, _mark *big.Int, hash_result_exam string) (*types.Transaction, error) {
	return _EduManage.contract.Transact(opts, "submitGrade", stu_id, course_id, _mark, hash_result_exam)
}

// SubmitGrade is a paid mutator transaction binding the contract method 0x40a7497d.
//
// Solidity: function submitGrade(uint256 stu_id, uint256 course_id, uint256 _mark, string hash_result_exam) returns(bool)
func (_EduManage *EduManageSession) SubmitGrade(stu_id *big.Int, course_id *big.Int, _mark *big.Int, hash_result_exam string) (*types.Transaction, error) {
	return _EduManage.Contract.SubmitGrade(&_EduManage.TransactOpts, stu_id, course_id, _mark, hash_result_exam)
}

// SubmitGrade is a paid mutator transaction binding the contract method 0x40a7497d.
//
// Solidity: function submitGrade(uint256 stu_id, uint256 course_id, uint256 _mark, string hash_result_exam) returns(bool)
func (_EduManage *EduManageTransactorSession) SubmitGrade(stu_id *big.Int, course_id *big.Int, _mark *big.Int, hash_result_exam string) (*types.Transaction, error) {
	return _EduManage.Contract.SubmitGrade(&_EduManage.TransactOpts, stu_id, course_id, _mark, hash_result_exam)
}

// EduManageCertificateMintedIterator is returned from FilterCertificateMinted and is used to iterate over the raw logs and unpacked data for CertificateMinted events raised by the EduManage contract.
type EduManageCertificateMintedIterator struct {
	Event *EduManageCertificateMinted // Event containing the contract specifics and raw log

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
func (it *EduManageCertificateMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EduManageCertificateMinted)
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
		it.Event = new(EduManageCertificateMinted)
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
func (it *EduManageCertificateMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EduManageCertificateMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EduManageCertificateMinted represents a CertificateMinted event raised by the EduManage contract.
type EduManageCertificateMinted struct {
	Recipient common.Address
	TokenId   *big.Int
	Name      string
	Course    string
	Date      *big.Int
	CerType   string
	ImageUri  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCertificateMinted is a free log retrieval operation binding the contract event 0xa4e45d12c1f3fdec083d18975dfe779a14e4ccdc4d29ce3a24761bebac3bf4f8.
//
// Solidity: event CertificateMinted(address recipient, uint256 tokenId, string name, string course, uint256 date, string cerType, string imageUri)
func (_EduManage *EduManageFilterer) FilterCertificateMinted(opts *bind.FilterOpts) (*EduManageCertificateMintedIterator, error) {

	logs, sub, err := _EduManage.contract.FilterLogs(opts, "CertificateMinted")
	if err != nil {
		return nil, err
	}
	return &EduManageCertificateMintedIterator{contract: _EduManage.contract, event: "CertificateMinted", logs: logs, sub: sub}, nil
}

// WatchCertificateMinted is a free log subscription operation binding the contract event 0xa4e45d12c1f3fdec083d18975dfe779a14e4ccdc4d29ce3a24761bebac3bf4f8.
//
// Solidity: event CertificateMinted(address recipient, uint256 tokenId, string name, string course, uint256 date, string cerType, string imageUri)
func (_EduManage *EduManageFilterer) WatchCertificateMinted(opts *bind.WatchOpts, sink chan<- *EduManageCertificateMinted) (event.Subscription, error) {

	logs, sub, err := _EduManage.contract.WatchLogs(opts, "CertificateMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EduManageCertificateMinted)
				if err := _EduManage.contract.UnpackLog(event, "CertificateMinted", log); err != nil {
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

// ParseCertificateMinted is a log parse operation binding the contract event 0xa4e45d12c1f3fdec083d18975dfe779a14e4ccdc4d29ce3a24761bebac3bf4f8.
//
// Solidity: event CertificateMinted(address recipient, uint256 tokenId, string name, string course, uint256 date, string cerType, string imageUri)
func (_EduManage *EduManageFilterer) ParseCertificateMinted(log types.Log) (*EduManageCertificateMinted, error) {
	event := new(EduManageCertificateMinted)
	if err := _EduManage.contract.UnpackLog(event, "CertificateMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
