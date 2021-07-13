/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package ecdsa

import (
	"chainmaker.org/chainmaker-go/common/crypto"
	"crypto/sha256"
	"encoding/pem"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var msg = "js"

func TestSM2(t *testing.T) {
	h := sha256.Sum256([]byte(msg))
	priv, err := New(crypto.SM2)
	require.Nil(t, err)

	privDER, err := priv.Bytes()
	require.Nil(t, err)

	block := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privDER,
	}

	err = pem.Encode(os.Stdout, block)
	require.Nil(t, err)

	buf, err := priv.String()
	require.Nil(t, err)
	fmt.Println(buf)

	sign, err := priv.Sign(h[:])
	require.Nil(t, err)
	require.NotEqual(t, nil, sign)
	pub := priv.PublicKey()
	buf, err = pub.String()
	require.Nil(t, err)
	fmt.Println(buf)

	b, err := pub.Verify(h[:], sign)
	require.Nil(t, err)
	require.True(t, b)
}

func TestP256(t *testing.T) {
	h := sha256.Sum256([]byte(msg))
	priv, err := New(crypto.ECC_NISTP256)
	require.Nil(t, err)

	buf, err := priv.String()
	require.Nil(t, err)
	fmt.Println(buf)

	sign, err := priv.Sign(h[:])
	require.Nil(t, err)
	require.NotEqual(t, nil, sign)

	pub := priv.PublicKey()
	buf, err = pub.String()
	require.Nil(t, err)
	fmt.Println(buf)

	b, err := pub.Verify(h[:], sign)
	require.Nil(t, err)
	require.True(t, b)
}

func TestSecp256k1(t *testing.T) {
	h := sha256.Sum256([]byte(msg))
	priv, err := New(crypto.ECC_Secp256k1)
	require.Nil(t, err)

	buf, err := priv.String()
	require.Nil(t, err)
	fmt.Println(buf)

	sign, err := priv.Sign(h[:])
	require.Nil(t, err)
	require.NotEqual(t, nil, sign)

	pub := priv.PublicKey()
	buf, err = pub.String()
	require.Nil(t, err)
	fmt.Println(buf)

	b, err := pub.Verify(h[:], sign)
	require.Nil(t, err)
	require.True(t, b)
}
