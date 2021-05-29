.DEFAULT_GOAL := bindings
.PHONY: clean bindings

clean:  ## Remove build artifacts
	rm -rf build abigenBindings

bindings:  ## Create the Go bindings
	yarn truffle compile
	yarn truffle run abigen Erc20 Erc165 Erc721 Erc777 ERc1155

	abigen --bin=abigenBindings/bin/Erc20.bin --abi=abigenBindings/abi/Erc20.abi --pkg=erc20 --out=erc20/erc20.go
	abigen --bin=abigenBindings/bin/Erc165.bin --abi=abigenBindings/abi/Erc165.abi --pkg=erc165 --out=erc165/erc165.go
	abigen --bin=abigenBindings/bin/Erc721.bin --abi=abigenBindings/abi/Erc721.abi --pkg=erc721 --out=erc721/erc721.go
	abigen --bin=abigenBindings/bin/Erc777.bin --abi=abigenBindings/abi/Erc777.abi --pkg=erc777 --out=erc777/erc777.go
	abigen --bin=abigenBindings/bin/Erc1155.bin --abi=abigenBindings/abi/Erc1155.abi --pkg=erc1155 --out=erc1155/erc1155.go
