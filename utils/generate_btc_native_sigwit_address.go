package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func GenerateNativeSegWitAddress(publicKeyHex string) (string, error) {
	// 解码 16 进制公钥
	pubKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode public key: %v", err)
	}
	// 计算 RIPEMD160(SHA256(pubkey))
	pubKeyHash := btcutil.Hash160(pubKeyBytes)
	// 使用 mainnet 参数创建地址
	address, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
	if err != nil {
		return "", fmt.Errorf("failed to create address: %v", err)
	}
	// 返回 Bech32 编码的地址
	return address.EncodeAddress(), nil
}
