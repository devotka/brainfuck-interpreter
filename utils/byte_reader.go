package utils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type ByteReader interface {
	ReadByte() (byte, error)
}

type FileReader interface {
	ByteReader
	Close() error
}

func NewFileReader(filepath string) (*fileReader, error) {
	f, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	fileReader := fileReader{
		file: f,
		reader:bufio.NewReader(f),
	}
	return &fileReader, nil
}

type fileReader struct {
	file *os.File
	reader *bufio.Reader
}

func (f *fileReader) ReadByte() (byte, error) {
	for {
		b, err := f.reader.ReadByte()
		if err!= nil {
			return 0, err
		}

		if err != nil {
			return 0, err
		}
		if strings.Contains("+-<>.,[]1234567890", string(b)) {
			return b, nil
		}
	}

	return 0, errors.New("")
}

func (f *fileReader) Close() error {
	return f.file.Close()
}
