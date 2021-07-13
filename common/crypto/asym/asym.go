/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package asym

import "C"
import (
	"chainmaker.org/chainmaker-go/common/crypto"
	"chainmaker.org/chainmaker-go/common/crypto/asym/ecdsa"
	"chainmaker.org/chainmaker-go/common/crypto/asym/rsa"
	crypto2 "crypto"
	ecdsa2 "crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	rsa2 "crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/tjfoc/gmsm/sm2"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"strings"
)

const pemBegin = "-----BEGIN"

// 生成签名公私钥对
func GenerateKeyPair(keyType crypto.KeyType) (crypto.PrivateKey, error) {
	switch keyType {
	case crypto.SM2, crypto.ECC_NISTP256, crypto.ECC_NISTP384, crypto.ECC_NISTP521, crypto.ECC_Secp256k1:
		return ecdsa.New(keyType)
	case crypto.RSA512, crypto.RSA1024, crypto.RSA2048, crypto.RSA3072:
		return rsa.New(keyType)
	case crypto.ECC_Ed25519:
		return nil, fmt.Errorf("unsupport signature algorithm")
	default:
		return nil, fmt.Errorf("wrong signature algorithm type")
	}
}

func GenerateKeyPairBytes(keyType crypto.KeyType) (sk, pk []byte, err error) {
	var priv crypto.PrivateKey
	if priv, err = GenerateKeyPair(keyType); err != nil {
		return
	}

	if sk, err = priv.Bytes(); err != nil {
		return
	}

	if pk, err = priv.PublicKey().Bytes(); err != nil {
		return
	}

	return sk, pk, nil
}

func GenerateKeyPairPEM(keyType crypto.KeyType) (sk string, pk string, err error) {
	var priv crypto.PrivateKey
	if priv, err = GenerateKeyPair(keyType); err != nil {
		return "", "", err
	}

	// Serialization for bitcoin signature key: encode ECC numbers with hex
	if sk, err = priv.String(); err != nil {
		return "", "", err
	}

	// Serialization for bitcoin signature key: encode ECC numbers with hex
	if pk, err = priv.PublicKey().String(); err != nil {
		return "", "", err
	}

	return sk, pk, nil
}

// Generate public-private key pair for encryption
func GenerateEncKeyPair(keyType crypto.KeyType) (crypto.DecryptKey, error) {
	switch keyType {
	case crypto.SM2:
		key, err := ecdsa.New(keyType)
		if err != nil {
			return nil, err
		}
		return key.(crypto.DecryptKey), nil
	case crypto.RSA512, crypto.RSA1024, crypto.RSA2048, crypto.RSA3072:
		return rsa.NewDecryptionKey(keyType)
	default:
		return nil, fmt.Errorf("unsupported encryption algorithm type")
	}
}

// ParsePrivateKey parse bytes to a private key.
func ParsePrivateKey(der []byte) (crypto2.PrivateKey, error) {
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}

	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return key, nil
	}

	if key, err := sm2.ParsePKCS8UnecryptedPrivateKey(der); err == nil {
		return key, nil
	}

	// Serialization for bitcoin signature key: encode ECC numbers with hex
	Secp256k1Key, _ := btcec.PrivKeyFromBytes(btcec.S256(), der)
	key := Secp256k1Key.ToECDSA()
	return key, nil
}

func ParsePublicKey(der []byte) (crypto2.PublicKey, error) {
	if key, err := x509.ParsePKCS1PublicKey(der); err == nil {
		return key, nil
	}

	if key, err := x509.ParsePKIXPublicKey(der); err == nil {
		return key, nil
	}

	if key, err := sm2.ParseSm2PublicKey(der); err == nil {
		return key, nil
	}

	// Serialization for bitcoin signature key: encode ECC numbers with hex
	if key, err := btcec.ParsePubKey(der, btcec.S256()); err == nil {
		return key.ToECDSA(), nil
	}

	return nil, errors.New("failed to parse public key")
}

