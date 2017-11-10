package golang_composer

import (
	"sync"
	"sync/atomic"
)

type composer struct {
	sync.WaitGroup

	flag uint64
}

var (
	instance *composer
	once     sync.Once
)

func GetComposer() *composer {
	once.Do(func() {
		instance = &composer{}
	})

	return instance
}

// Checks rather to wait or does not need
func (c *composer) NeedWait() {
	c.Wait()
}

// Says all of goroutins to resume execution
func (c *composer) Play() {
	if atomic.LoadUint64(&c.flag) == 1 {
		c.Done()
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Says all of goroutins to pause execution
func (c *composer) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.Add(1)
	}
}
