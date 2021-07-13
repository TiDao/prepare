/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/


package tee

import (
	"encoding/asn1"
	"fmt"
)

const (
	KLV_LENGTH_SIZE = 4
)

var (
	OidKeyBag = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 12, 10, 1, 1}
)

func BinaryToUint32(in []byte) (uint32, error) {
	if len(in) != KLV_LENGTH_SIZE {
		return 0, fmt.Errorf("input is not an uint32: %v", in)
	}

	result := uint32(in[0]) * (1 << 24) + uint32(in[1]) * (1 << 16) + uint32(in[2]) * (1 << 8) + uint32(in[3])
	return result, nil
}

func Uint32ToBinary(in uint32) []byte {
	out := make([]byte, 4)
	out[0] = byte(in / (1 << 24))
	out[1] = byte((in % (1 << 24)) / (1 << 16))
	out[2] = byte((in % (1 << 16)) / (1 << 8))
	out[3] = byte(in % (1 << 8))
	return out
}
