package main

import (
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/params"
	"encoding/json"
	"fmt"
)

func convertGoGenesisToParityConfig() {
	genesisBlock := core.DefaultGenesisBlock()
	parityChainSpec, _ := newParityChainSpec("mainnet", genesisBlock, params.MainnetBootnodes)
	b, err := json.Marshal(parityChainSpec)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return;
	}
	fmt.Println(string(b))
}

//func main() {
//	convertGoGenesisToParityConfig()
//}
