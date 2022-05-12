# How to deploy your smart contract using gas model

This guide aims to give you an overview of some mainstream tools that you can use to deploy Smart Contracts.

## Truffle

[Truffle](https://www.trufflesuite.com/docs/truffle/overview "Truffle Overview") is basically a development environment where you could easily develop smart contracts with it’s built-in testing framework, smart contract compilation and deployment, interactive console, and many more features.

### Install

First, let's install truffle.

>`npm install -g truffle`

>`truffle version`

Now we can create our project folder, which we will name **MyDapp**.

>`mkdir MyDapp`

>`cd MyDapp`

For this tutorial, we will start from scratch, so we execute the following command in our MyApp directory:

>`truffle init`

This command creates a bare Truffle project. After doing so, you should have the following files and folders:

* contracts/: Directory for Solidity contracts
* migrations/: Directory for scriptable deployment
* test/: Directory for test files for testing your application and contracts
* truffle-config.js: Truffle configuration file

### Contract Compilation

Before anything else, let's create a very simple smart contract named **MyContract.sol** and store it in the contracts folder. All smart contracts you create should be stored there.

Our smart contract will contain code that's as simple as this:

```js
    // We will be using Solidity version 0.5.12 
    pragma solidity 0.5.12;

    contract MyContract {
        string private message = "My First Smart Contract";

        function getMessage() public view returns(string memory) {
            return message;
        }

        function setMessage(string memory newMessage) public {
            message = newMessage;
        }
    }
```
Basically, our smart contract has a variable named `message`, which contains a little message that is initilized as `My First Smart Contract`. Also we have two functions that can set or get that variable `message`

To compile the smart contract, execute the command:

>`truffle compile`

### Contract Deployment

#### Prerequisites

First, we need to install the [truffle hdwallet-provider](https://github.com/trufflesuite/truffle/tree/develop/packages/hdwallet-provider) according to [Using Hyperledger Besu with Truffle](https://besu.hyperledger.org/en/stable/HowTo/Develop-Dapps/Truffle "Truffle with Besu") to be able to deploy contracts and send transactiones with truffle:

>`npm install -g @truffle/hdwallet-provider@1.2.2`

Now, remove "1_initial_migrations.js" file and create a new file in the **migrations** directory. Create a new file named **1_deploy_contracts.js**, and write the following code:

```js
    var MyDapp = artifacts.require("MyContract");

    const nodeAddress = "0xd00e6624a73f88b39f82ab34e8bf2b4d226fd768"   //change this address by your node address
    const expiration = 1736394529

    module.exports = function(deployer){
        deployer.deploy(MyDapp, nodeAddress, expiration);
    };
```
Next, since the GAS model needs to obtain two parameters(nodeAddress, expiration), it is necessary to add these as parameters to the ABI of the contract. Please, add these two parameters as inputs at constructor function.

```json
"contractName": "MyContract",
  "abi": [
    {
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "_nodeAddress",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "_expiration",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
```

Next, we need to edit the Truffle configuration (**truffle-config.js**).

To briefly describe the parts that make up the configuration:

* networks: Will hold the configuration of our Ethereum client where we will be deploying our contracts
* compilers: Will hold the configuration of Solc compiler

Type your private key, network address IP node and RPC port in the networks part:

```js
    const HDWalletProvider = require("@truffle/hdwallet-provider");
    const privateKey = "<PRIVATE_KEY>";
    const privateKeyProvider = new HDWalletProvider(privateKey, "http://<ADDRESS_IP_NODE>:<PORT_RPC_NODE>");

    module.exports = {
        networks: {
            development: {
                host: "127.0.0.1",
                port: 7545,
                network_id: "*"
            },
            lacchain: {
                provider: privateKeyProvider,
                network_id: "648530",  //change this value by network id 
                gasPrice: 0
            }
        }
    };
```
***NOTE: This is just an example. NEVER hard code production private keys in your code or commit them to git. They should always be loaded from environment variables or a secure secret management system.***

Truffle migrations are scripts that help us deploy our smart contract to the LACCHAIN network. Let's deploy it:

>`truffle migrate -network lacchain`

Finally you get the deployment report where you can see the address contract similar to this:

```json
    Deploying 'MyDapp'
    --------------------
    transaction hash:0x31d91fa2524953e49cfc4c433ac939b56df8d9371fdde74c56a75634efcf823d
    Blocks: 0            Seconds: 0
    contract address:    0xFA3F403BeC6D3dd2eF9008cf8D21e3CA0FD1B9C4
    block number:        4006082
    block timestamp:     1574190784
    account:             0xbcEda2Ba9aF65c18C7992849C312d1Db77cF008E
    balance:             0
    gas used:            340697
    gas price:           0 gwei
    value sent:          0 ETH
    total cost:          0 ETH
```

## Ethers

To deploy using ethers library you are able check this [example](https://github.com/lacchain/lacchain-did-registry/blob/master/deploy.js)