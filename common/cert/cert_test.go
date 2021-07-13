/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package cert

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"chainmaker.org/chainmaker-go/common/crypto"
	"github.com/stretchr/testify/require"
)

const (
	pathPrefix = "/tmp/js"
	c          = "CN"
	l          = "Beijing"
	p          = "Beijing"
	ou         = "chainmaker.org-OU"
	o          = "chainmaker.org-O"
	cn         = "jasonruan"
	expireYear = 8
)

var (
	sans = []string{"127.0.0.1", "localhost", "chainmaker.org", "8.8.8.8"}
)

func TestCreateCACertificate(t *testing.T) {
	createCACertificate(t, crypto.SM2)
	//createCACertificate(t, crypto.RSA512)
	//createCACertificate(t, crypto.RSA1024)
	//createCACertificate(t, crypto.RSA2048)
	//createCACertificate(t, crypto.ECC_NISTP256)
	//createCACertificate(t, crypto.ECC_NISTP384)
	//createCACertificate(t, crypto.ECC_NISTP521)
}

func TestCreateCSR(t *testing.T) {
	createCSR(t, crypto.SM2)
	//createCSR(t, crypto.RSA512)
	//createCSR(t, crypto.RSA1024)
	//createCSR(t, crypto.RSA2048)
	//createCSR(t, crypto.ECC_NISTP256)
	//createCSR(t, crypto.ECC_NISTP384)
	//createCSR(t, crypto.ECC_NISTP521)
}

func TestIssueCertificate(t *testing.T) {
	issueCertificate(t, crypto.SM2)
	//issueCertificate(t, crypto.RSA512)
	//issueCertificate(t, crypto.RSA1024)
	//issueCertificate(t, crypto.RSA2048)
	//issueCertificate(t, crypto.ECC_NISTP256)
	//issueCertificate(t, crypto.ECC_NISTP384)
	//issueCertificate(t, crypto.ECC_NISTP521)
}

func TestParseCertificateToString(t *testing.T) {
	certStr, err := ParseCertificateToJson(filepath.Join(pathPrefix, "ecc_nistp384_issued.crt"))
	require.Nil(t, err)
	fmt.Println(certStr)

	fmt.Println("\n\n===============================================================")

	certStr, err = ParseCertificateToJson(filepath.Join(pathPrefix, "rsa2048_ca.crt"))
	require.Nil(t, err)
	fmt.Println(certStr)
}

func createCACertificate(t *testing.T, keyType crypto.KeyType) {
	keyName, ok := crypto.KeyType2NameMap[keyType]
	require.Equal(t, true, ok)
	keyName = strings.ToLower(keyName)

	privKey, err := CreatePrivKey(keyType, pathPrefix, keyName+"_ca.key")
	require.Nil(t, err)

	err = CreateCACertificate(&CACertificateConfig{privKey, crypto.HASH_TYPE_SHA256, pathPrefix, keyName + "_ca.crt", c, l, p, ou, o, cn + "_ca", expireYear, sans})
	require.Nil(t, err)
}

func createCSR(t *testing.T, keyType crypto.KeyType) {
	keyName, ok := crypto.KeyType2NameMap[keyType]
	require.Equal(t, true, ok)
	keyName = strings.ToLower(keyName)

	privKey, err := CreatePrivKey(keyType, pathPrefix, keyName+"_csr.key")
	require.Nil(t, err)

	err = CreateCSR(&CSRConfig{privKey, pathPrefix, keyName + ".csr", c, l, p, ou, o, cn + "_csr"})
	require.Nil(t, err)

}

func issueCertificate(t *testing.T, keyType crypto.KeyType) {
	keyName, ok := crypto.KeyType2NameMap[keyType]
	require.Equal(t, true, ok)
	keyName = strings.ToLower(keyName)

	issuerPrivKeyFilePath := filepath.Join(pathPrefix, keyName+"_ca.key")
	issuerCertFilePath := filepath.Join(pathPrefix, keyName+"_ca.crt")
	csrFilePath := filepath.Join(pathPrefix, keyName+".csr")

	//err := IssueCertificate(crypto.HASH_TYPE_SHA256, false, issuerPrivKeyFilePath, issuerCertFilePath, nil,
	//	csrFilePath, pathPrefix, keyName+"_issued.crt", expireYear, sans, uuid.UUID())
	err := IssueCertificate(&IssueCertificateConfig{crypto.HASH_TYPE_SHA256, false, issuerPrivKeyFilePath, issuerCertFilePath, nil,
		csrFilePath, pathPrefix, keyName + "_issued.crt", expireYear, sans, ""})
	require.Nil(t, err)
}
