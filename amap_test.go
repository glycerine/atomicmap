package amap

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestAtomicMapIntStringGetSet(t *testing.T) {
	cv.Convey("AtomicMapIntString should Get and Set", t, func() {
		am := NewAtomicMapIntString()
		am.Set(1, "one")
		am.Set(2, "two")

		g1, already := am.Get2(1)
		cv.So(g1, cv.ShouldEqual, "one")
		cv.So(already, cv.ShouldEqual, true)

		g2 := am.Get(2)
		cv.So(g2, cv.ShouldEqual, "two")

		// non-existant keys should be indicated
		// as per regular map
		cv.So(am.Get(99), cv.ShouldEqual, "")

		g3, already := am.Get2(99)
		cv.So(g3, cv.ShouldEqual, "")
		cv.So(already, cv.ShouldEqual, false)

		// Del should be available
		am.Del(2)
	})
}
