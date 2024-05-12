// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract CertificateNFT is ERC721Enumerable, AccessControl, Ownable {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIdCounter;

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    struct Certificate {
        string name;
        string course;
        string date;
        string cerType;
        string imageUri;
    }

    mapping(uint256 => Certificate) private _certificates;

    constructor() ERC721("CertificateNFT", "CERNFT") Ownable(msg.sender) {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);
        _tokenIdCounter.increment(); // Ensure we start token IDs from 1
    }

    function mintCertificate(
        address recipient, 
        string memory name,
        string memory course,
        string memory date,
        string memory cerType,
        string memory imageUri
    )
        external
        onlyRole(MINTER_ROLE)
        returns (uint256)
    {
        uint256 newTokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();

        _safeMint(recipient, newTokenId); // Mint to the specified recipient's address
        _certificates[newTokenId] = Certificate(name, course, date, cerType, imageUri);
        return newTokenId;
    }


    function getCertificateDetails(uint256 tokenId)
        external
        view
        returns (
            string memory,
            string memory,
            string memory,
            string memory,
            string memory
        )
    {
        // Ensure the token exists
        try this.ownerOf(tokenId) returns (address) {
            Certificate memory cert = _certificates[tokenId];
            return (cert.name, cert.course, cert.date, cert.cerType, cert.imageUri);
        } catch {
            revert("Token ID does not exist");
        }
    }

    function addMinter(address minterAddress) external onlyOwner {
        grantRole(MINTER_ROLE, minterAddress);
    }

    function removeMinter(address minterAddress) external onlyOwner {
        revokeRole(MINTER_ROLE, minterAddress);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(ERC721Enumerable, AccessControl)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}