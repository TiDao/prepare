//+build bulletproofs

/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package bulletproofs

func init() {
	Helper = func() Provider {
		return new(provider)
	}
}
