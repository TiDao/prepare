/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package ecdsa

import (
	"bytes"
	"chainmaker.org/chainmaker-go/common/crypto"
	"chainmaker.org/chainmaker-go/common/crypto/hash"
	crypto2 "crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/tjfoc/gmsm/sm2"
	"math/big"
)

type PrivateKey struct {
	K crypto2.PrivateKey
}

type Sig struct {
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`
}

func (sk *PrivateKey) Bytes() ([]byte, error) {
	if sk.K == nil {
		return nil, fmt.Errorf("private key is nil")
	}

	key := sk.K
	switch key := key.(type) {
	case *ecdsa.PrivateKey:
		if sk.Type() == crypto.ECC_Secp256k1 {
			rawKey := (*btcec.PrivateKey)(key).Serialize()
			return rawKey, nil
		} else {
			return x509.MarshalECPrivateKey(key)
		}
	case *sm2.PrivateKey:
		return MarshalPKCS8PrivateKey(key)
	}

	return nil, errors.New("unsupport private key")
}

func (sk *PrivateKey) PublicKey() crypto.PublicKey {
	key := sk.K
	switch key := key.(type) {
	case *ecdsa.PrivateKey:
		return &PublicKey{K: &key.PublicKey}
	case *sm2.PrivateKey:
		return &PublicKey{K: &key.PublicKey}
	}

	return nil
}

func (sk *PrivateKey) Sign(digest []byte) ([]byte, error) {
	var (
		r, s *big.Int
		err  error
	)

	key := sk.K
	switch key := key.(type) {
	case *ecdsa.PrivateKey:
		r, s, err = ecdsa.Sign(rand.Reader, key, digest[:])
	case *sm2.PrivateKey:
		r, s, err = sm2.Sign(key, digest[:])
	}

	if err != nil {
		return nil, err
	}

	return asn1.Marshal(Sig{R: r, S: s})
}

func (sk *PrivateKey) SignWithOpts(msg []byte, opts *crypto.SignOpts) ([]byte, error) {
	if opts == nil {
		return sk.Sign(msg)
	}
	if opts.Hash == crypto.HASH_TYPE_SM3 && sk.Type() == crypto.SM2 {
		skSM2, ok := sk.K.(*sm2.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("SM2 private key does not match the type it claims")
		}
		uid := opts.UID
		if len(uid) == 0 {
			uid = crypto.CRYPTO_DEFAULT_UID
		}

		r, s, err := sm2.Sm2Sign(skSM2, msg, []byte(uid))
		if err != nil {
			return nil, fmt.Errorf("fail to sign with SM2-SM3: [%v]", err)
		}

		return asn1.Marshal(Sig{R: r, S: s})
	} else {
		dgst, err := hash.Get(opts.Hash, msg)
		if err != nil {
			return nil, err
		}
		return sk.Sign(dgst)
	}
}

func (sk *PrivateKey) Type() crypto.KeyType {
	return sk.PublicKey().Type()
}

func (sk *PrivateKey) String() (string, error) {
	skDER, err := sk.Bytes()
	if err != nil {
		return "", err
	}

	if sk.Type() == crypto.ECC_Secp256k1 {
		return hex.EncodeToString(skDER), nil
	}

	block := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: skDER,
	}

	switch sk.K.(type) {
	case *sm2.PrivateKey:
		block.Type = "PRIVATE KEY"
	}

	buf := new(bytes.Buffer)
	if err = pem.Encode(buf, block); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (sk *PrivateKey) ToStandardKey() crypto2.PrivateKey {
	return sk.K
}

func New(keyType crypto.KeyType) (crypto.PrivateKey, error) {
	switch keyType {
	case crypto.ECC_Secp256k1:
		pri, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
		if err != nil {
			return nil, err
		}

		return &PrivateKey{K: pri}, nil
	case crypto.ECC_NISTP256:
		pri, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return nil, err
		}

		return &PrivateKey{K: pri}, nil
	case crypto.ECC_NISTP384:
		pri, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		if err != nil {
			return nil, err
		}

		return &PrivateKey{K: pri}, nil
	case crypto.ECC_NISTP521:
		pri, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
		if err != nil {
			return nil, err
		}

		return &PrivateKey{K: pri}, nil
	case crypto.SM2:
		pri, err := sm2.GenerateKey()
		if err != nil {
			return nil, err
		}

		return &PrivateKey{K: pri}, nil
	}
	return nil, fmt.Errorf("wrong curve option")
}
