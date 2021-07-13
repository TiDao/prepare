/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"time"

	bcx509 "chainmaker.org/chainmaker-go/common/crypto/x509"
	"github.com/tjfoc/gmsm/sm2"

	"chainmaker.org/chainmaker-go/common/crypto"
	"chainmaker.org/chainmaker-go/common/crypto/asym"
	"chainmaker.org/chainmaker-go/common/crypto/hash"
)

const (
	defaultCountry            = "CN"
	defaultLocality           = "Beijing"
	defaultProvince           = "Beijing"
	defaultOrganizationalUnit = "ChainMaker"
	defaultOrganization       = "ChainMaker"
	defaultCommonName         = "chainmaker.org"

	createFileFailedErrorTemplate       = "create file failed, %s"
	parseCertificateFailedErrorTemplate = "ParseCertificateRequest failed, %s"
)

const (
	defaultExpireYear = 10
)

// CACertificateConfig contains necessary parameters for creating private key.
type CACertificateConfig struct {
	PrivKey            crypto.PrivateKey
	HashType           crypto.HashType
	CertPath           string
	CertFileName       string
	Country            string
	Locality           string
	Province           string
	OrganizationalUnit string
	Organization       string
	CommonName         string
	ExpireYear         int32
	Sans               []string
}

// CreatePrivKey - create private key file
func CreatePrivKey(keyType crypto.KeyType, privKeyPath, privKeyFileName string) (crypto.PrivateKey, error) {
	algoName, ok := crypto.KeyType2NameMap[keyType]
	if !ok {
		return nil, fmt.Errorf("unknown key algo type [%d]", keyType)
	}

	privKey, err := asym.GenerateKeyPair(keyType)
	if err != nil {
		return nil, fmt.Errorf("generate key pair [%s] failed, %s", algoName, err.Error())
	}

	privKeyPEM, err := privKey.String()
	if err != nil {
		return nil, fmt.Errorf("key to pem failed, %s", err.Error())
	}

	if privKeyPath != "" {
		if err = os.MkdirAll(privKeyPath, os.ModePerm); err != nil {
			return nil, fmt.Errorf("mk privKey dir failed, %s", err.Error())
		}

		if err = ioutil.WriteFile(filepath.Join(privKeyPath, privKeyFileName),
			[]byte(privKeyPEM), 0600); err != nil {
			return nil, fmt.Errorf("save privKey to file [%s] failed, %s", privKeyPath, err.Error())
		}
	}

	return privKey, nil
}

// CreateCACertificate - create ca cert file
func CreateCACertificate(cfg *CACertificateConfig) error {
	template, err := GenerateCertTemplate(&GenerateCertTemplateConfig{
		PrivKey:            cfg.PrivKey,
		IsCA:               true,
		Country:            cfg.Country,
		Locality:           cfg.Locality,
		Province:           cfg.Province,
		OrganizationalUnit: cfg.OrganizationalUnit,
		Organization:       cfg.Organization,
		CommonName:         cfg.CommonName,
		ExpireYear:         cfg.ExpireYear,
		Sans:               cfg.Sans,
	})
	if err != nil {
		return fmt.Errorf("generateCertTemplate failed, %s", err.Error())
	}

	template.SubjectKeyId, err = ComputeSKI(cfg.HashType, cfg.PrivKey.PublicKey().ToStandardKey())
	if err != nil {
		return fmt.Errorf("create CA cert compute SKI failed, %s", err.Error())
	}

	err = createCertificate(cfg.PrivKey, template, template, cfg.CertPath, cfg.CertFileName)
	if err != nil {
		return fmt.Errorf("createCertificate failed, %s", err.Error())
	}

	return nil
}

// CSRConfig contains necessary parameters for creating csr.
type CSRConfig struct {
	PrivKey            crypto.PrivateKey
	CsrPath            string
	CsrFileName        string
	Country            string
	Locality           string
	Province           string
	OrganizationalUnit string
	Organization       string
	CommonName         string
}

