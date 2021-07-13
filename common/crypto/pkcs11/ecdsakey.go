/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package pkcs11

import (
	bccrypto "chainmaker.org/chainmaker-go/common/crypto"
	bcecdsa "chainmaker.org/chainmaker-go/common/crypto/asym/ecdsa"
	"chainmaker.org/chainmaker-go/common/crypto/hash"
	"crypto"
	"encoding/asn1"
	"fmt"
	"github.com/miekg/pkcs11"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"math/big"
)

type ECDSAPrivateKey struct {
	p11Ctx *P11Handle
	pubKey bccrypto.PublicKey
	ski    []byte
}

func (sk *ECDSAPrivateKey) Type() bccrypto.KeyType {
	return sk.PublicKey().Type()
}

func (sk *ECDSAPrivateKey) Bytes() ([]byte, error) {
	return sk.ski, nil
}

func (sk *ECDSAPrivateKey) String() (string, error) {
	return string(sk.ski), nil
}

func (sk *ECDSAPrivateKey) PublicKey() bccrypto.PublicKey {
	return sk.pubKey
}

func (sk *ECDSAPrivateKey) Sign(data []byte) ([]byte, error) {
	session, err := sk.p11Ctx.getSession()
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to get session [%s]", err)
	}
	defer sk.p11Ctx.returnSession(session)

	privateKey, err := findPrivateKey(sk.p11Ctx.ctx, session, sk.ski)
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to find private key [%s]", err)
	}

	var mech uint
	switch sk.Type() {
	case bccrypto.SM2:
		// test needed to verify correctness
		mech = CKM_SM2_SIGN_NO_DER
	case bccrypto.ECC_Secp256k1, bccrypto.ECC_NISTP256, bccrypto.ECC_NISTP384, bccrypto.ECC_NISTP521:
		mech = pkcs11.CKM_ECDSA
	}
	err = sk.p11Ctx.ctx.SignInit(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(mech, nil)}, *privateKey)
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to initiate signing procedure [%s]", err)
	}

	sig, err := sk.p11Ctx.ctx.Sign(session, data)
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to sign [%s]", err)
	}

	sigLen := len(sig)
	R := new(big.Int)
	S := new(big.Int)
	R.SetBytes(sig[0 : sigLen/2])
	S.SetBytes(sig[sigLen/2:])

	return asn1.Marshal(bcecdsa.Sig{R: R, S: S})
}

func (sk *ECDSAPrivateKey) SignWithOpts(msg []byte, opts *bccrypto.SignOpts) ([]byte, error) {
	if opts == nil {
		return sk.Sign(msg)
	}
	if opts.Hash == bccrypto.HASH_TYPE_SM3 && sk.Type() == bccrypto.SM2 {
		pkSM2, ok := sk.PublicKey().ToStandardKey().(*sm2.PublicKey)
		if !ok {
			return nil, fmt.Errorf("SM2 private key does not match the type it claims")
		}
		uid := opts.UID
		if len(uid) == 0 {
			uid = bccrypto.CRYPTO_DEFAULT_UID
		}

		za, err := sm2.ZA(pkSM2, []byte(uid))
		if err != nil {
			return nil, fmt.Errorf("PKCS11 error: fail to create SM3 digest for msg [%v]", err)
		}
		e := sm3.New()
		e.Write(za)
		e.Write(msg)
		dgst := e.Sum(nil)[:32]

		return sk.Sign(dgst)
	} else {
		dgst, err := hash.Get(opts.Hash, msg)
		if err != nil {
			return nil, err
		}
		return sk.Sign(dgst)
	}
}

func (sk *ECDSAPrivateKey) ToStandardKey() crypto.PrivateKey {
	return sk
}
