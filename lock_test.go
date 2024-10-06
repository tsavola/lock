// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lock_test

import (
	"fmt"
	"sync"
	"testing"

	"import.name/lock"
)

func ExampleGuard() {
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

func ExampleTagMutex() {
	type countLocked struct{}

	var (
		mu    lock.TagMutex[countLocked]
		count int
	)

	incrementWithLock := func(l countLocked) {
		count++
	}

	readWithLock := func(l countLocked) int {
		return count
	}

	lock.GuardTag(&mu, func(l countLocked) {
		incrementWithLock(l)
	})

	n := lock.GuardTagged(&mu, func(l countLocked) int {
		return readWithLock(l)
	})
	fmt.Println(n)
}

func TestLock(t *testing.T) {
	ExampleGuard()
	ExampleTagMutex()
}
