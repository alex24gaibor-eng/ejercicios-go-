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

func (m *Medicine) SetManufacturer(v string) error {
	if v == "" {
		return errors.New("el fabricante no puede estar vacío")
	}
	m.manufacturer = v
	return nil
}

func (m *Medicine) SetManufactureDate(v string) error {
	d, err := time.Parse("2006-01-02", v)
	if err != nil {
		return errors.New("fecha inválida")
	}
	m.manufactureDate = d
	return nil
}

func (m *Medicine) SetShelfLife(v int) error {
	if v <= 0 {
		return errors.New("la vida útil debe ser positiva")
	}
	m.shelfLife = v
	return nil
}

func NewMedicine(name, manufacturer, date string, shelfLife int) (Medicine, error) {
	m := Medicine{}
	if err := m.SetName(name); err != nil {
		return m, err
	}
	if err := m.SetManufacturer(manufacturer); err != nil {
		return m, err
	}
	if err := m.SetManufactureDate(date); err != nil {
		return m, err
	}
	if err := m.SetShelfLife(shelfLife); err != nil {
		return m, err
	}
	return m, nil
}

func (m *Medicine) ExpirationDate() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}
