package engine

type LocalScopeClose interface {
	Command
	Close() error
}
