package amap

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestAtomicMapGetSet(t *testing.T) {
	cv.Convey("AtomicMap should Get and Set", t, func() {
		am := NewAtomicMap()
		am.Set(1, "one")
		am.Set(2, "two")

		g1, already := am.Get(1)
		cv.So(g1, cv.ShouldEqual, "one")
		cv.So(already, cv.ShouldEqual, true)

		g2, already := am.Get(2)
		cv.So(g2, cv.ShouldEqual, "two")
		cv.So(already, cv.ShouldEqual, true)
	})
}
