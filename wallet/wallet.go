package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/magutae/kumayacoin/utils"
)

// 1) hash the msg
// "message" -> hash(x) -> "hashed_message"

// 2) generate key pair
// KeyPair (privateKey, publicKey) -> (save priv to a file: wallet)

// 3) sign the hash
// ("hashed_message" + privateKey) -> "signature"

// 4) verify
// ("hashed_message", + "segnature" + publicKey) -> true / false

func Start() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	message := "i love you"
	hashedMessage := utils.Hash(message)

	fmt.Println(hashedMessage)

	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	utils.HandleErr(err)

	ok := ecdsa.Verify(&privateKey.PublicKey, hashAsBytes, r, s)
	fmt.Println("ok:", ok)
}