func PrivateKeyFromDER(der []byte) (crypto.PrivateKey, error) {
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return &rsa.PrivateKey{K: key}, nil
	}

	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return &ecdsa.PrivateKey{K: key}, nil
	}

	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		return &ecdsa.PrivateKey{K: key}, nil
	}

	if key, err := sm2.ParsePKCS8UnecryptedPrivateKey(der); err == nil {
		return &ecdsa.PrivateKey{K: key}, nil
	}

	Secp256k1Key, _ := btcec.PrivKeyFromBytes(btcec.S256(), der)
	key := Secp256k1Key.ToECDSA()
	return &ecdsa.PrivateKey{K: key}, nil
}

func PublicKeyFromDER(der []byte) (crypto.PublicKey, error) {
	if key, err := x509.ParsePKCS1PublicKey(der); err == nil {
		return &rsa.PublicKey{K: key}, nil
	}

	if key, err := x509.ParsePKIXPublicKey(der); err == nil {
		switch key.(type) {
		case *rsa2.PublicKey:
			return &rsa.PublicKey{K: key.(*rsa2.PublicKey)}, nil
		case *ecdsa2.PublicKey:
			return &ecdsa.PublicKey{K: key.(*ecdsa2.PublicKey)}, nil
		case *sm2.PublicKey:
			return &ecdsa.PublicKey{K: key.(*sm2.PublicKey)}, nil
		default:
			return nil, fmt.Errorf("unsupported public key type")
		}
	}

	if key, err := sm2.ParseSm2PublicKey(der); err == nil {
		return &ecdsa.PublicKey{K: key}, nil
	}

	if key, err := btcec.ParsePubKey(der, btcec.S256()); err == nil {
		return &ecdsa.PublicKey{K: key.ToECDSA()}, nil
	}

	return nil, errors.New("failed to parse public key")
}

func PrivateKeyFromPEM(raw []byte, pwd []byte) (crypto.PrivateKey, error) {
	var err error

	if len(raw) <= 0 {
		return nil, errors.New("PEM is nil")
	}

	if !strings.Contains(string(raw), pemBegin) {
		keyBytes, err := hex.DecodeString(string(raw))
		if err != nil {
			return nil, fmt.Errorf("fail to decode Secp256k1 public key: [%v]", err)
		}
		return PrivateKeyFromDER(keyBytes)
	}

	block, _ := pem.Decode(raw)
	if block == nil {
		return PrivateKeyFromDER(raw)
	}

	plain := block.Bytes
	if x509.IsEncryptedPEMBlock(block) {
		if len(pwd) <= 0 {
			return nil, errors.New("missing password for encrypted PEM")
		}

		plain, err = x509.DecryptPEMBlock(block, pwd)
		if err != nil {
			return nil, fmt.Errorf("fail to decrypt PEM: [%s]", err)
		}
	}

	return PrivateKeyFromDER(plain)
}

func PublicKeyFromPEM(raw []byte) (crypto.PublicKey, error) {
	if len(raw) <= 0 {
		return nil, errors.New("PEM is nil")
	}

	if !strings.Contains(string(raw), pemBegin) {
		keyBytes, err := hex.DecodeString(string(raw))
		if err != nil {
			return nil, fmt.Errorf("fail to decode Secp256k1 public key: [%v]", err)
		}
		return PublicKeyFromDER(keyBytes)
	}

	block, _ := pem.Decode(raw)
	if block == nil {
		return PublicKeyFromDER(raw)
	}

	return PublicKeyFromDER(block.Bytes)
}

