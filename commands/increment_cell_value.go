package commands

import (
	"brainfuck-interpreter/engine"
)

var _ engine.Command = (*incrementCellValue)(nil)

func NewIncrementCellValue(memoryAccess engine.MemoryAccess) *incrementCellValue {
	return &incrementCellValue{
		memoryAccess: memoryAccess,
	}
}

type incrementCellValue struct {
	memoryAccess engine.MemoryAccess
}

func (s *incrementCellValue) Execute() error {
	s.memoryAccess.IncrementCellValue()
	return nil
}
