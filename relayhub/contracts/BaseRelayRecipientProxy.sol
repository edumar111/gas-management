// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.8.0 <0.9.0;

contract BaseRelayRecipientProxy {

    address relayHubAddress;
    address owner;

    constructor(address _newRelayHub) {
        owner = msg.sender;
        relayHubAddress = _newRelayHub;
    }

    modifier onlyOwner(){
        require(msg.sender == owner, "Only owner can execute this mehtod");
        _;
    }

    function getMsgSender() external returns (address){
        bytes memory bytesSender;
        (,bytesSender) = relayHubAddress.call(msg.data);

        return abi.decode(bytesSender, (address));
    }

    function setRelayHub(address _newRelayHub) external onlyOwner {
        relayHubAddress = _newRelayHub;
    }

    function getRelayHub() external view returns (address) {
        return relayHubAddress;    
    }
}