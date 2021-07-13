/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package queue

type Iterator struct {
	current *node
}

func (i *Iterator) Value() Element {
	if i.current == nil {
		return nil
	}
	return i.current.value
}

func (i *Iterator) Next() *Iterator {
	if i.current == nil {
		return nil
	}
	i.current = i.current.next
	return i
}
