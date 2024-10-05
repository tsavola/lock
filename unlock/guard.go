// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package unlock provides unlocking helpers.
package unlock

import (
	"sync"
)

// Guard unlocks the lock, invokes f, and locks the lock again.
func Guard(lock sync.Locker, f func()) {
	lock.Unlock()
	defer lock.Lock()
	f()
}

// Guarded unlocks the lock, invokes f, and locks the lock again.  The
// return value is passed through.
func Guarded[T any](lock sync.Locker, f func() T) T {
	lock.Unlock()
	defer lock.Lock()
	return f()
}
