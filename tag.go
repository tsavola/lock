// Copyright (c) 2024 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lock

import (
	"sync"
)

type TagLocker[Tag any] interface {
	Lock() Tag
	Unlock()
}

func GuardTag[Tag any](ock TagLocker[Tag], f func(Tag)) {
	tag := ock.Lock()
	defer ock.Unlock()
	f(tag)
}

func GuardTagged[Tag any, Val any](lock TagLocker[Tag], f func(Tag) Val) Val {
	tag := lock.Lock()
	defer lock.Unlock()
	return f(tag)
}

type TagMutex[Tag any] struct {
	sync.Mutex
}

func (mu *TagMutex[Tag]) Lock() Tag {
	mu.Mutex.Lock()
	var tag Tag
	return tag
}
