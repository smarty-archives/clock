// package clock is a near drop-in replacement for time.Now().UTC() and time.Sleep().
// The structs defined here are intended to be used as pointer fields on struct types.
// When nil, these references forward to the corresponding functions in the
// standard time package. When not nil they perform behavior that facilitates
// unit testing when accessing the current time or sleeping is involved.
// The main advantage to this approach is that it is not necessary to
// provide a non-nil instance in 'contructor' functions or wireup for
// production code. It is also still trivial to set a non-nil reference
// in test code.
package clock

import "time"

// Clock is meant be included as a pointer field on a struct. Leaving the
// instance as a nil reference will cause any calls on the *Clock to forward
// to the corresponding functions in the standard time package. This is meant
// to be the behavior in production. In testing, set the field to a non-nil
// instance of a *Clock to provide a frozen time instant whenever UTCNow()
// is called.
type Clock struct {
	instants []time.Time
	index    int
}

// Freeze creates a new *Clock instance with an internal time instant.
// This function is meant to be called from test code. See the godoc for the
// Clock struct for details.
func Freeze(instants ...time.Time) *Clock {
	return &Clock{instants: instants}
}

// UTCNow() -> time.Now().UTC() // (unless frozen)
func (this *Clock) UTCNow() time.Time {
	if this == nil || len(this.instants) == 0 {
		return UTCNow()
	}
	defer this.next()
	return this.instants[this.index]
}

func (this *Clock) next() {
	this.index++
	if this.index == len(this.instants) {
		this.index = 0
	}
}

// Analogous to time.Since(instant) // (unless frozen)
func (this *Clock) TimeSince(instant time.Time) time.Duration {
	return this.UTCNow().Sub(instant)
}

///////////////////////////////////////////////////

// UTCNow() -> time.Now().UTC()
func UTCNow() time.Time {
	return time.Now().UTC()
}

///////////////////////////////////////////////////

// Sleeper is meant be included as a pointer field on a struct. Leaving the
// instance as a nil reference will cause any calls on the *Sleeper to forward
// to the corresponding functions in the standard time package. This is meant
// to be the behavior in production. In testing, set the field to a non-nil
// instance of a *Sleeper to record sleep durations for later inspection.
type Sleeper struct {
	Naps []time.Duration
}

// StayAwake creates a new *Sleeper instance with an internal duration slice.
// This function is meant to be called from test code. See the godoc for the
// Sleeper struct for details.
func StayAwake() *Sleeper {
	return &Sleeper{}
}

// Sleep -> time.Sleep
func (this *Sleeper) Sleep(duration time.Duration) {
	if this == nil {
		time.Sleep(duration)
	} else {
		this.Naps = append(this.Naps, duration)
	}
}

////////////////////////////////////////////////////
