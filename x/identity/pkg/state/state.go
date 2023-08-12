// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package state

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

// IStateGistProof is an auto generated low-level Go binding around an user-defined struct.
type IStateGistProof struct {
	Root         *big.Int
	Existence    bool
	Siblings     [64]*big.Int
	Index        *big.Int
	Value        *big.Int
	AuxExistence bool
	AuxIndex     *big.Int
	AuxValue     *big.Int
}

// IStateGistRootInfo is an auto generated low-level Go binding around an user-defined struct.
type IStateGistRootInfo struct {
	Root                *big.Int
	ReplacedByRoot      *big.Int
	CreatedAtTimestamp  *big.Int
	ReplacedAtTimestamp *big.Int
	CreatedAtBlock      *big.Int
	ReplacedAtBlock     *big.Int
}

// IStateStateInfo is an auto generated low-level Go binding around an user-defined struct.
type IStateStateInfo struct {
	Id                  *big.Int
	State               *big.Int
	ReplacedByState     *big.Int
	CreatedAtTimestamp  *big.Int
	ReplacedAtTimestamp *big.Int
	CreatedAtBlock      *big.Int
	ReplacedAtBlock     *big.Int
}

// StateMetaData contains all meta data concerning the State contract.
var StateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gistRoot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"StateTransited\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStateTransitionVerifier\",\"name\":\"verifierContractAddr_\",\"type\":\"address\"}],\"name\":\"__State_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"getGISTProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber_\",\"type\":\"uint256\"}],\"name\":\"getGISTProofByBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root_\",\"type\":\"uint256\"}],\"name\":\"getGISTProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp_\",\"type\":\"uint256\"}],\"name\":\"getGISTProofByTime\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGISTRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length_\",\"type\":\"uint256\"}],\"name\":\"getGISTRootHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGISTRootHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root_\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber_\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfoByBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp_\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfoByTime\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"getStateInfoById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state_\",\"type\":\"uint256\"}],\"name\":\"getStateInfoByIdAndState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length_\",\"type\":\"uint256\"}],\"name\":\"getStateInfoHistoryById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.StateInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"getStateInfoHistoryLengthById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"name\":\"idExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVerifierAddr_\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state_\",\"type\":\"uint256\"}],\"name\":\"stateExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldState_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newState_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOldStateGenesis_\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"a_\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b_\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c_\",\"type\":\"uint256[2]\"}],\"name\":\"transitState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StateABI is the input ABI used to generate the binding from.
// Deprecated: Use StateMetaData.ABI instead.
var StateABI = StateMetaData.ABI

// State is an auto generated Go binding around an Ethereum contract.
type State struct {
	StateCaller     // Read-only binding to the contract
	StateTransactor // Write-only binding to the contract
	StateFilterer   // Log filterer for contract events
}

// StateCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateSession struct {
	Contract     *State            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateCallerSession struct {
	Contract *StateCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateTransactorSession struct {
	Contract     *StateTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateRaw struct {
	Contract *State // Generic contract binding to access the raw methods on
}

// StateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateCallerRaw struct {
	Contract *StateCaller // Generic read-only contract binding to access the raw methods on
}

// StateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateTransactorRaw struct {
	Contract *StateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewState creates a new instance of State, bound to a specific deployed contract.
func NewState(address common.Address, backend bind.ContractBackend) (*State, error) {
	contract, err := bindState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &State{StateCaller: StateCaller{contract: contract}, StateTransactor: StateTransactor{contract: contract}, StateFilterer: StateFilterer{contract: contract}}, nil
}

// NewStateCaller creates a new read-only instance of State, bound to a specific deployed contract.
func NewStateCaller(address common.Address, caller bind.ContractCaller) (*StateCaller, error) {
	contract, err := bindState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateCaller{contract: contract}, nil
}

// NewStateTransactor creates a new write-only instance of State, bound to a specific deployed contract.
func NewStateTransactor(address common.Address, transactor bind.ContractTransactor) (*StateTransactor, error) {
	contract, err := bindState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateTransactor{contract: contract}, nil
}

// NewStateFilterer creates a new log filterer instance of State, bound to a specific deployed contract.
func NewStateFilterer(address common.Address, filterer bind.ContractFilterer) (*StateFilterer, error) {
	contract, err := bindState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateFilterer{contract: contract}, nil
}

// bindState binds a generic wrapper to an already deployed contract.
func bindState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_State *StateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _State.Contract.StateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_State *StateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.Contract.StateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_State *StateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _State.Contract.StateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_State *StateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _State.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_State *StateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_State *StateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _State.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_State *StateCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_State *StateSession) VERSION() (string, error) {
	return _State.Contract.VERSION(&_State.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_State *StateCallerSession) VERSION() (string, error) {
	return _State.Contract.VERSION(&_State.CallOpts)
}

// GetGISTProof is a free data retrieval call binding the contract method 0x3025bb8c.
//
// Solidity: function getGISTProof(uint256 id_) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_State *StateCaller) GetGISTProof(opts *bind.CallOpts, id_ *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getGISTProof", id_)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProof is a free data retrieval call binding the contract method 0x3025bb8c.
//
// Solidity: function getGISTProof(uint256 id_) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_State *StateSession) GetGISTProof(id_ *big.Int) (IStateGistProof, error) {
	return _State.Contract.GetGISTProof(&_State.CallOpts, id_)
}

// GetGISTProof is a free data retrieval call binding the contract method 0x3025bb8c.
//
// Solidity: function getGISTProof(uint256 id_) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_State *StateCallerSession) GetGISTProof(id_ *big.Int) (IStateGistProof, error) {
	return _State.Contract.GetGISTProof(&_State.CallOpts, id_)
}

// GetGISTProofByBlock is a free data retrieval call binding the contract method 0x046ff140.
//
// Solidity: function getGISTProofByBlock(uint256 id_, uint256 blockNumber_) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_State *StateCaller) GetGISTProofByBlock(opts *bind.CallOpts, id_ *big.Int, blockNumber_ *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getGISTProofByBlock", id_, blockNumber_)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProofByBlock is a free data retrieval call binding the contract method 0x046ff140.
//
// Solidity: function getGISTProofByBlock(uint256 id_, uint256 blockNumber_) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_State *StateSession) GetGISTProofByBlock(id_ *big.Int, blockNumber_ *big.Int) (IStateGistProof, error) {
	return _State.Contract.GetGISTProofByBlock(&_State.CallOpts, id_, blockNumber_)
}

// GetGISTProofByBlock is a free data retrieval call binding the contract method 0x046ff140.
//
// Solidity: function getGISTProofByBlock(uint256 id_, uint256 blockNumber_) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_State *StateCallerSession) GetGISTProofByBlock(id_ *big.Int, blockNumber_ *big.Int) (IStateGistProof, error) {
	return _State.Contract.GetGISTProofByBlock(&_State.CallOpts, id_, blockNumber_)
}

// GetGISTProofByRoot is a free data retrieval call binding the contract method 0xe12a36c0.
//
// Solidity: function getGISTProofByRoot(uint256 id_, uint256 root_) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_State *StateCaller) GetGISTProofByRoot(opts *bind.CallOpts, id_ *big.Int, root_ *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getGISTProofByRoot", id_, root_)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProofByRoot is a free data retrieval call binding the contract method 0xe12a36c0.
//
// Solidity: function getGISTProofByRoot(uint256 id_, uint256 root_) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_State *StateSession) GetGISTProofByRoot(id_ *big.Int, root_ *big.Int) (IStateGistProof, error) {
	return _State.Contract.GetGISTProofByRoot(&_State.CallOpts, id_, root_)
}

// GetGISTProofByRoot is a free data retrieval call binding the contract method 0xe12a36c0.
//
// Solidity: function getGISTProofByRoot(uint256 id_, uint256 root_) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_State *StateCallerSession) GetGISTProofByRoot(id_ *big.Int, root_ *big.Int) (IStateGistProof, error) {
	return _State.Contract.GetGISTProofByRoot(&_State.CallOpts, id_, root_)
}

