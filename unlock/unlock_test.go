// Copyright (c) 2024 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unlock_test

import (
	"sync"
	"testing"
	"time"

	"import.name/lock/unlock"
)

func ExampleGuard() {
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	unlock.Guard(&mu, func() {
		time.Sleep(time.Microsecond)
	})
}

func TestUnlock(t *testing.T) {
	ExampleGuard()
}
