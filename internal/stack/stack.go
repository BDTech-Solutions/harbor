package stack

// Stack defines the contract that every Harbor-managed stack must implement.
// Adding a new stack (Django, Rails, plain PHP, etc.) means implementing
// this interface — nothing else in the codebase needs to change.
type Stack interface {
	// Init creates a brand-new project in the given directory.
	// The directory state preconditions (empty, existing, etc.)
	// are the responsibility of each implementation.
	Init(dir string) error

	// Up starts the development environment for an existing project.
	Up(dir string) error

	// Down stops the development environment.
	Down(dir string) error
}
