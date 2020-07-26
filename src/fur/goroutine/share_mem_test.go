package sharemem

import (
	"sync"
	"testing"
	"time"
)

// TestCounter is xxx
func TestCounter(t *testing.T) {
	counter := 0
	var mut sync.Mutex
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock() // 需要在对counter调用前加锁
			counter++
		}()
	}
	time.Sleep(1 * time.Second) // 可能携程还没有执行完，主进程执行完，因此Sleep一会儿
	t.Log(counter)
}
func TestCounterWaitGroup(t *testing.T) {
	counter := 0
	var wg sync.WaitGroup
	var mut sync.Mutex
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
				wg.Done()
			}()
			mut.Lock() // 需要在对counter调用前加锁
			counter++
		}()
	}
	wg.Wait()
	t.Log(counter)
}
