# clock
--
    import "github.com/smartystreets/clock"


## Usage

```go
var Now = time.Now
```
Now forwards to time.Now.

```go
var Sleep = time.Sleep
```
Sleep forwards to time.Sleep

#### func  Freeze

```go
func Freeze(times ...time.Time)
```
Freeze uses the times provided as cyclic return values for the Now func. It is
intended to be called from test code in order to mock calls to Now in production
code.

#### func  Restore

```go
func Restore()
```
Restore discards any values provided to Freeze by assigning time.Now back to
Now. It is intended to be called from test code as cleanup after the actions
under test have been invoked.

#### type Sleeper

```go
type Sleeper struct {
	Naps []time.Duration
}
```

Sleeper stores calls to Sleep in its Naps slice.

#### func  FakeSleep

```go
func FakeSleep() *Sleeper
```
FakeSleep returns a *Sleeper instance and replace Sleep with *Sleeper.Sleep. It
is intended to be called from test code in order to mock calls to Now in
production code.

#### func (*Sleeper) Restore

```go
func (this *Sleeper) Restore()
```
Restore assigns time.Sleep as the value for Sleep. It is intended to be called
from test code as cleanup after the actions under test have been invoked.
