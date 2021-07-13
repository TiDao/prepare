// +build paillier

/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package paillier implements paillier interface.
//
// GNU multi precision algorithm library (GMP) is used for basic number theory operations.
// You need to install GMP before using Paillier.
// GMP relies on M4, so you should install M4 first.
//
// Installation procedure:
//
// M4
// 1. wget http://mirrors.kernel.org/gnu/m4/m4-1.4.18.tar.gz
// 2. tar -xzvf m4-1.4.18.tar.gz
// 3. cd m4-1.4.18
// 4. ./configure --prefix=/usr/local
// 5. make && make install
// check: m4 --version
//
// GMP
// 1. wget https://gmplib.org/download/gmp/gmp-6.2.1.tar.bz2
// 2. tar -jxvf gmp-6.2.1.tar.bz2
// 3. cd gmp-6.2.1
// 4. ./configure --prefix=/usr/local
// 5. make && make check && make install
package paillier

//#cgo CFLAGS: -I./c_include
//#cgo LDFLAGS: -L/usr/bin/ld -lgmp
//#cgo LDFLAGS: -L./c_lib -lpaillier_0_8
//#cgo LDFLAGS: -L./c_lib -lpaillier_extension
//#include <stdlib.h>
//#include <stdio.h>
//#include <string.h>
//#include <paillier_extension.h>
import "C"
import (
	"crypto/sha256"
	"errors"
	"math/big"
	"reflect"
	"unsafe"
)

const (
	defaultKeySize      = 256
	defaultChecksumSize = 5
)

var (
	ErrInvalidCiphertext   = errors.New("paillier: invalid ciphertext")
	ErrInvalidCiphertextFg = errors.New("paillier: invalid ciphertext Checksum")
	ErrInvalidPlaintext    = errors.New("paillier: invalid plaintext")
	ErrInvalidPublicKey    = errors.New("paillier: invalid public key")
	ErrInvalidPrivateKey   = errors.New("paillier: invalid private key")
	ErrUnknown             = errors.New("paillier: unknown error")
	ErrInvalidMismatch     = errors.New("paillier: key mismatch")
)

// PubKey paillier public key
type PubKey struct {
	Key *C.paillier_pubkey_t
}

// PrvKey paillier private key
type PrvKey struct {
	*PubKey
	Key *C.paillier_prvkey_t
}

type Ciphertext struct {
	Ct       *big.Int
	Checksum []byte
}

type provider struct{}

// NewPubKey new pub key paillier
func (*provider) NewPubKey() Pub {
	return new(PubKey)
}

// NewPrvKey new pub prv paillier
func (*provider) NewPrvKey() Prv {
	return new(PrvKey)
}

// NewCiphertext new ciphertext
func (*provider) NewCiphertext() Ct {
	return new(Ciphertext)
}

func (*provider) NewKeyGenerator() KeyGenerator {
	return new(keyGenerator)
}

type keyGenerator struct{}

func (*keyGenerator) GenKey() (Prv, error) {
	prv := new(PrvKey)
	prv.PubKey = new(PubKey)

	C.paillier_key_gen(defaultKeySize, &prv.PubKey.Key, &prv.Key, nil)

	if err := validatePrvKey(prv); err != nil {
		return nil, errors.New("paillier: generates key failed")
	}

	return prv, nil
}

// Encrypt converts the provided plaintext to ciphertext, using the provided public key.
func (key *PubKey) Encrypt(plaintext *big.Int) (Ct, error) {
	if err := validatePubKey(key); err != nil {
		return nil, err
	}

	if err := validatePlaintext(plaintext); err != nil {
		return nil, err
	}

	pt := plaintext.String()

	ciphertext := C.paillier_encrypt(key.Key, C.CString(pt), nil)
	defer C.free(unsafe.Pointer(ciphertext))

	ciphertextStr := C.GoString(ciphertext)
	result, ok := new(big.Int).SetString(ciphertextStr, 10)
	if !ok {
		return nil, ErrUnknown
	}

	Checksum, err := key.bindingCtPubKey(result.Bytes())
	ct := &Ciphertext{
		Ct:       result,
		Checksum: Checksum,
	}
	return ct, err
}

// Decrypt decrypt the supplied ciphertext into plaintext using the private key provided
func (key *PrvKey) Decrypt(ciphertext Ct) (*big.Int, error) {
	if err := validatePrvKey(key); err != nil {
		return nil, err
	}

	if err := validateCiphertext(ciphertext); err != nil {
		return nil, err
	}

	ciphertextStr, err := ciphertext.GetCtStr()
	if err != nil {
		return nil, err
	}

	plaintextCStr := C.paillier_decrypt(key.PubKey.Key, key.Key, C.CString(ciphertextStr))
	defer C.free(unsafe.Pointer(plaintextCStr))

	plaintextGoStr := C.GoString(plaintextCStr)

	plaintext, ok := new(big.Int).SetString(plaintextGoStr, 10)
	if !ok {
		return nil, ErrUnknown
	}
	return plaintext, nil
}

