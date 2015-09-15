package clock

import (
	"testing"
	"time"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestNilReferencesForwardToStandardBehavior(t *testing.T) {
	t.Parallel()

	thing := new(ThingUnderTest)
	now := time.Now().UTC()
	now2 := thing.CurrentTime()
	thing.Sleep(time.Millisecond * 100)
	now3 := thing.CurrentTime()
	assertions.New(t).So(now, should.HappenWithin, time.Millisecond, now2)
	assertions.New(t).So(now3, should.HappenOnOrAfter, now2.Add(time.Millisecond*100))
}

func TestActualSleeperInstanceIsUsefulForTesting(t *testing.T) {
	t.Parallel()

	thing := new(ThingUnderTest)
	now := time.Now().UTC()
	thing.sleeper = StayAwake()
	thing.Sleep(time.Hour)
	now2 := thing.CurrentTime()
	assertions.New(t).So(now, should.HappenWithin, time.Millisecond, now2)
	assertions.New(t).So(thing.sleeper.Naps, should.Resemble, []time.Duration{time.Hour})
}

func TestActualClockInstanceIsUsefulForTesting(t *testing.T) {
	t.Parallel()

	thing := new(ThingUnderTest)
	now := time.Now().UTC()
	thing.clock = Freeze(now)
	now2 := thing.CurrentTime()
	assertions.New(t).So(now, should.Resemble, now2)
}

func TestCyclicNatureOfFrozenClock(t *testing.T) {
	t.Parallel()

	thing := new(ThingUnderTest)
	now1 := time.Now()
	now2 := now1.Add(time.Second)
	now3 := now2.Add(time.Second)

	thing.clock = Freeze(now1, now2, now3)

	now1a := thing.CurrentTime()
	now2a := thing.CurrentTime()
	now3a := thing.CurrentTime()

	assertions.New(t).So(now1, should.Resemble, now1a)
	assertions.New(t).So(now2, should.Resemble, now2a)
	assertions.New(t).So(now3, should.Resemble, now3a)

	now1b := thing.CurrentTime()
	now2b := thing.CurrentTime()
	now3b := thing.CurrentTime()

	assertions.New(t).So(now1, should.Resemble, now1b)
	assertions.New(t).So(now2, should.Resemble, now2b)
	assertions.New(t).So(now3, should.Resemble, now3b)
}

////////////////////////////////////////////////////////////////////////////////////////////

type ThingUnderTest struct {
	clock   *Clock
	sleeper *Sleeper
}

func (this *ThingUnderTest) CurrentTime() time.Time {
	return this.clock.UTCNow()
}

func (this *ThingUnderTest) Sleep(duration time.Duration) {
	this.sleeper.Sleep(duration)
}
