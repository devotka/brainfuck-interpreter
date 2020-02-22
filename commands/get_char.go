package commands

import (
	"brainfuck-interpreter/engine"
	"brainfuck-interpreter/utils"
)

var _ engine.Command = (*getChar)(nil)

func NewGetChar(memoryAccess engine.MemoryAccess, writer utils.ByteWriter) *getChar {
	return &getChar{
		memoryAccess: memoryAccess,
		byteWriter:   writer,
	}
}

type getChar struct {
	memoryAccess engine.MemoryAccess
	byteWriter   utils.ByteWriter
}

func (s *getChar) Execute() error {
	value := s.memoryAccess.GetCellValue()
	return s.byteWriter.WriteByte(value)
}
