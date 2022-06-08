# LACChain in Go


**Disclaimer:** The LACChain platform and ecosystem is constantly evolving so I apologize if any link in the following documentation is missing, has been archived or has been deprecated. If you have any question please submit a issue in this Github so we can review it together.

To make your Go application interact with the LACChain network using the Gas Model proposed [here](https://github.com/lacchain/gas-management/blob/master/docs/OVERVIEW.md), you would need to modify some go-ethereum package code in order to correctly run your app.

## Previous requirements

* Install the required go modules using: ```go install``` in the go directory

* Add the following function in your **go-ethereum** library code. It needs to be in the ~/go/pkg/mod/github.com/ethereum/go-ethereum@v1.10.17/ethclient/ethclient.go (preferably in line 526 after the *SendTransaction* function declaration):
```
func (ec *Client) SendTransactionRelay(ctx context.Context, tx *types.Transaction) (interface{}, error) {
	data, err := tx.MarshalBinary()
	if err != nil {
		return nil, err
	}
	var bytes interface{}
	return bytes, ec.c.CallContext(ctx, &bytes, "eth_sendRawTransaction", hexutil.Encode(data))
}
```




## Code

Go projects can take advantage of go-ethereum binding feature to get ready-to-use functions that would interact with their smart contract. To compile and bind a specific smart contract to your Go application use: ```abigen -sol contracts/ContractBytecodeFromRemix.sol -pkg controller --out controller/contractBytecodeFromRemix.go``` (it should be run in the root directory of this repository). This will create a .go file in the controller directory containing a **controller** package with all functions related to your specific smart contract.

The main.go file is divided in:

1. **Configuration variables you need:** like the private key (you get the public key from it) and the LACChain required values (authorized node address and expiration time).

1. **Deploying contract code:** to deploy a smart contract written in Solidity and already binded using the abigen tool that go-ethereum provides. The code specifies the transaction information (value, gas limit and gas price) and also the constructor arguments. It compiles the arguments separately of the LACChain values so no dynamic variables (such as string or bytes) mess up with the input arguments. Both arguments are then concatenate and signed using the private key. The signed raw transaction is send using a special function called **SendTransactionRelay** that differs from **SendTransaction** in that the former one receives the information sent by the LACChain client. This step is performed because the transaction hash coming from the client is the real one. To get the receipt (and therefore the contract address), we need to use the transaction hash in the client response.

5. **Sending transaction code:**

6. **Check new state of smart contract:**