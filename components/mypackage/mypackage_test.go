package mypackage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBar(t *testing.T) {

	Convey("should return sting", t, func() {
		value := Bar()
		So(value, ShouldEqual, "bar")
	})
}
