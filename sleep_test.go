package clock

import (
	"testing"
	"time"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestFakeSleep(t *testing.T) {
	assert := assertions.New(t)
	start := time.Now()

	sleeper := FakeSleep()
	defer sleeper.Restore()

	Sleep(time.Nanosecond)
	Sleep(time.Millisecond)
	Sleep(time.Second)
	Sleep(time.Minute)
	Sleep(time.Hour)

	t.Log("After FakeSleep() but before sleeper.Restore(), calls to Sleep(...) do not block, except to save the provided value.")
	assert.So(time.Since(start), should.BeLessThan, time.Duration(time.Millisecond))

	t.Log("All values passed to sleeper.Sleep(...) should be saved.")
	assert.So(sleeper.Naps, should.Resemble, []time.Duration{
		time.Nanosecond,
		time.Millisecond,
		time.Second,
		time.Minute,
		time.Hour,
	})
}
