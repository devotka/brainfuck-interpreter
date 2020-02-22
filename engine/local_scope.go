package engine

type LocalScope interface {
	Command
	Append(step Command)
}
