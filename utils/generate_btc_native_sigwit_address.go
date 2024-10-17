package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
	"golang.org/x/crypto/ripemd160"
)

// Hash160 计算 SHA256 + RIPEMD160
func hash160(data []byte) []byte {
	sha256Hash := sha256.Sum256(data)
	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(sha256Hash[:])
	return ripemd160Hasher.Sum(nil)
}

func GenerateNativeSegWitAddress(publicKeyHex string) (string, error) {
	// Step 1: 解码 16 进制公钥
	pubKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode public key: %v", err)
	}

	// Step 2: 计算公钥哈希
	pubKeyHash := hash160(pubKeyBytes)

	// Step 3: 准备 SegWit 地址编码数据
	witnessVersion := byte(0x00) // P2WPKH 的版本号为 0

	// Step 4: 使用 bech32 库进行编码
	data, err := bech32.ConvertBits(pubKeyHash, 8, 5, true) // 将 8-bit 转换为 5-bit
	if err != nil {
		return "", fmt.Errorf("failed to convert bits: %v", err)
	}

	// 将版本号和转换后的数据合并
	combinedData := append([]byte{witnessVersion}, data...)

	// Bech32 编码
	address, err := bech32.Encode("bc", combinedData)
	if err != nil {
		return "", fmt.Errorf("failed to encode Bech32 address: %v", err)
	}

	// 返回 Bech32 编码的地址
	return address, nil
}

// 使用自带的工具函数可以便捷生成
//func GenerateNativeSegWitAddress(publicKeyHex string) (string, error) {
//	// 解码 16 进制公钥
//	pubKeyBytes, err := hex.DecodeString(publicKeyHex)
//	if err != nil {
//		return "", fmt.Errorf("failed to decode public key: %v", err)
//	}
//	// 计算 RIPEMD160(SHA256(pubkey))
//	pubKeyHash := btcutil.Hash160(pubKeyBytes)
//	// 使用 mainnet 参数创建地址
//	address, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
//	if err != nil {
//		return "", fmt.Errorf("failed to create address: %v", err)
//	}
//	// 返回 Bech32 编码的地址
//	return address.EncodeAddress(), nil
//}
