//+build paillier

/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package paillier

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	// plaintext
	p1          *big.Int
	p10         *big.Int
	p_10        *big.Int
	p_15        *big.Int
	p20         *big.Int
	p_20        *big.Int
	p0          *big.Int
	pMax        *big.Int
	pMaxPlusOne *big.Int
	pMin        *big.Int
	pMixSubOne  *big.Int
	bigOne      *big.Int
	bigTwo      *big.Int
	bigNegOne   *big.Int
	bigZero     *big.Int

	// ciphertext
	c10          Ct
	c_10         Ct
	c_15         Ct
	c20          Ct
	c0           Ct
	cMax         Ct
	cMin         Ct
	cBigOne      Ct
	cBigTwo      Ct
	cBigNegOne   Ct
	cBigZero     Ct
	cBigMax      Ct
	c10_testKey2 Ct
	// ciphertext
	cEmpty = ""

	// key
	testKey             Prv
	pubKey              Pub
	testKey2            Prv
	prvKeyFromUnmarshal Prv
	pubKeyFromUnmarshal Pub
	// keyBytes
	prvBytes  []byte
	prv2Bytes []byte
	pubBytes  []byte
)

func TestPaillier(t *testing.T) {
	fmt.Printf("=================================init=================================\n")
	testInitPlaintext(t)

	// Construct Key
	fmt.Printf("=============================Gen Key & Marshal & Unmarshal=============================\n")
	testGenKey(t)

	fmt.Printf("====================================Encrypt2Ciphertext Marshal & Unmarshal====================================\n")
	testEncrypt(t)

	fmt.Printf("====================================Decrypt test====================================\n")
	testDecrypt(t)

	fmt.Printf("====================================AddCiphertext test====================================\n")
	testAddCiphertext(t)

	fmt.Printf("====================================AddPlaintext test====================================\n")
	testAddPlaintext(t)

	fmt.Printf("====================================SubCiphertext test====================================\n")
	testSubCiphertext(t)

	fmt.Printf("====================================SubPlaintext test====================================\n")
	testSubPlaintext(t)

	fmt.Printf("====================================NumMul test====================================\n")
	testNumMul(t)

	fmt.Printf("====================================Boundary test====================================\n")
	testBoundary(t)

	fmt.Printf("===================================== Bug test ====================================\n")
	testBug(t)

}

func testInitPlaintext(t *testing.T) {
	p10, _ = new(big.Int).SetString("10", 10)
	p_10, _ = new(big.Int).SetString("-10", 10)
	p_15, _ = new(big.Int).SetString("-15", 10)
	p20, _ = new(big.Int).SetString("20", 10)
	p_20, _ = new(big.Int).SetString("-20", 10)
	p0, _ = new(big.Int).SetString("0", 10)
	pMax, _ = new(big.Int).SetString("9223372036854775807", 10)
	pMaxPlusOne, _ = new(big.Int).SetString("9223372036854775808", 10)
	pMin, _ = new(big.Int).SetString("-9223372036854775808", 10)
	pMixSubOne, _ = new(big.Int).SetString("-9223372036854775809", 10)
	p1, _ = new(big.Int).SetString("1", 10)
	bigOne, _ = new(big.Int).SetString("1", 10)
	bigTwo, _ = new(big.Int).SetString("2", 10)
	bigNegOne, _ = new(big.Int).SetString("-1", 10)
	bigZero, _ = new(big.Int).SetString("0", 10)
}

