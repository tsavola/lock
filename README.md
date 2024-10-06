### Example

```go
import "import.name/lock"
import "import.name/lock/unlock"

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

func ExampleTagLock() {
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
```
