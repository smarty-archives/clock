package clock

import (
	"testing"
	"time"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestClockFixture(t *testing.T) {
	gunit.Run(new(ClockFixture), t)
}

type ClockFixture struct {
	*gunit.Fixture

	now    time.Time
	thawed *ThingUnderTest
	frozen *ThingUnderTest
}

func (it *ClockFixture) Setup() {
	it.now = UTCNow()
	it.thawed = new(ThingUnderTest)
	it.frozen = new(ThingUnderTest)
	it.frozen.clock = Freeze(it.now)
	it.frozen.sleeper = StayAwake()
}

func (it *ClockFixture) TestNilReferencesForwardToStandardBehavior() {
	now2 := it.thawed.CurrentTime()
	it.thawed.Sleep(time.Millisecond * 100)
	now3 := it.thawed.CurrentTime()
	it.So(it.now, should.HappenWithin, time.Millisecond, now2)
	it.So(now3, should.HappenOnOrAfter, now2.Add(time.Millisecond*100))
	it.So(it.thawed.TimeSince(it.now), should.BeGreaterThan, time.Millisecond*100)
}

func (it *ClockFixture) TestActualSleeperInstanceIsUsefulForTesting() {
	it.frozen.Sleep(time.Hour)
	now2 := it.frozen.CurrentTime()
	it.So(it.now, should.HappenWithin, time.Millisecond, now2)
	it.So(it.frozen.sleeper.Naps, should.Resemble, []time.Duration{time.Hour})
}

func (it *ClockFixture) TestActualClockInstanceIsUsefulForTesting() {
	now2 := it.frozen.CurrentTime()
	it.So(it.now, should.Resemble, now2)
}

func (it *ClockFixture) TestTimeSinceWhenFrozen() {
	it.So(it.frozen.TimeSince(it.now.Add(-time.Second)), should.Equal, time.Second)
}

func (it *ClockFixture) TestCyclicNatureOfFrozenClock() {
	now1 := time.Now()
	now2 := now1.Add(time.Second)
	now3 := now2.Add(time.Second)

	it.thawed.clock = Freeze(now1, now2, now3)

	now1a := it.thawed.CurrentTime()
	now2a := it.thawed.CurrentTime()
	now3a := it.thawed.CurrentTime()

	it.So(now1, should.Resemble, now1a)
	it.So(now2, should.Resemble, now2a)
	it.So(now3, should.Resemble, now3a)

	now1b := it.thawed.CurrentTime()
	now2b := it.thawed.CurrentTime()
	now3b := it.thawed.CurrentTime()

	it.So(now1, should.Resemble, now1b)
	it.So(now2, should.Resemble, now2b)
	it.So(now3, should.Resemble, now3b)
}

////////////////////////////////////////////////////////////////////////////////////////////

type ThingUnderTest struct {
	clock   *Clock
	sleeper *Sleeper
}

func (it *ThingUnderTest) CurrentTime() time.Time {
	return it.clock.UTCNow()
}

func (it *ThingUnderTest) Sleep(duration time.Duration) {
	it.sleeper.Sleep(duration)
}

func (it *ThingUnderTest) TimeSince(instant time.Time) time.Duration {
	return it.clock.TimeSince(instant)
}
