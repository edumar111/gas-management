(async () => {
    /* This is the code that will be used to deploy the contract and send a transaction to it. */
    const Web3 = require('web3')
    const ethTx = require('ethereumjs-tx')
    const metadata = require("./ContractBytecodeFromRemix.json")
    // Change initial parameters like RPC URL, Node Address, Expiration Time, Private key
    const web3 = new Web3('http://34.69.22.82:80')
    const account = web3.eth.accounts.privateKeyToAccount("0x919b7e0e4095ce8a2cb22cea25a4d5888981d29d03cbdc714ed4b5f58313fdc6");
    const privateKey = Buffer.from(account.privateKey.substr(2), 'hex')
    const valueLAC = web3.eth.abi.encodeParameters(["address", "uint256"], ["0xd00e6624a73f88b39f82ab34e8bf2b4d226fd768", Math.floor(Date.now()/1000 + 1200)]) // 20 minutes
    let txCount, txData, txObject, tx, receipt



    console.log("DEPLOYING CONTRACT  ======>")
    txCount = await web3.eth.getTransactionCount(account.address)
    // Change the contract construction arguments
    const constructArgs = [
        account.address
    ]
    let deployContract = new web3.eth.Contract(metadata.abi)
    deployContract = deployContract.deploy({
        data: metadata.data.bytecode.object,
        arguments: constructArgs
    })
    txData = deployContract.encodeABI() + valueLAC.substr(2)
    txObject = {
        nonce: web3.utils.toHex(txCount),
        gasPrice: web3.utils.toHex(0),
        gasLimit: web3.utils.toHex(4000000),
        data: web3.utils.toHex(txData)
    }
    tx = new ethTx(txObject)
    tx.sign(privateKey)
    receipt = await web3.eth.sendSignedTransaction('0x' + tx.serialize().toString('hex'))
    const addressTo = receipt.contractAddress



    console.log("SENDING TRANSACTION ======>")
    txCount = await web3.eth.getTransactionCount(account.address)
    // Change the function name, types and parameters
    const answers = "This is working"
    let value = web3.eth.abi.encodeParameters(
        ["string"],
        [answers]
    )
    txData = web3.eth.abi.encodeFunctionSignature('vote(string)') + value.substr(2) + valueLAC.substr(2)
    txObject = {
        nonce: web3.utils.toHex(txCount),
        gasPrice: web3.utils.toHex(0),
        gasLimit: web3.utils.toHex(4000000),
        to: addressTo,
        data: web3.utils.toHex(txData)
    }
    tx = new ethTx(txObject)
    tx.sign(privateKey)
    await web3.eth.sendSignedTransaction('0x' + tx.serialize().toString('hex'))



    console.log("CHECKING CHANGE     ======>")
    let contract = new web3.eth.Contract(metadata.abi, addressTo)
    // Change the checking function / attribute in the contract
    receipt = await contract.methods.results(0).call();
    console.log(receipt)

})();