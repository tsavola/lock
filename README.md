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
```