func testGenKey(t *testing.T) {
	var err error
	keyGenerator := Helper().NewKeyGenerator()
	keyGenerator2 := Helper().NewKeyGenerator()

	testKey, err = keyGenerator.GenKey()
	require.Nil(t, err)
	pubKey, err = testKey.GetPubKey()
	require.Nil(t, err)

	testKey2, err = keyGenerator.GenKey()
	require.Nil(t, err)

	_, err = Helper().NewPrvKey().GetPubKey()
	require.EqualError(t, err, ErrInvalidPrivateKey.Error())

	pubBytes, err = pubKey.Marshal()
	require.Nil(t, err)

	_, err = Helper().NewPubKey().Marshal()
	require.EqualError(t, err, ErrInvalidPublicKey.Error())

	fmt.Printf("pubBytes -> %s\n", pubBytes)
	prvBytes, err = testKey.Marshal()
	require.Nil(t, err)
	fmt.Printf("prvBytes -> %s\n", prvBytes)

	_, err = Helper().NewPrvKey().Marshal()
	require.EqualError(t, err, ErrInvalidPrivateKey.Error())

	testKey2, err = keyGenerator2.GenKey()
	require.Nil(t, err)
	prv2Bytes, err = testKey2.Marshal()
	require.Nil(t, err)

	pubKeyFromUnmarshal = Helper().NewPubKey()
	err = pubKeyFromUnmarshal.Unmarshal(pubBytes)
	require.Nil(t, err)

	err = pubKeyFromUnmarshal.Unmarshal(nil)
	require.EqualError(t, err, ErrInvalidPublicKey.Error())

	err = Helper().NewPubKey().Unmarshal(pubBytes)
	require.Nil(t, err)

	prvKeyFromUnmarshal = Helper().NewPrvKey()

	err = prvKeyFromUnmarshal.Unmarshal(prvBytes)
	require.Nil(t, err)

	err = Helper().NewPrvKey().Unmarshal(nil)
	require.EqualError(t, err, ErrInvalidPrivateKey.Error())

	err = Helper().NewPrvKey().Unmarshal([]byte("啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊aaaaaaaa"))
	require.EqualError(t, err, ErrInvalidPublicKey.Error())

	err = Helper().NewPrvKey().Unmarshal(prvBytes)
	require.Nil(t, err)
	//require.EqualError(t, err, ErrInvalidPrivateKey.Error())

	new(big.Int).Bytes()
	new(big.Int).SetBytes(nil)
}

func testEncrypt(t *testing.T) {
	var err error
	// Encrypt 10, -15, 20, 0
	c10 = Helper().NewCiphertext()
	c10, err = testKey.Encrypt(p10)
	require.Nil(t, err)
	c_10, err = testKey.Encrypt(p_10)
	require.Nil(t, err)
	c_15, err = testKey.Encrypt(p_15)
	require.Nil(t, err)
	c20, err = testKey.Encrypt(p20)
	require.Nil(t, err)
	c0, err = testKey.Encrypt(p0)
	require.Nil(t, err)
	cMax, err = testKey.Encrypt(pMax)
	require.Nil(t, err)
	cMin, err = testKey.Encrypt(pMin)
	// Encrypt BigInt
	var bigZero *big.Int
	var bigOne *big.Int
	var bigTwo *big.Int
	var bigNegOne *big.Int
	var bigMax *big.Int
	bigZero, _ = new(big.Int).SetString("0", 10)
	bigOne, _ = new(big.Int).SetString("1", 10)
	bigTwo, _ = new(big.Int).SetString("2", 10)
	bigNegOne, _ = new(big.Int).SetString("-1", 10)
	bigMax, _ = new(big.Int).SetString("9223372036854775807", 10)
	cBigZero, err = testKey.Encrypt(bigZero)
	require.Nil(t, err)
	cBigOne, err = testKey.Encrypt(bigOne)
	require.Nil(t, err)
	cBigTwo, err = testKey.Encrypt(bigTwo)
	require.Nil(t, err)
	cBigNegOne, err = testKey.Encrypt(bigNegOne)
	require.Nil(t, err)
	cBigMax, err = testKey.Encrypt(bigMax)

	c10Bytes, err := c10.Marshal()
	require.Nil(t, err)

	c_ := Helper().NewCiphertext()
	_, err = c_.Marshal()
	require.EqualError(t, err, "paillier: invalid ciphertext")

	err = c10.Unmarshal(c10Bytes)
	require.Nil(t, err)

	err = c10.Unmarshal(nil)
	require.EqualError(t, err, ErrInvalidCiphertext.Error())

	_, err = Helper().NewCiphertext().GetChecksum()
	require.EqualError(t, err, ErrInvalidCiphertext.Error())

	_, err = Helper().NewCiphertext().GetCtBytes()
	require.EqualError(t, err, ErrInvalidCiphertext.Error())

	_, err = Helper().NewCiphertext().GetCtStr()
	require.EqualError(t, err, ErrInvalidCiphertext.Error())

	cBytes, err := c10.Marshal()
	require.Nil(t, err)
	err = Helper().NewCiphertext().Unmarshal(cBytes)
	require.Nil(t, err)

	//var pubKeyBad Pub
	tests := []struct {
		pubKey  Pub
		pt      *big.Int
		wantErr string
	}{
		{
			pubKey:  Helper().NewPubKey(),
			pt:      p10,
			wantErr: ErrInvalidPublicKey.Error(),
		},
		{
			pubKey:  pubKey,
			pt:      nil,
			wantErr: ErrInvalidPlaintext.Error(),
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			_, err = tt.pubKey.Encrypt(tt.pt)
			require.EqualError(t, err, tt.wantErr)
		})
	}
}