// GetGISTProofByTime is a free data retrieval call binding the contract method 0xd51afebf.
//
// Solidity: function getGISTProofByTime(uint256 id_, uint256 timestamp_) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_State *StateCaller) GetGISTProofByTime(opts *bind.CallOpts, id_ *big.Int, timestamp_ *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getGISTProofByTime", id_, timestamp_)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProofByTime is a free data retrieval call binding the contract method 0xd51afebf.
//
// Solidity: function getGISTProofByTime(uint256 id_, uint256 timestamp_) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_State *StateSession) GetGISTProofByTime(id_ *big.Int, timestamp_ *big.Int) (IStateGistProof, error) {
	return _State.Contract.GetGISTProofByTime(&_State.CallOpts, id_, timestamp_)
}

// GetGISTProofByTime is a free data retrieval call binding the contract method 0xd51afebf.
//
// Solidity: function getGISTProofByTime(uint256 id_, uint256 timestamp_) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_State *StateCallerSession) GetGISTProofByTime(id_ *big.Int, timestamp_ *big.Int) (IStateGistProof, error) {
	return _State.Contract.GetGISTProofByTime(&_State.CallOpts, id_, timestamp_)
}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_State *StateCaller) GetGISTRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getGISTRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_State *StateSession) GetGISTRoot() (*big.Int, error) {
	return _State.Contract.GetGISTRoot(&_State.CallOpts)
}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_State *StateCallerSession) GetGISTRoot() (*big.Int, error) {
	return _State.Contract.GetGISTRoot(&_State.CallOpts)
}

// GetGISTRootHistory is a free data retrieval call binding the contract method 0x2f7670e4.
//
// Solidity: function getGISTRootHistory(uint256 start_, uint256 length_) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_State *StateCaller) GetGISTRootHistory(opts *bind.CallOpts, start_ *big.Int, length_ *big.Int) ([]IStateGistRootInfo, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getGISTRootHistory", start_, length_)

	if err != nil {
		return *new([]IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStateGistRootInfo)).(*[]IStateGistRootInfo)

	return out0, err

}

// GetGISTRootHistory is a free data retrieval call binding the contract method 0x2f7670e4.
//
// Solidity: function getGISTRootHistory(uint256 start_, uint256 length_) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_State *StateSession) GetGISTRootHistory(start_ *big.Int, length_ *big.Int) ([]IStateGistRootInfo, error) {
	return _State.Contract.GetGISTRootHistory(&_State.CallOpts, start_, length_)
}

// GetGISTRootHistory is a free data retrieval call binding the contract method 0x2f7670e4.
//
// Solidity: function getGISTRootHistory(uint256 start_, uint256 length_) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_State *StateCallerSession) GetGISTRootHistory(start_ *big.Int, length_ *big.Int) ([]IStateGistRootInfo, error) {
	return _State.Contract.GetGISTRootHistory(&_State.CallOpts, start_, length_)
}

// GetGISTRootHistoryLength is a free data retrieval call binding the contract method 0xdccbd57a.
//
// Solidity: function getGISTRootHistoryLength() view returns(uint256)
func (_State *StateCaller) GetGISTRootHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getGISTRootHistoryLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGISTRootHistoryLength is a free data retrieval call binding the contract method 0xdccbd57a.
//
// Solidity: function getGISTRootHistoryLength() view returns(uint256)
func (_State *StateSession) GetGISTRootHistoryLength() (*big.Int, error) {
	return _State.Contract.GetGISTRootHistoryLength(&_State.CallOpts)
}

