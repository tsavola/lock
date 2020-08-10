// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.18
// +build go1.18

package lock_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"import.name/lock"
	"import.name/lock/unlock"
)

func ExampleLock() {
	var (
		mu    sync.Mutex
		count int
	)

	lock.Guard(&mu, func() {
		count++
	})

	n := lock.Guarded(&mu, func() int {
		return count
	})
	fmt.Println(n)
}

func ExampleUnlock() {
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	unlock.Guard(&mu, func() {
		time.Sleep(time.Second)
	})
}

func Test(t *testing.T) {
	ExampleLock()
	ExampleUnlock()
}