func testDecrypt(t *testing.T) {
	// Decrypt 10, -15, 20, 0
	d10, err := prvKeyFromUnmarshal.Decrypt(c10)
	require.Nil(t, err)
	d_15, err := prvKeyFromUnmarshal.Decrypt(c_15)
	require.Nil(t, err)
	d20, err := prvKeyFromUnmarshal.Decrypt(c20)
	require.Nil(t, err)
	d0, err := prvKeyFromUnmarshal.Decrypt(c0)
	require.Nil(t, err)
	dMax, err := prvKeyFromUnmarshal.Decrypt(cMax)
	require.Nil(t, err)
	dMin, err := prvKeyFromUnmarshal.Decrypt(cMin)
	require.Nil(t, err)
	fmt.Printf("[ decrypt ciphertext 10 ] %d\n", d10)
	fmt.Printf("[ decrypt ciphertext _15 ] %d\n", d_15)
	fmt.Printf("[ decrypt ciphertext 20 ] %d\n", d20)
	fmt.Printf("[ decrypt ciphertext 0 ] %d\n", d0)
	fmt.Printf("[ decrypt ciphertext Max ] %d\n", dMax)
	fmt.Printf("[ decrypt ciphertext Min ] %d\n", dMin)

	// Decrypt big.Int
	dBigOne, err := prvKeyFromUnmarshal.Decrypt(cBigOne)
	require.Nil(t, err)
	dBigTwo, err := prvKeyFromUnmarshal.Decrypt(cBigTwo)
	require.Nil(t, err)
	fmt.Printf("[ decrytp ciphertext bigInt 1 ] %d\n", dBigOne)
	fmt.Printf("[ decrytp ciphertext bigInt 2 ] %d\n", dBigTwo)

	tests := []struct {
		prvKey  Prv
		ct      Ct
		wantErr string
	}{
		{
			prvKey:  Helper().NewPrvKey(),
			ct:      c10,
			wantErr: ErrInvalidPrivateKey.Error(),
		},
		{
			prvKey:  testKey,
			ct:      nil,
			wantErr: ErrInvalidCiphertext.Error(),
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			_, err = tt.prvKey.Decrypt(tt.ct)
			require.EqualError(t, err, tt.wantErr)
		})
	}
}

