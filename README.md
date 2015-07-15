# clock
--
    import "github.com/smartystreets/clock"


## Usage

```go
var Now = time.Now
```
Now is a proxy for time.Now.

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
