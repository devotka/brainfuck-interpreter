package commands

import (
	"errors"
	"brainfuck-interpreter/engine"
)

func NewDecrementPointer(service engine.MemoryAccess) *decrementPointer {
	return &decrementPointer{
		memory: service,
	}
}

type decrementPointer struct {
	memory engine.MemoryAccess
}

func (s *decrementPointer) Execute() error {
	err := s.memory.DecrementPointer()
	if err != nil {
		return errors.New("command failure: " + err.Error())
	}

	return nil
}
