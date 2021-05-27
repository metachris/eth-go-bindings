# Go bindings for Ethereum smart contracts

* [ERC721: NFT](https://eips.ethereum.org/EIPS/eip-721) with [ERC721 Metadata](https://docs.openzeppelin.com/contracts/4.x/api/token/erc721#IERC721Metadata) and [ERC721 Enumerable](https://docs.openzeppelin.com/contracts/4.x/api/token/erc721#IERC721Enumerable) extensions.
* [EIP165: Standard Interface Detection](https://eips.ethereum.org/EIPS/eip-165)

Source: [contracts/Erc721.sol](https://github.com/metachris/erc721-go-bindings/blob/master/contracts/Erc721.sol) (based on [OpenZeppelin ERC721 contract](https://docs.openzeppelin.com/contracts/4.x/erc721)).

## Quickstart

In Go:

```go
package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metachris/erc721-go-bindings/erc165"
	"github.com/metachris/erc721-go-bindings/erc721"
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

	supportsMetadata, err := token.SupportsInterface(nil, erc165.InterfaceIdErc721Metadata)
	if err != nil {
		log.Fatalf("Failed to retrieve supportsInterface: %v", err)
	}
	fmt.Println("Supports ERC721Metadata extension:", supportsMetadata)
}
```

Get dependencies and run:

```bash
go mod tidy
go run .
```

## Smart contract & building

Source: `Erc721.sol`

Build the contract:

```bash
# Install dependencies
yarn init -y
yarn add truffle @openzeppelin/contracts @chainsafe/truffle-plugin-abigen
yarn truffle compile
yarn truffle run abigen Erc165 Erc721
abigen --bin=abigenBindings/bin/Erc165.bin --abi=abigenBindings/abi/Erc165.abi --pkg=erc165 --out=erc165/erc165.go
abigen --bin=abigenBindings/bin/Erc721.bin --abi=abigenBindings/abi/Erc721.abi --pkg=erc721 --out=erc721/erc721.go
```


Versions used to build the bindings:

* `go1.16.3`
* `abigen version 1.10.4-unstable`
* `solc: 0.8.4+commit.c7e474f2.Emscripten.clang`