func CreateCSR(cfg *CSRConfig) error {

	templateX509 := GenerateCSRTemplate(cfg.PrivKey, cfg.Country, cfg.Locality, cfg.Province, cfg.OrganizationalUnit, cfg.Organization, cfg.CommonName)

	template, err := bcx509.X509CertCsrToChainMakerCertCsr(templateX509)
	if err != nil {
		return fmt.Errorf("generate csr failed, %s", err.Error())
	}

	data, err := bcx509.CreateCertificateRequest(rand.Reader, template, cfg.PrivKey.ToStandardKey())
	if err != nil {
		return fmt.Errorf("CreateCertificateRequest failed, %s", err.Error())
	}

	if err = os.MkdirAll(cfg.CsrPath, os.ModePerm); err != nil {
		return fmt.Errorf("mk csr dir failed, %s", err.Error())
	}

	path := filepath.Join(cfg.CsrPath, cfg.CsrFileName)
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf(createFileFailedErrorTemplate, err.Error())
	}
	defer f.Close()

	return pem.Encode(f, &pem.Block{Type: "CSR", Bytes: data})
}

// IssueCertificateConfig contains necessary parameters for issuing cert.
type IssueCertificateConfig struct {
	HashType              crypto.HashType
	IsCA                  bool
	IssuerPrivKeyFilePath string
	IssuerCertFilePath    string
	IssuerPrivKeyPwd      []byte
	CsrFilePath           string
	CertPath              string
	CertFileName          string
	ExpireYear            int32
	Sans                  []string
	Uuid                  string
}

// IssueCertificate - issue certification
func IssueCertificate(cfg *IssueCertificateConfig) error {
	privKey, issuerCert, csr, sn, err := issueCertificatePrepare(cfg)
	if err != nil {
		return err
	}

	basicConstraintsValid := false
	if cfg.IsCA {
		basicConstraintsValid = true
	}

	expireYear := cfg.ExpireYear
	if expireYear <= 0 {
		expireYear = defaultExpireYear
	}

	dnsName, ipAddrs := dealSANS(cfg.Sans)

	var extraExtensions []pkix.Extension
	if cfg.Uuid != "" {
		extSubjectAltName := pkix.Extension{}
		extSubjectAltName.Id = bcx509.OidNodeId
		extSubjectAltName.Critical = false
		extSubjectAltName.Value = []byte(cfg.Uuid)

		extraExtensions = append(extraExtensions, extSubjectAltName)
	}

	notBefore := time.Now().Add(-10 * time.Minute).UTC()
	template := &x509.Certificate{
		Signature:             csr.Signature,
		SignatureAlgorithm:    x509.SignatureAlgorithm(csr.SignatureAlgorithm),
		PublicKey:             csr.PublicKey,
		PublicKeyAlgorithm:    x509.PublicKeyAlgorithm(csr.PublicKeyAlgorithm),
		SerialNumber:          sn,
		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(time.Duration(expireYear) * 365 * 24 * time.Hour).UTC(),
		BasicConstraintsValid: basicConstraintsValid,
		IsCA:                  cfg.IsCA,
		Issuer:                issuerCert.Subject,
		KeyUsage: x509.KeyUsageDigitalSignature |
			x509.KeyUsageKeyEncipherment |
			x509.KeyUsageCertSign |
			x509.KeyUsageCRLSign,
		ExtKeyUsage:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		IPAddresses:     ipAddrs,
		DNSNames:        dnsName,
		ExtraExtensions: extraExtensions,
		Subject:         csr.Subject,
	}

	if issuerCert.SubjectKeyId != nil {
		template.AuthorityKeyId = issuerCert.SubjectKeyId
	} else {
		template.AuthorityKeyId, err = ComputeSKI(cfg.HashType, issuerCert.PublicKey)
		if err != nil {
			return fmt.Errorf("issue cert compute issuer cert SKI failed, %s", err.Error())
		}
	}

	template.SubjectKeyId, err = ComputeSKI(cfg.HashType, csr.PublicKey.ToStandardKey())
	if err != nil {
		return fmt.Errorf("issue cert compute csr SKI failed, %s", err.Error())
	}

	x509certEncode, err := bcx509.CreateCertificate(rand.Reader, template, issuerCert,
		csr.PublicKey.ToStandardKey(), privKey.ToStandardKey())
	if err != nil {
		return fmt.Errorf("issue certificate failed, %s", err)
	}

	if err = os.MkdirAll(cfg.CertPath, os.ModePerm); err != nil {
		return fmt.Errorf("mk cert dir failed, %s", err.Error())
	}

	f, err := os.Create(filepath.Join(cfg.CertPath, cfg.CertFileName))
	if err != nil {
		return fmt.Errorf(createFileFailedErrorTemplate, err.Error())
	}
	defer f.Close()

	return pem.Encode(f, &pem.Block{Type: "CERTIFICATE", Bytes: x509certEncode})
}

