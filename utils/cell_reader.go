package utils

import (
	"fmt"
)

type ByteWriter interface {
	WriteByte(b int) error
}

func NewConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}

type consoleWriter struct{}

func (wr *consoleWriter) WriteByte(b int) error {
	fmt.Print(string(b))

	return nil
}

func init() {
}