func testAddCiphertext(t *testing.T) {
	pubKey2FromUnmarshal := Helper().NewPubKey()

	rAddC10C20, err := testKey.AddCiphertext(c10, c20)
	require.Nil(t, err)

	checkPrvKey := Helper().NewPrvKey()
	err = checkPrvKey.Unmarshal(prvBytes)
	require.Nil(t, err)

	err = checkPrvKey.Unmarshal(prvBytes)
	require.Nil(t, err)

	err = checkPrvKey.Unmarshal(prv2Bytes)
	require.Nil(t, err)

	dRAddC10C20, err := testKey.Decrypt(rAddC10C20)
	require.Nil(t, err)
	fmt.Printf("[ c10 + c20] %d\n", dRAddC10C20)

	// operand verify
	rAddC10C_15, err := testKey.AddCiphertext(c10, c_15)
	require.Nil(t, err)

	ok := pubKey2FromUnmarshal.ChecksumVerify(rAddC10C_15)
	require.False(t, ok)

	ok = testKey.ChecksumVerify(nil)
	require.False(t, ok)

	dRAddC10C_15, err := testKey.Decrypt(rAddC10C_15)
	require.Nil(t, err)
	fmt.Printf("[ c10 + c(-15)] %d\n", dRAddC10C_15)

	// AddCiphertext
	rAddCBig, err := testKey.AddCiphertext(cBigOne, cBigTwo)
	require.Nil(t, err)
	dRAddCBig, err := testKey.Decrypt(rAddCBig)
	require.Nil(t, err)
	fmt.Printf("[ c1 + c2 ] %s\n", dRAddCBig.String())

	// 0
	// AddCiphertext
	rB3, err := testKey.AddCiphertext(cBigOne, cBigNegOne)
	require.Nil(t, err)
	dRB3, err := testKey.Decrypt(rB3)
	require.Nil(t, err)
	fmt.Printf("[ 0 ]%s\n", dRB3.String())
	// AddCiphertextStr
	rB4, err := testKey.AddCiphertext(c10, c_10)
	require.Nil(t, err)
	dRB4, err := testKey.Decrypt(rB4)
	require.Nil(t, err)
	fmt.Printf("[ 0 ]%d\n", dRB4)

	_, err = Helper().NewPubKey().AddCiphertext(c10, c20)
	require.EqualError(t, err, ErrInvalidPublicKey.Error())

	c10_testKey2, err = testKey2.Encrypt(p10)
	require.Nil(t, err)

	_, err = testKey.AddCiphertext(c10, c10_testKey2)
	require.EqualError(t, err, ErrInvalidMismatch.Error())

}

func testAddPlaintext(t *testing.T) {
	// AddPlaintext
	rAddC10P20, err := testKey.AddPlaintext(c10, p20)
	require.Nil(t, err)
	dRAddC10P20, err := testKey.Decrypt(rAddC10P20)
	require.Nil(t, err)
	fmt.Printf("[ c10 + p20] %d\n", dRAddC10P20)
	rAddC10P_15, _ := testKey.AddPlaintext(c10, p_15)
	dRAddC10P_15, err := testKey.Decrypt(rAddC10P_15)
	require.Nil(t, err)
	fmt.Printf("[ c10 + p(-15)] %d\n", dRAddC10P_15)

	// AddPlaintext
	rAddPlainCBig, err := testKey.AddPlaintext(cBigOne, bigTwo)
	require.Nil(t, err)
	dRAddPlainCBig, err := testKey.Decrypt(rAddPlainCBig)
	require.Nil(t, err)
	fmt.Printf("[ c1 + p2 ] %s\n", dRAddPlainCBig.String())

	// 0
	// AddPlaintextInt
	rB1, _ := testKey.AddPlaintext(c10, p_10)
	dRB1, err := testKey.Decrypt(rB1)
	require.Nil(t, err)
	fmt.Printf("[ 0 ]%d\n", dRB1)
	// AddPlaintext
	fmt.Printf("[ -1 ]%s\n", bigNegOne.String())
	rB2, _ := testKey.AddPlaintext(cBigOne, bigNegOne)
	dRB2, err := testKey.Decrypt(rB2)
	require.Nil(t, err)
	fmt.Printf("[ 0 ]%s\n", dRB2.String())

	_, err = Helper().NewPrvKey().AddPlaintext(nil, nil)
	require.EqualError(t, err, ErrInvalidPublicKey.Error())
	_, err = testKey.AddPlaintext(nil, p10)
	require.EqualError(t, err, ErrInvalidCiphertext.Error())
	_, err = testKey.AddPlaintext(c_10, nil)
	require.EqualError(t, err, ErrInvalidPlaintext.Error())
	_, err = testKey.AddPlaintext(c10_testKey2, p10)
	require.EqualError(t, err, ErrInvalidMismatch.Error())

}

