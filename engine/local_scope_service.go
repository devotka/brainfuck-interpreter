package engine

func NewLocalScopeService() *localScopeService {
	return &localScopeService{}
}

type LocalScopeService interface {
	Push(scope LocalScope)
	Pop() LocalScope
	Get() LocalScope
}

type scopeEntry struct {
	current  LocalScope
	previous *scopeEntry
}

type localScopeService struct {
	currentScope *scopeEntry
}

func (s *localScopeService) Push(scope LocalScope) {
	newScope := &scopeEntry{
		current:  scope,
		previous: s.currentScope,
	}

	s.currentScope = newScope
}

func (s *localScopeService) Pop() LocalScope {
	scope := s.currentScope
	if scope != nil {
		s.currentScope = s.currentScope.previous

		return scope.current
	}

	return nil
}

func (s *localScopeService) Get() LocalScope {
	if s.currentScope == nil {
		return nil
	}

	return s.currentScope.current
}
