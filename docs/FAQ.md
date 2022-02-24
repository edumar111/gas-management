# Frequently Asked Questions

This document is an attempt to collect some of the questions frequently asked by developers.

## Table of Contents

1.	[Frequently Asked Questions](#frequently-asked-questions)
	1.	[Table of Contents](#table-of-contents)
	2.	[General Questions](#general-questions)
		1.	[How much Gas is my node assigned on the network?](#how-much-gas-is-my-node-assigned-on-the-network)
		2.	[How much Gas was used in last blocks?](#how-much-gas-was-used-in-last-blocks)
		3.	[How much Gas does my transaction consume?](#how-much-gas-does-my-transaction-consume)
		4.	[How much is the price of gas in LACNET?](#how-much-is-the-price-of-gas-in-lacnet)             
         5.  [How to implement custom account permissioning](#how-to-implement-custom-account-permissioning)

## General Questions

### How much GAS is my node assigned on the network?

To find out how much GAS your node has been assigned on the network. You can make a call to the getGasLimit method:
```
function getGasLimit() external view returns (uint256){
...
}
```
The following [API](https://besu.hyperledger.org/en/stable/Reference/API-Methods/#eth_call) can be used.
```
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_call","params":[{"from":"0x971bb94d235a4ba42d53ab6fb0a86b12c73ba460", "to":"0x7a4363E55Ef04e9144a2B187ACA804631A3155B5","data":"0x1a93d1c3"},"latest"],"id":53}' http://127.0.0.1:4545
```
At this example we want to know how much GAS assigned has this address 0x971bb94d235a4ba42d53ab6fb0a86b12c73ba460, in your case change this parameter for your node address.

The response will be:
```
{
  "jsonrpc" : "2.0",
  "id" : 53,
  "result" : "0x0000000000000000000000000000000000000000000000000000000005b01a0a"
```
Where result has the GAS assigned in hexadecimal. Transforming this value to decimal you get your GAS assigned.

### How much Gas was used in last blocks?

To find out how much GAS was used on the network in the last N blocks. You can make a call to the getGasUsedLastBlocks method:
```
function getGasUsedLastBlocks() external view returns (uint256){
...
}
```
The following [API](https://besu.hyperledger.org/en/stable/Reference/API-Methods/#eth_call) can be used.
```
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_call","params":[{"from":"0x971bb94d235a4ba42d53ab6fb0a86b12c73ba460", "to":"0x7a4363E55Ef04e9144a2B187ACA804631A3155B5","data":"0xd03ce2db"},"latest"],"id":53}' http://127.0.0.1:4545
```
The response will be:
```
{
  "jsonrpc" : "2.0",
  "id" : 53,
  "result" : "0x0000000000000000000000000000000000000000000000000000000000034ab"
```
Where result has the GAS used in hexadecimal. Transforming this value to decimal you get your GAS used in the last blocks.

### How much gas does my transaction consume?

To find out how much gas your transaction consumes, you need to get the transaction Id, which is displayed after submitting your transaction to the network, then search for your transaction using the following [API](https://besu.hyperledger.org/en/stable/Reference/API-Methods/#eth_gettransactionbyhash). As a response you will get a **"gas" parameter** with value in hexadecimal, transforming this value to decimal you can know how much GAS your transaction consumed.

### How much is the price of gas in LACNET?

Gas price is 0. You do not need to buy a token or pay a transaction fee to deploy contracts or send transactions.

### How to implement custom account permissioning

To implement a custom account permissiong in your node you can take this [example](https://github.com/LACNet-Networks/gas-management/blob/master/samples/custom-permissioning-contracts/AccountRules.sol). This example to allow only certain accounts to send transactions to a specific contract destination. You can add account using addAccount method. Is important to mentio that you can change whatever you want, but **keep the inheritance of AccountRulesProxy.sol** in your contract. You can put your own logic in **transactionAllowed** function