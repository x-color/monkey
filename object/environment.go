package object

// Environment is store of object generated in executing
type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnvironment returns new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// NewEnclosedEnvironment returns new environment enclosing given environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get returns object stored in environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set stores object in environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