// GetGISTRootHistoryLength is a free data retrieval call binding the contract method 0xdccbd57a.
//
// Solidity: function getGISTRootHistoryLength() view returns(uint256)
func (_State *StateCallerSession) GetGISTRootHistoryLength() (*big.Int, error) {
	return _State.Contract.GetGISTRootHistoryLength(&_State.CallOpts)
}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root_) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateCaller) GetGISTRootInfo(opts *bind.CallOpts, root_ *big.Int) (IStateGistRootInfo, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getGISTRootInfo", root_)

	if err != nil {
		return *new(IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistRootInfo)).(*IStateGistRootInfo)

	return out0, err

}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root_) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateSession) GetGISTRootInfo(root_ *big.Int) (IStateGistRootInfo, error) {
	return _State.Contract.GetGISTRootInfo(&_State.CallOpts, root_)
}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root_) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateCallerSession) GetGISTRootInfo(root_ *big.Int) (IStateGistRootInfo, error) {
	return _State.Contract.GetGISTRootInfo(&_State.CallOpts, root_)
}

// GetGISTRootInfoByBlock is a free data retrieval call binding the contract method 0x5845e530.
//
// Solidity: function getGISTRootInfoByBlock(uint256 blockNumber_) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateCaller) GetGISTRootInfoByBlock(opts *bind.CallOpts, blockNumber_ *big.Int) (IStateGistRootInfo, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getGISTRootInfoByBlock", blockNumber_)

	if err != nil {
		return *new(IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistRootInfo)).(*IStateGistRootInfo)

	return out0, err

}

// GetGISTRootInfoByBlock is a free data retrieval call binding the contract method 0x5845e530.
//
// Solidity: function getGISTRootInfoByBlock(uint256 blockNumber_) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateSession) GetGISTRootInfoByBlock(blockNumber_ *big.Int) (IStateGistRootInfo, error) {
	return _State.Contract.GetGISTRootInfoByBlock(&_State.CallOpts, blockNumber_)
}

// GetGISTRootInfoByBlock is a free data retrieval call binding the contract method 0x5845e530.
//
// Solidity: function getGISTRootInfoByBlock(uint256 blockNumber_) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateCallerSession) GetGISTRootInfoByBlock(blockNumber_ *big.Int) (IStateGistRootInfo, error) {
	return _State.Contract.GetGISTRootInfoByBlock(&_State.CallOpts, blockNumber_)
}

// GetGISTRootInfoByTime is a free data retrieval call binding the contract method 0x0ef6e65b.
//
// Solidity: function getGISTRootInfoByTime(uint256 timestamp_) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateCaller) GetGISTRootInfoByTime(opts *bind.CallOpts, timestamp_ *big.Int) (IStateGistRootInfo, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getGISTRootInfoByTime", timestamp_)

	if err != nil {
		return *new(IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistRootInfo)).(*IStateGistRootInfo)

	return out0, err

}

// GetGISTRootInfoByTime is a free data retrieval call binding the contract method 0x0ef6e65b.
//
// Solidity: function getGISTRootInfoByTime(uint256 timestamp_) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateSession) GetGISTRootInfoByTime(timestamp_ *big.Int) (IStateGistRootInfo, error) {
	return _State.Contract.GetGISTRootInfoByTime(&_State.CallOpts, timestamp_)
}

// GetGISTRootInfoByTime is a free data retrieval call binding the contract method 0x0ef6e65b.
//
// Solidity: function getGISTRootInfoByTime(uint256 timestamp_) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateCallerSession) GetGISTRootInfoByTime(timestamp_ *big.Int) (IStateGistRootInfo, error) {
	return _State.Contract.GetGISTRootInfoByTime(&_State.CallOpts, timestamp_)
}

// GetStateInfoById is a free data retrieval call binding the contract method 0xb4bdea55.
//
// Solidity: function getStateInfoById(uint256 id_) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateCaller) GetStateInfoById(opts *bind.CallOpts, id_ *big.Int) (IStateStateInfo, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getStateInfoById", id_)

	if err != nil {
		return *new(IStateStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateStateInfo)).(*IStateStateInfo)

	return out0, err

}

// GetStateInfoById is a free data retrieval call binding the contract method 0xb4bdea55.
//
// Solidity: function getStateInfoById(uint256 id_) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateSession) GetStateInfoById(id_ *big.Int) (IStateStateInfo, error) {
	return _State.Contract.GetStateInfoById(&_State.CallOpts, id_)
}

