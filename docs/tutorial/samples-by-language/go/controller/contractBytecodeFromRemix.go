// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package controller

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
)

// BaseRelayRecipientMetaData contains all meta data concerning the BaseRelayRecipient contract.
var BaseRelayRecipientMetaData = &bind.MetaData{
	ABI: "[]",
}

// BaseRelayRecipientABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseRelayRecipientMetaData.ABI instead.
var BaseRelayRecipientABI = BaseRelayRecipientMetaData.ABI

// BaseRelayRecipient is an auto generated Go binding around an Ethereum contract.
type BaseRelayRecipient struct {
	BaseRelayRecipientCaller     // Read-only binding to the contract
	BaseRelayRecipientTransactor // Write-only binding to the contract
	BaseRelayRecipientFilterer   // Log filterer for contract events
}

// BaseRelayRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseRelayRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseRelayRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseRelayRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseRelayRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseRelayRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseRelayRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseRelayRecipientSession struct {
	Contract     *BaseRelayRecipient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BaseRelayRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseRelayRecipientCallerSession struct {
	Contract *BaseRelayRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BaseRelayRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseRelayRecipientTransactorSession struct {
	Contract     *BaseRelayRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BaseRelayRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseRelayRecipientRaw struct {
	Contract *BaseRelayRecipient // Generic contract binding to access the raw methods on
}

// BaseRelayRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseRelayRecipientCallerRaw struct {
	Contract *BaseRelayRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// BaseRelayRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseRelayRecipientTransactorRaw struct {
	Contract *BaseRelayRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseRelayRecipient creates a new instance of BaseRelayRecipient, bound to a specific deployed contract.
func NewBaseRelayRecipient(address common.Address, backend bind.ContractBackend) (*BaseRelayRecipient, error) {
	contract, err := bindBaseRelayRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseRelayRecipient{BaseRelayRecipientCaller: BaseRelayRecipientCaller{contract: contract}, BaseRelayRecipientTransactor: BaseRelayRecipientTransactor{contract: contract}, BaseRelayRecipientFilterer: BaseRelayRecipientFilterer{contract: contract}}, nil
}

// NewBaseRelayRecipientCaller creates a new read-only instance of BaseRelayRecipient, bound to a specific deployed contract.
func NewBaseRelayRecipientCaller(address common.Address, caller bind.ContractCaller) (*BaseRelayRecipientCaller, error) {
	contract, err := bindBaseRelayRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseRelayRecipientCaller{contract: contract}, nil
}

// NewBaseRelayRecipientTransactor creates a new write-only instance of BaseRelayRecipient, bound to a specific deployed contract.
func NewBaseRelayRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseRelayRecipientTransactor, error) {
	contract, err := bindBaseRelayRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseRelayRecipientTransactor{contract: contract}, nil
}

// NewBaseRelayRecipientFilterer creates a new log filterer instance of BaseRelayRecipient, bound to a specific deployed contract.
func NewBaseRelayRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseRelayRecipientFilterer, error) {
	contract, err := bindBaseRelayRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseRelayRecipientFilterer{contract: contract}, nil
}

// bindBaseRelayRecipient binds a generic wrapper to an already deployed contract.
func bindBaseRelayRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseRelayRecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseRelayRecipient *BaseRelayRecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseRelayRecipient.Contract.BaseRelayRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseRelayRecipient *BaseRelayRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseRelayRecipient.Contract.BaseRelayRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseRelayRecipient *BaseRelayRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseRelayRecipient.Contract.BaseRelayRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseRelayRecipient *BaseRelayRecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseRelayRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseRelayRecipient *BaseRelayRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseRelayRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseRelayRecipient *BaseRelayRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseRelayRecipient.Contract.contract.Transact(opts, method, params...)
}

