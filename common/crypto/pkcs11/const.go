/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package pkcs11

//
const (
	CKM_VENDOR_DEFINED  = 0x80000000
	CKM_SM2             = CKM_VENDOR_DEFINED + 0x8000
	CKM_SM3_SM2         = CKM_SM2 + 0x00000100 // SM2-SM3 sign with ASN1 encoding
	CKM_SM2_SIGN        = CKM_SM2 + 0x00000104 // SM2 sign with ASN1 encoding (no SM3 hash)
	CKM_SM2_SIGN_NO_DER = CKM_SM2 + 0x00000105 // SM2 sign with plain R|S concatenation (no SM3 hash)
)
