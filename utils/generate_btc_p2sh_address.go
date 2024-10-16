package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

func GenerateP2SHAddress(pubKeyHex string) (string, error) {
	// 将公钥的十六进制字符串转成字节数组
	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return "", err
	}

	// 对公钥进行 SHA-256 哈希运算
	hasherSHA256 := sha256.New()
	hasherSHA256.Write(pubKeyBytes)
	pubKeySHA256 := hasherSHA256.Sum(nil)

	// 对 SHA-256 哈希值进行 RIPEMD-160 哈希运算
	hasherRIPEMD160 := ripemd160.New()
	hasherRIPEMD160.Write(pubKeySHA256)
	pubKeyHash := hasherRIPEMD160.Sum(nil)

	// 构建 P2WPKH 的见证程序（Witness Program）
	// 其中 0x00 表示版本号，0x14 表示长度（20 字节）
	witnessProgram := append([]byte{0x00, 0x14}, pubKeyHash...)

	// 对见证程序进行 SHA-256 哈希运算
	hasherSHA256.Reset()
	hasherSHA256.Write(witnessProgram)
	witnessProgramSHA256 := hasherSHA256.Sum(nil)

	// 对 SHA-256 哈希值进行 RIPEMD-160 哈希运算
	hasherRIPEMD160.Reset()
	hasherRIPEMD160.Write(witnessProgramSHA256)
	witnessProgramHash := hasherRIPEMD160.Sum(nil)

	// 添加 P2SH 前缀
	version := []byte{0x05}
	versionedHash := append(version, witnessProgramHash...)

	// 计算校验和
	hasherDoubleSHA256 := sha256.New()
	hasherDoubleSHA256.Write(versionedHash)
	checksum := hasherDoubleSHA256.Sum(nil)
	hasherDoubleSHA256.Reset()
	hasherDoubleSHA256.Write(checksum)
	checksum = hasherDoubleSHA256.Sum(nil)
	finalChecksum := checksum[:4]

	// 拼接版本前缀、赎回脚本哈希和校验和
	fullHash := append(versionedHash, finalChecksum...)

	// 转换到 Base58
	address := base58.Encode(fullHash)

	return address, nil
}
