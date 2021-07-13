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
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/tjfoc/gmsm/sm2"
)

type PublicKey struct {
	K crypto2.PublicKey
}

func (pk *PublicKey) Bytes() ([]byte, error) {
	if pk.K == nil {
		return nil, fmt.Errorf("public key is nil")
	}

	key := pk.K
	switch key := key.(type) {
	case *ecdsa.PublicKey:
		if pk.Type() == crypto.ECC_Secp256k1 {
			rawKey := (*btcec.PublicKey)(key).SerializeCompressed()
			return rawKey, nil
		} else {
			return x509.MarshalPKIXPublicKey(key)
		}
	case *sm2.PublicKey:
		return sm2.MarshalPKIXPublicKey(key)
	}

	return nil, errors.New("unsupport public key")
}

func (pk *PublicKey) Verify(digest []byte, sig []byte) (bool, error) {
	if sig == nil {
		return false, fmt.Errorf("nil signature")
	}

	sigStruct := &Sig{}
	if _, err := asn1.Unmarshal(sig, sigStruct); err != nil {
		return false, fmt.Errorf("fail to decode signature: [%v]", err)
	}

	key := pk.K
	switch key := key.(type) {
	case *ecdsa.PublicKey:
		if !ecdsa.Verify(key, digest, sigStruct.R, sigStruct.S) {
			return false, fmt.Errorf("struct invalid ecdsa signature")
		}

	case *sm2.PublicKey:
		if !sm2.Verify(key, digest, sigStruct.R, sigStruct.S) {
			return false, fmt.Errorf("invalid sm2 signature")
		}
	}

	return true, nil
}

func (pk *PublicKey) VerifyWithOpts(msg []byte, sig []byte, opts *crypto.SignOpts) (bool, error) {
	if opts == nil {
		return pk.Verify(msg, sig)
	}
	if opts.Hash == crypto.HASH_TYPE_SM3 && pk.Type() == crypto.SM2 {
		pkSM2, ok := pk.K.(*sm2.PublicKey)
		if !ok {
			return false, fmt.Errorf("SM2 public key does not match the type it claims")
		}
		uid := opts.UID
		if len(uid) == 0 {
			uid = crypto.CRYPTO_DEFAULT_UID
		}

		if sig == nil {
			return false, fmt.Errorf("nil signature")
		}

		sigStruct := &Sig{}
		if _, err := asn1.Unmarshal(sig, sigStruct); err != nil {
			return false, fmt.Errorf("fail to decode signature: [%v]", err)
		}

		return sm2.Sm2Verify(pkSM2, msg, []byte(uid), sigStruct.R, sigStruct.S), nil
	} else {
		dgst, err := hash.Get(opts.Hash, msg)
		if err != nil {
			return false, err
		}
		return pk.Verify(dgst, sig)
	}
}

func (pk *PublicKey) Type() crypto.KeyType {
	if pk.K != nil {
		key := pk.K
		switch key := key.(type) {
		case *ecdsa.PublicKey:
			switch key.Curve {
			case elliptic.P256():
				return crypto.ECC_NISTP256
			case elliptic.P384():
				return crypto.ECC_NISTP384
			case elliptic.P521():
				return crypto.ECC_NISTP521
			case btcec.S256():
				return crypto.ECC_Secp256k1
			}
		case *sm2.PublicKey:
			switch key.Curve {
			case sm2.P256Sm2():
				return crypto.SM2
			}
		}
	}

	return -1
}

func (pk *PublicKey) String() (string, error) {

	pkDER, err := pk.Bytes()
	if err != nil {
		return "", err
	}

	if pk.Type() == crypto.ECC_Secp256k1 {
		return hex.EncodeToString(pkDER), nil
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pkDER,
	}

	buf := new(bytes.Buffer)
	if err = pem.Encode(buf, block); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (pk *PublicKey) ToStandardKey() crypto2.PublicKey {
	return pk.K
}
