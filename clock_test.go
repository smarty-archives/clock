package clock

import (
	"testing"
	"time"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestNow(t *testing.T) {
	actual := Now()
	expected := time.Now()
	assertions.New(t).So(actual, should.HappenWithin, time.Millisecond, expected)

	actual = UTCNow()
	expected = time.Now().UTC()
	assertions.New(t).So(actual, should.HappenWithin, time.Millisecond, expected)
}

func TestFreezeWithNoTimesGiven(t *testing.T) {
	freeze := func() { Freeze() }
	assertions.New(t).So(freeze, should.Panic)
}

func TestFreezeAndRestore(t *testing.T) {
	So := assertions.New(t).So
	a := time.Now().Add(-time.Second)
	b := a.Add(time.Second)
	Freeze(a, b)

	t.Log("Now() should forever cycle through a and b...")
	for x := 0; x < 100; x++ {
		So(Now(), should.Resemble, a)
		So(Now(), should.Resemble, b)
	}

	Restore()

	t.Log("Now() should now give back the current, incrementing time...")
	c := Now()
	d := Now()

	So(b, should.HappenBefore, c)
	So(c, should.HappenBefore, d)
}

func TestUTCFreezeAndRestore(t *testing.T) {
	So := assertions.New(t).So
	a := time.Now().UTC().Add(-time.Second)
	b := a.Add(time.Second)
	Freeze(a, b)

	t.Log("Now() should forever cycle through a and b...")
	for x := 0; x < 100; x++ {
		So(UTCNow(), should.Resemble, a)
		So(UTCNow(), should.Resemble, b)
	}

	Restore()

	t.Log("Now() should now give back the current, incrementing time...")
	c := UTCNow()
	d := UTCNow()

	So(b, should.HappenBefore, c)
	So(c, should.HappenBefore, d)
}
