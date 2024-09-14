package verse

import (
	"fmt"
	"strings"
)

var DaysAsNumber []string = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var CountAsNumber []string = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve"}
var Possessions map[int]string = map[int]string{
	2:  "turtle doves",
	3:  "french hens",
	4:  "calling birds",
	5:  "golden rings",
	6:  "geese a-laying",
	7:  "swans a-swimming",
	8:  "maids a-milking",
	9:  "ladies dancing",
	10: "lords a-leaping",
	11: "pipers piping",
	12: "drummers drumming",
}

type Verse struct {
	days   int
	header string
	body   string
	footer string
}

func NewVerseBuilder(days int) VerseInterface {
	return &Verse{days: days}
}

type VerseInterface interface {
	Build() *Verse
	Print() string
}

func (v *Verse) addHeader() {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("On the %v day of Christmas\n", DaysAsNumber[v.days-1]))
	builder.WriteString("My true love gave to me:\n")

	v.header = builder.String()
}

func (v *Verse) addBody() {
	var builder strings.Builder

	for day := v.days; day > 1; day-- {
		builder.WriteString(CountAsNumber[day-1])
		builder.WriteString(" ")
		builder.WriteString(Possessions[day])
		if day == 2 {
			builder.WriteString(" and")
		}
		builder.WriteString("\n")
	}
	v.body = builder.String()
}

func (v *Verse) addFooter() {
	v.footer = "A partridge in a pear tree."
}

func (v *Verse) Build() *Verse {
	v.addHeader()
	v.addBody()
	v.addFooter()
	return v
}

func (v *Verse) Print() string {
	var builder strings.Builder
	builder.WriteString(v.header)
	builder.WriteString(v.body)
	builder.WriteString(v.footer)
	return builder.String()
}
