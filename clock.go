package clock

import (
	"sync"
	"time"
)

// Now is a proxy for time.Now.
var Now = time.Now

// Freeze uses the times provided as cyclic return values for the Now func.
// It is intended to be called from test code in order to mock calls to Now
// in production code.
func Freeze(times ...time.Time) {
	if len(times) == 0 {
		panic("You must provide at least one time value.")
	}
	Now = new(times).now
}

// Restore discards any values provided to Freeze by assigning time.Now back to Now.
// It is intended to be called from test code as cleanup after the actions under test
// have been invoked.
func Restore() {
	Now = time.Now
}

type clock struct {
	index int
	times []time.Time
	lock  *sync.Mutex
}

func new(times []time.Time) *clock { return &clock{times: times, lock: &sync.Mutex{}} }
func (this *clock) now() time.Time {
	defer this.incrementAndUnlock()
	this.lock.Lock()
	return this.times[this.index]
}

func (this *clock) incrementAndUnlock() {
	defer this.lock.Unlock()

	this.index++
	if this.index >= len(this.times) {
		this.index = 0
	}
}
