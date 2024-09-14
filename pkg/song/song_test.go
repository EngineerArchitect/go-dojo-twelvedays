package song

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func printSongLines(title string, s *Song) {
	fmt.Println(title)
	for _, line := range strings.Split(s.Print(), "\n") {
		fmt.Println(line)
	}
}

func Test_SongCreator(t *testing.T) {
	Convey("Given Song Creator", t, func() {
		Convey("When Song is built for 2 days", func() {
			song := NewSongBuilder().SetDays(2).Build()

			Convey("Should create 2 verses", func() {
				So(song, ShouldNotBeNil)

				printSongLines("Song with 2 verses", song)
				So(song.Print(), ShouldEqual, "On the first day of Christmas\n"+
					"My true love gave to me:\n"+
					"A partridge in a pear tree.\n\n"+
					"On the second day of Christmas\n"+
					"My true love gave to me:\n"+
					"Two turtle doves and\n"+
					"A partridge in a pear tree.")
			})
		})
	})
}
