/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package pkcs11

import (
	bccrypto "chainmaker.org/chainmaker-go/common/crypto"
	bcecdsa "chainmaker.org/chainmaker-go/common/crypto/asym/ecdsa"
	bcrsa "chainmaker.org/chainmaker-go/common/crypto/asym/rsa"
	"chainmaker.org/chainmaker-go/common/crypto/hash"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"github.com/miekg/pkcs11"
	"github.com/tjfoc/gmsm/sm2"
	"math/big"
)

type P11Handle struct {
	ctx              *pkcs11.Ctx
	sessions         chan pkcs11.SessionHandle
	slot             uint
	sessionCacheSize int
	hash             string
}

func New(lib string, label string, password string, sessionCacheSize int, hash string) (*P11Handle, error) {
	if lib == "" {
		return nil, fmt.Errorf("PKCS11 error: empty library path")
	}

	if password == "" {
		return nil, fmt.Errorf("PKCS11 error: no pin provided")
	}

	if sessionCacheSize <= 0 {
		sessionCacheSize = 10
	}

	ctx := pkcs11.New(lib)
	if ctx == nil {
		return nil, fmt.Errorf("PKCS11 error: fail to initialize [%s]", lib)
	}

	ctx.Initialize()
	slots, err := ctx.GetSlotList(true)
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to get slot list [%v]", err)
	}

	found := false
	var slot uint
	slot, found = findSlot(ctx, slots, label)
	if !found {
		return nil, fmt.Errorf("PKCS11 error: fail to find token with label [%s]", label)
	}

	var session pkcs11.SessionHandle
	for i := 0; i < 10; i++ {
		session, err = ctx.OpenSession(slot, pkcs11.CKF_SERIAL_SESSION|pkcs11.CKF_RW_SESSION)
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to open session [%v]", err)
	}

	err = ctx.Login(session, pkcs11.CKU_USER, password)
	if err != nil {
		return nil, fmt.Errorf("PKCS11 error: fail to login session [%v]", err)
	}

	sessions := make(chan pkcs11.SessionHandle, sessionCacheSize)
	p11Handle := &P11Handle{
		ctx:              ctx,
		sessions:         sessions,
		slot:             slot,
		sessionCacheSize: sessionCacheSize,
		hash:             hash,
	}
	p11Handle.returnSession(session)

	return p11Handle, nil
}

func findSlot(ctx *pkcs11.Ctx, slots []uint, label string) (uint, bool) {
	var slot uint
	var found bool
	for _, s := range slots {
		info, err := ctx.GetTokenInfo(s)
		if err != nil {
			continue
		}
		if info.Label == label {
			found = true
			slot = s
			break
		}
	}
	return slot, found
}

func (p11 *P11Handle) getSession() (pkcs11.SessionHandle, error) {
	var session pkcs11.SessionHandle
	select {
	case session = <-p11.sessions:
		return session, nil
	default:
		var err error
		for i := 0; i < 10; i++ {
			session, err = p11.ctx.OpenSession(p11.slot, pkcs11.CKF_SERIAL_SESSION|pkcs11.CKF_RW_SESSION)
			if err == nil {
				return session, nil
			}
		}
		return 0, fmt.Errorf("PKCS11 error: fail to open session [%v]", err)
	}
}

func (p11 *P11Handle) returnSession(session pkcs11.SessionHandle) {
	select {
	case p11.sessions <- session:
		return
	default:
		p11.ctx.CloseSession(session)
		return
	}
}

func findPrivateKey(ctx *pkcs11.Ctx, session pkcs11.SessionHandle, ski []byte) (*pkcs11.ObjectHandle, error) {
	template := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PRIVATE_KEY),
		pkcs11.NewAttribute(pkcs11.CKA_ID, ski),
	}
	if err := ctx.FindObjectsInit(session, template); err != nil {
		return nil, err
	}

	objs, _, err := ctx.FindObjects(session, 1)
	if err != nil {
		return nil, err
	}
	if err = ctx.FindObjectsFinal(session); err != nil {
		return nil, err
	}

	if len(objs) <= 0 {
		return nil, fmt.Errorf("key not found [%s]", hex.Dump(ski))
	}

	return &objs[0], nil
}

// rsaPublicKey reflects the ASN.1 structure of a PKCS#1 public key.
type rsaPublicKeyASN struct {
	N *big.Int
	E int
}

func (p11 *P11Handle) GetPublicKeySKI(pk bccrypto.PublicKey) ([]byte, error) {
	if pk == nil {
		return nil, fmt.Errorf("public key is nil")
	}

	var pkBytes []byte
	var err error
	switch pk.(type) {
	case *bcecdsa.PublicKey:
		pubKey := pk.ToStandardKey()
		switch k := pubKey.(type) {
		case *ecdsa.PublicKey:
			pkBytes = elliptic.Marshal(k.Curve, k.X, k.Y)
		case *sm2.PublicKey:
			pkBytes = elliptic.Marshal(k.Curve, k.X, k.Y)
		default:
			return nil, fmt.Errorf("unknown public key type")
		}
	case *bcrsa.PublicKey:
		k := pk.ToStandardKey().(*rsa.PublicKey)
		pkBytes, err = asn1.Marshal(rsaPublicKeyASN{
			N: k.N,
			E: k.E,
		})
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown public key type")
	}

	return hash.GetByStrType(p11.hash, pkBytes)
}

func NewPrivateKey(p11 *P11Handle, pk bccrypto.PublicKey) (bccrypto.PrivateKey, error) {
	ski, err := p11.GetPublicKeySKI(pk)
	if err != nil {
		return nil, err
	}
	switch pk.(type) {
	case *bcecdsa.PublicKey:
		return &ECDSAPrivateKey{
			p11Ctx: p11,
			pubKey: pk,
			ski:    ski,
		}, nil
	case *bcrsa.PublicKey:
		return &RSAPrivateKey{
			p11Ctx: p11,
			pubKey: pk,
			ski:    ski,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported public key type: %v", pk.Type())
	}
}
