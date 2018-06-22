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

func (this *ClockFixture) Setup() {
	this.now = UTCNow()
	this.thawed = new(ThingUnderTest)
	this.frozen = new(ThingUnderTest)
	this.frozen.clock = Freeze(this.now)
	this.frozen.sleeper = StayAwake()
}

func (this *ClockFixture) TestNilReferencesForwardToStandardBehavior() {
	now2 := this.thawed.CurrentTime()
	this.thawed.Sleep(time.Millisecond * 100)
	now3 := this.thawed.CurrentTime()
	this.So(this.now, should.HappenWithin, time.Millisecond, now2)
	this.So(now3, should.HappenOnOrAfter, now2.Add(time.Millisecond*100))
	this.So(this.thawed.TimeSince(this.now), should.BeGreaterThan, time.Millisecond*100)
}

func (this *ClockFixture) TestActualSleeperInstanceIsUsefulForTesting() {
	this.frozen.Sleep(time.Hour)
	now2 := this.frozen.CurrentTime()
	this.So(this.now, should.HappenWithin, time.Millisecond, now2)
	this.So(this.frozen.sleeper.Naps, should.Resemble, []time.Duration{time.Hour})
}

func (this *ClockFixture) TestActualClockInstanceIsUsefulForTesting() {
	now2 := this.frozen.CurrentTime()
	this.So(this.now, should.Resemble, now2)
}

func (this *ClockFixture) TestTimeSinceWhenFrozen() {
	this.So(this.frozen.TimeSince(this.now.Add(-time.Second)), should.Equal, time.Second)
}

func (this *ClockFixture) TestCyclicNatureOfFrozenClock() {
	now1 := time.Now()
	now2 := now1.Add(time.Second)
	now3 := now2.Add(time.Second)

	this.thawed.clock = Freeze(now1, now2, now3)

	now1a := this.thawed.CurrentTime()
	now2a := this.thawed.CurrentTime()
	now3a := this.thawed.CurrentTime()

	this.So(now1, should.Resemble, now1a)
	this.So(now2, should.Resemble, now2a)
	this.So(now3, should.Resemble, now3a)

	now1b := this.thawed.CurrentTime()
	now2b := this.thawed.CurrentTime()
	now3b := this.thawed.CurrentTime()

	this.So(now1, should.Resemble, now1b)
	this.So(now2, should.Resemble, now2b)
	this.So(now3, should.Resemble, now3b)
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

func (this *ThingUnderTest) TimeSince(instant time.Time) time.Duration {
	return this.clock.TimeSince(instant)
}