// GetStateInfoById is a free data retrieval call binding the contract method 0xb4bdea55.
//
// Solidity: function getStateInfoById(uint256 id_) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateCallerSession) GetStateInfoById(id_ *big.Int) (IStateStateInfo, error) {
	return _State.Contract.GetStateInfoById(&_State.CallOpts, id_)
}

// GetStateInfoByIdAndState is a free data retrieval call binding the contract method 0x53c87312.
//
// Solidity: function getStateInfoByIdAndState(uint256 id_, uint256 state_) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateCaller) GetStateInfoByIdAndState(opts *bind.CallOpts, id_ *big.Int, state_ *big.Int) (IStateStateInfo, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getStateInfoByIdAndState", id_, state_)

	if err != nil {
		return *new(IStateStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateStateInfo)).(*IStateStateInfo)

	return out0, err

}

// GetStateInfoByIdAndState is a free data retrieval call binding the contract method 0x53c87312.
//
// Solidity: function getStateInfoByIdAndState(uint256 id_, uint256 state_) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateSession) GetStateInfoByIdAndState(id_ *big.Int, state_ *big.Int) (IStateStateInfo, error) {
	return _State.Contract.GetStateInfoByIdAndState(&_State.CallOpts, id_, state_)
}

// GetStateInfoByIdAndState is a free data retrieval call binding the contract method 0x53c87312.
//
// Solidity: function getStateInfoByIdAndState(uint256 id_, uint256 state_) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_State *StateCallerSession) GetStateInfoByIdAndState(id_ *big.Int, state_ *big.Int) (IStateStateInfo, error) {
	return _State.Contract.GetStateInfoByIdAndState(&_State.CallOpts, id_, state_)
}

// GetStateInfoHistoryById is a free data retrieval call binding the contract method 0xe99858fe.
//
// Solidity: function getStateInfoHistoryById(uint256 id_, uint256 startIndex_, uint256 length_) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_State *StateCaller) GetStateInfoHistoryById(opts *bind.CallOpts, id_ *big.Int, startIndex_ *big.Int, length_ *big.Int) ([]IStateStateInfo, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getStateInfoHistoryById", id_, startIndex_, length_)

	if err != nil {
		return *new([]IStateStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStateStateInfo)).(*[]IStateStateInfo)

	return out0, err

}

// GetStateInfoHistoryById is a free data retrieval call binding the contract method 0xe99858fe.
//
// Solidity: function getStateInfoHistoryById(uint256 id_, uint256 startIndex_, uint256 length_) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_State *StateSession) GetStateInfoHistoryById(id_ *big.Int, startIndex_ *big.Int, length_ *big.Int) ([]IStateStateInfo, error) {
	return _State.Contract.GetStateInfoHistoryById(&_State.CallOpts, id_, startIndex_, length_)
}

// GetStateInfoHistoryById is a free data retrieval call binding the contract method 0xe99858fe.
//
// Solidity: function getStateInfoHistoryById(uint256 id_, uint256 startIndex_, uint256 length_) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_State *StateCallerSession) GetStateInfoHistoryById(id_ *big.Int, startIndex_ *big.Int, length_ *big.Int) ([]IStateStateInfo, error) {
	return _State.Contract.GetStateInfoHistoryById(&_State.CallOpts, id_, startIndex_, length_)
}

