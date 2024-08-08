package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/cbergoon/merkletree"
	"golang.org/x/crypto/sha3"
	"log"
)

type Content struct {
	x string
}

// CalculateHash calculates the hash of the Content
func (c Content) CalculateHash() ([]byte, error) {
	hash := sha3.NewLegacyKeccak256()
	_, err := hash.Write([]byte(c.x))
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (c Content) Equals(other merkletree.Content) (bool, error) {
	return c.x == other.(Content).x, nil
}

// HashFunc Hash function used for Merkle tree
func HashFunc(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

// VerifyMerkleProof manually verifies the Merkle proof
func VerifyMerkleProof(rootHash []byte, leafHash []byte, proof [][]byte, isLeft []bool) bool {
	currentHash := leafHash
	for i, p := range proof {
		if isLeft[i] {
			// Hash (currentHash || p)
			currentHash = HashFunc(append(currentHash, p...))
		} else {
			// Hash (p || currentHash)
			currentHash = HashFunc(append(p, currentHash...))
		}
	}
	return hex.EncodeToString(currentHash) == hex.EncodeToString(rootHash)
}

func VerifyMerkle() string {
	//
	whitelistAddresses := []string{
		"0x1234567890123456789012345678901234567890",
		"0x2345678901234567890123456789012345678901",
	}

	// Create leaf nodes
	var list []merkletree.Content
	for _, addr := range whitelistAddresses {
		list = append(list, Content{x: addr})
	}

	// Create a new Merkle Tree
	tree, err := merkletree.NewTreeWithHashStrategy(list, sha3.NewLegacyKeccak256)
	if err != nil {
		log.Fatal(err)
	}

	// Get the Merkle Root
	root := tree.MerkleRoot()
	fmt.Printf("Merkle Root: %s\n", hex.EncodeToString(root))

	// Generate proof for a specific address
	address := "0x1234567890123456789012345678901234567890"
	targetContent := Content{x: address}
	proof, isLeftInt64, err := tree.GetMerklePath(targetContent)
	if err != nil {
		log.Fatalf("Error getting Merkle path: %v", err)
	}

	// Convert []int64 to []bool
	isLeft := make([]bool, len(isLeftInt64))
	for i, v := range isLeftInt64 {
		isLeft[i] = v != 0
	}

	// Print the proof
	fmt.Printf("Proof for %s: ", address)
	for _, p := range proof {
		fmt.Printf("%s ", hex.EncodeToString(p))
	}
	fmt.Println()

	// Calculate the leaf hash
	leafHash, err := targetContent.CalculateHash()
	if err != nil {
		log.Fatalf("Error calculating leaf hash: %v", err)
	}

	// Verify the proof
	valid := VerifyMerkleProof(root, leafHash, proof, isLeft)
	if valid {
		fmt.Printf("Proof is valid for address %s\n", address)
		return address
	} else {
		fmt.Printf("Proof is invalid for address %s\n", address)
		return address
	}
}
