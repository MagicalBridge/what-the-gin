package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/tyler-smith/go-bip39"
	"strconv"
	"strings"
)

// GenerateEntropy generates random entropy of the specified bit size.
func GenerateEntropy(bitSize int) ([]byte, error) {
	// Define a map of valid bit sizes
	validBitSizes := map[int]bool{128: true, 160: true, 192: true, 224: true, 256: true}
	// Check if the provided bitSize is valid.
	if !validBitSizes[bitSize] {
		return nil, errors.New("Invalid entropy bit size, should be one of 128, 160, 192, 224, or 256")
	}

	// Calculate the byte size by dividing the bit size by 8.
	// 将位转换为字节
	byteSize := bitSize / 8

	// Create a byte slice of the calculated size.
	entropy := make([]byte, byteSize)

	// Generate random bytes to fill the entropy slice.
	_, err := rand.Read(entropy)

	if err != nil {
		return nil, err
	}
	return entropy, nil
}

func CalculateCheckBits(entropy []byte) uint8 {
	// 计算熵的长度（bit数）
	ENT := len(entropy) * 8
	// 计算校验和位数（checksum length）。这里等于总熵长度除以32
	CS := ENT / 32

	// 使用 SHA-256 哈希算法对熵进行哈希，得到一个hash值
	hash := sha256.Sum256(entropy)

	// 获取哈希的第一个字节
	firstByte := hash[0]
	shiftAmount := 8 - uint(CS) // 计算需要右移的位数

	// 无符号右移操作
	var result uint8
	if shiftAmount == 0 {
		result = firstByte
	} else {
		result = firstByte >> shiftAmount
	}
	return result
}

// CombineEntropyAndCheckBitsToBinary 将熵和校验位组合成二进制字符串
func CombineEntropyAndCheckBitsToBinary(entropy []byte, checkBits uint8) string {
	// 初始化一个空的二进制字符串
	binaryString := ""

	// 将熵中的每个字节转换为二进制字符串，并连接起来
	for _, byteIndex := range entropy {
		// 将字节转换为8位二进制字符串，不足8位的用0填充
		binaryString += fmt.Sprintf("%08b", byteIndex)
	}

	// 计算校验和位数（checksum length）
	CS := (len(entropy) * 8) / 32

	// 将校验位转换为二进制字符串，不足CS位的用0填充
	binaryString += fmt.Sprintf("%0*b", CS, checkBits)
	return binaryString
}

func SplitIntoIndices(bits string) ([]int, error) {
	var indices []int
	totalBits := len(bits)
	wordCount := totalBits / 11

	for i := 0; i < totalBits; i += 11 {
		index, err := strconv.ParseInt(bits[i:i+11], 2, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse bits: %v", err)
		}
		indices = append(indices, int(index))
	}

	if len(indices) != wordCount {
		return nil, fmt.Errorf("invalid number of indices generated. Expected %d, but got %d", wordCount, len(indices))
	}

	return indices, nil
}

func IndicesToMnemonic(indices []int) (string, error) {
	// 获取 BIP-39 英文单词列表
	wordlist := bip39.GetWordList()
	if len(wordlist) == 0 {
		return "", fmt.Errorf("failed to get wordlist")
	}

	var mnemonicWords []string

	for _, index := range indices {
		if index < 0 || index >= len(wordlist) {
			return "", fmt.Errorf("index out of bounds: %d", index)
		}
		mnemonicWords = append(mnemonicWords, wordlist[index])
	}

	return strings.Join(mnemonicWords, " "), nil
}

func GenerateMnemonic(entropyBitSize int) (string, error) {
	entropy, err := GenerateEntropy(entropyBitSize)
	if err != nil {
		return "", err
	}
	checkBits := CalculateCheckBits(entropy)
	combinedBits := CombineEntropyAndCheckBitsToBinary(entropy, checkBits)
	indices, err := SplitIntoIndices(combinedBits)
	return IndicesToMnemonic(indices)
}
