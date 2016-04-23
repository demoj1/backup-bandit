package robber_test

import (
	"testing"

	"github.com/diman3241/backupbandit/robber"
	. "github.com/smartystreets/goconvey/convey"
)

func TestToolsSetting(t *testing.T) {
	Convey("Test tools setting struct", t, func() {
		tool := robber.ToolsSetting{}

		Convey("Append tool to tools setting", func() {
			tool.Tools = append(tool.Tools, robber.Tool{
				Path: "/usr/bin/uptime",
				Groups: []string{
					"up",
					"users",
					"avr",
				},
				Regex: `up (?P<up>.*),\s*(?P<users>\d+) users,\s*load average: (?P<avr>.*\d)`,
			})

			So(len(tool.Tools), ShouldEqual, 1)
			m, err := tool.Tools[0].ParseOut()

			So(err, ShouldBeNil)
			So(m, ShouldNotBeNil)

			Convey("Correct parse", func() {
				So(len(m), ShouldEqual, 1)
				So(m[0]["up"], ShouldNotBeNil)
				So(m[0]["users"], ShouldNotBeNil)
				So(m[0]["avr"], ShouldNotBeNil)
			})
		})
	})
}
