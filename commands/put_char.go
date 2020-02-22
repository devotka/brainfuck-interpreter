package commands

import (
	"brainfuck-interpreter/engine"
	"brainfuck-interpreter/utils"
)

var _ engine.Command = (*PutChar)(nil)

func NewPutChar(memoryAccess engine.MemoryAccess, reader utils.ByteReader) *PutChar {


	return &PutChar{
		memoryAccess: memoryAccess,
		byteReader:   reader,
	}
}

type PutChar struct {
	value        byte
	memoryAccess engine.MemoryAccess
	byteReader   utils.ByteReader
}

func (s *PutChar) Execute() error {
	value, err := s.byteReader.ReadByte()
	if err != nil {
		return err
	}
	s.memoryAccess.SetCellValue(int(value))
	return nil
}
