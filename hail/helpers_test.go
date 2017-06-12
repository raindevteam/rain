package hail

import (
	"bytes"
	"io"
	"os"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	output := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		output <- buf.String()
	}()
	w.Close()
	os.Stdout = old
	return <-output
}
