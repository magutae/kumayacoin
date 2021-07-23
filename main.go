package main

import (
	"fmt"

	"github.com/magutae/kumayacoin/blockchain"
)

// import (
// 	"crypto/sha256"
// 	"fmt"
// )

// type block struct {
// 	data     string
// 	hash     string
// 	prevHash string
// }

// type blockchain struct {
// 	blocks []block
// }

// func (b *blockchain) getLastHash() string {
// 	if len(b.blocks) > 0 {
// 		return b.blocks[len(b.blocks)-1].hash
// 	}
// 	return ""
// }

// func (b *blockchain) addBlcok(data string) {
// 	newBlock := block{data, "", b.getLastHash()}
// 	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
// 	newBlock.hash = fmt.Sprintf("%x", hash)

// 	b.blocks = append(b.blocks, newBlock)
// }

// func (b *blockchain) listBlocks() {
// 	for _, block := range b.blocks {
// 		fmt.Printf("Data: %s\n", block.data)
// 		fmt.Printf("Hash: %s\n", block.hash)
// 		fmt.Printf("Prev Hash: %s\n\n", block.prevHash)
// 	}

// }

func main() {
	// genesisBlock := block{"Genesis Block", "", ""}
	// hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	// hexHash := fmt.Sprintf("%x", hash)
	// genesisBlock.hash = hexHash
	// fmt.Println(genesisBlock)
	// chain := blockchain{}
	// chain.addBlcok("Genesis Block")
	// chain.addBlcok("Second Block")
	// chain.addBlcok("Third Block")
	// chain.listBlocks()
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Prev Hash: %s\n", block.PrevHash)
	}
}
