package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"lacchain-tools-go/controller"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Change initial parameters like RPC URL, Node Address, Expiration Time, Private key
	web3, err := ethclient.Dial("http://34.69.22.82:80")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("919b7e0e4095ce8a2cb22cea25a4d5888981d29d03cbdc714ed4b5f58313fdc6")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// transaction information
	value := big.NewInt(0)      // in wei (0 eth)
	gasLimit := uint64(4000000) // in units
	gasPrice := big.NewInt(0)   // in wei 0 gas price
	// arguments
	addressArg, err := abi.NewType("address", "", nil)
	if err != nil {
		log.Fatal(err)
	}
	uint256Arg, err := abi.NewType("uint256", "", nil)
	if err != nil {
		log.Fatal(err)
	}
	// LACChain parameters
	lacchainAddress := common.HexToAddress("0xd00e6624a73f88b39f82ab34e8bf2b4d226fd768")
	lacchainExpiration := new(big.Int).SetInt64(time.Now().Add(time.Minute * time.Duration(20)).Unix())
	argumentsLAC := abi.Arguments{
		{Type: addressArg},
		{Type: uint256Arg},
	}
	valueLAC, err := argumentsLAC.Pack(
		lacchainAddress,
		lacchainExpiration,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DEPLOYING CONTRACT  ======>")
	nonce, err := web3.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// contract
	contractBytecode := controller.ContractBytecodeFromRemixBin
	// inputs
	// Change the contract construction arguments
	voter := common.HexToAddress("0xbcEda2Ba9aF65c18C7992849C312d1Db77cF008E")
	argumentsContract := abi.Arguments{
		{Type: addressArg},
	}
	bytesContractArguments, err := argumentsContract.Pack(
		voter,
	)
	if err != nil {
		log.Fatal(err)
	}
	// send transaction
	deployContract := contractBytecode + hex.EncodeToString(bytesContractArguments) + hex.EncodeToString(valueLAC)
	txData, err := hexutil.Decode(deployContract)
	if err != nil {
		log.Fatal(err)
	}
	txObject := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     txData,
	})
	tx, err := types.SignTx(txObject, types.FrontierSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	result, err := web3.SendTransactionRelay(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	txHash := fmt.Sprintf("%v", result)
	time.Sleep(5 * time.Second)
	// Get the transaction receipt
	receipt, err := web3.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Fatal(err)
	}
	addressTo := receipt.ContractAddress

	fmt.Println("SENDING TRANSACTION ======>")
	nonce, err = web3.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	abiContract, err := abi.JSON(strings.NewReader(controller.ContractBytecodeFromRemixABI))
	if err != nil {
		log.Fatal(err)
	}
	bytesContractArguments, err = abiContract.Pack("vote", "This is working")
	if err != nil {
		log.Fatal(err)
	}
	txContract := "0x" + hex.EncodeToString(bytesContractArguments) + hex.EncodeToString(valueLAC)
	txData, err = hexutil.Decode(txContract)
	if err != nil {
		log.Fatal(err)
	}
	txObject = types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &addressTo,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     txData,
	})
	tx, err = types.SignTx(txObject, types.FrontierSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	result, err = web3.SendTransactionRelay(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	txHash = fmt.Sprintf("%v", result)
	time.Sleep(5 * time.Second)
	// Get the transaction receipt
	receipt, err = web3.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("CHECKING CHANGE     ======>")
	contract, err := controller.NewContractBytecodeFromRemix(addressTo, web3)
	if err != nil {
		log.Fatal(err)
	}
	// Use the contract object attributes for easy read actions
	checkResults, err := contract.Results(nil, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(checkResults)
}
