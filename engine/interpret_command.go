package engine

import (
	"errors"
)

func InterpretCommand(command Command, scopeService LocalScopeService) error {
	scope := scopeService.Get()

	switch value := command.(type) {
	case LocalScope:
		if scope != nil {
			scope.Append(command)
		}
		scopeService.Push(value)

	case LocalScopeClose:
		if scope != nil {
			return scopeService.Pop().Execute()
		}

		return errors.New("wrong input - missing loop opening bracket")

	default:
		if scope != nil {
			scope.Append(command)
		}
		if canExecuteCommand(scope) {
			err := command.Execute()
			if err != nil {
				return err
			}
		}


	}

	return nil
}

func canExecuteCommand(scope LocalScope) bool {
	if scope == nil {
		return true
	}

	return scope.CanExecutePartially()
}