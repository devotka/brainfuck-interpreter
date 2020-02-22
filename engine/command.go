package engine

type Command interface {
	Execute() error
}
