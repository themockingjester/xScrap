package machine

import (
	"sync"
)

type DynamicLimiter struct {
	mu        sync.Mutex
	cond      *sync.Cond
	active    int
	maxActive int
}

func NewDynamicLimiter(initial int) *DynamicLimiter {
	dl := &DynamicLimiter{maxActive: initial}
	dl.cond = sync.NewCond(&dl.mu)
	return dl
}

func (dl *DynamicLimiter) Acquire() {
	dl.mu.Lock()
	defer dl.mu.Unlock()
	for dl.active >= dl.maxActive {
		dl.cond.Wait()
	}
	dl.active++
}

func (dl *DynamicLimiter) Release() {
	dl.mu.Lock()
	defer dl.mu.Unlock()
	dl.active--
	dl.cond.Signal()
}
func (dl *DynamicLimiter) UpdateLimit(newLimit int, relative bool) {
	dl.mu.Lock()
	defer dl.mu.Unlock()

	if relative {
		dl.maxActive += newLimit
	} else {
		dl.maxActive = newLimit
	}

	if dl.maxActive < 1 {
		dl.maxActive = 1
	}
	dl.cond.Broadcast()
}
