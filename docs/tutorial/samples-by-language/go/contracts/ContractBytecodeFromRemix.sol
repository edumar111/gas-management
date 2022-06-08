// / SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.0 <0.9.0;
import "./BaseRelayRecipient.sol";

contract ContractBytecodeFromRemix is BaseRelayRecipient {
    struct Voter {
        bool canVote;
        bool hasVoted;
        string packsPath;
    }
    mapping( address => Voter ) public voters;
    struct Result {
        string packsPath;
        string votesPath;
    }
    Result[] public results;
    bool public checkChange = true;

    function vote(string memory answers) public returns (bool)  {
        checkChange = true;
        address voterAddress = _msgSender();
        Voter memory person = voters[voterAddress];

        require(person.canVote, "no_permission");
        require(!person.hasVoted, "already_voted");

        voters[voterAddress].hasVoted = true;
        results.push( Result(person.packsPath, answers) );
        return true;
    }



    constructor( address _voter ){
        voters[_voter].canVote = true;
        voters[_voter].hasVoted = false;
        voters[_voter].packsPath = "IPFSLink";
        checkChange = false;
    }
}