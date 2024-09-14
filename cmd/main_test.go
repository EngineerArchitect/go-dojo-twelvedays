package main

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"strings"

	"io"
	"log"
	"os"
	"sync"
	"testing"
)

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		if _, err := io.Copy(&buf, reader); err != nil {
			panic(err)
		}
		out <- buf.String()
	}()
	wg.Wait()
	f()
	if err := writer.Close(); err != nil {
		panic(err)
	}
	return <-out
}

func countVerses(content string) int {
	paragraphs := strings.Split(content, "\n\n")
	return len(paragraphs)
}

func TestRun(t *testing.T) {

	Convey("When executing run()", t, func() {
		re := captureOutput(func() {
			main()
		})

		Convey("Should return content", func() {
			So(re, ShouldNotBeEmpty)
		})

		Convey("Should have 12 verses", func() {
			So(countVerses(re), ShouldEqual, 12)
		})

	})

}
