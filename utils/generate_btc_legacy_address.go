package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

// HexStringToBytes 将 16 进制的字符串转为字节数组
func HexStringToBytes(hexStr string) ([]byte, error) {
	return hex.DecodeString(hexStr)
}

func Generate_btc_legacy_address(publicKeyHex string) (string, error) {
	// 1. 将公钥从16进制字符串转换为字节数组
	publicKeyBytes, err := HexStringToBytes(publicKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid public key: %v", err)
	}

	// 2. 进行 SHA-256 哈希
	sha256Hash := sha256.New()
	sha256Hash.Write(publicKeyBytes)
	publicKeySHA256 := sha256Hash.Sum(nil)

	// 3. 进行 RIPEMD-160 哈希
	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(publicKeySHA256)
	publicKeyHash := ripemd160Hasher.Sum(nil)

	// 4. 添加比特币 P2PKH 地址的前缀 0x00
	version := []byte{0x00}
	versionedPayload := append(version, publicKeyHash...)

	// 5. 进行双重 SHA-256 哈希，用于生成校验和
	firstSHA := sha256.Sum256(versionedPayload)
	secondSHA := sha256.Sum256(firstSHA[:]) // 将 [32]byte 转换为切片
	checksumBytes := secondSHA[:4]          // 取前4个字节作为校验和

	// 6. 将 version, publicKeyHash 和校验和组合在一起
	fullPayload := append(versionedPayload, checksumBytes...)

	// 7. 使用 Base58 编码
	address := base58.Encode(fullPayload)

	return address, nil
}