func (key *PrvKey) GetPubKey() (Pub, error) {
	if err := validatePrvKey(key); err != nil {
		return nil, err
	}
	return key.PubKey, nil
}

func (ct *Ciphertext) Marshal() ([]byte, error) {
	if err := validateCiphertext(ct); err != nil {
		return nil, ErrInvalidCiphertext
	}

	ctBytes := ct.Ct.Bytes()
	return append(ct.Checksum, ctBytes...), nil
}

func (ct *Ciphertext) Unmarshal(ctBytes []byte) error {
	if ctBytes == nil {
		return ErrInvalidCiphertext
	}

	if ct.Ct == nil {
		ct.Ct = new(big.Int)
	}

	ct.Ct.SetBytes(ctBytes[defaultChecksumSize:])
	ct.Checksum = ctBytes[:defaultChecksumSize]
	return nil
}

func (ct *Ciphertext) GetChecksum() ([]byte, error) {
	if err := validateCiphertext(ct); err != nil {
		return nil, err
	}

	return ct.Checksum, nil
}

func (ct *Ciphertext) GetCtBytes() ([]byte, error) {
	if err := validateCiphertext(ct); err != nil {
		return nil, err
	}

	return ct.Ct.Bytes(), nil
}

func (ct *Ciphertext) GetCtStr() (string, error) {
	if err := validateCiphertext(ct); err != nil {
		return "", err
	}

	return ct.Ct.String(), nil
}

// Marshal encodes the PubKey as a byte slice.
func (key *PubKey) Marshal() ([]byte, error) {
	if err := validatePubKey(key); err != nil {
		return nil, err
	}

	str := C.paillier_pubkey_to_hex(key.Key)
	defer C.free(unsafe.Pointer(str))
	goStr := C.GoString(str)
	return []byte(goStr), nil
}

// Unmarshal recovers the PubKey from an encoded byte slice.
func (key *PubKey) Unmarshal(pubKeyBytes []byte) error {
	if pubKeyBytes == nil {
		return ErrInvalidPublicKey
	}

	key.Key = C.paillier_pubkey_from_hex(C.CString(string(pubKeyBytes)))

	if err := validatePubKey(key); err != nil {
		return err
	}

	return nil
}

// Marshal encodes the PrvKey as a byte slice.
func (key *PrvKey) Marshal() ([]byte, error) {
	if err := validatePrvKey(key); err != nil {
		return nil, err
	}

	pubKeyBytes, err := key.PubKey.Marshal()
	if err != nil {
		return nil, err
	}

	str := C.paillier_prvkey_to_hex(key.Key)
	defer C.free(unsafe.Pointer(str))
	goStr := C.GoString(str)

	return append(pubKeyBytes, []byte(goStr)...), nil
}

// Unmarshal recovers the PrvKey from an encoded byte slice.
func (key *PrvKey) Unmarshal(prvKeyBytes []byte) error {
	if prvKeyBytes == nil || len(prvKeyBytes) != 128 {
		return ErrInvalidPrivateKey
	}

	if key.PubKey == nil {
		key.PubKey = new(PubKey)
	}

	key.PubKey.Key = C.paillier_pubkey_from_hex(C.CString(string(prvKeyBytes[:64])))

	if err := validatePubKey(key.PubKey); err != nil {
		return err
	}

	key.Key = C.paillier_prvkey_from_hex(C.CString(string(prvKeyBytes[64:])), key.PubKey.Key)

	if err := validatePrvKey(key); err != nil {
		return err
	}

	return nil
}

// AddCiphertext sets a new Ciphertext to the sum ciphertext1+ciphertext2 of type big.Int and return the result
func (key *PubKey) AddCiphertext(ciphertext1, ciphertext2 Ct) (Ct, error) {
	if err := validatePubKey(key); err != nil {
		return nil, err
	}

	if err := validateCiphertext(ciphertext1, ciphertext2); err != nil {
		return nil, err
	}

	if !key.checkOperand(ciphertext1, ciphertext2) {
		return nil, ErrInvalidMismatch
	}

	ciphertext1Str, err := ciphertext1.GetCtStr()
	if err != nil {
		return nil, err
	}

	ciphertext2Str, err := ciphertext2.GetCtStr()
	if err != nil {
		return nil, err
	}

	resultCStr := C.paillier_add_cipher(key.Key,
		C.CString(ciphertext1Str),
		C.CString(ciphertext2Str))
	defer C.free(unsafe.Pointer(resultCStr))

	resultGoStr := C.GoString(resultCStr)

	return key.constructCiphertext(resultGoStr)
}

// AddPlaintext uses the key.PubKey.Key to add ciphertext1, ciphertext2 of type big.Int and return the result
func (key *PubKey) AddPlaintext(ciphertext Ct, plaintext *big.Int) (Ct, error) {
	if err := validatePubKey(key); err != nil {
		return nil, err
	}

	if err := validateCiphertext(ciphertext); err != nil {
		return nil, err
	}

	if err := validatePlaintext(plaintext); err != nil {
		return nil, err
	}

	if !key.checkOperand(ciphertext) {
		return nil, ErrInvalidMismatch
	}

	ciphertextStr, err := ciphertext.GetCtStr()
	if err != nil {
		return nil, err
	}

	resultCStr := C.paillier_add_plain(key.Key,
		C.CString(ciphertextStr),
		C.CString(plaintext.String()))
	defer C.free(unsafe.Pointer(resultCStr))

	resultGoStr := C.GoString(resultCStr)

	return key.constructCiphertext(resultGoStr)
}

