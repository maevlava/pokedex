package tests

import (
	"bytes"
	"os"
)

func CaptureStdOutput() (std *os.File, read *os.File, write *os.File) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	return oldStdout, r, w
}
func RestoreStdOutput(std *os.File, read *os.File, write *os.File) bytes.Buffer {
	write.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(read)
	os.Stdout = std
	return buf
}
