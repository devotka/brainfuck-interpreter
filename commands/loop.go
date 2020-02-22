package commands

import (
	"brainfuck-interpreter/engine"
)

var _ engine.LocalScope = (*loop)(nil)

func NewLoop(memoryAccess engine.MemoryAccess) *loop {
	return &loop{
		memoryAccess: memoryAccess,
	}
}

type loop struct {
	memoryAccess engine.MemoryAccess
	steps        []engine.Command
}

func (w *loop) Execute() error {
	for w.memoryAccess.GetCellValue() != 0 {
		for _, c := range w.steps {
			err := c.Execute()

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *loop) Append(s engine.Command) {
	w.steps = append(w.steps, s)
}
