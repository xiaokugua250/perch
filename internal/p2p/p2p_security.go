package p2p

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/libp2p/go-libp2p-core/crypto"
	"os"
	"strings"
)

/**
根據指定字符串生成publickey，privatekey
*/
func GenSecurekeysByStr(str string) (crypto.PrivKey, crypto.PubKey, error) {

	h := sha512.New()
	h.Write([]byte(str))
	hashed := h.Sum(nil)
	hash := hex.EncodeToString(hashed)
	r := strings.NewReader(hash)
	privateKey, publicKey, err := crypto.GenerateKeyPairWithReader(crypto.Ed25519, 1024, r)
	return privateKey, publicKey, err

}

/**
根據指定字符串生成publickey，privatekey
*/
func GenSecurekeysByHostname() (crypto.PrivKey, crypto.PubKey, error) {

	str, err := os.Hostname()
	if err != nil {
		return nil, nil, err
	}
	h := sha512.New()
	h.Write([]byte(str))
	hashed := h.Sum(nil)
	hash := hex.EncodeToString(hashed)
	r := strings.NewReader(hash)
	privateKey, publicKey, err := crypto.GenerateKeyPairWithReader(crypto.Ed25519, 1024, r)
	return privateKey, publicKey, err

}