func testSubCiphertext(t *testing.T) {
	// SubCiphertextWithString
	rSubC10C20, _ := testKey.SubCiphertext(c10, c20)
	dRSubC10C20, err := testKey.Decrypt(rSubC10C20)
	require.Nil(t, err)
	fmt.Printf("[ c10 - c20] %d\n", dRSubC10C20)

	rSubC10C_15, _ := testKey.SubCiphertext(c10, c_15)
	dRSubC10C_15, err := testKey.Decrypt(rSubC10C_15)
	require.Nil(t, err)
	fmt.Printf("[ c10 - c(-15)] %d\n", dRSubC10C_15)

	rSubC10C10, _ := testKey.SubCiphertext(c10, c10)
	dRSubC10C10, err := testKey.Decrypt(rSubC10C10)
	require.Nil(t, err)
	fmt.Printf("[ c10 - c100] %d\n", dRSubC10C10)

	// SubCiphertext
	rSubCBig, err := testKey.SubCiphertext(cBigOne, cBigTwo)
	require.Nil(t, err)
	dRSubCBig, err := testKey.Decrypt(rSubCBig)
	require.Nil(t, err)
	fmt.Printf("[ c1 - c2 ] %s\n", dRSubCBig.String())

	// 0
	// SubCiphertext
	rB5, err := testKey.SubCiphertext(cBigOne, cBigOne)
	require.Nil(t, err)
	dRB5, err := testKey.Decrypt(rB5)
	require.Nil(t, err)
	fmt.Printf("[ 0 ]%s\n", dRB5.String())
	// SubCiphertextStr
	rB6, err := testKey.SubCiphertext(c10, c10)
	require.Nil(t, err)
	dRB6, err := testKey.Decrypt(rB6)
	require.Nil(t, err)
	fmt.Printf("[ 0 ]%d\n", dRB6)

	_, err = Helper().NewPubKey().SubCiphertext(c10, c20)
	require.EqualError(t, err, ErrInvalidPublicKey.Error())

	_, err = testKey.SubCiphertext(c10, nil)
	require.EqualError(t, err, ErrInvalidCiphertext.Error())

	c10_testKey2, err = testKey2.Encrypt(p10)
	require.Nil(t, err)

	_, err = testKey.SubCiphertext(c10, c10_testKey2)
	require.EqualError(t, err, ErrInvalidMismatch.Error())

}

func testSubPlaintext(t *testing.T) {
	// SunPlaintext
	rSubC10P20, _ := testKey.SubPlaintext(c10, p20)
	//rSubC10P20, _ := testKey.SubPlaintextInt64(c10, 1)
	dRSubC10P20, err := testKey.Decrypt(rSubC10P20)
	require.Nil(t, err)
	fmt.Printf("[ c10 - p20] %d\n", dRSubC10P20)

	rSubC10P_15, _ := testKey.SubPlaintext(c10, p_15)
	dRSubC10P_15, err := testKey.Decrypt(rSubC10P_15)
	require.Nil(t, err)
	fmt.Printf("[ c10 - p(-15)] %d\n", dRSubC10P_15)

	rSubC10p10, _ := testKey.SubPlaintext(c10, p10)
	dRSubC10p10, err := testKey.Decrypt(rSubC10p10)
	require.Nil(t, err)
	fmt.Printf("[ c10 - p10] %d\n", dRSubC10p10)

	// SubPlaintext
	rSubPlainCBig, err := testKey.SubPlaintext(cBigOne, bigTwo)
	require.Nil(t, err)
	dRSubPlainCBig, err := testKey.Decrypt(rSubPlainCBig)
	require.Nil(t, err)
	fmt.Printf("[ c1 - p2 ] %s\n", dRSubPlainCBig.String())

	// 0
	// SubPlaintext
	rB7, err := testKey.SubPlaintext(cBigOne, bigOne)
	require.Nil(t, err)
	dRB7, err := testKey.Decrypt(rB7)
	require.Nil(t, err)
	fmt.Printf("[ 0 ]%s\n", dRB7.String())
	// SubPlaintextStr
	rB8, err := testKey.SubPlaintext(c10, p10)
	require.Nil(t, err)
	dRB8, err := testKey.Decrypt(rB8)
	require.Nil(t, err)
	fmt.Printf("[ 0 ]%d\n", dRB8)

	_, err = Helper().NewPrvKey().SubPlaintext(nil, nil)
	require.EqualError(t, err, ErrInvalidPublicKey.Error())
	_, err = testKey.SubPlaintext(nil, p10)
	require.EqualError(t, err, ErrInvalidCiphertext.Error())
	_, err = testKey.SubPlaintext(c_10, nil)
	require.EqualError(t, err, ErrInvalidPlaintext.Error())
	_, err = testKey.SubPlaintext(c10_testKey2, p10)
	require.EqualError(t, err, ErrInvalidMismatch.Error())
}

