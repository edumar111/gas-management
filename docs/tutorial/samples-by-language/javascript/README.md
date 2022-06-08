# LACChain in Javascript

**Disclaimer:** The LACChain platform and ecosystem is constantly evolving so I apologize if any link in the following documentation is missing, has been archived or has been deprecated. If you have any question please submit a issue in this Github so we can review it together.



To make your Javascript application interact with the LACChain network using the Gas Model proposed [here](https://github.com/lacchain/gas-management/blob/master/docs/OVERVIEW.md), you can use the ```DeploSendCheck.sj``` code as a use case and tutorial. The official tutorial can be found [here](https://github.com/LACNetNetworks/gas-management/blob/master/docs/How_adapt_your_Dapp.md) but I think that this tutorial might help you as well.

## Previous requirements

* Install the required node modules using: ```npm install``` in the javascript directory

## Code

The file is divided in:

1. **Initial libraries you need:** ```web3```, ```ethereumjs-tx``` (keep in mind that the compatible version is 1.3.7. The [repository](https://github.com/ethereumjs/ethereumjs-tx) has been deprecated but it still works. Please report if you find a vulnerability)

2. **Configuration variables you need:** the private key in string format (you get the public key from it) and the LACChain required values: *authorized node address* and *expiration time*

3. **Deploying contract code:** to launch a smart contract written in Solidity and already compiled (remember to put the ```.json``` file in the ```metadata``` variable that Remix or any other compiler returns you after compiling the Solidity source code). You can use the ```contract.deploy``` function to get the hexadecimal bytecode to be used.

4. **Sending transaction code:** to explicitly change the recently deployed smart contract's state and assure that it is working as expected. You need to encode the function, argument types and arguments themselves as bytecode represented in hexadecimal.

5. **Check new state of smart contract:** to reassure that the previous transaction was executed correctly. In the example, we check the public variable results but you can implement any changing state feature in the contract just to confirm that the transaction works and executes correctly in the [Ethereum Virtual Machine](https://ethereum.org/en/developers/docs/evm/).


## Remember that:

* The LACChain values need to be added at the end of the bytecode right before it is signed.

* When you have a undefined length in some arguments (string, bytes, etc), you have to consider that the encoding would put a fixed-size pointer where the variable goes and append the whole variable value at the end of the bytecode. This means that the encoding and addition of the LACChain nodes have to be performed after encoding the function itself (checkout that I encode the LAC values separate of the function encoding).

* The LACChain ecosystem is constantly evolving and so far there has been a lot of tutorials. I recommend you to look into:
    * The oficial repositories [here](https://github.com/LACNetNetworks)
    * The past repositories [here](https://github.com/lacchain/)
    * This amazing online course [here](https://aula.blockchainacademy.cl/p/introduccion-a-lacchain)

