package models

import (
	"fmt"
)

type Member struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// NewMember creates a new member given a id, password and name that can't be empty.
func NewMember(id string, password string, name string) (*Member, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id")
	}
	if password == "" {
		return nil, fmt.Errorf("empty password")
	}
	if name == "" {
		return nil, fmt.Errorf("empty name")
	}

	return &Member{id, password, name}, nil
}

func (m *Member) CheckPassword(password string) bool {
	if m.Password == password {
		return true
	}

	return false
}
