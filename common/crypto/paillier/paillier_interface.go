/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package paillier

import (
	"math/big"
)

var Helper func() Provider

// Provider .
type Provider interface {
	// NewPubKey new pub key
	NewPubKey() Pub
	// NewPrvKey new prv key
	NewPrvKey() Prv
	// NewCiphertext new ciphertext
	NewCiphertext() Ct
	// NewKeyGenerator new generator
	NewKeyGenerator() KeyGenerator
}

type Serializer interface {
	// Marshal returns the byte encoding.
	Marshal() ([]byte, error)

	// Unmarshal parses the encoded data and stores the result
	Unmarshal([]byte) error
}

// Ct paillier ciphertext
type Ct interface {
	Serializer
	GetChecksum() ([]byte, error)
	GetCtBytes() ([]byte, error)
	GetCtStr() (string, error)
}

// Pub paillier public key
type Pub interface {
	Serializer
	// ChecksumVerify is used to verify hash fingerprint
	ChecksumVerify(ciphertext Ct) bool

	// Encrypt converts the provided plaintext to ciphertext, using the provided public key
	Encrypt(plaintext *big.Int) (Ct, error)

	// AddCiphertext sets a new Ciphertext to the sum ciphertext1+ciphertext2 of type big.Int and return the result
	AddCiphertext(ciphertext1, ciphertext2 Ct) (Ct, error)

	// AddPlaintext uses the key.PubKey.Key to add ciphertext1, ciphertext2 of type big.Int and return the result
	AddPlaintext(ciphertext Ct, plaintext *big.Int) (Ct, error)

	// SubCiphertext sets a new Ciphertext to the sum ciphertext1-ciphertext2 and return the result
	SubCiphertext(ciphertext1, ciphertext2 Ct) (Ct, error)

	// SubPlaintext sets a new Ciphertext to the sum ciphertext1-ciphertext2 and return the result
	SubPlaintext(ciphertext Ct, plaintext *big.Int) (Ct, error)

	// NumMul sets a new Ciphertext to the sum ciphertext1-ciphertext2 and return the result
	NumMul(ciphertext Ct, plaintext *big.Int) (Ct, error)
}

// KeyGenerator paillier key generator
type KeyGenerator interface {
	// GenKey generates public key and private key
	GenKey() (Prv, error)
}

// Prv paillier private key
type Prv interface {
	Pub
	// Decrypt decrypt the supplied ciphertext into plaintext using the private key provided
	Decrypt(ciphertext Ct) (*big.Int, error)

	// GetPubKey Get public key
	GetPubKey() (Pub, error)
}
