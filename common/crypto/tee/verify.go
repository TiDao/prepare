/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/


package tee

import (
	"bytes"
	bccrypto "chainmaker.org/chainmaker-go/common/crypto"
	"chainmaker.org/chainmaker-go/common/crypto/asym"
	"chainmaker.org/chainmaker-go/common/crypto/asym/rsa"
	bcx509 "chainmaker.org/chainmaker-go/common/crypto/x509"
	"encoding/pem"
	"fmt"
)

type TEEProof struct {
	VerificationKey    bccrypto.PublicKey
	VerificationKeyPEM []byte
	EncryptionKey      bccrypto.EncryptKey
	EncryptionKeyPEM   []byte
	Certificate        *bcx509.Certificate
	CertificateDER     []byte
	Report             []byte
	Challenge          []byte
	Signature          []byte
}

func AttestationVerify(proof []byte, certOpts bcx509.VerifyOptions, reportFromChain []byte) (bool, *TEEProof, error) {
	challengeLen, err := BinaryToUint32(proof[0:KLV_LENGTH_SIZE])
	if err != nil {
		return false, nil, fmt.Errorf("invalid input: %v", err)
	}
	challenge := proof[KLV_LENGTH_SIZE : challengeLen+KLV_LENGTH_SIZE]

	reportLen, err := BinaryToUint32(proof[challengeLen+KLV_LENGTH_SIZE : challengeLen+KLV_LENGTH_SIZE*2])
	if err != nil {
		return false, nil, fmt.Errorf("invalid input: %v", err)
	}
	report := proof[challengeLen+KLV_LENGTH_SIZE*2 : challengeLen+reportLen+KLV_LENGTH_SIZE*2]

	certLen, err := BinaryToUint32(proof[challengeLen+reportLen+KLV_LENGTH_SIZE*2 : challengeLen+reportLen+KLV_LENGTH_SIZE*3])
	if err != nil {
		return false, nil, fmt.Errorf("invalid input: %v", err)
	}
	certDER := proof[challengeLen+reportLen+KLV_LENGTH_SIZE*3 : challengeLen+reportLen+certLen+KLV_LENGTH_SIZE*3]

	sigLen, err := BinaryToUint32(proof[challengeLen+reportLen+certLen+KLV_LENGTH_SIZE*3 : challengeLen+reportLen+certLen+KLV_LENGTH_SIZE*4])
	if err != nil {
		return false, nil, fmt.Errorf("invalid input: %v", err)
	}
	sig := proof[challengeLen+reportLen+certLen+KLV_LENGTH_SIZE*4 : challengeLen+reportLen+certLen+sigLen+KLV_LENGTH_SIZE*4]

	certificate, err := bcx509.ParseCertificate(certDER)
	if err != nil {
		return false, nil, fmt.Errorf("fail to parse TEE certificate: %v", err)
	}

	verificationKey := certificate.PublicKey

	encryptionKeyPEM, err := bcx509.GetExtByOid(OidKeyBag, certificate.Extensions)
	if err != nil {
		encryptionKeyPEM, err = bcx509.GetExtByOid(OidKeyBag, certificate.ExtraExtensions)
		if err != nil {
			return false, nil, fmt.Errorf("fail to get encryption key: %v", err)
		}
	}

	//encryptionKeyBlock, _ := pem.Decode(encryptionKeyPEM)
	//if encryptionKeyBlock == nil {
	//	return false, nil, fmt.Errorf("fail to decode encryption key")
	//}
	encryptionKeyInterface, err := asym.PublicKeyFromPEM(encryptionKeyPEM)
	if err != nil {
		return false, nil, fmt.Errorf("fail to parse TEE encryption key: %v", err)
	}

	var encryptionKey bccrypto.EncryptKey
	switch k := encryptionKeyInterface.(type) {
	case *rsa.PublicKey:
		encryptionKey = k
	default:
		return false, nil, fmt.Errorf("unrecognized encryption key type")
	}

	msg := proof[0 : challengeLen+reportLen+certLen+KLV_LENGTH_SIZE*3]
	isValid, err := verificationKey.VerifyWithOpts(msg, sig, &bccrypto.SignOpts{
		Hash:         bccrypto.HASH_TYPE_SHA256,
		UID:          "",
		EncodingType: rsa.RSA_PSS,
	})
	if err != nil {
		return false, nil, fmt.Errorf("invalid signature: %v", err)
	}
	if !isValid {
		return false, nil, fmt.Errorf("invalid signature")
	}

	certChains, err := certificate.Verify(certOpts)
	if err != nil || certChains == nil {
		return false, nil, fmt.Errorf("untrusted certificate: %v", err)
	}

	fmt.Printf("###### report = %s\n", string(report))
	fmt.Printf("###### report from chain = %s\n", string(reportFromChain));
	if !bytes.Equal(report, reportFromChain) {
		return false, nil, fmt.Errorf("report does not match, reportFromChain: %s, report: %s",
			reportFromChain, report)
	}

	verificationKeyPEM, err := verificationKey.String()
	if err != nil {
		return false, nil, fmt.Errorf("fail to serialize verification key")
	}

	teeProof := &TEEProof{
		VerificationKey:    verificationKey,
		VerificationKeyPEM: []byte(verificationKeyPEM),
		EncryptionKey:      encryptionKey,
		EncryptionKeyPEM:   encryptionKeyPEM,
		Certificate:        certificate,
		CertificateDER:     certDER,
		Report:             report,
		Challenge:          challenge,
		Signature:          sig,
	}

	return true, teeProof, nil
}

