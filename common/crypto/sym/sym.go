/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sym

import (
	"chainmaker.org/chainmaker-go/common/crypto"
	"chainmaker.org/chainmaker-go/common/crypto/sym/aes"
	"chainmaker.org/chainmaker-go/common/crypto/sym/sm4"
	"encoding/hex"
	"errors"
	"fmt"
)

var (
	aesKeyLengthError = errors.New("aes key len must be 128bit，192bit，256bit")
	sm4KeyLengthError = errors.New("sm4 key len must be 128bit")
)

func GenerateSymKeyStr(keyType crypto.KeyType, keyHex string) (crypto.SymmetricKey, error) {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return nil, err
	}

	return GenerateSymKey(keyType, key)
}

func GenerateSymKey(keyType crypto.KeyType, key []byte) (crypto.SymmetricKey, error) {
	bits := len(key) * 8

	switch keyType {
	case crypto.AES:
		if bits != int(crypto.BITS_SIZE_128) &&
			bits != int(crypto.BITS_SIZE_192) && bits != int(crypto.BITS_SIZE_256) {
			return nil, aesKeyLengthError
		}
		return &aes.AESKey{Key: key}, nil
	case crypto.SM4:
		if bits != int(crypto.BITS_SIZE_128) {
			return nil, sm4KeyLengthError
		}
		return &sm4.SM4Key{Key: key}, nil
	default:
		return nil, fmt.Errorf("unsupport symmetric algorithm")
	}
}
