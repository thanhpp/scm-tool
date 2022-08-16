//Contract based on [https://docs.openzeppelin.com/contracts/3.x/erc721](https://docs.openzeppelin.com/contracts/3.x/erc721)
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../openzeppelin-contracts/contracts/token/ERC721/ERC721.sol";
import "../openzeppelin-contracts/contracts/utils/Counters.sol";
import "../openzeppelin-contracts/contracts/access/Ownable.sol";
import "../openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract ScmNFT is ERC721URIStorage {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    constructor() ERC721("ScmNFT", "SCM") {}

    function selfMintNFT(string memory tokenURI) public returns (uint256) {
        address recipient = msg.sender;

        _tokenIds.increment();

        uint256 newItemId = _tokenIds.current();
        _mint(recipient, newItemId);
        _setTokenURI(newItemId, tokenURI);

        return newItemId;
    }

    function multiSafeTransferFrom(
        address from,
        address to,
        uint256[] memory tokenIds
    ) public {
        for (uint256 i = 0; i < tokenIds.length; i++) {
            safeTransferFrom(from, to, tokenIds[i]);
        }
    }
}
