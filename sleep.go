package clock

import (
	"sync"
	"time"
)

// Sleep forwards to time.Sleep
var Sleep = time.Sleep

// FakeSleep returns a *Sleeper instance and replace Sleep with *Sleeper.Sleep.
// It is intended to be called from test code in order to mock calls to Now
// in production code.
func FakeSleep() *Sleeper {
	sleeper := &Sleeper{
		Naps: []time.Duration{},
		lock: &sync.Mutex{},
	}
	Sleep = sleeper.sleep
	return sleeper
}

// Sleeper stores calls to Sleep in its Naps slice.
type Sleeper struct {
	Naps []time.Duration
	lock *sync.Mutex
}

func (this *Sleeper) sleep(nap time.Duration) {
	defer this.lock.Unlock()
	this.lock.Lock()
	this.Naps = append(this.Naps, nap)
}

// Restore assigns time.Sleep as the value for Sleep.
// It is intended to be called from test code as cleanup after the actions under test
// have been invoked.
func (this *Sleeper) Restore() {
	Sleep = time.Sleep
}