func issueCertificatePrepare(cfg *IssueCertificateConfig) (privKey crypto.PrivateKey, issuerCert *x509.Certificate, csr *bcx509.CertificateRequest, sn *big.Int, err error) {

	privKeyRaw, err := ioutil.ReadFile(cfg.IssuerPrivKeyFilePath)
	if err != nil {
		err = fmt.Errorf("read private key file [%s] failed, %s", cfg.IssuerPrivKeyFilePath, err)
		return
	}

	privKey, err = asym.PrivateKeyFromPEM(privKeyRaw, cfg.IssuerPrivKeyPwd)
	if err != nil {
		err = fmt.Errorf("PrivateKeyFromPEM failed, %s", err)
		return
	}

	issuerCert, err = ParseCertificate(cfg.IssuerCertFilePath)
	if err != nil {
		err = fmt.Errorf("ParseCertificate cert failed, %s", err)
		return
	}

	csrOriginal, err := ParseCertificateRequest(cfg.CsrFilePath)
	if err != nil {
		err = fmt.Errorf(parseCertificateFailedErrorTemplate, err)
		return
	}

	csr, err = bcx509.X509CertCsrToChainMakerCertCsr(csrOriginal)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf(parseCertificateFailedErrorTemplate, err)
	}

	if err = csr.CheckSignature(); err != nil {
		return nil, nil, nil, nil, fmt.Errorf("csr CheckSignature failed, %s", err)
	}

	sn, err = rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("get sn failed, %s", err)
	}
	return
}

// ParseCertificate - parse certification
func ParseCertificate(certFilePath string) (*x509.Certificate, error) {
	certRaw, err := ioutil.ReadFile(certFilePath)
	if err != nil {
		return nil, fmt.Errorf("read cert file [%s] failed, %s", certFilePath, err)
	}

	block, _ := pem.Decode(certRaw)
	cert, err := bcx509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("ParseCertificate cert failed, %s", err)
	}

	return bcx509.ChainMakerCertToX509Cert(cert)
}

// ParseCertificateRequest - parse certification request
func ParseCertificateRequest(csrFilePath string) (*x509.CertificateRequest, error) {
	csrRaw, err := ioutil.ReadFile(csrFilePath)
	if err != nil {
		return nil, fmt.Errorf("read csr file [%s] failed, %s", csrFilePath, err)
	}

	block, _ := pem.Decode(csrRaw)
	csrBC, err := bcx509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf(parseCertificateFailedErrorTemplate, err)
	}

	return bcx509.ChainMakerCertCsrToX509CertCsr(csrBC)
}

func ParseCertificateToJson(certFilePath string) (string, error) {
	cert, err := ParseCertificate(certFilePath)
	if err != nil {
		return "", err
	}

	ret, err := json.Marshal(cert)
	if err != nil {
		return "", fmt.Errorf("json marshal cert failed, %s", err)
	}

	return string(ret), nil
}

type subjectPublicKeyInfo struct {
	Algorithm        pkix.AlgorithmIdentifier
	SubjectPublicKey asn1.BitString
}

func ComputeSKI(hashType crypto.HashType, pub interface{}) ([]byte, error) {
	encodedPub, err := bcx509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}

	var subPKI subjectPublicKeyInfo
	_, err = asn1.Unmarshal(encodedPub, &subPKI)
	if err != nil {
		return nil, err
	}

	pubHash, err := hash.Get(hashType, subPKI.SubjectPublicKey.Bytes)
	if err != nil {
		return nil, err
	}

	return pubHash[:], nil
}

func createCertificate(privKey crypto.PrivateKey, template, parent *x509.Certificate,
	certPath, certFileName string) error {

	x509certEncode, err := bcx509.CreateCertificate(rand.Reader, template, parent,
		privKey.PublicKey().ToStandardKey(), privKey.ToStandardKey())
	if err != nil {
		return err
	}

	if err = os.MkdirAll(certPath, os.ModePerm); err != nil {
		return fmt.Errorf("mk cert dir failed, %s", err.Error())
	}

	f, err := os.Create(filepath.Join(certPath, certFileName))
	if err != nil {
		return fmt.Errorf(createFileFailedErrorTemplate, err.Error())
	}
	defer f.Close()

	return pem.Encode(f, &pem.Block{Type: "CERTIFICATE", Bytes: x509certEncode})
}

