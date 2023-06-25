package state

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

/*
以太坊地址生成的过程

第一步：私钥 (private key)

　　伪随机数产生的256bit私钥示例(256bit  16进制32字节)

　　18e14a7b6a307f426a94f8114701e7c8e774e7f9a47e2c2035db29a206321725

第二步：公钥 (public key)

　　1. 采用椭圆曲线数字签名算法ECDSA-secp256k1将私钥（32字节）映射成公钥（65字节）（前缀04+X公钥+Y公钥）：

　　04
　　50863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b2352
　　2cd470243453a299fa9e77237716103abc11a1df38855ed6f2ee187e9c582ba6

       2. 拿公钥（非压缩公钥）来hash，计算公钥的 Keccak-256 哈希值（32bytes）：

　　fc12ad814631ba689f7abe671016f75c54c607f082ae6b0881fac0abeda21781

       3. 取上一步结果取后20bytes即以太坊地址：

　　1016f75c54c607f082ae6b0881fac0abeda21781

第三步：地址 (address)

　　0x1016f75c54c607f082ae6b0881fac0abeda21781
*/

func TestAccount(t *testing.T) {
	// 生成随机的以太坊私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Failed to generate private key:", err)
		return
	}

	// 从私钥生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Failed to derive public key")
		return
	}

	//从公钥生成以太坊地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	addressHex := address.Hex()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyHex := hex.EncodeToString(publicKeyBytes)

	fmt.Println("Private key:", privateKeyHex)
	fmt.Println("Public key:", publicKeyHex)
	fmt.Println("Address:", addressHex)
	fmt.Println(common.Address{})
}
