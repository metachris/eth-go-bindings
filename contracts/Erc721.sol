// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";

contract Erc721 is ERC721Enumerable {
    constructor() ERC721("MyNFT", "MNFT") {}
}
