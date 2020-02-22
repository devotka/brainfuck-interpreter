package main

import (
	"siiMod/commands"
	"siiMod/engine"
	"siiMod/utils"
)

type commandSelector struct {
	writer       utils.ByteWriter
	reader       utils.ByteReader
	memoryAccess engine.MemoryAccess
}

func (s commandSelector) selectCommand(symbol string) engine.Command {
	switch symbol {

	case ">":
		return commands.NewIncrementPointer(s.memoryAccess)

	case "<":
		return commands.NewDecrementPointer(s.memoryAccess)

	case "+":
		return commands.NewIncrementCellValue(s.memoryAccess)

	case "-":
		return commands.NewDecrementCellValue(s.memoryAccess)

	case ".":
		return commands.NewGetChar(s.memoryAccess, s.writer)

	case ",":
		return commands.NewPutChar(s.memoryAccess, s.reader)

	case "[":
		return commands.NewLoop(s.memoryAccess)

	case "]":
		return commands.NewLoopClose()
	}

	return nil
}