func AttestationVerifyComponents(challenge, signature, report []byte, certificate *bcx509.Certificate, verificationKey bccrypto.PublicKey, encryptionKey bccrypto.EncryptKey, certOpts bcx509.VerifyOptions) (bool, *TEEProof, error) {
	challengeLen := Uint32ToBinary(uint32(len(challenge)))
	reportLen := Uint32ToBinary(uint32(len(report)))
	certLen := Uint32ToBinary(uint32(len(certificate.Raw)))
	msg := append(challengeLen, challenge...)
	msg = append(msg, reportLen...)
	msg = append(msg, report...)
	msg = append(msg, certLen...)
	msg = append(msg, certificate.Raw...)

	isValid, err := verificationKey.VerifyWithOpts(msg, signature, &bccrypto.SignOpts{
		Hash:         bccrypto.HASH_TYPE_SHA256,
		UID:          "",
		EncodingType: rsa.RSA_PSS,
	})
	if err != nil {
		return false, nil, fmt.Errorf("invalid signature: %v", err)
	}
	if !isValid {
		return false, nil, fmt.Errorf("invalid signature")
	}

	certChains, err := certificate.Verify(certOpts)
	if err != nil || certChains == nil {
		return false, nil, fmt.Errorf("untrusted certificate: %v", err)
	}

	verificationKeyPEM, err := verificationKey.String()
	if err != nil {
		return false, nil, fmt.Errorf("fail to serialize verification key")
	}
	verificationKeyDER, err := verificationKey.Bytes()
	if err != nil {
		return false, nil, fmt.Errorf("fail to serialize verification key")
	}

	verificationKeyDERFromCert, err := certificate.PublicKey.Bytes()
	if err != nil {
		return false, nil, fmt.Errorf("fail to serialize verification key in certificate")
	}

	if !bytes.Equal(verificationKeyDER, verificationKeyDERFromCert) {
		return false, nil, fmt.Errorf("verification key do not match")
	}

	encryptionKeyPEM, err := encryptionKey.String()
	if err != nil {
		return false, nil, fmt.Errorf("fail to serialize encryption key")
	}
	encryptionKeyDER, err := encryptionKey.Bytes()
	if err != nil {
		return false, nil, fmt.Errorf("fail to serialize encryption key")
	}

	encryptionKeyPEMFromCert, err := bcx509.GetExtByOid(OidKeyBag, certificate.Extensions)
	if err != nil {
		encryptionKeyPEMFromCert, err = bcx509.GetExtByOid(OidKeyBag, certificate.ExtraExtensions)
		if err != nil {
			return false, nil, fmt.Errorf("fail to get encryption key: %v", err)
		}
	}

	encryptionKeyBlockFromCert, _ := pem.Decode(encryptionKeyPEMFromCert)
	if encryptionKeyBlockFromCert == nil {
		return false, nil, fmt.Errorf("fail to decode encryption key")
	}

	if !bytes.Equal(encryptionKeyDER, encryptionKeyBlockFromCert.Bytes) {
		return false, nil, fmt.Errorf("encryption key do not match")
	}

	teeProof := &TEEProof{
		VerificationKey:    verificationKey,
		VerificationKeyPEM: []byte(verificationKeyPEM),
		EncryptionKey:      encryptionKey,
		EncryptionKeyPEM:   []byte(encryptionKeyPEM),
		Certificate:        certificate,
		CertificateDER:     certificate.Raw,
		Report:             report,
		Challenge:          challenge,
		Signature:          signature,
	}

	return true, teeProof, nil
}
