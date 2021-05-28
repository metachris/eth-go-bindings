# Go bindings for Ethereum smart contracts

Currently implemented:

* [ERC20: Token Standard](https://eips.ethereum.org/EIPS/eip-20)
* [ERC165: Standard Interface Detection](https://eips.ethereum.org/EIPS/eip-165), several [interface IDs](https://github.com/metachris/eth-go-bindings/blob/master/erc165/interfaceids.go)
* [ERC721: Non-Fungible Token Standard (NFT)](https://eips.ethereum.org/EIPS/eip-721) with [ERC721 Metadata](https://docs.openzeppelin.com/contracts/4.x/api/token/erc721#IERC721Metadata) and [ERC721 Enumerable](https://docs.openzeppelin.com/contracts/4.x/api/token/erc721#IERC721Enumerable) extensions.

Example usage: [examples/erc20-erc721.go](https://github.com/metachris/eth-go-bindings/blob/master/examples/erc20-erc721.go)

Notes:

* Inspired by [github.com/fxfactorial/defi-abigen](https://github.com/fxfactorial/defi-abigen) (which has bindings for Aave, Chainlink price feed, Compound, Erc20, Onesplit and Uniswap)
* Based on [OpenZeppelin contracts](https://docs.openzeppelin.com/contracts/4.x/)

Versions used to build the bindings:

* `go1.16.3`
* `solc: 0.8.4+commit.c7e474f2.Emscripten.clang`
* `abigen version 1.10.4-unstable`
* `go-ethereum: v1.10.3-56-g0703ef62d`

Next contract bindings to add:

* ERC621, ERC721x, ERC777, ERC884, ERC918, ERC1155
* Popular Applications


## Quickstart

Accessing an ERC-721 (NFT) smart contract in Go:

```go
package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/metachris/eth-go-bindings/erc165"
    "github.com/metachris/eth-go-bindings/erc721"
)

func main() {
    // Connect to a geth node (when using Infura, you need to use your own API key)
    conn, err := ethclient.Dial("https://mainnet.infura.io/v3/API_KEY")
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    // Instantiate the ERC721 contract for Uniswap V3: Positions NFT
    address := common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88")
    token, err := erc721.NewErc721(address, conn)
    if err != nil {
        log.Fatalf("Failed to instantiate a Token contract: %v", err)
    }

    // Access token properties
    name, err := token.Name(nil)
    if err != nil {
        log.Fatalf("Failed to retrieve token name: %v", err)
    }
    fmt.Println("Token name:", name)

    // Invoke the ERC165 SupportsInterface method
    supportsMetadata, err := token.SupportsInterface(nil, erc165.InterfaceIdErc721Metadata)
    if err != nil {
        log.Fatalf("Failed to retrieve supportsInterface: %v", err)
    }
    fmt.Println("Supports ERC721Metadata extension:", supportsMetadata)
}
```

## Smart contracts & building them

Sources: https://github.com/metachris/eth-go-bindings/tree/master/contracts

Building:

```bash
# Install dependencies
yarn init -y
yarn add truffle @openzeppelin/contracts @chainsafe/truffle-plugin-abigen

# Compile and create ABIs
yarn truffle compile
yarn truffle run abigen Erc20 Erc165 Erc721

# Generate Go bindings
abigen --bin=abigenBindings/bin/Erc20.bin --abi=abigenBindings/abi/Erc20.abi --pkg=erc20 --out=erc20/erc20.go
abigen --bin=abigenBindings/bin/Erc165.bin --abi=abigenBindings/abi/Erc165.abi --pkg=erc165 --out=erc165/erc165.go
abigen --bin=abigenBindings/bin/Erc721.bin --abi=abigenBindings/abi/Erc721.abi --pkg=erc721 --out=erc721/erc721.go
```
