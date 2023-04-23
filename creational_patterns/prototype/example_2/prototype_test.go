package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()

	return <-out
}

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) SetupTest() {}

func (s *TestSuite) TestFilePrint() {
	type parameter struct {
		f *file
	}

	testCases := []struct {
		name   string
		args   parameter
		expect string
	}{
		{
			name:   "Abnormal",
			args:   parameter{f: &file{name: ""}},
			expect: "-\n",
		},
		{
			name:   "Normal 1",
			args:   parameter{f: &file{name: "file1"}},
			expect: "-file1\n",
		},
		{
			name:   "Normal 2",
			args:   parameter{f: &file{name: "file2"}},
			expect: "-file2\n",
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			out := captureOutput(func() {
				tc.args.f.print("-")
			})
			assert.Equal(s.T(), tc.expect, out)
		})
	}
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
