package main

import (
	"errors"
	"fmt"
	"log"
)

type truck struct {
	ID    string
	Cargo int
}

var (
	ErrTruckNotFound  = errors.New("Truck not found")
	ErrNotImplemented = errors.New("Truck not implemented")
)

func (t *truck) LoadCargo() error {
	return nil
}
func (t *truck) UnLoadCargo() error {
	return nil
}

func processTruck(truck truck) error {
	fmt.Printf("processing Truck: %s\n", truck.ID)
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}
	if err := truck.UnLoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}
	return ErrNotImplemented
}
func main() {
	trucks := []truck{
		{ID: "Truck 1"},
		{ID: "Truck 2"},
		{ID: "Truck 3"},
		{ID: "Truck 4"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.ID)
		if err := processTruck(truck); err != nil {
			if errors.Is(err, ErrTruckNotFound) {
				log.Fatalf("Truck not found: %s", err)

			}
			log.Fatalf("error processing truck: %s", err)
		}
	}

}