// GenerateCertTemplateConfig contains necessary parameters for creating private key.
type GenerateCertTemplateConfig struct {
	PrivKey            crypto.PrivateKey
	IsCA               bool
	Country            string
	Locality           string
	Province           string
	OrganizationalUnit string
	Organization       string
	CommonName         string
	ExpireYear         int32
	Sans               []string
}

func GenerateCertTemplate(cfg *GenerateCertTemplateConfig) (*x509.Certificate, error) {
	sn, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return nil, err
	}
	notBefore := time.Now().Add(-10 * time.Minute).UTC()

	c := cfg.Country
	if c == "" {
		c = defaultCountry
	}

	l := cfg.Locality
	if l == "" {
		l = defaultLocality
	}

	p := cfg.Province
	if p == "" {
		p = defaultProvince
	}

	ou := cfg.OrganizationalUnit
	if ou == "" {
		ou = defaultOrganizationalUnit
	}

	o := cfg.Organization
	if o == "" {
		o = defaultOrganization
	}

	cn := cfg.CommonName
	if cn == "" {
		cn = defaultCommonName
	}

	basicConstraintsValid := false
	if cfg.IsCA {
		basicConstraintsValid = true
	}

	expireYear := cfg.ExpireYear
	if expireYear <= 0 {
		expireYear = defaultExpireYear
	}

	signatureAlgorithm := getSignatureAlgorithm(cfg.PrivKey)
	dnsName, ipAddrs := dealSANS(cfg.Sans)

	template := &x509.Certificate{
		SignatureAlgorithm:    signatureAlgorithm,
		SerialNumber:          sn,
		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(time.Duration(expireYear) * 365 * 24 * time.Hour).UTC(),
		BasicConstraintsValid: basicConstraintsValid,
		IsCA:                  cfg.IsCA,
		KeyUsage: x509.KeyUsageDigitalSignature |
			x509.KeyUsageKeyEncipherment |
			x509.KeyUsageCertSign |
			x509.KeyUsageCRLSign,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		IPAddresses: ipAddrs,
		DNSNames:    dnsName,
		Subject: pkix.Name{
			Country:            []string{c},
			Locality:           []string{l},
			Province:           []string{p},
			OrganizationalUnit: []string{ou},
			Organization:       []string{o},
			CommonName:         cn,
		},
	}

	return template, nil
}

func GenerateCSRTemplate(privKey crypto.PrivateKey,
	country, locality, province, organizationalUnit, organization, commonName string) *x509.CertificateRequest {
	c := country
	if c == "" {
		c = defaultCountry
	}

	l := locality
	if l == "" {
		l = defaultLocality
	}

	p := province
	if p == "" {
		p = defaultProvince
	}

	ou := organizationalUnit
	if ou == "" {
		ou = defaultOrganizationalUnit
	}

	o := organization
	if o == "" {
		o = defaultOrganization
	}

	cn := commonName
	if cn == "" {
		cn = defaultCommonName
	}

	signatureAlgorithm := getSignatureAlgorithm(privKey)

	template := &x509.CertificateRequest{
		SignatureAlgorithm: signatureAlgorithm,
		Subject: pkix.Name{
			Country:            []string{c},
			Locality:           []string{l},
			Province:           []string{p},
			OrganizationalUnit: []string{ou},
			Organization:       []string{o},
			CommonName:         cn,
		},
	}

	return template
}

func getSignatureAlgorithm(privKey crypto.PrivateKey) x509.SignatureAlgorithm {
	signatureAlgorithm := x509.ECDSAWithSHA256
	switch privKey.PublicKey().ToStandardKey().(type) {
	case *rsa.PublicKey:
		signatureAlgorithm = x509.SHA256WithRSA
	case *sm2.PublicKey:
		signatureAlgorithm = x509.SignatureAlgorithm(bcx509.SM3WithSM2)
	}

	return signatureAlgorithm
}

func dealSANS(sans []string) ([]string, []net.IP) {

	var dnsName []string
	var ipAddrs []net.IP

	for _, san := range sans {
		ip := net.ParseIP(san)
		if ip != nil {
			ipAddrs = append(ipAddrs, ip)
		} else {
			dnsName = append(dnsName, san)
		}
	}

	return dnsName, ipAddrs
}
