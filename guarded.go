// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.18
// +build go1.18

package lock

import (
	"sync"
)

// Guarded invokes f while keeping the lock locked.  The return value is
// passed through.
func Guarded[T any](lock sync.Locker, f func() T) T {
	lock.Lock()
	defer lock.Unlock()
	return f()
}
