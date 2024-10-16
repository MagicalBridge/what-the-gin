package utils

import (
	"crypto/sha256"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	"log"
)

// 生成私钥和公钥
func generateKeyPair() (*btcec.PrivateKey, []byte) {

	privateKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		log.Fatal(err)
	}

	// 公钥采用压缩格式
	publicKey := privateKey.PubKey().SerializeCompressed()

	return privateKey, publicKey
}

// 计算SHA256 -> RIPE-MD160
func hash160(data []byte) []byte {
	// 先对数据进行SHA-256
	sha256Hash := sha256.Sum256(data)

	// 再对结果进行RIPE-MD-160
	ripemd160Hasher := ripemd160.New()
	_, err := ripemd160Hasher.Write(sha256Hash[:])
	if err != nil {
		log.Fatal(err)
	}

	return ripemd160Hasher.Sum(nil)
}

// 计算双SHA-256并取前4个字节作为校验和
func checksum(payload []byte) []byte {
	// 第一次SHA-256
	firstHash := sha256.Sum256(payload)

	// 第二次SHA-256
	secondHash := sha256.Sum256(firstHash[:])

	// 取前4个字节作为校验和
	return secondHash[:4]
}

func GenerateP2PKHAddress() string {
	// 1. 生成私钥和公钥
	_, publicKey := generateKeyPair()

	// 2. 公钥哈希: 先SHA-256，再RIPEMD-160
	publicKeyHash := hash160(publicKey)

	// 3. 添加网络前缀，0x00表示比特币主网地址
	versionedPayload := append([]byte{0x00}, publicKeyHash...)

	// 4. 计算校验和，取前4个字节
	checksum := checksum(versionedPayload)

	// 5. 拼接校验和，生成完整的payload
	fullPayload := append(versionedPayload, checksum...)

	// 6. Base58Check编码
	address := base58.Encode(fullPayload)

	return address
}
