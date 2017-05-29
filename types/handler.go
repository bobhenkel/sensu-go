package types

import "errors"

// A Handler is a handler specification.
type Handler struct {
	// Name is the unique identifier for a handler.
	Name string `json:"name"`

	// Type is the handler type, i.e. pipe.
	Type string `json:"type"`

	// Mutator is the handler event data mutator.
	Mutator string `json:"mutator,omitempty"`

	// Command is the command to be executed for a pipe handler.
	Command string `json:"command,omitempty"`

	// Timeout is the handler timeout in seconds.
	Timeout int `json:"timeout"`

	Socket HandlerSocket `json:"socket,omitempty"`

	// Handlers is a list of handlers for a handler set.
	Handlers []string `json:"handlers,omitempty"`

	// Env is a list of environment variables to use with command execution
	Env []string `json:"environment,omitempty"`
}

// HandlerSocket contains configuration for a TCP or UDP handler.
type HandlerSocket struct {
	// Host is the socket peer address.
	Host string `json:"host"`

	// Port is the socket peer port.
	Port int `json:"port"`
}

// Validate returns an error if the handler does not pass validation tests.
func (h *Handler) Validate() error {
	err := validateName(h.Name)
	if err != nil {
		return errors.New("handler name " + err.Error())
	}

	err = validateHandlerType(h.Type)
	if err != nil {
		return errors.New("handler type " + err.Error())
	}

	return nil
}

// FixtureHandler returns a Handler fixture for testing.
func FixtureHandler(name string) *Handler {
	return &Handler{
		Name:    name,
		Type:    "pipe",
		Command: "command",
	}
}