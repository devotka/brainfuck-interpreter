package commands

import (
	"siiMod/engine"
)

var _ engine.LocalScopeClose = (*loopClose)(nil)

func NewLoopClose() *loopClose {
	return &loopClose{}
}

type loopClose struct{}

func (s *loopClose) Execute() error {
	return nil
}

func (s *loopClose) Close() error {
	return nil
}
