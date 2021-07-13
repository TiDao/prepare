/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package helper

import (
	"chainmaker.org/chainmaker-go/common/helper/libp2pcrypto"
	"chainmaker.org/chainmaker-go/common/helper/libp2ppeer"
	gocrypto "crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/multiformats/go-multiaddr"
	"github.com/tjfoc/gmsm/sm2"
	"strings"
)

// GetNodeUidFromAddr get the unique id of node from a addr. 从地址中截取出节点ID
func GetNodeUidFromAddr(addr string) (string, error) {
	addrInfo := strings.Split(addr, "/")
	l := len(addrInfo)
	if l < 2 {
		return "", errors.New("wrong address")
	}
	return addrInfo[l-1], nil
}

// GetLibp2pPeerIdFromCert create a peer.ID with pubKey that contains in cert. 根据证书中的公钥生成一个libp2p的peer.ID。
func GetLibp2pPeerIdFromCert(certPemBytes []byte) (string, error) {
	var block *pem.Block
	block, _ = pem.Decode(certPemBytes)
	if block == nil {
		return "", errors.New("empty pem block")
	}
	if block.Type != "CERTIFICATE" || len(block.Headers) != 0 {
		return "", errors.New("not certificate pem")
	}

	cert, err := sm2.ParseCertificate(block.Bytes)
	if err != nil {
		return "", err
	}

	pubKey, err := parsePublicKeyToPubKey(cert.PublicKey)
	if err != nil {
		return "", err
	}
	pid, err := libp2ppeer.IDFromPublicKey(pubKey)
	if err != nil {
		return "", err
	}
	return pid.Pretty(), err
}

func parsePublicKeyToPubKey(publicKey gocrypto.PublicKey) (libp2pcrypto.PubKey, error) {
	switch p := publicKey.(type) {
	case *ecdsa.PublicKey:
		if p.Curve == sm2.P256Sm2() {
			b, err := sm2.MarshalPKIXPublicKey(p)
			if err != nil {
				return nil, err
			}
			pub, err := sm2.ParseSm2PublicKey(b)
			if err != nil {
				return nil, err
			}
			return libp2pcrypto.NewSM2PublicKey(pub), nil
		}
		return libp2pcrypto.NewECDSAPublicKey(p), nil
	case *sm2.PublicKey:
		return libp2pcrypto.NewSM2PublicKey(p), nil
	case *rsa.PublicKey:
		return libp2pcrypto.NewRsaPublicKey(*p), nil
	}
	return nil, errors.New("unsupported public key type")
}

// P2pAddressFormatVerify verify a node address format.
func P2pAddressFormatVerify(address string) bool {
	ma, err := multiaddr.NewMultiaddr(address)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = libp2ppeer.AddrInfoFromP2pAddr(ma)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