func Sign(sk interface{}, data []byte) ([]byte, error) {
	var (
		err        error
		r, s       *big.Int
		keyBytes   []byte
		signedData []byte
	)

	switch sk.(type) {
	case string:
		if strings.Contains(sk.(string), pemBegin) {
			der, _ := pem.Decode([]byte(sk.(string)))
			keyBytes = der.Bytes
		} else {
			keyBytes, err = hex.DecodeString(sk.(string))
			if err != nil {
				return nil, err
			}
		}
	case []byte:
		keyBytes = sk.([]byte)
	default:
		return nil, errors.New("invalid sk format")
	}

	key, err := ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}

	switch key := key.(type) {
	case *ecdsa2.PrivateKey:
		if r, s, err = ecdsa2.Sign(rand.Reader, key, data); err != nil {
			return nil, err
		}
		return asn1.Marshal(ecdsa.Sig{R: r, S: s})

	case *sm2.PrivateKey:
		if r, s, err = sm2.Sign(key, data); err != nil {
			return nil, err
		}
		return asn1.Marshal(ecdsa.Sig{R: r, S: s})

	case *rsa2.PrivateKey:
		hashed := sha256.Sum256(data)
		if signedData, err = rsa2.SignPKCS1v15(rand.Reader, key, crypto2.SHA256, hashed[:]); err != nil {
			return nil, err
		}
		return signedData, nil
	default:
		return nil, fmt.Errorf("fail to sign: unsupported algorithm")
	}
}

func Verify(pk interface{}, data, sig []byte) (bool, error) {
	if sig == nil {
		return false, fmt.Errorf("nil signature")
	}

	var (
		err      error
		keyBytes []byte
	)

	keyBytes, err = loadKeyBytesWithPK(pk)
	if err != nil {
		return false, err
	}

	key, err := ParsePublicKey(keyBytes)
	if err != nil {
		return false, err
	}

	if err := verifyDataSignWithPubKey(key, data, sig); err != nil {
		return false, err
	}

	return true, nil
}

func loadKeyBytesWithPK(pk interface{}) ([]byte, error) {
	var keyBytes []byte
	switch pk.(type) {
	case string:
		if strings.Contains(pk.(string), pemBegin) {
			der, _ := pem.Decode([]byte(pk.(string)))
			keyBytes = der.Bytes
		} else {
			var err error
			keyBytes, err = hex.DecodeString(pk.(string))
			if err != nil {
				return nil, err
			}
		}
	case []byte:
		keyBytes = pk.([]byte)
	default:
		return nil, errors.New("invalid sk format")
	}
	return keyBytes, nil
}

func verifyDataSignWithPubKey(key crypto2.PublicKey, data, sig []byte) error {
	switch key := key.(type) {
	case *ecdsa2.PublicKey:
		sigStruct := &ecdsa.Sig{}
		if _, err := asn1.Unmarshal(sig, sigStruct); err != nil {
			return err
		}

		if !ecdsa2.Verify(key, data, sigStruct.R, sigStruct.S) {
			return fmt.Errorf("string invalid ecdsa signature")
		}
	case *sm2.PublicKey:
		sigStruct := &ecdsa.Sig{}
		if _, err := asn1.Unmarshal(sig, sigStruct); err != nil {
			return err
		}

		if !sm2.Verify(key, data, sigStruct.R, sigStruct.S) {
			return fmt.Errorf("invalid sm2 signature")
		}
	case *rsa2.PublicKey:
		hashed := sha256.Sum256(data)
		err := rsa2.VerifyPKCS1v15(key, crypto2.SHA256, hashed[:], sig)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("fail to verify: unsupported algorithm")
	}
	return nil
}

func WriteFile(keyType crypto.KeyType, filePath string) error {
	sk, pk, err := GenerateKeyPairPEM(keyType)
	if err != nil {
		return err
	}

	skPath := filepath.Join(filePath, "node.key")
	if err = ioutil.WriteFile(skPath, []byte(sk), 0644); err != nil {
		return fmt.Errorf("save sk failed, %s", err)
	}

	pkPath := filepath.Join(filePath, "node.crt")
	if err = ioutil.WriteFile(pkPath, []byte(pk), 0644); err != nil {
		return fmt.Errorf("save pk failed, %s", err)
	}

	return nil
}

func ParseSM2PublicKey(asn1Data []byte) (*sm2.PublicKey, error) {
	if asn1Data == nil {
		return nil, errors.New("fail to unmarshal public key: public key is empty")
	}

	x, y := elliptic.Unmarshal(sm2.P256Sm2(), asn1Data)
	if x == nil {
		return nil, errors.New("x509: failed to unmarshal elliptic curve point")
	}

	pk := sm2.PublicKey{
		Curve: sm2.P256Sm2(),
		X:     x,
		Y:     y,
	}
	return &pk, nil
}
