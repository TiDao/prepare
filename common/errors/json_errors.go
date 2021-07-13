/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package errors

import (
	"chainmaker.org/chainmaker-go/common/json"
	"fmt"
)

var (
	// json
	// -32768 to -32000 is the reserved predefined error code
	ErrParseError     = JsonError{Code: -32700, Message: "Parse error"}      // the server received an invalid JSON. The error is sent to the server trying to parse the JSON text
	ErrInvalidRequest = JsonError{Code: -32600, Message: "Invalid request"}  // the sent JSON is not a valid request object
	ErrMethodNotFound = JsonError{Code: -32601, Message: "Method not found"} // the method does not exist or is invalid
	ErrInvalidParams  = JsonError{Code: -32602, Message: "Invalid params"}   // invalid method parameter
	ErrInternalError  = JsonError{Code: -32603, Message: "Internal error"}   // json-rpc internal error.
	// -32000 to -32099	is the server error reserved for customization

	// txPool
	ErrStructEmpty      = JsonError{Code: -31100, Message: "Struct is nil"}                      // the object is nil
	ErrTxIdExist        = JsonError{Code: -31105, Message: "TxId exist"}                         // tx-id already exists
	ErrTxIdExistDB      = JsonError{Code: -31106, Message: "TxId exist in DB"}                   // tx-id already exists in DB
	ErrTxTimeout        = JsonError{Code: -31108, Message: "TxTimestamp error"}                  // tx-timestamp out of range
	ErrTxPoolLimit      = JsonError{Code: -31110, Message: "TxPool is full"}                     // transaction pool is full
	ErrTxSource         = JsonError{Code: -31112, Message: "TxSource is err"}                    // tx-source is error
	ErrTxHadOnTheChain  = JsonError{Code: -31113, Message: "The tx had been on the block chain"} // The tx had been on the block chain
	ErrTxPoolHasStopped = JsonError{Code: -31114, Message: "The tx pool has stopped"}            // The txPool service has stopped
	ErrTxPoolHasStarted = JsonError{Code: -31114, Message: "The tx pool has started"}            // The txPool service has started

	// core
	ErrBlockHadBeenCommited = JsonError{Code: -31200, Message: "Block had been committed err"} // block had been committed
	ErrConcurrentVerify     = JsonError{Code: -31201, Message: "Block concurrent verify err"}  // block concurrent verify error
	ErrRepeatedVerify       = JsonError{Code: -31202, Message: "Block had been verified err"}  // block had been verified

	// sync
	ErrSyncServiceHasStarted = JsonError{Code: -33000, Message: "The sync service has been started"} // The sync service has been started
	ErrSyncServiceHasStoped  = JsonError{Code: -33001, Message: "The sync service has been stoped"}  // The sync service has been stoped
)

type JsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (err JsonError) String() string {
	marshal, _ := json.Marshal(err)
	return string(marshal)
}

func (err JsonError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("error %d", err.Code)
	}
	return err.Message
}

func (err JsonError) ErrorCode() int {
	return err.Code
}
