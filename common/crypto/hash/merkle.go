/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package hash

import (
	"crypto/sha256"
	"math"

	"chainmaker.org/chainmaker-go/common/crypto"
)

var h = sha256.New()

func GetMerkleRoot(hashType string, hashes [][]byte) ([]byte, error) {
	if hashes == nil || len(hashes) == 0 {
		return nil, nil
	}

	merkleTree, err := BuildMerkleTree(hashType, hashes)
	if err != nil {
		return nil, err
	}
	return merkleTree[len(merkleTree)-1], nil
}

// take leaf node hash array and build merkle tree
func BuildMerkleTree(hashType string, hashes [][]byte) ([][]byte, error) {
	var hasher = Hash{
		hashType: crypto.HashAlgoMap[hashType],
	}

	var err error
	if hashes == nil || len(hashes) == 0 {
		return nil, nil
	}

	// use array to store merkle tree entries
	nextPowOfTwo := getNextPowerOfTwo(len(hashes))
	arraySize := nextPowOfTwo*2 - 1
	merkelTree := make([][]byte, arraySize)

	// 1. copy hashes first
	copy(merkelTree[:len(hashes)], hashes[:])

	// 2. compute merkle step by step
	offset := nextPowOfTwo
	for i := 0; i < arraySize-1; i += 2 {
		switch {
		case merkelTree[i] == nil:
			// parent node is nil if left is nil
			merkelTree[offset] = nil
		case merkelTree[i+1] == nil:
			// hash(left, left) if right is nil
			merkelTree[offset], err = hashMerkleBranches(hasher, merkelTree[i], merkelTree[i])
			if err != nil {
				return nil, err
			}
		default:
			// default hash(left||right)
			merkelTree[offset], err = hashMerkleBranches(hasher, merkelTree[i], merkelTree[i+1])
			if err != nil {
				return nil, err
			}
		}
		offset++
	}

	return merkelTree, nil
}

func getNextPowerOfTwo(n int) int {
	if n&(n-1) == 0 {
		return n
	}

	exponent := uint(math.Log2(float64(n))) + 1
	return 1 << exponent
}

func hashMerkleBranches(hasher Hash, left []byte, right []byte) ([]byte, error) {
	data := make([]byte, len(left)+len(right))
	copy(data[:len(left)], left)
	copy(data[len(left):], right)
	return hasher.Get(data)
}

func getNextPowerOfTen(n int) (int, int) {
	//if n&(n-1) == 0 {
	//	return n, 0
	//}
	if n == 1 {
		return 1, 0
	}

	exponent := int(math.Log10(float64(n-1))) + 1
	rootsSize := 0
	for i := 0; i < exponent; i++ {
		rootsSize += int(math.Pow10(i))
	}
	return int(math.Pow10(exponent)), rootsSize
}
