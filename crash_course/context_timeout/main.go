package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type ContextKey string

var UserIdKey ContextKey = "userId"

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

func processTruck(ctx context.Context, truck Truck) error {
	fmt.Printf("processing truck %+v \n", truck)

	fmt.Println("userId", ctx.Value(UserIdKey))

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	// simulate processing time
	delay := time.Second * 3
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo %w", err)
	}
	if err := truck.UnLoadCargo(); err != nil {
		return fmt.Errorf("Error unloading cargo %w", err)
	}
	fmt.Printf("finished processing truck %+v \n", truck)
	return nil
}

func processFleet(ctx context.Context, trucks []Truck) error {
	var wg sync.WaitGroup
	for _, t := range trucks {
		wg.Add(1)
		go func(t Truck) {
			if err := processTruck(ctx, t); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(t)
	}
	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIdKey, 42)
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(ctx, fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
	}

	fmt.Println("Done processing all fleets")
	// time.Sleep(time.Second * 5)
}
