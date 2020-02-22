package commands

import (
	"brainfuck-interpreter/engine"
)

var _ engine.LocalScope = (*loop)(nil)

func NewLoop(memoryAccess engine.MemoryAccess) *loop {
	return &loop{
		canExecutePartially: memoryAccess.GetCellValue() != 0,
		memoryAccess: memoryAccess,
	}
}

type loop struct {
	memoryAccess engine.MemoryAccess
	steps        []engine.Command
	canExecutePartially bool
}

func (l *loop) Execute() error {
	for l.memoryAccess.GetCellValue() != 0 {
		for _, c := range l.steps {
			err := c.Execute()

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (l *loop) Append(s engine.Command) {
	l.steps = append(l.steps, s)
}
func (l *loop) CanExecutePartially() bool {
	return l.canExecutePartially
}
