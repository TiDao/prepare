/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package pkcs11

import (
	bccrypto "chainmaker.org/chainmaker-go/common/crypto"
	"crypto"
	"fmt"
	"github.com/miekg/pkcs11"
)

type RSAPrivateKey struct {
	p11Ctx *P11Handle
	pubKey bccrypto.PublicKey
	ski    []byte
}

func (sk *RSAPrivateKey) Type() bccrypto.KeyType {
	return sk.PublicKey().Type()
}

func (sk *RSAPrivateKey) Bytes() ([]byte, error) {
	return sk.ski, nil
}

func (sk *RSAPrivateKey) String() (string, error) {
	return string(sk.ski), nil
}

func (sk *RSAPrivateKey) PublicKey() bccrypto.PublicKey {
	return sk.pubKey
}

func (sk *RSAPrivateKey) Sign(data []byte) ([]byte, error) {
	session, err := sk.p11Ctx.getSession()
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to get session [%s]", err)
	}
	defer sk.p11Ctx.returnSession(session)

	privateKey, err := findPrivateKey(sk.p11Ctx.ctx, session, sk.ski)
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to find private key [%s]", err)
	}

	mech := uint(pkcs11.CKM_SHA256_RSA_PKCS)
	err = sk.p11Ctx.ctx.SignInit(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(mech, nil)}, *privateKey)
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to initiate signing procedure [%s]", err)
	}

	sig, err := sk.p11Ctx.ctx.Sign(session, data)
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to sign [%s]", err)
	}

	return sig, nil
}

func (sk *RSAPrivateKey) SignWithOpts(msg []byte, opts *bccrypto.SignOpts) ([]byte, error) {
	if opts == nil {
		return sk.Sign(msg)
	}

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
	switch opts.Hash {
	case bccrypto.HASH_TYPE_SHA256:
		mech = pkcs11.CKM_SHA256_RSA_PKCS
	case bccrypto.HASH_TYPE_SHA3_256:
		mech = pkcs11.CKM_SHA3_256_RSA_PKCS
	default:
		return nil, fmt.Errorf("PKCS11 error: unsupported hash type [%v]", opts.Hash)
	}

	err = sk.p11Ctx.ctx.SignInit(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(mech, nil)}, *privateKey)
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to initiate signing procedure [%s]", err)
	}

	sig, err := sk.p11Ctx.ctx.Sign(session, msg)
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to sign [%s]", err)
	}

	return sig, nil
}

func (sk *RSAPrivateKey) ToStandardKey() crypto.PrivateKey {
	return sk
}
