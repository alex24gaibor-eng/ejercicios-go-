package main

import (
	"errors"
	"fmt"
	"time"
)

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int
}

func (m *Medicine) Name() string               { return m.name }
func (m *Medicine) Manufacturer() string       { return m.manufacturer }
func (m *Medicine) ManufactureDate() time.Time { return m.manufactureDate }
func (m *Medicine) ShelfLife() int             { return m.shelfLife }

func (m *Medicine) SetName(v string) error {
	if v == "" {
		return errors.New("el nombre no puede estar vacío")
	}
	m.name = v
	return nil
}
