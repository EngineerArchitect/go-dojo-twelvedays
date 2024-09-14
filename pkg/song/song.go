package song

import (
	"strings"
	. "twelvedays/pkg/verse"
)

type Song struct {
	verses []Verse
	days   int
}

type SongInterface interface {
	SetDays(days int) *Song
	Build() *Song
	Print() string
}

func NewSongBuilder() SongInterface {
	return &Song{}
}

func (s *Song) SetDays(days int) *Song {
	s.days = days
	return s
}

func (s *Song) Build() *Song {
	s.verses = make([]Verse, s.days)

	for day := 1; day <= s.days; day++ {
		var verse = NewVerseBuilder(day).Build()
		s.verses[day-1] = *verse
	}

	return s
}

func (s *Song) Print() string {
	var builder strings.Builder

	for i, verse := range s.verses {
		builder.WriteString(verse.Print())

		if i != len(s.verses)-1 {
			builder.WriteString("\n")
			builder.WriteString("\n")
		}
	}

	return builder.String()
}
