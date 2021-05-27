Go bindings for ERC721 (NFT) Ethereum smart contract.

Source: `Erc721.sol`

Build the contract:

```bash
# Install dependencies
yarn init -y
yarn add truffle @openzeppelin/contracts @chainsafe/truffle-plugin-abigen
yarn truffle compile
yarn truffle run abigen Erc721
abigen --bin=abigenBindings/bin/Erc721.bin --abi=abigenBindings/abi/Erc721.abi --pkg=erc721 --out=erc721.go
```


Versions used to build the bindings:

* `go1.16.3`
* `abigen version 1.10.4-unstable`
* `solc: 0.8.4+commit.c7e474f2.Emscripten.clang`
