/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package helper

import (
	"fmt"
	"testing"
)

func TestGetLibp2pPeerIdFromCert(t *testing.T) {
	certBytes := []byte("-----BEGIN CERTIFICATE-----\nMIICHzCCAcSgAwIBAgIRAMR9Zia8ue5OEB/mEJ0B5jYwCgYIKoEcz1UBg3UwYDEL\nMAkGA1UEBhMCQ04xCzAJBgNVBAgTAkdEMQswCQYDVQQHEwJTWjEZMBcGA1UEChMQ\nb3JnMS5leGFtcGxlLmNvbTEcMBoGA1UEAxMTY2Eub3JnMS5leGFtcGxlLmNvbTAe\nFw0yMDA1MjkxMDMwNDJaFw0zMDA1MjcxMDMwNDJaMGAxCzAJBgNVBAYTAkNOMQsw\nCQYDVQQIEwJHRDELMAkGA1UEBxMCU1oxGTAXBgNVBAoTEG9yZzEuZXhhbXBsZS5j\nb20xHDAaBgNVBAMTE2NhLm9yZzEuZXhhbXBsZS5jb20wWTATBgcqhkjOPQIBBggq\ngRzPVQGCLQNCAAQWXBhGZrChTwqPDfhxeXr930tjVWaiF+bToVSAHpYYAOzAI/7S\nB/MMp82P71BDTp+dua4N0VhWWZNYtJRMravvo18wXTAOBgNVHQ8BAf8EBAMCAaYw\nDwYDVR0lBAgwBgYEVR0lADAPBgNVHRMBAf8EBTADAQH/MCkGA1UdDgQiBCA48Q7H\nPVM6G837SCKsNuxA4VsoeLKxs4//8a65NUiNDzAKBggqgRzPVQGDdQNJADBGAiEA\nkSQyih4ax6A7UWiWyzBTv7oNdUL2BGG6I3N5BDZ/040CIQCGlW38vfSntJe1Vvgg\n5ctBDSRW9ophuyCuUX6Gx99Ogw==\n-----END CERTIFICATE-----\n")
	nodeUid, err := GetLibp2pPeerIdFromCert(certBytes)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(nodeUid)
}

func TestGetNodeUidFromAddr(t *testing.T) {
	addr := "/ip4/0.0.0.0/tcp/6666/p2p/QmTrsVrof7hvU79LmAMnJrmhTCUdaBoVNYDhHMUGaVQa6m"
	nodeUid, err := GetNodeUidFromAddr(addr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(nodeUid)
}

func TestP2pAddressFormatVerify(_ *testing.T) {
	fmt.Println(P2pAddressFormatVerify("/ip4/0.0.0.0/tcp/6666/p2p/QmTrsVrof7hvU79LmAMnJrmhTCUdaBoVNYDhHMUGaVQa6m"))
	fmt.Println(P2pAddressFormatVerify("/ip4/0.0.0.0/tcp/6666"))
	fmt.Println(P2pAddressFormatVerify("0.0.0.0:6666"))
}
