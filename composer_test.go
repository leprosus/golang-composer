package golang_composer

import (
	"testing"
	"time"
	"sync"
)

var wg sync.WaitGroup

func TestComposer(t *testing.T) {
	d := make(map[uint64]uint64)

	d[0] = 0
	d[1] = 0
	d[2] = 0

	GetComposer().Pause()

	var i uint64
	for i = 0; i < 3; i++ {
		go func(i uint64) {
			wg.Add(1)

			println("goroutine #", i, "is waited")

			GetComposer().NeedWait()

			println("goroutine #", i, "is running")

			d[i] = 1

			wg.Done()
		}(i)
	}

	if d[0] != 0 || d[1] != 0 || d[2] != 0 {
		t.Fatal("Can't stop goroutines by composer")
	}

	time.Sleep(3 * time.Second)

	GetComposer().Play()

	wg.Wait()

	if d[0] != 1 || d[1] != 1 || d[2] != 1 {
		t.Fatal("Can't play goroutines by composer")
	}
}
