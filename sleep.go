package clock

import (
	"sync"
	"time"
)

// Sleep forwards to time.Sleep
var Sleep = time.Sleep

type Sleeper struct {
	Naps []time.Duration
	lock *sync.Mutex
}

func (this *Sleeper) Sleep(nap time.Duration) {
	defer this.lock.Unlock()
	this.lock.Lock()
	this.Naps = append(this.Naps, nap)
}

func FakeSleep() *Sleeper {
	sleeper := &Sleeper{
		Naps: []time.Duration{},
		lock: &sync.Mutex{},
	}
	Sleep = sleeper.Sleep
	return sleeper
}

func RestoreSleep() {
	Sleep = time.Sleep
}
