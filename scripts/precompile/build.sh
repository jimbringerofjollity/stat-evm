#!/bin/bash
solc --abi --bin --overwrite Test.sol -o ./
abigen --pkg main --abi Test.abi --bin Test.bin > wrap.go
# abigen --sol Test.sol --pkg main --out wrap.go
