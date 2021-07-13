/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package ca

import (
	"chainmaker.org/chainmaker-go/common/log"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmtls"
	"github.com/tjfoc/gmtls/gmcredentials"
	"google.golang.org/grpc/credentials"
)

type CAClient struct {
	ServerName string
	CaPaths    []string
	CaCerts    []string
	CertFile   string
	KeyFile    string
	CertBytes  []byte
	KeyBytes   []byte
	Logger     log.LoggerInterface
}

func (c *CAClient) GetCredentialsByCA() (*credentials.TransportCredentials, error) {
	var (
		cert tls.Certificate
		gmCert gmtls.Certificate
		err error
	)

	if c.CertBytes != nil && c.KeyBytes != nil {
		cert, err = tls.X509KeyPair(c.CertBytes, c.KeyBytes)
	} else {
		cert, err = tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
	}
	if err == nil {
		return c.getCredentialsByCA(&cert)
	}

	if c.CertBytes != nil && c.KeyBytes != nil {
		gmCert, err = gmtls.X509KeyPair(c.CertBytes, c.KeyBytes)
	} else {
		gmCert, err = gmtls.LoadX509KeyPair(c.CertFile, c.KeyFile)
	}
	if err == nil {
		return c.getGMCredentialsByCA(&gmCert)
	}

	return nil, fmt.Errorf("load X509 key pair failed, %s", err.Error())
}

func (c *CAClient) getCredentialsByCA(cert *tls.Certificate) (*credentials.TransportCredentials, error) {
	certPool := x509.NewCertPool()
	if len(c.CaCerts) != 0 {
		c.appendCertsToCertPool(certPool)
	} else {
		if err := c.addTrustCertsToCertPool(certPool); err != nil {
			return nil, err
		}
	}

	clientTLS := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{*cert},
		ServerName:         c.ServerName,
		RootCAs:            certPool,
		InsecureSkipVerify: false,
	})

	return &clientTLS, nil
}

func (c *CAClient) appendCertsToCertPool(certPool *x509.CertPool) {
	for _, caCert := range c.CaCerts {
		if caCert != "" {
			certPool.AppendCertsFromPEM([]byte(caCert))
		}
	}
}

func (c *CAClient) addTrustCertsToCertPool(certPool *x509.CertPool) error {
	certs, err := loadCerts(c.CaPaths)
	if err != nil {
		errMsg := fmt.Sprintf("load trust certs failed, %s", err.Error())
		return errors.New(errMsg)
	}

	if len(certs) == 0 {
		errMsg := fmt.Sprintf("trust certs dir is empty!")
		return errors.New(errMsg)
	}

	for _, cert := range certs {
		err := addTrust(certPool, cert)
		if err != nil {
			c.Logger.Warnf("ignore invalid cert [%s], %s", cert, err.Error())
			continue
		}
	}
	return nil
}

func (c *CAClient) getGMCredentialsByCA(cert *gmtls.Certificate) (*credentials.TransportCredentials, error) {
	certPool := sm2.NewCertPool()
	if len(c.CaCerts) != 0 {
		c.appendCertsToSM2CertPool(certPool)
	} else {
		if err := c.addTrustCertsToSM2CertPool(certPool); err != nil {
			return nil, err
		}
	}

	clientTLS := gmcredentials.NewTLS(&gmtls.Config{
		Certificates:       []gmtls.Certificate{*cert},
		ServerName:         c.ServerName,
		RootCAs:            certPool,
		InsecureSkipVerify: false,
	})

	return &clientTLS, nil
}

func (c *CAClient) appendCertsToSM2CertPool(certPool *sm2.CertPool) {
	for _, caCert := range c.CaCerts {
		if caCert != "" {
			certPool.AppendCertsFromPEM([]byte(caCert))
		}
	}
}

func (c *CAClient) addTrustCertsToSM2CertPool(certPool *sm2.CertPool) error {
	certs, err := loadCerts(c.CaPaths)
	if err != nil {
		errMsg := fmt.Sprintf("load trust certs failed, %s", err.Error())
		return errors.New(errMsg)
	}

	if len(certs) == 0 {
		errMsg := fmt.Sprintf("trust certs dir is empty!")
		return errors.New(errMsg)
	}

	for _, cert := range certs {
		err := addGMTrust(certPool, cert)
		if err != nil {
			c.Logger.Warnf("ignore invalid cert [%s], %s", cert, err.Error())
			continue
		}
	}
	return nil
}
