/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/
package hibe

import (
	"chainmaker.org/chainmaker-go/common/crypto"
	"chainmaker.org/chainmaker-go/common/crypto/hash"
	"chainmaker.org/chainmaker-go/common/crypto/sym"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/samkumar/hibe"
	"math/big"
	"reflect"
	"strings"
)

// ValidateId is used to validate id format
func ValidateId(id string) error {
	if id == "" {
		return errors.New("invalid parameters, id is nil")
	}

	idStrList := strings.Split(id, "/")

	for _, s := range idStrList {
		if s == "" {
			return fmt.Errorf("invalid parameters, id: %s, format error, only like : \"A/B/C\" can be used", id)
		}

		if strings.Contains(s, " ") {
			return fmt.Errorf("invalid parameters, id: %s, format error, only like : \"A/B/C\" can be used", id)
		}
	}
	return nil
}

// refineIdsAndParams Remove redundant data (Id)
// Only the lowest level ID is reserved for the ID with superior subordinate relationship
func refineIdsAndParams(ids []string, paramsList []*hibe.Params) ([]string, []*hibe.Params, error) {
	refinedReceiverIds := make([]string, 1)
	refinedReceiverIds[0] = ids[0]
	refinedParamsList := make([]*hibe.Params, 1)
	refinedParamsList[0] = paramsList[0]

	for i := 1; i < len(ids); i++ {
		matched := false
		for j, rId := range refinedReceiverIds {
			// 1. refinedReceiverIds[j] has prefix ids[i]
			if strings.HasPrefix(rId, ids[i]) {
				// 1.1 they have same parameters, break inner loop and continue
				if reflect.DeepEqual(paramsList[i].Marshal(), refinedParamsList[j].Marshal()) {
					matched = true
					break
				} else {
					// 1.2 they have different parameters, return error
					return nil, nil, fmt.Errorf("ID [%s] match, but Params are different, please check it",
						ids[i])
				}
			}

			// 2. ids[i] has prefix refinedReceiverIds[j]
			if strings.HasPrefix(ids[i], rId) {
				// 2.1 they have same parameters, replace the current receiverId in rId (always keep the longer one)
				if reflect.DeepEqual(paramsList[i].Marshal(), refinedParamsList[j].Marshal()) {
					refinedReceiverIds[j] = ids[i]
					matched = true
					break
				} else {
					//  2.2 they have different parameters, return error
					return nil, nil, fmt.Errorf("ID [%s] is matched, but Params are different, please check it",
						ids[i])
				}
			}
		}

		if !matched {
			refinedReceiverIds = append(refinedReceiverIds, ids[i])
			refinedParamsList = append(refinedParamsList, paramsList[i])
		}
	}

	return refinedReceiverIds, refinedParamsList, nil
}

// idStrList2HibeIds construct HibeId list according to id list
func idStrList2HibeIds(idList []string) [][]*big.Int {
	// construct ids []*big.Int
	hibeIds := make([][]*big.Int, len(idList))

	for i, idStr := range idList {
		_, hibeIds[i] = IdStr2HibeId(idStr)
	}

	return hibeIds
}

// IdStr2HibeId construct HibeId according to id
func IdStr2HibeId(id string) ([]string, []*big.Int) {
	// idStr eg: "org1/ou1" -> ["org1", "ou1"]
	strId := strings.Split(id, "/")

	// idsStr -> hibeId
	hibeIds := make([]*big.Int, len(strId))
	for i, value := range strId {
		hashedStrId := sha256.Sum256([]byte(value))
		bigIdBytes := hashedStrId[:]
		bigId := new(big.Int)
		bigId.SetBytes(bigIdBytes)
		hibeIds[i] = bigId
	}

	return strId, hibeIds
}

// generateSymKeyFromGtBytes is used to generate symmetric key according to gtBytes and symKeyType
func generateSymKeyFromGtBytes(gtBytes []byte, symKeyType crypto.KeyType) (crypto.SymmetricKey, error) {
	var symKey crypto.SymmetricKey

	if symKeyType == crypto.AES {
		// not gm
		gtBytesHash, err := hash.Get(crypto.HASH_TYPE_SHA256, gtBytes)
		if err != nil {
			return nil, err
		}
		symKey, err = sym.GenerateSymKey(crypto.AES, gtBytesHash)
		if err != nil {
			return nil, err
		}

	} else if symKeyType == crypto.SM4 {
		// gm
		gtBytesHash, err := hash.Get(crypto.HASH_TYPE_SM3, gtBytes)
		if err != nil {
			return nil, err
		}
		symKey, err = sym.GenerateSymKey(crypto.SM4, gtBytesHash[:16])
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("invalid parameters, unsupported symmetric encryption algorithm type : %d", symKeyType)
	}

	return symKey, nil
}