// ContractBytecodeFromRemixMetaData contains all meta data concerning the ContractBytecodeFromRemix contract.
var ContractBytecodeFromRemixMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"checkChange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"results\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"packsPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"votesPath\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"answers\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"canVote\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"packsPath\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c9669182": "checkChange()",
		"1b0c27da": "results(uint256)",
		"fc36e15b": "vote(string)",
		"a3ec138d": "voters(address)",
	},
	Bin: "0x6080604052600080546001600160a01b031916733b62e51e37d090453600395ff1f9bdf4d73984041790556003805460ff1916600117905534801561004357600080fd5b50604051610a7a380380610a7a83398101604081905261006291610168565b6001600160a01b0381166000818152600160208181526040808420805461ffff19168417815581518083019092526008825267495046534c696e6b60c01b8284019081529590945290829052516100be939290910191906100cf565b50506003805460ff191690556101d2565b8280546100db90610198565b90600052602060002090601f0160209004810192826100fd5760008555610143565b82601f1061011657805160ff1916838001178555610143565b82800160010185558215610143579182015b82811115610143578251825591602001919060010190610128565b5061014f929150610153565b5090565b5b8082111561014f5760008155600101610154565b60006020828403121561017a57600080fd5b81516001600160a01b038116811461019157600080fd5b9392505050565b600181811c908216806101ac57607f821691505b6020821081036101cc57634e487b7160e01b600052602260045260246000fd5b50919050565b610899806101e16000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80631b0c27da14610051578063a3ec138d1461007b578063c96691821461009d578063fc36e15b146100ba575b600080fd5b61006461005f366004610627565b6100cd565b60405161007292919061069c565b60405180910390f35b61008e6100893660046106e2565b610211565b60405161007293929190610706565b6003546100aa9060ff1681565b6040519015158152602001610072565b6100aa6100c836600461073f565b6102c4565b600281815481106100dd57600080fd5b9060005260206000209060020201600091509050806000018054610100906107f0565b80601f016020809104026020016040519081016040528092919081815260200182805461012c906107f0565b80156101795780601f1061014e57610100808354040283529160200191610179565b820191906000526020600020905b81548152906001019060200180831161015c57829003601f168201915b50505050509080600101805461018e906107f0565b80601f01602080910402602001604051908101604052809291908181526020018280546101ba906107f0565b80156102075780601f106101dc57610100808354040283529160200191610207565b820191906000526020600020905b8154815290600101906020018083116101ea57829003601f168201915b5050505050905082565b60016020819052600091825260409091208054918101805460ff80851694610100900416929190610241906107f0565b80601f016020809104026020016040519081016040528092919081815260200182805461026d906107f0565b80156102ba5780601f1061028f576101008083540402835291602001916102ba565b820191906000526020600020905b81548152906001019060200180831161029d57829003601f168201915b5050505050905083565b6003805460ff191660011790556000806102dc6104e8565b6001600160a01b03811660009081526001602081815260408084208151606081018352815460ff808216151583526101009091041615159381019390935292830180549596509394919390840191610333906107f0565b80601f016020809104026020016040519081016040528092919081815260200182805461035f906107f0565b80156103ac5780601f10610381576101008083540402835291602001916103ac565b820191906000526020600020905b81548152906001019060200180831161038f57829003601f168201915b505050505081525050905080600001516103fd5760405162461bcd60e51b815260206004820152600d60248201526c3737afb832b936b4b9b9b4b7b760991b60448201526064015b60405180910390fd5b80602001511561043f5760405162461bcd60e51b815260206004820152600d60248201526c185b1c9958591e57dd9bdd1959609a1b60448201526064016103f4565b6001600160a01b0382166000908152600160208181526040808420805461ff0019166101001790558051808201825290850151815280820188905260028054938401815593849052805180519194939093027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01926104c292849291019061058e565b5060208281015180516104db926001850192019061058e565b5060019695505050505050565b6000805460408051600481526024810182526020810180516001600160e01b0316637a6ce2e160e01b17905290516060926001600160a01b03169161052c9161082a565b6000604051808303816000865af19150503d8060008114610569576040519150601f19603f3d011682016040523d82523d6000602084013e61056e565b606091505b508051909250610588915082016020908101908301610846565b91505090565b82805461059a906107f0565b90600052602060002090601f0160209004810192826105bc5760008555610602565b82601f106105d557805160ff1916838001178555610602565b82800160010185558215610602579182015b828111156106025782518255916020019190600101906105e7565b5061060e929150610612565b5090565b5b8082111561060e5760008155600101610613565b60006020828403121561063957600080fd5b5035919050565b60005b8381101561065b578181015183820152602001610643565b8381111561066a576000848401525b50505050565b60008151808452610688816020860160208601610640565b601f01601f19169290920160200192915050565b6040815260006106af6040830185610670565b82810360208401526106c18185610670565b95945050505050565b6001600160a01b03811681146106df57600080fd5b50565b6000602082840312156106f457600080fd5b81356106ff816106ca565b9392505050565b831515815282151560208201526060604082015260006106c16060830184610670565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561075157600080fd5b813567ffffffffffffffff8082111561076957600080fd5b818401915084601f83011261077d57600080fd5b81358181111561078f5761078f610729565b604051601f8201601f19908116603f011681019083821181831017156107b7576107b7610729565b816040528281528760208487010111156107d057600080fd5b826020860160208301376000928101602001929092525095945050505050565b600181811c9082168061080457607f821691505b60208210810361082457634e487b7160e01b600052602260045260246000fd5b50919050565b6000825161083c818460208701610640565b9190910192915050565b60006020828403121561085857600080fd5b81516106ff816106ca56fea2646970667358221220b60a68399f2b14179f6f0ea3700935b69faf2e64eea9606bf6825fa2100bcbae64736f6c634300080d0033",
}

// ContractBytecodeFromRemixABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractBytecodeFromRemixMetaData.ABI instead.
var ContractBytecodeFromRemixABI = ContractBytecodeFromRemixMetaData.ABI

// Deprecated: Use ContractBytecodeFromRemixMetaData.Sigs instead.
// ContractBytecodeFromRemixFuncSigs maps the 4-byte function signature to its string representation.
var ContractBytecodeFromRemixFuncSigs = ContractBytecodeFromRemixMetaData.Sigs

// ContractBytecodeFromRemixBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractBytecodeFromRemixMetaData.Bin instead.
var ContractBytecodeFromRemixBin = ContractBytecodeFromRemixMetaData.Bin

// DeployContractBytecodeFromRemix deploys a new Ethereum contract, binding an instance of ContractBytecodeFromRemix to it.
func DeployContractBytecodeFromRemix(auth *bind.TransactOpts, backend bind.ContractBackend, _voter common.Address) (common.Address, *types.Transaction, *ContractBytecodeFromRemix, error) {
	parsed, err := ContractBytecodeFromRemixMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBytecodeFromRemixBin), backend, _voter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractBytecodeFromRemix{ContractBytecodeFromRemixCaller: ContractBytecodeFromRemixCaller{contract: contract}, ContractBytecodeFromRemixTransactor: ContractBytecodeFromRemixTransactor{contract: contract}, ContractBytecodeFromRemixFilterer: ContractBytecodeFromRemixFilterer{contract: contract}}, nil
}

// ContractBytecodeFromRemix is an auto generated Go binding around an Ethereum contract.
type ContractBytecodeFromRemix struct {
	ContractBytecodeFromRemixCaller     // Read-only binding to the contract
	ContractBytecodeFromRemixTransactor // Write-only binding to the contract
	ContractBytecodeFromRemixFilterer   // Log filterer for contract events
}

// ContractBytecodeFromRemixCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractBytecodeFromRemixCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBytecodeFromRemixTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractBytecodeFromRemixTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBytecodeFromRemixFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractBytecodeFromRemixFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBytecodeFromRemixSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractBytecodeFromRemixSession struct {
	Contract     *ContractBytecodeFromRemix // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractBytecodeFromRemixCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractBytecodeFromRemixCallerSession struct {
	Contract *ContractBytecodeFromRemixCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractBytecodeFromRemixTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractBytecodeFromRemixTransactorSession struct {
	Contract     *ContractBytecodeFromRemixTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractBytecodeFromRemixRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractBytecodeFromRemixRaw struct {
	Contract *ContractBytecodeFromRemix // Generic contract binding to access the raw methods on
}

// ContractBytecodeFromRemixCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractBytecodeFromRemixCallerRaw struct {
	Contract *ContractBytecodeFromRemixCaller // Generic read-only contract binding to access the raw methods on
}

// ContractBytecodeFromRemixTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractBytecodeFromRemixTransactorRaw struct {
	Contract *ContractBytecodeFromRemixTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractBytecodeFromRemix creates a new instance of ContractBytecodeFromRemix, bound to a specific deployed contract.
