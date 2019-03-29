package greeter

import (
	"errors"
	"fmt"
)

// ErrInvalidName is the error returned if name is empty string.
var ErrInvalidName = errors.New("greeter: name required")

// NewGreeter allocates and returns new Greeter.
func NewGreeter() *Greeter {
	return &Greeter{}
}

// Greeter is for greeting.
type Greeter struct {
}

// SayHello says hello.
func (g *Greeter) SayHello(name string) (string, error) {
	if len(name) == 0 {
		return "", ErrInvalidName
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}
