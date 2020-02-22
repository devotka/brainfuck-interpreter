package engine

func InterpretCommand(command Command, scopeService LocalScopeService, ma MemoryAccess) error {
	scope := scopeService.Get()

	switch value := command.(type) {
	case LocalScope:
		if scope != nil {
			scope.Append(command)
		}
		scopeService.Push(value)

	case LocalScopeClose:
		err := scopeService.Pop().Execute()
		if err != nil {
			return err
		}
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