// GetStateInfoHistoryLengthById is a free data retrieval call binding the contract method 0x676d5b5a.
//
// Solidity: function getStateInfoHistoryLengthById(uint256 id_) view returns(uint256)
func (_State *StateCaller) GetStateInfoHistoryLengthById(opts *bind.CallOpts, id_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getStateInfoHistoryLengthById", id_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStateInfoHistoryLengthById is a free data retrieval call binding the contract method 0x676d5b5a.
//
// Solidity: function getStateInfoHistoryLengthById(uint256 id_) view returns(uint256)
func (_State *StateSession) GetStateInfoHistoryLengthById(id_ *big.Int) (*big.Int, error) {
	return _State.Contract.GetStateInfoHistoryLengthById(&_State.CallOpts, id_)
}

// GetStateInfoHistoryLengthById is a free data retrieval call binding the contract method 0x676d5b5a.
//
// Solidity: function getStateInfoHistoryLengthById(uint256 id_) view returns(uint256)
func (_State *StateCallerSession) GetStateInfoHistoryLengthById(id_ *big.Int) (*big.Int, error) {
	return _State.Contract.GetStateInfoHistoryLengthById(&_State.CallOpts, id_)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_State *StateCaller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_State *StateSession) GetVerifier() (common.Address, error) {
	return _State.Contract.GetVerifier(&_State.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_State *StateCallerSession) GetVerifier() (common.Address, error) {
	return _State.Contract.GetVerifier(&_State.CallOpts)
}

// IdExists is a free data retrieval call binding the contract method 0x0b8a295a.
//
// Solidity: function idExists(uint256 id_) view returns(bool)
func (_State *StateCaller) IdExists(opts *bind.CallOpts, id_ *big.Int) (bool, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "idExists", id_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IdExists is a free data retrieval call binding the contract method 0x0b8a295a.
//
// Solidity: function idExists(uint256 id_) view returns(bool)
func (_State *StateSession) IdExists(id_ *big.Int) (bool, error) {
	return _State.Contract.IdExists(&_State.CallOpts, id_)
}

// IdExists is a free data retrieval call binding the contract method 0x0b8a295a.
//
// Solidity: function idExists(uint256 id_) view returns(bool)
func (_State *StateCallerSession) IdExists(id_ *big.Int) (bool, error) {
	return _State.Contract.IdExists(&_State.CallOpts, id_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_State *StateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_State *StateSession) Owner() (common.Address, error) {
	return _State.Contract.Owner(&_State.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_State *StateCallerSession) Owner() (common.Address, error) {
	return _State.Contract.Owner(&_State.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_State *StateCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_State *StateSession) PendingOwner() (common.Address, error) {
	return _State.Contract.PendingOwner(&_State.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_State *StateCallerSession) PendingOwner() (common.Address, error) {
	return _State.Contract.PendingOwner(&_State.CallOpts)
}

// StateExists is a free data retrieval call binding the contract method 0x233a4d23.
//
// Solidity: function stateExists(uint256 id_, uint256 state_) view returns(bool)
func (_State *StateCaller) StateExists(opts *bind.CallOpts, id_ *big.Int, state_ *big.Int) (bool, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "stateExists", id_, state_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StateExists is a free data retrieval call binding the contract method 0x233a4d23.
//
// Solidity: function stateExists(uint256 id_, uint256 state_) view returns(bool)
func (_State *StateSession) StateExists(id_ *big.Int, state_ *big.Int) (bool, error) {
	return _State.Contract.StateExists(&_State.CallOpts, id_, state_)
}

// StateExists is a free data retrieval call binding the contract method 0x233a4d23.
//
// Solidity: function stateExists(uint256 id_, uint256 state_) view returns(bool)
func (_State *StateCallerSession) StateExists(id_ *big.Int, state_ *big.Int) (bool, error) {
	return _State.Contract.StateExists(&_State.CallOpts, id_, state_)
}

// StateInit is a paid mutator transaction binding the contract method 0x3ed5b500.
//
// Solidity: function __State_init(address verifierContractAddr_) returns()
func (_State *StateTransactor) StateInit(opts *bind.TransactOpts, verifierContractAddr_ common.Address) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "__State_init", verifierContractAddr_)
}

// StateInit is a paid mutator transaction binding the contract method 0x3ed5b500.
//
// Solidity: function __State_init(address verifierContractAddr_) returns()
func (_State *StateSession) StateInit(verifierContractAddr_ common.Address) (*types.Transaction, error) {
	return _State.Contract.StateInit(&_State.TransactOpts, verifierContractAddr_)
}

// StateInit is a paid mutator transaction binding the contract method 0x3ed5b500.
//
// Solidity: function __State_init(address verifierContractAddr_) returns()
func (_State *StateTransactorSession) StateInit(verifierContractAddr_ common.Address) (*types.Transaction, error) {
	return _State.Contract.StateInit(&_State.TransactOpts, verifierContractAddr_)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_State *StateTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_State *StateSession) AcceptOwnership() (*types.Transaction, error) {
	return _State.Contract.AcceptOwnership(&_State.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_State *StateTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _State.Contract.AcceptOwnership(&_State.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_State *StateTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_State *StateSession) RenounceOwnership() (*types.Transaction, error) {
	return _State.Contract.RenounceOwnership(&_State.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_State *StateTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _State.Contract.RenounceOwnership(&_State.TransactOpts)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifierAddr_) returns()
func (_State *StateTransactor) SetVerifier(opts *bind.TransactOpts, newVerifierAddr_ common.Address) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "setVerifier", newVerifierAddr_)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifierAddr_) returns()
func (_State *StateSession) SetVerifier(newVerifierAddr_ common.Address) (*types.Transaction, error) {
	return _State.Contract.SetVerifier(&_State.TransactOpts, newVerifierAddr_)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifierAddr_) returns()
func (_State *StateTransactorSession) SetVerifier(newVerifierAddr_ common.Address) (*types.Transaction, error) {
	return _State.Contract.SetVerifier(&_State.TransactOpts, newVerifierAddr_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_State *StateTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_State *StateSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _State.Contract.TransferOwnership(&_State.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_State *StateTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _State.Contract.TransferOwnership(&_State.TransactOpts, newOwner)
}

// TransitState is a paid mutator transaction binding the contract method 0x28f88a65.
//
// Solidity: function transitState(uint256 id_, uint256 oldState_, uint256 newState_, bool isOldStateGenesis_, uint256[2] a_, uint256[2][2] b_, uint256[2] c_) returns()
func (_State *StateTransactor) TransitState(opts *bind.TransactOpts, id_ *big.Int, oldState_ *big.Int, newState_ *big.Int, isOldStateGenesis_ bool, a_ [2]*big.Int, b_ [2][2]*big.Int, c_ [2]*big.Int) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "transitState", id_, oldState_, newState_, isOldStateGenesis_, a_, b_, c_)
}

// TransitState is a paid mutator transaction binding the contract method 0x28f88a65.
//
// Solidity: function transitState(uint256 id_, uint256 oldState_, uint256 newState_, bool isOldStateGenesis_, uint256[2] a_, uint256[2][2] b_, uint256[2] c_) returns()
func (_State *StateSession) TransitState(id_ *big.Int, oldState_ *big.Int, newState_ *big.Int, isOldStateGenesis_ bool, a_ [2]*big.Int, b_ [2][2]*big.Int, c_ [2]*big.Int) (*types.Transaction, error) {
	return _State.Contract.TransitState(&_State.TransactOpts, id_, oldState_, newState_, isOldStateGenesis_, a_, b_, c_)
}

// TransitState is a paid mutator transaction binding the contract method 0x28f88a65.
//
// Solidity: function transitState(uint256 id_, uint256 oldState_, uint256 newState_, bool isOldStateGenesis_, uint256[2] a_, uint256[2][2] b_, uint256[2] c_) returns()
func (_State *StateTransactorSession) TransitState(id_ *big.Int, oldState_ *big.Int, newState_ *big.Int, isOldStateGenesis_ bool, a_ [2]*big.Int, b_ [2][2]*big.Int, c_ [2]*big.Int) (*types.Transaction, error) {
	return _State.Contract.TransitState(&_State.TransactOpts, id_, oldState_, newState_, isOldStateGenesis_, a_, b_, c_)
}

// StateInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the State contract.
type StateInitializedIterator struct {
	Event *StateInitialized // Event containing the contract specifics and raw log

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
func (it *StateInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateInitialized)
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
		it.Event = new(StateInitialized)
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
func (it *StateInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateInitialized represents a Initialized event raised by the State contract.
type StateInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_State *StateFilterer) FilterInitialized(opts *bind.FilterOpts) (*StateInitializedIterator, error) {

	logs, sub, err := _State.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StateInitializedIterator{contract: _State.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_State *StateFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StateInitialized) (event.Subscription, error) {

	logs, sub, err := _State.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateInitialized)
				if err := _State.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_State *StateFilterer) ParseInitialized(log types.Log) (*StateInitialized, error) {
	event := new(StateInitialized)
	if err := _State.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the State contract.
type StateOwnershipTransferStartedIterator struct {
	Event *StateOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *StateOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateOwnershipTransferStarted)
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
		it.Event = new(StateOwnershipTransferStarted)
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
func (it *StateOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the State contract.
type StateOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_State *StateFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StateOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _State.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StateOwnershipTransferStartedIterator{contract: _State.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_State *StateFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *StateOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _State.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateOwnershipTransferStarted)
				if err := _State.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_State *StateFilterer) ParseOwnershipTransferStarted(log types.Log) (*StateOwnershipTransferStarted, error) {
	event := new(StateOwnershipTransferStarted)
	if err := _State.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the State contract.
type StateOwnershipTransferredIterator struct {
	Event *StateOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StateOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateOwnershipTransferred)
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
		it.Event = new(StateOwnershipTransferred)
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
func (it *StateOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateOwnershipTransferred represents a OwnershipTransferred event raised by the State contract.
type StateOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_State *StateFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StateOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _State.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StateOwnershipTransferredIterator{contract: _State.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_State *StateFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StateOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _State.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateOwnershipTransferred)
				if err := _State.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_State *StateFilterer) ParseOwnershipTransferred(log types.Log) (*StateOwnershipTransferred, error) {
	event := new(StateOwnershipTransferred)
	if err := _State.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateStateTransitedIterator is returned from FilterStateTransited and is used to iterate over the raw logs and unpacked data for StateTransited events raised by the State contract.
type StateStateTransitedIterator struct {
	Event *StateStateTransited // Event containing the contract specifics and raw log

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
func (it *StateStateTransitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateStateTransited)
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
		it.Event = new(StateStateTransited)
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
func (it *StateStateTransitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateStateTransitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateStateTransited represents a StateTransited event raised by the State contract.
type StateStateTransited struct {
	GistRoot    *big.Int
	Id          *big.Int
	State       *big.Int
	Timestamp   *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStateTransited is a free log retrieval operation binding the contract event 0x819388a8ff4f5591c6f06c7c9e4cc1d5d7faf9e97e7faede4d1db8932b48be63.
//
// Solidity: event StateTransited(uint256 gistRoot, uint256 indexed id, uint256 state, uint256 timestamp, uint256 blockNumber)
func (_State *StateFilterer) FilterStateTransited(opts *bind.FilterOpts, id []*big.Int) (*StateStateTransitedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _State.contract.FilterLogs(opts, "StateTransited", idRule)
	if err != nil {
		return nil, err
	}
	return &StateStateTransitedIterator{contract: _State.contract, event: "StateTransited", logs: logs, sub: sub}, nil
}

// WatchStateTransited is a free log subscription operation binding the contract event 0x819388a8ff4f5591c6f06c7c9e4cc1d5d7faf9e97e7faede4d1db8932b48be63.
//
// Solidity: event StateTransited(uint256 gistRoot, uint256 indexed id, uint256 state, uint256 timestamp, uint256 blockNumber)
func (_State *StateFilterer) WatchStateTransited(opts *bind.WatchOpts, sink chan<- *StateStateTransited, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _State.contract.WatchLogs(opts, "StateTransited", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateStateTransited)
				if err := _State.contract.UnpackLog(event, "StateTransited", log); err != nil {
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

// ParseStateTransited is a log parse operation binding the contract event 0x819388a8ff4f5591c6f06c7c9e4cc1d5d7faf9e97e7faede4d1db8932b48be63.
//
// Solidity: event StateTransited(uint256 gistRoot, uint256 indexed id, uint256 state, uint256 timestamp, uint256 blockNumber)
func (_State *StateFilterer) ParseStateTransited(log types.Log) (*StateStateTransited, error) {
	event := new(StateStateTransited)
	if err := _State.contract.UnpackLog(event, "StateTransited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