func testNumMul(t *testing.T) {
	// NumMulSimple test
	rMulC10P20, _ := testKey.NumMul(c10, p20)
	dRMulC10P20, err := testKey.Decrypt(rMulC10P20)
	require.Nil(t, err)
	fmt.Printf("[ c10 * p20] %d\n", dRMulC10P20)

	rMulC_15P20, _ := testKey.NumMul(c_15, p20)
	dRMulC_15P20, err := testKey.Decrypt(rMulC_15P20)
	require.Nil(t, err)
	fmt.Printf("[ c(-15) * p20] %d\n", dRMulC_15P20)

	rMulC_15P_20, _ := testKey.NumMul(c_15, p_20)
	dRMulC_15P_20, err := testKey.Decrypt(rMulC_15P_20)
	require.Nil(t, err)
	fmt.Printf("[ c(-15) * p(-20)] %d\n", dRMulC_15P_20)

	// NumMul
	rMulCBig, err := testKey.NumMul(cBigOne, bigTwo)
	require.Nil(t, err)
	dRMulCBig, err := testKey.Decrypt(rMulCBig)
	require.Nil(t, err)
	fmt.Printf("[ c1 * p2 ] %s\n", dRMulCBig.String())

	rMulCMax, err := testKey.NumMul(cBigMax, pMax)
	require.Nil(t, err)
	dRMulCMax, err := testKey.Decrypt(rMulCMax)
	require.Nil(t, err)
	fmt.Printf("[ cMax * cMax ] %s\n", dRMulCMax.String())

	// 0
	// NumMul
	rB9, err := testKey.NumMul(cBigOne, bigZero)
	require.Nil(t, err)
	dRB9, err := testKey.Decrypt(rB9)
	require.Nil(t, err)
	fmt.Printf("[ 0 ] %s\n", dRB9.String())
	rB11, err := testKey.NumMul(cBigZero, bigTwo)
	require.Nil(t, err)
	dRB11, err := testKey.Decrypt(rB11)
	require.Nil(t, err)
	fmt.Printf("[ 0 ] %s\n", dRB11.String())
	// NumMulInt64
	rB10, _ := testKey.NumMul(c_15, p0)
	dRB10, err := testKey.Decrypt(rB10)
	require.Nil(t, err)
	fmt.Printf("[ 0 ] %d\n", dRB10)

	// ciphertext * ciphertext
	rB12, err := testKey.NumMul(cBigTwo, p10)
	require.Nil(t, err)
	dRB12, err := testKey.Decrypt(rB12)
	require.Nil(t, err)
	fmt.Printf("[ c2 * p10] %s\n", dRB12.String())

	_, err = Helper().NewPrvKey().NumMul(nil, nil)
	require.EqualError(t, err, ErrInvalidPublicKey.Error())
	_, err = testKey.NumMul(nil, p10)
	require.EqualError(t, err, ErrInvalidCiphertext.Error())
	_, err = testKey.NumMul(c_10, nil)
	require.EqualError(t, err, ErrInvalidPlaintext.Error())
	_, err = testKey.NumMul(c10_testKey2, p10)
	require.EqualError(t, err, ErrInvalidMismatch.Error())
}

func testBoundary(t *testing.T) {
	var err error
	rMulMax, _ := testKey.AddPlaintext(cMax, p1)
	_, err = testKey.Decrypt(rMulMax)
	require.Nil(t, err)

	rMulMin, _ := testKey.SubPlaintext(cMin, p1)
	_, err = testKey.Decrypt(rMulMin)
	require.Nil(t, err)
}

func testBug(t *testing.T) {
	var err error
	// empty ciphertext
	_, err = testKey.AddCiphertext(c10, nil)
	require.EqualError(t, err, "paillier: invalid ciphertext")

	_, err = testKey.Decrypt(nil)
	require.EqualError(t, err, "paillier: invalid ciphertext")

	for _, item := range []byte("啊") {
		pubBytes = append(pubBytes, item)
	}

	err = pubKeyFromUnmarshal.Unmarshal(pubBytes)
	require.EqualError(t, err, "paillier: invalid public key")
}
