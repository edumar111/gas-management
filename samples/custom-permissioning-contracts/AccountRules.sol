pragma solidity ^0.8.0;

import "./AccountRulesProxy.sol";
import "./AccountRulesList.sol";
import "./access/AccessControl.sol";

contract AccountRules is AccountRulesProxy, AccountRulesList, AccessControl {

    // in read-only mode rules can't be added/removed
    // this will be used to protect data when upgrading contracts
    bool private readOnlyMode = false;
    // version of this contract: semver like 1.2.14 represented like 001002014
    uint private version = 1000000;

    AccountIngress private ingressContract;

    modifier onlyOnEditMode() {
        require(!readOnlyMode, "In read only mode: rules cannot be modified");
        _;
    }

    modifier onlyAdmin(){
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "Caller is not Admin");
        _;
    }

    constructor (AccountIngress _ingressContract) public {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        add(msg.sender);
    }

    // VERSION
    function getContractVersion() public view returns (uint) {
        return version;
    }

    // READ ONLY MODE
    function isReadOnly() public view returns (bool) {
        return readOnlyMode;
    }

    function enterReadOnly() public onlyAdmin returns (bool) {
        require(readOnlyMode == false, "Already in read only mode");
        readOnlyMode = true;
        return true;
    }

    function exitReadOnly() public onlyAdmin returns (bool) {
        require(readOnlyMode == true, "Not in read only mode");
        readOnlyMode = false;
        return true;
    }

    function transactionAllowed(
        address sender,
        address, // target
        uint256, // value
        uint256, // gasPrice
        uint256, // gasLimit
        bytes memory // payload
    ) public view returns (bool) {
        if (
            accountPermitted (sender) && (target == "0x1234567890123456789012345678901234567890")
        ) {
            return true;
        } else {
            return false;
        }
    }

    function accountPermitted(
        address _account
    ) public view returns (bool) {
        return exists(_account);
    }

    function addAccount(
        address account
    ) public onlyAdmin onlyOnEditMode returns (bool) {
        bool added = add(account);
        emit AccountAdded(added, account);
        return added;
    }

    function removeAccount(
        address account
    ) public onlyAdmin onlyOnEditMode returns (bool) {
        bool removed = remove(account);
        emit AccountRemoved(removed, account);
        return removed;
    }

    function getSize() public view returns (uint) {
        return size();
    }

    function getByIndex(uint index) public view returns (address account) {
        return allowlist[index];
    }

    function getAccounts() public view returns (address[] memory){
        return allowlist;
    }

    function addAccounts(address[] memory accounts) public onlyAdmin returns (bool) {
        return addAll(accounts);
    }
}
