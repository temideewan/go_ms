package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck interface {
	LoadCargo() error
	UnLoadCargo() error
}
type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 2
	return nil
}
func (t *NormalTruck) UnLoadCargo() error {
	t.cargo = 0
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery -= 1
	return nil
}
func (e *ElectricTruck) UnLoadCargo() error {
	e.cargo = 0
	e.battery -= 1
	return nil
}

func processTruck(truck Truck) error {
	fmt.Printf("processing truck %+v \n", truck)
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo %w", err)
	}
	if err := truck.UnLoadCargo(); err != nil {
		return fmt.Errorf("Error unloading cargo %w", err)
	}
	return nil
}

func main() {
	t := NormalTruck{cargo: 0}

	fillTruckCargo(&t)
	log.Println(t)
}

func fillTruckCargo(t *NormalTruck) {
	t.cargo = 100
}
