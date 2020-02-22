package commands

import (
	"brainfuck-interpreter/engine"
)

var _ engine.Command = (*decrementCellValue)(nil)

func NewDecrementCellValue(memoryAccess engine.MemoryAccess) *decrementCellValue {
	return &decrementCellValue{
		memoryAccess: memoryAccess,
	}
}

type decrementCellValue struct {
	memoryAccess engine.MemoryAccess
}

func (s *decrementCellValue) Execute() error {
	s.memoryAccess.DecrementCellValue()
	return nil
}
