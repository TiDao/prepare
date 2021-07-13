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

type CAServer struct {
	CaPaths  []string
	CaCerts  []string
	CertFile string
	KeyFile  string
	Logger   log.LoggerInterface
}

func (s *CAServer) GetCredentialsByCA(checkClientAuth bool) (*credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(s.CertFile, s.KeyFile)
	if err == nil {
		return s.getCredentialsByCA(checkClientAuth, &cert)
	}

	gmCert, err := gmtls.LoadX509KeyPair(s.CertFile, s.KeyFile)
	if err == nil {
		return s.getGMCredentialsByCA(checkClientAuth, &gmCert)
	}

	return nil, fmt.Errorf("load X509 key pair failed, %s", err.Error())
}

func (s *CAServer) getCredentialsByCA(checkClientAuth bool, cert *tls.Certificate) (*credentials.TransportCredentials, error) {
	var (
		clientAuth tls.ClientAuthType
		clientCAs  *x509.CertPool
	)

	if checkClientAuth {

		certPool := x509.NewCertPool()

		if len(s.CaCerts) > 0 {
			if err := s.addCertsToCertPool(certPool); err != nil {
				return nil, err
			}
		} else {
			if err := s.addTrustCertsToCertPool(certPool); err != nil {
				return nil, err
			}
		}

		clientAuth = tls.RequireAndVerifyClientCert
		clientCAs = certPool
	} else {
		clientAuth = tls.NoClientCert
		clientCAs = nil
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{*cert},
		ClientAuth:         clientAuth,
		ClientCAs:          clientCAs,
		InsecureSkipVerify: false,
	})

	return &c, nil
}

func (s *CAServer) addCertsToCertPool(certPool *x509.CertPool) error {
	for _, caCert := range s.CaCerts {
		if caCert != "" {
			err := addCertPool(certPool, caCert)
			if err != nil {
				s.Logger.Warnf("ignore invalid cert [%s], %s", caCert, err.Error())
				continue
			}
		}
	}
	return nil
}

func (s *CAServer) addTrustCertsToCertPool(certPool *x509.CertPool) error {
	caCerts, err := loadCerts(s.CaPaths)
	if err != nil {
		errMsg := fmt.Sprintf("load trust certs failed, %s", err.Error())
		return errors.New(errMsg)
	}

	if len(caCerts) == 0 {
		errMsg := fmt.Sprintf("trust certs dir is empty!")
		return errors.New(errMsg)
	}

	for _, caCert := range caCerts {
		err := addTrust(certPool, caCert)
		if err != nil {
			s.Logger.Warnf("ignore invalid cert [%s], %s", caCert, err.Error())
			continue
		}
	}
	return nil
}

func (s *CAServer) getGMCredentialsByCA(checkClientAuth bool, cert *gmtls.Certificate) (*credentials.TransportCredentials, error) {
	var clientAuth gmtls.ClientAuthType
	var clientCAs *sm2.CertPool

	if checkClientAuth {

		certPool := sm2.NewCertPool()

		if len(s.CaCerts) > 0 {
			if err := s.addCertsToSM2CertPool(certPool); err != nil {
				return nil, err
			}
		} else {
			if err := s.addTrustCertsToSM2CertPool(certPool); err != nil {
				return nil, err
			}
		}

		clientAuth = gmtls.RequireAndVerifyClientCert
		clientCAs = certPool
	} else {
		clientAuth = gmtls.NoClientCert
		clientCAs = nil
	}

	c := gmcredentials.NewTLS(&gmtls.Config{
		Certificates:       []gmtls.Certificate{*cert},
		ClientAuth:         clientAuth,
		ClientCAs:          clientCAs,
		InsecureSkipVerify: false,
	})

	return &c, nil
}

func (s *CAServer) addCertsToSM2CertPool(certPool *sm2.CertPool) error {
	for _, caCert := range s.CaCerts {
		if caCert != "" {
			err := addSM2CertPool(certPool, caCert)
			if err != nil {
				s.Logger.Warnf("ignore invalid cert [%s], %s", caCert, err.Error())
				continue
			}
		}
	}
	return nil
}

func (s *CAServer) addTrustCertsToSM2CertPool(certPool *sm2.CertPool) error {
	caCerts, err := loadCerts(s.CaPaths)
	if err != nil {
		errMsg := fmt.Sprintf("load trust certs failed, %s", err.Error())
		return errors.New(errMsg)
	}

	if len(caCerts) == 0 {
		errMsg := fmt.Sprintf("trust certs dir is empty!")
		return errors.New(errMsg)
	}

	for _, caCert := range caCerts {
		err := addGMTrust(certPool, caCert)
		if err != nil {
			s.Logger.Warnf("ignore invalid cert [%s], %s", caCert, err.Error())
			continue
		}
	}
	return nil
}
