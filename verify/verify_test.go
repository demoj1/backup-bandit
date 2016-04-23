package verify_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/diman3241/backupbandit/verify"
	. "github.com/smartystreets/goconvey/convey"
)

func TestVerify(t *testing.T) {
	Convey("Test verify func", t, func() {
		path := verify.PathVerifyInfo{
			Path:      "/tmp/__verify_test",
			ValidSize: 0,
		}

		So(path, ShouldNotBeNil)

		Convey("Create tmp folder", func() {
			err := os.Mkdir("/tmp/__verify_test", 0777)

			So(err, ShouldNotBeNil)
		})

		Convey("Test null size folder", func() {
			res := verify.Verify(&path)

			So(res, ShouldBeTrue)
		})

		Convey("Create tmp file", func() {
			path.ValidSize = 6 * 1024

			Convey("Write data to file", func() {
				err := ioutil.WriteFile("/tmp/__verify_test/test_file", []byte("Hello world"), 777)
				So(err, ShouldBeNil)
			})

			Convey("Test verify path", func() {
				res := verify.Verify(&path)

				So(res, ShouldBeTrue)
			})

			file := verify.FileVerifyInfo{
				FilePattern: ".*_file",
				Size:        11,
			}

			So(file, ShouldNotBeNil)

			Convey("Test verify path in included file", func() {
				Convey("Add new file info", func() {
					path.FileVerifyInfos = append(path.FileVerifyInfos, file)

					So(len(path.FileVerifyInfos), ShouldEqual, 1)
				})

				res := verify.Verify(&path)

				So(path.FileVerifyInfos[0].Info, ShouldBeNil)
				So(res, ShouldBeTrue)
			})
		})
	})
}
