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

