# clock
--
    import "github.com/smartystreets/clock"

package clock is a drop replacement for time.Now().UTC() and time.Sleep(). The
structs defined here are intened to be used as pointer fields on structs. When
nil, these references forward to the corresponding functions in the standard
time package. When not nil they perform behavior that facilitates unit testing
when accessing the current time or sleeping is involved. The main advantage to
this approach is that it is not necessary to provide a non-nil instance in
'contructor' functions or wireup for production code. It is also still trivial
to set a non-nil reference in test code.

## Usage

#### type Clock

```go
type Clock struct {
}
```

Clock is meant be included as a pointer field on a struct. Leaving the instance
as a nil reference will cause any calls on the *Clock to forward to the
corresponding functions in the standard time package. This is meant to be the
behavior in production. In testing, set the field to a non-nil instance of a
*Clock to provide a frozen time instant whenever UTCNow() is called.

#### func  Freeze

```go
func Freeze(instant time.Time) *Clock
```
Freeze creates a new *Clock instance with an internal time instant. This
function is meant to be called from test code. See the godoc for the Clock
struct for details.

#### func (*Clock) UTCNow

```go
func (this *Clock) UTCNow() time.Time
```
UTCNow() -> time.Now().UTC()

#### type Sleeper

```go
type Sleeper struct {
	Naps []time.Duration
}
```

Sleeper is meant be included as a pointer field on a struct. Leaving the
instance as a nil reference will cause any calls on the *Sleeper to forward to
the corresponding functions in the standard time package. This is meant to be
the behavior in production. In testing, set the field to a non-nil instance of a
*Sleeper to record sleep durations for later inspection.

#### func  StayAwake

```go
func StayAwake() *Sleeper
```
StayAwake creates a new *Sleeper instance with an internal duration slice. This
function is meant to be called from test code. See the godoc for the Sleeper
struct for details.

#### func (*Sleeper) Sleep

```go
func (this *Sleeper) Sleep(duration time.Duration)
```
Sleep -> time.Sleep
