// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lock provides locking helpers.
package lock

import (
	"sync"
)

// Guard invokes f while keeping the lock locked.
func Guard(lock sync.Locker, f func()) {
	lock.Lock()
	defer lock.Unlock()
	f()
}
