/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sortedmap

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestIntKeySortedMap(_ *testing.T) {
	m := NewIntKeySortedMap()
	m.Put(1, "a")
	fmt.Println("put 2")
	m.Put(2, "b")
	fmt.Println("put 1")
	m.Put(3, "c")
	fmt.Println("put 3")

	fmt.Println(fmt.Sprintf("contains 1: %v", m.Contains(1)))
	fmt.Println(fmt.Sprintf("contains 2: %v", m.Contains(2)))
	fmt.Println(fmt.Sprintf("contains 3: %v", m.Contains(3)))
	fmt.Println(fmt.Sprintf("contains 4: %v", m.Contains(4)))

	m.Range(func(val interface{}) (isContinue bool) {
		if val.(string) != "c" {
			isContinue = true
		}
		fmt.Println(fmt.Sprintf("range v: %s, continue: %v", val, isContinue))
		return
	})

	v, ok := m.Remove(2)
	fmt.Println(fmt.Sprintf("remove 2: v: %s, ok: %v", v, ok))
	v, ok = m.Remove(4)
	fmt.Println(fmt.Sprintf("remove 4: v: %s, ok: %v", v, ok))

	fmt.Println(fmt.Sprintf("contains 1: %v", m.Contains(1)))
	fmt.Println(fmt.Sprintf("contains 2: %v", m.Contains(2)))
	fmt.Println(fmt.Sprintf("contains 3: %v", m.Contains(3)))
	fmt.Println(fmt.Sprintf("contains 4: %v", m.Contains(4)))

}

type lockMap struct {
	l sync.RWMutex
	m map[int]struct{}
}

func (m *lockMap) Put(i int) {
	m.l.Lock()
	m.m[i] = struct{}{}
	m.l.Unlock()
}

func (m *lockMap) Remove(i int) {
	m.l.Lock()
	delete(m.m, i)
	m.l.Unlock()
}

func (m *lockMap) Get(i int) struct{} {
	m.l.RLock()
	defer m.l.RUnlock()
	return m.m[i]
}

func TestMapPerformance(_ *testing.T) {
	m := &lockMap{
		l: sync.RWMutex{},
		m: make(map[int]struct{}),
	}
	var num int = 200000
	var goroutine int = 2000
	fmt.Println("times:", num)
	fmt.Println("goroutine num:", goroutine)
	wg := sync.WaitGroup{}

	startTime := time.Now().UnixNano()
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(start int) {
			for j := start; j < start+goroutine; j++ {
				m.Put(j)
			}
			wg.Done()
		}(num / goroutine * i)
	}

	wg.Wait()
	useTime := time.Now().UnixNano() - startTime
	fmt.Println("lock map put , use time:", useTime, "ns")

	wg = sync.WaitGroup{}
	startTime = time.Now().UnixNano()
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(start int) {
			for j := start; j < start+goroutine; j++ {
				m.Get(j)
			}
			wg.Done()
		}(num / goroutine * i)
	}

	wg.Wait()
	useTime = time.Now().UnixNano() - startTime
	fmt.Println("lock map get , use time:", useTime, "ns")

	wg = sync.WaitGroup{}
	startTime = time.Now().UnixNano()
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(start int) {
			for j := start; j < start+goroutine; j++ {
				m.Remove(j)
			}
			wg.Done()
		}(num / goroutine * i)
	}

	wg.Wait()
	useTime = time.Now().UnixNano() - startTime
	fmt.Println("lock map del , use time:", useTime, "ns")

	m2 := sync.Map{}
	wg = sync.WaitGroup{}
	startTime = time.Now().UnixNano()
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(start int) {
			for j := start; j < start+goroutine; j++ {
				m2.Store(j, struct {
				}{})
			}
			wg.Done()
		}(num / goroutine * i)
	}

	wg.Wait()
	useTime = time.Now().UnixNano() - startTime
	fmt.Println("sync map put , use time:", useTime, "ns")

	wg = sync.WaitGroup{}
	startTime = time.Now().UnixNano()
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(start int) {
			for j := start; j < start+goroutine; j++ {
				m2.Load(j)
			}
			wg.Done()
		}(num / goroutine * i)
	}

	wg.Wait()
	useTime = time.Now().UnixNano() - startTime
	fmt.Println("sync map get , use time:", useTime, "ns")

	wg = sync.WaitGroup{}
	startTime = time.Now().UnixNano()
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(start int) {
			for j := start; j < start+goroutine; j++ {
				m2.Delete(j)
			}
			wg.Done()
		}(num / goroutine * i)
	}

	wg.Wait()
	useTime = time.Now().UnixNano() - startTime
	fmt.Println("sync map del , use time:", useTime, "ns")

	m3 := NewIntKeySortedMap()
	wg = sync.WaitGroup{}
	startTime = time.Now().UnixNano()
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(start int) {
			for j := start; j < start+goroutine; j++ {
				m3.Put(j, struct {
				}{})
			}
			wg.Done()
		}(num / goroutine * i)
	}

	wg.Wait()
	useTime = time.Now().UnixNano() - startTime
	fmt.Println("sort map put , use time:", useTime, "ns")

	wg = sync.WaitGroup{}
	startTime = time.Now().UnixNano()
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(start int) {
			for j := start; j < start+goroutine; j++ {
				m3.Get(j)
			}
			wg.Done()
		}(num / goroutine * i)
	}

	wg.Wait()
	useTime = time.Now().UnixNano() - startTime
	fmt.Println("sort map get , use time:", useTime, "ns")

	wg = sync.WaitGroup{}
	startTime = time.Now().UnixNano()
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(start int) {
			for j := start; j < start+goroutine; j++ {
				m3.Remove(j)
			}
			wg.Done()
		}(num / goroutine * i)
	}

	wg.Wait()
	useTime = time.Now().UnixNano() - startTime
	fmt.Println("sort map del , use time:", useTime, "ns")
}
