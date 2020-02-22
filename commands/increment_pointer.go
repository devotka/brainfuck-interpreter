package commands

import (
	"siiMod/engine"
)

var _ engine.Command = (*incrementPointer)(nil)

func NewIncrementPointer(memoryAccess engine.MemoryAccess) *incrementPointer {
	return &incrementPointer{
		memoryAccess: memoryAccess,
	}
}

type incrementPointer struct {
	memoryAccess engine.MemoryAccess
}

func (s *incrementPointer) Execute() error {
	s.memoryAccess.IncrementPointer()
	return nil
}
