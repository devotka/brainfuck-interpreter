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
		scope := scopeService.Pop()
		if scopeService.Get() == nil {

			err := scope.Execute()
			if err != nil {
				return err
			}
		}
	default:
		if scope != nil {
			scope.Append(command)
			return nil
		}

		err := command.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}
