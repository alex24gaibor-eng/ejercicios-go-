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

type Tablet struct {
	Medicine
	dosePerTablet  float64
	isPrescription bool
}

func (t *Tablet) DosePerTablet() float64 { return t.dosePerTablet }
func (t *Tablet) IsPrescription() bool   { return t.isPrescription }

func (t *Tablet) SetDosePerTablet(v float64) error {
	if v <= 0 {
		return errors.New("la dosis debe ser mayor a 0")
	}
	t.dosePerTablet = v
	return nil
}

func NewTablet(name, manufacturer, date string, shelfLife int, dose float64, prescription bool) (*Tablet, error) {
	med, err := NewMedicine(name, manufacturer, date, shelfLife)
	if err != nil {
		return nil, err
	}
	t := &Tablet{Medicine: med, isPrescription: prescription}
	if err := t.SetDosePerTablet(dose); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Tablet) ShowDetails() {
	fmt.Println("Tableta:", t.name, "|", t.manufacturer, "|", t.manufactureDate.Format("2006-01-02"),
		"|", t.shelfLife, "meses |", t.dosePerTablet, "mg | Receta:", t.isPrescription)
}

type Syrup struct {
	Medicine
	volume float64
	flavor string
}

func (s *Syrup) Volume() float64 { return s.volume }
func (s *Syrup) Flavor() string  { return s.flavor }

func (s *Syrup) SetVolume(v float64) error {
	if v <= 0 {
		return errors.New("el volumen debe ser mayor a 0")
	}
	s.volume = v
	return nil
}

func NewSyrup(name, manufacturer, date string, shelfLife int, volume float64, flavor string) (*Syrup, error) {
	med, err := NewMedicine(name, manufacturer, date, shelfLife)
	if err != nil {
		return nil, err
	}
	s := &Syrup{Medicine: med, flavor: flavor}
	if err := s.SetVolume(volume); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Syrup) ShowDetails() {
	fmt.Println("Jarabe:", s.name, "|", s.manufacturer, "|", s.manufactureDate.Format("2006-01-02"),
		"|", s.shelfLife, "meses |", s.volume, "ml | Sabor:", s.flavor)
}

func main() {
	t1, _ := NewTablet("Paracetamol", "Genfar", "2025-01-15", 24, 500, false)
	t2, _ := NewTablet("Amoxicilina", "Bayer", "2025-06-01", 18, 250, true)
	tablets := []*Tablet{t1, t2}

	s1, _ := NewSyrup("Tos Stop", "Pfizer", "2025-03-10", 12, 120, "Fresa")
	s2, _ := NewSyrup("Vitaminas", "Roche", "2025-08-20", 36, 200, "Naranja")
	syrups := []*Syrup{s1, s2}

	fmt.Println("--- Inventario ---")
	for _, t := range tablets {
		t.ShowDetails()
	}
	for _, s := range syrups {
		s.ShowDetails()
	}

	fmt.Println("--- Actualización ---")
	tablets[0].SetName("Paracetamol Forte")
	fmt.Println("Nuevo nombre:", tablets[0].Name())

	fmt.Println("--- Caducidad ---")
	for _, t := range tablets {
		fmt.Println(t.Name(), "caduca:", t.ExpirationDate().Format("2006-01-02"))
	}
	for _, s := range syrups {
		fmt.Println(s.Name(), "caduca:", s.ExpirationDate().Format("2006-01-02"))
	}
}