package state

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Address() common.Address {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Failed to generate private key:", err)
		return common.Address{}
	}

	// 从私钥生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Failed to derive public key")
		return common.Address{}
	}

	//从公钥生成以太坊地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	//addressHex := address.Hex()
	//privateKeyBytes := crypto.FromECDSA(privateKey)
	//privateKeyHex := hex.EncodeToString(privateKeyBytes)
	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//publicKeyHex := hex.EncodeToString(publicKeyBytes)

	//fmt.Println("Private key:", privateKeyHex)
	//fmt.Println("Public key:", publicKeyHex)
	//fmt.Println("Address:", addressHex)

	return address
}
