package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/magutae/kumayacoin/db"
	"github.com/magutae/kumayacoin/utils"
	"github.com/magutae/kumayacoin/wallet"
)

var ErrNotFound = errors.New("block not found")

type Block struct {
	Hash         string `json:"hash"`
	PrevHash     string `json:"prevHash,omitempty"`
	Height       int    `json:"height"`
	Difficulty   int    `json:"difficulty"`
	Nonce        int    `json:"nonce"`
	Timestamp    int    `json:"timestamp"`
	Transactions []*Tx  `json:"transations"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		fmt.Printf("Hash %s\nTarget: %s\nNonce: %d\n\n", hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(prevHash string, height, diff int) *Block {
	block := &Block{
		Hash:         "",
		PrevHash:     prevHash,
		Height:       height,
		Difficulty:   diff,
		Nonce:        0,
		Transactions: []*Tx{makeCoinbaseTx(wallet.Wallet().Address)},
	}
	block.mine()
	block.Transactions = Mempool.TxToConfirm()
	block.persist()
	return block
}
