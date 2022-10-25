package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/LACNetNetworks/gas-relay-signer/audit"
	"github.com/LACNetNetworks/gas-relay-signer/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const DATA_CALL_RELAYHUB = "0x"

func isPoolEmpty(rpcURL string, id json.RawMessage) (bool, error) {
	data := fmt.Sprintf(`{"jsonrpc":"2.0","method":"txpool_besuTransactions",
	"params":[], "id":"%s"}`, id)

	requestBody := []byte(data)

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", rpcURL, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")

	if err != nil {
		return false, err
	}

	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return false, err
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(body))

	var rpcMessage rpc.JsonrpcMessage

	err = json.NewDecoder(rdr1).Decode(&rpcMessage)
	if err != nil {
		return false, err
	}

	audit.GeneralLogger.Println("Transactions in pool:", rpcMessage.String())

	var v []json.RawMessage
	err = json.Unmarshal(rpcMessage.Result, &v)
	if err != nil {
		return false, err
	}

	if len(v) > 0 {
		return false, nil
	}

	return true, nil
}

func getRelayHubContractAddress(rpcURL string, id string, relayHubProxyAddress string, _timeout int) (*common.Address, error) {
	data := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_call","params":[{"to":"%s","data":"%s"},"latest"], "id":"%s"}`, relayHubProxyAddress, DATA_CALL_RELAYHUB, id)

	fmt.Println(data)

	requestBody := []byte(data)

	timeout := time.Duration(time.Duration(_timeout) * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", rpcURL, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(body))

	var rpcMessage rpc.JsonrpcMessage

	err = json.NewDecoder(rdr1).Decode(&rpcMessage)
	if err != nil {
		return nil, err
	}

	var resultData string = string(rpcMessage.Result)

	responseData := common.Hex2Bytes(resultData[3 : len(resultData)-1])

	addressPacked, err := abi.NewType("address", "", nil)
	if err != nil {
		fmt.Println(err)
	}

	resultPayloadPacked := abi.Arguments{
		{Type: addressPacked},
	}

	addressUnpacked, err := resultPayloadPacked.UnpackValues(responseData)
	if err != nil {
		fmt.Println(err)
	}

	relayHubAddress := addressUnpacked[0].(common.Address)

	return &relayHubAddress, nil
}
