package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"siiMod/engine"
	"siiMod/utils"
)

func main() {
	var err error
	args := os.Args[1:]

	option := ""
	if len(args) > 0 {
		option = args[0]
	}

	switch option {
	case "-c":
		err = readFromConsole()
	case "-f":

		if len(args) < 2 {
			fmt.Println("Please give a file path")
			return
		}

		file := args[1]

		err = readFromFile(file)
	default:
		loadTestData()
	}

	if err != nil {
		fmt.Println(err)
	}
}

func readFromConsole() error {
	ma := engine.NewMemoryAccess()
	stdinReader := bufio.NewReader(os.Stdin)
	scopeService := engine.NewLocalScopeService()
	commandSelector := commandSelector{
		memoryAccess: ma,
		writer:       utils.NewConsoleWriter(),
		reader:       stdinReader,
	}

	for {
		c, err := stdinReader.ReadByte()
		if err == io.EOF {
			break
		}

		symbol := string(c)
		command := commandSelector.selectCommand(symbol)
		if command != nil {
			if err := engine.InterpretCommand(command, scopeService, ma); err != nil {
				return err
			}
		}
	}

	return nil
}

func readFromFile(path string) error {
	ma := engine.NewMemoryAccess()
	scopeService := engine.NewLocalScopeService()
	commandSelector := commandSelector{
		memoryAccess: ma,
		writer:       utils.NewConsoleWriter(),
		reader:       bufio.NewReader(os.Stdin),
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	for _, c := range content {

		symbol := string(c)
		command := commandSelector.selectCommand(symbol)
		if command != nil {
			if err := engine.InterpretCommand(command, scopeService, ma); err != nil {
				return err
			}
		}
	}
	return nil
}
