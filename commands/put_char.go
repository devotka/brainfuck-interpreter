package commands

import (
	"siiMod/engine"
	"siiMod/utils"
)

var _ engine.Command = (*putChar)(nil)

func NewPutChar(memoryAccess engine.MemoryAccess, reader utils.ByteReader) *putChar {
	return &putChar{
		memoryAccess: memoryAccess,
		byteReader:   reader,
	}
}

type putChar struct {
	memoryAccess engine.MemoryAccess
	byteReader   utils.ByteReader
}

func (s *putChar) Execute() error {
	readVal, err := s.byteReader.ReadByte()
	if err != nil {
		return err
	}

	s.memoryAccess.SetCellValue(int(readVal))
	return nil
}
