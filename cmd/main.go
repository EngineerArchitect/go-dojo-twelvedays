package main

import (
	"fmt"
	"twelvedays/pkg/song"
)

func main() {
	mysong := song.NewSongBuilder().SetDays(12).Build()
	fmt.Println(mysong.Print())
}
