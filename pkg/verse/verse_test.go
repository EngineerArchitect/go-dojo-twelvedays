package verse

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func splitString(v *Verse) []string {
	return strings.Split(v.Print(), "\n")
}

func printVerseLines(title string, v *Verse) {
	fmt.Println(title)
	for _, line := range strings.Split(v.Print(), "\n") {
		fmt.Println(line)
	}
}

func Test_VerseCreator(t *testing.T) {
	Convey("Builder should add 1st verse line", t, func() {
		v := NewVerseBuilder(1).
			Build()

		So(v, ShouldNotBeNil)
		line1 := splitString(v)[0]
		So(line1, ShouldEqual, "On the first day of Christmas")
	})

	Convey("Builder should add 2nd verse line", t, func() {
		v := NewVerseBuilder(1).
			Build()

		So(v, ShouldNotBeNil)
		line2 := splitString(v)[1]
		So(line2, ShouldEqual, "My true love gave to me:")
	})

	Convey("Builder should footer at end", t, func() {
		v := NewVerseBuilder(1).
			Build()

		So(v, ShouldNotBeNil)
		lines := splitString(v)

		lastLine := lines[len(lines)-1]
		So(lastLine, ShouldEqual, "A partridge in a pear tree.")
	})

	Convey("Builder should Not add Body for day 1", t, func() {
		v := NewVerseBuilder(1).
			Build()

		So(v, ShouldNotBeNil)
		lines := splitString(v)
		So(len(lines), ShouldEqual, 3)
	})

	Convey("Builder should add Body for day2", t, func() {
		v := NewVerseBuilder(2).
			Build()

		So(v, ShouldNotBeNil)
		lines := splitString(v)
		So(len(lines), ShouldEqual, 4)
	})

	Convey("Builder should create day 3", t, func() {
		v := NewVerseBuilder(3).
			Build()

		printVerseLines("Verse for Day 3", v)

		So(v, ShouldNotBeNil)
		So(v.Print(), ShouldEqual,
			"On the third day of Christmas\n"+
				"My true love gave to me:\n"+
				"Three french hens\n"+
				"Two turtle doves and\n"+
				"A partridge in a pear tree.")
	})

}
