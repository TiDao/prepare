// +build paillier

/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package paillier

import (
	"math/big"
)

// validatePrvKey is used to validate the private key
func validatePrvKey(key *PrvKey) error {
	if key == nil || key.Key == nil {
		return ErrInvalidPrivateKey
	}

	return validatePubKey(key.PubKey)
}

// validatePubKey is used to validate the public key
func validatePubKey(key *PubKey) error {
	if key == nil || key.Key == nil || key.Key.bits != defaultKeySize {
		return ErrInvalidPublicKey
	}

	return nil
}

// validateCiphertext is used to validate Ciphertext
func validateCiphertext(cts ...Ct) error {
	for _, ct := range cts {
		if ct == nil || ct.(*Ciphertext).Ct == nil || ct.(*Ciphertext).Checksum == nil {
			return ErrInvalidCiphertext
		}
	}

	return nil
}

// validatePlaintext is used to validate the paillier Ciphertext of type big.Int
func validatePlaintext(paillierTexts ...*big.Int) error {
	for _, paillierText := range paillierTexts {
		if paillierText == nil {
			return ErrInvalidPlaintext
		}
	}
	return nil
}
