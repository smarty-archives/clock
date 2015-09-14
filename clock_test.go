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
	thing.Sleeper = StayAwake()
	thing.Sleep(time.Hour)
	now2 := thing.CurrentTime()
	assertions.New(t).So(now, should.HappenWithin, time.Millisecond, now2)
	assertions.New(t).So(thing.Sleeper.Naps, should.Resemble, []time.Duration{time.Hour})
}

func TestActualClockInstanceIsUsefulForTesting(t *testing.T) {
	t.Parallel()

	thing := new(ThingUnderTest)
	now := time.Now().UTC()
	thing.Clock = Freeze(now)
	now2 := thing.CurrentTime()
	assertions.New(t).So(now, should.Resemble, now2)
}

////////////////////////////////////////////////////////////////////////////////////////////

type ThingUnderTest struct {
	*Clock
	*Sleeper
}

func (this *ThingUnderTest) CurrentTime() time.Time {
	return this.Clock.UTCNow()
}

func (this *ThingUnderTest) Sleep(duration time.Duration) {
	this.Sleeper.Sleep(duration)
}
