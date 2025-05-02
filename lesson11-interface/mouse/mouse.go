package mouse

import (
	"errors"
	"strings"

	"quoctuan.com/hoc-golang/service"
)

type Mouse struct {
	Name string `json:"name"`
}

// Create a constructor for Mouse
func New(name string) (service.Animal, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if len(name) > 50 {
		return nil, errors.New("name is too long (max 50 characters)")
	}

	return &Mouse{
		Name: name,
	}, nil
}

func (m *Mouse) GetName() string {
	return m.Name
}

func (m *Mouse) Speak() string {
	return "Chit chit!"
}

func (m *Mouse) Run() string {
	return "Rat nhanh"
}