func NewContractBytecodeFromRemix(address common.Address, backend bind.ContractBackend) (*ContractBytecodeFromRemix, error) {
	contract, err := bindContractBytecodeFromRemix(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractBytecodeFromRemix{ContractBytecodeFromRemixCaller: ContractBytecodeFromRemixCaller{contract: contract}, ContractBytecodeFromRemixTransactor: ContractBytecodeFromRemixTransactor{contract: contract}, ContractBytecodeFromRemixFilterer: ContractBytecodeFromRemixFilterer{contract: contract}}, nil
}

// NewContractBytecodeFromRemixCaller creates a new read-only instance of ContractBytecodeFromRemix, bound to a specific deployed contract.
func NewContractBytecodeFromRemixCaller(address common.Address, caller bind.ContractCaller) (*ContractBytecodeFromRemixCaller, error) {
	contract, err := bindContractBytecodeFromRemix(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBytecodeFromRemixCaller{contract: contract}, nil
}

// NewContractBytecodeFromRemixTransactor creates a new write-only instance of ContractBytecodeFromRemix, bound to a specific deployed contract.
func NewContractBytecodeFromRemixTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractBytecodeFromRemixTransactor, error) {
	contract, err := bindContractBytecodeFromRemix(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBytecodeFromRemixTransactor{contract: contract}, nil
}

// NewContractBytecodeFromRemixFilterer creates a new log filterer instance of ContractBytecodeFromRemix, bound to a specific deployed contract.
func NewContractBytecodeFromRemixFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractBytecodeFromRemixFilterer, error) {
	contract, err := bindContractBytecodeFromRemix(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractBytecodeFromRemixFilterer{contract: contract}, nil
}

// bindContractBytecodeFromRemix binds a generic wrapper to an already deployed contract.
func bindContractBytecodeFromRemix(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractBytecodeFromRemixABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBytecodeFromRemix.Contract.ContractBytecodeFromRemixCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBytecodeFromRemix.Contract.ContractBytecodeFromRemixTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBytecodeFromRemix.Contract.ContractBytecodeFromRemixTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBytecodeFromRemix.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBytecodeFromRemix.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBytecodeFromRemix.Contract.contract.Transact(opts, method, params...)
}

// CheckChange is a free data retrieval call binding the contract method 0xc9669182.
//
// Solidity: function checkChange() view returns(bool)
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixCaller) CheckChange(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractBytecodeFromRemix.contract.Call(opts, &out, "checkChange")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckChange is a free data retrieval call binding the contract method 0xc9669182.
//
// Solidity: function checkChange() view returns(bool)
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixSession) CheckChange() (bool, error) {
	return _ContractBytecodeFromRemix.Contract.CheckChange(&_ContractBytecodeFromRemix.CallOpts)
}

// CheckChange is a free data retrieval call binding the contract method 0xc9669182.
//
// Solidity: function checkChange() view returns(bool)
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixCallerSession) CheckChange() (bool, error) {
	return _ContractBytecodeFromRemix.Contract.CheckChange(&_ContractBytecodeFromRemix.CallOpts)
}

// Results is a free data retrieval call binding the contract method 0x1b0c27da.
//
// Solidity: function results(uint256 ) view returns(string packsPath, string votesPath)
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixCaller) Results(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PacksPath string
	VotesPath string
}, error) {
	var out []interface{}
	err := _ContractBytecodeFromRemix.contract.Call(opts, &out, "results", arg0)

	outstruct := new(struct {
		PacksPath string
		VotesPath string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PacksPath = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.VotesPath = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Results is a free data retrieval call binding the contract method 0x1b0c27da.
//
// Solidity: function results(uint256 ) view returns(string packsPath, string votesPath)
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixSession) Results(arg0 *big.Int) (struct {
	PacksPath string
	VotesPath string
}, error) {
	return _ContractBytecodeFromRemix.Contract.Results(&_ContractBytecodeFromRemix.CallOpts, arg0)
}

// Results is a free data retrieval call binding the contract method 0x1b0c27da.
//
// Solidity: function results(uint256 ) view returns(string packsPath, string votesPath)
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixCallerSession) Results(arg0 *big.Int) (struct {
	PacksPath string
	VotesPath string
}, error) {
	return _ContractBytecodeFromRemix.Contract.Results(&_ContractBytecodeFromRemix.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool canVote, bool hasVoted, string packsPath)
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	CanVote   bool
	HasVoted  bool
	PacksPath string
}, error) {
	var out []interface{}
	err := _ContractBytecodeFromRemix.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		CanVote   bool
		HasVoted  bool
		PacksPath string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CanVote = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.HasVoted = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.PacksPath = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool canVote, bool hasVoted, string packsPath)
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixSession) Voters(arg0 common.Address) (struct {
	CanVote   bool
	HasVoted  bool
	PacksPath string
}, error) {
	return _ContractBytecodeFromRemix.Contract.Voters(&_ContractBytecodeFromRemix.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool canVote, bool hasVoted, string packsPath)
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixCallerSession) Voters(arg0 common.Address) (struct {
	CanVote   bool
	HasVoted  bool
	PacksPath string
}, error) {
	return _ContractBytecodeFromRemix.Contract.Voters(&_ContractBytecodeFromRemix.CallOpts, arg0)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string answers) returns(bool)
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixTransactor) Vote(opts *bind.TransactOpts, answers string) (*types.Transaction, error) {
	return _ContractBytecodeFromRemix.contract.Transact(opts, "vote", answers)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string answers) returns(bool)
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixSession) Vote(answers string) (*types.Transaction, error) {
	return _ContractBytecodeFromRemix.Contract.Vote(&_ContractBytecodeFromRemix.TransactOpts, answers)
}

// Vote is a paid mutator transaction binding the contract method 0xfc36e15b.
//
// Solidity: function vote(string answers) returns(bool)
func (_ContractBytecodeFromRemix *ContractBytecodeFromRemixTransactorSession) Vote(answers string) (*types.Transaction, error) {
	return _ContractBytecodeFromRemix.Contract.Vote(&_ContractBytecodeFromRemix.TransactOpts, answers)
}
