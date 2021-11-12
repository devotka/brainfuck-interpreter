# Brainfuck interpreter

OOP / "interfaced" variation of brrainfuck interpreter interpreter.

## Application execution and commands

Program consumes console input and executes interprets commands. 
Available commands:
- `+` increments current memory cell
- `-` decrements current memory cell
- `<` increments memory current address
- `>` decrements memory current address
- `.` display current cell value
- `,` set current cell value
- `[` open loop
- `]` close loop

Each independent command is executed just after reading them from stream.


### Running application

```// build app
go build brainfuck-interpreter 

// running with console input
./brainfuck-interpreter -c`

// running with input from file
./brainfuck-interpreter -f <filepath>
```

### Example test data

Repo contains example input data, you can easily test app using commands like below:
```./brainfuck-interpreter -f test_data/hello_world.txt
./brainfuck-interpreter -f test_data/cells_32.txt
./brainfuck-interpreter -f test_data/beer_bottles.txt
./brainfuck-interpreter -f test_data/division.txt
./brainfuck-interpreter -f test_data/following_number_squares.txt
./brainfuck-interpreter -f test_data/triangle.txt
```

## Sources

https://esolangs.org/wiki/Brainfuck_algorithms
https://en.wikipedia.org/wiki/Brainfuck
https://copy.sh/brainfuck/
http://www.hevanet.com/cristofd/brainfuck/ 