// SubCiphertext sets a new Ciphertext to the sum ciphertext1-ciphertext2 and return the result
func (key *PubKey) SubCiphertext(ciphertext1, ciphertext2 Ct) (Ct, error) {
	if err := validatePubKey(key); err != nil {
		return nil, err
	}

	if err := validateCiphertext(ciphertext1, ciphertext2); err != nil {
		return nil, err
	}

	if !key.checkOperand(ciphertext1, ciphertext2) {
		return nil, ErrInvalidMismatch
	}

	ciphertext1Str, err := ciphertext1.GetCtStr()
	if err != nil {
		return nil, err
	}

	ciphertext2Str, err := ciphertext2.GetCtStr()
	if err != nil {
		return nil, err
	}

	resultCStr := C.paillier_sub_cipher(key.Key,
		C.CString(ciphertext1Str),
		C.CString(ciphertext2Str))
	defer C.free(unsafe.Pointer(resultCStr))

	resultGoStr := C.GoString(resultCStr)

	return key.constructCiphertext(resultGoStr)
}

// SubPlaintext sets a new Ciphertext to the sum ciphertext1-ciphertext2 and return the result
func (key *PubKey) SubPlaintext(ciphertext Ct, plaintext *big.Int) (Ct, error) {
	if err := validatePubKey(key); err != nil {
		return nil, err
	}

	if err := validateCiphertext(ciphertext); err != nil {
		return nil, err
	}

	if err := validatePlaintext(plaintext); err != nil {
		return nil, err
	}

	if !key.checkOperand(ciphertext) {
		return nil, ErrInvalidMismatch
	}

	ciphertextStr, err := ciphertext.GetCtStr()
	if err != nil {
		return nil, err
	}

	resultCStr := C.paillier_sub_plain(key.Key,
		C.CString(ciphertextStr),
		C.CString(plaintext.String()))
	defer C.free(unsafe.Pointer(resultCStr))

	resultGoStr := C.GoString(resultCStr)

	return key.constructCiphertext(resultGoStr)
}

// NumMul sets a new Ciphertext to the sum ciphertext1-ciphertext2 and return the result
func (key *PubKey) NumMul(ciphertext Ct, plaintext *big.Int) (Ct, error) {
	if err := validatePubKey(key); err != nil {
		return nil, err
	}

	if err := validateCiphertext(ciphertext); err != nil {
		return nil, err
	}

	if err := validatePlaintext(plaintext); err != nil {
		return nil, err
	}

	if !key.checkOperand(ciphertext) {
		return nil, ErrInvalidMismatch
	}

	ciphertextStr, err := ciphertext.GetCtStr()
	if err != nil {
		return nil, err
	}

	resultCStr := C.paillier_num_mul(key.Key,
		C.CString(ciphertextStr),
		C.CString(plaintext.String()))
	defer C.free(unsafe.Pointer(resultCStr))

	resultGoStr := C.GoString(resultCStr)

	return key.constructCiphertext(resultGoStr)
}

func (key *PubKey) constructCiphertext(resultGoStr string) (Ct, error) {
	result, ok := new(big.Int).SetString(resultGoStr, 10)
	if !ok {
		return nil, ErrUnknown
	}

	Checksum, err := key.bindingCtPubKey(result.Bytes())
	if err != nil {
		return nil, err
	}

	ct := &Ciphertext{
		Ct:       result,
		Checksum: Checksum,
	}

	return ct, nil
}

func (key *PubKey) checkOperand(cts ...Ct) bool {
	for _, ct := range cts {
		if !key.ChecksumVerify(ct) {
			return false
		}
	}
	return true
}

// ChecksumVerify verifying public key ciphertext pairs
func (key *PubKey) ChecksumVerify(ct Ct) bool {
	if err := validatePubKey(key); err != nil {
		return false
	}

	if err := validateCiphertext(ct); err != nil {
		return false
	}

	keyBytes, err := key.Marshal()
	if err != nil {
		return false
	}

	ctBytes, err := ct.GetCtBytes()
	if err != nil {
		return false
	}

	currentChecksum, err := ct.GetChecksum()
	if err != nil {
		return false
	}

	checksum := sha256.Sum256(append(keyBytes, ctBytes...))
	return reflect.DeepEqual(checksum[:defaultChecksumSize], currentChecksum)
}

func (key *PubKey) bindingCtPubKey(ciphertext []byte) ([]byte, error) {
	pubKeyBytes, err := key.Marshal()
	if ciphertext == nil {
		return nil, ErrInvalidCiphertext
	}

	if err != nil {
		return nil, err
	}

	Checksum := sha256.Sum256(append(pubKeyBytes, ciphertext[:]...))
	return Checksum[:defaultChecksumSize], nil
}
