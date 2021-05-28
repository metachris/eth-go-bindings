// Example for interacting with ERC20 and ERC721 contracts
package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/metachris/eth-go-bindings/erc165"
	"github.com/metachris/eth-go-bindings/erc20"
	"github.com/metachris/eth-go-bindings/erc721"
)

func Perror(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func TestErc20(conn *ethclient.Client) {
	// Instantiate the ERC20 contract for the Uniswap token
	address := common.HexToAddress("0x1f9840a85d5af5bf1d1762f925bdaddc4201f984")
	token, err := erc20.NewErc20(address, conn)
	Perror(err)

	// Access token properties
	name, err := token.Name(nil)
	Perror(err)
	fmt.Println("ERC20 token name:", name)

	decimals, err := token.Decimals(nil)
	Perror(err)
	fmt.Println("Token decimals:", decimals)
}

func TestErc721(conn *ethclient.Client) {
	// Instantiate the ERC721 contract for Uniswap V3: Positions NFT
	address := common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88")
	token, err := erc721.NewErc721(address, conn)
	Perror(err)

	// Access token properties
	name, err := token.Name(nil)
	Perror(err)
	fmt.Println("ER721 token name:", name)

	supportsMetadata, err := token.SupportsInterface(nil, erc165.InterfaceIdErc721Metadata)
	Perror(err)
	fmt.Println("Supports ERC721Metadata extension:", supportsMetadata)
}

func main() {
	// Connect to a geth node (when using Infura, you need to use your own API key)
	var conn, err = ethclient.Dial("https://mainnet.infura.io/v3/API_KEY")
	Perror(err)

	TestErc20(conn)
	TestErc721(conn)
}
