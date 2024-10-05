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

func TestLock(t *testing.T) {
	ExampleGuard()
}
