/*
 * Copyright (C) BABEC. All rights reserved.
 * Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package common

import (
	"crypto/sha256"
	"errors"
)

//GetSenderAccountId 获得交易的发起人的唯一账户标识，这个标识如果大于2048字节，则返回的是SHA256 Hash
func (m *Transaction) GetSenderAccountId() []byte {
	var accountId []byte
	if m != nil && m.Header != nil && m.Header.Sender != nil {
		accountId = m.Header.Sender.MemberInfo
	}
	if len(accountId) > 2048 {
		hash := sha256.Sum256(accountId)
		accountId = hash[:]
	}
	return accountId
}

func (m *Transaction) GetContractName() (string, error) {
	if m == nil || m.Header == nil {
		return "", errors.New("null point")
	}
	if m.Header.TxType == TxType_INVOKE_USER_CONTRACT {
		var payload = &TransactPayload{}
		err := payload.Unmarshal(m.RequestPayload)
		if err != nil {
			return "", err
		}
		return payload.ContractName, nil
	}
	if m.Header.TxType == TxType_MANAGE_USER_CONTRACT {
		return ContractName_SYSTEM_CONTRACT_CERT_MANAGE.String(), nil //TODO
	}
	if m.Header.TxType == TxType_UPDATE_CHAIN_CONFIG {
		return ContractName_SYSTEM_CONTRACT_CHAIN_CONFIG.String(), nil //TODO
	}
	if m.Header.TxType == TxType_INVOKE_SYSTEM_CONTRACT {
		var payload = &SystemContractPayload{}
		err := payload.Unmarshal(m.RequestPayload)
		if err != nil {
			return "", err
		}
		return payload.ContractName, nil
	}
	return "", errors.New("unknown tx type " + m.Header.TxType.String())
}
