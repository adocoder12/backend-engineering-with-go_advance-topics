package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrTruckNotFound  = errors.New("Truck not found")
	ErrNotImplemented = errors.New("Truck not implemented")
)

type TruckInterface interface {
	LoadCargo() error
	UnLoadCargo() error
}

type Truck struct {
	ID    string
	Cargo int
}

func (t *Truck) LoadCargo() error {
	t.Cargo += 1
	return nil
}
func (t *Truck) UnLoadCargo() error {
	t.Cargo = 0

	return nil
}

type ElectrictTruck struct {
	ID      string
	Cargo   int
	Battery float64
}

func (t *ElectrictTruck) LoadCargo() error {
	t.Cargo += 1
	t.Battery -= 1
	return nil
}
func (t *ElectrictTruck) UnLoadCargo() error {
	t.Cargo = 0

	return nil
}
func processTruck(truck TruckInterface) error {
	fmt.Printf("processing Truck: %+v\n", truck)
	if err := truck.LoadCargo(); err != nil {
		if errors.Is(err, ErrTruckNotFound) {
			log.Fatalf("Truck not found: %s", err)

		}
		return fmt.Errorf("error loading cargo: %w", err)
	}
	if err := truck.UnLoadCargo(); err != nil {
		if errors.Is(err, ErrTruckNotFound) {
			log.Fatalf("Truck not found: %s", err)

		}
		return fmt.Errorf("error loading cargo: %w", err)
	}
	return nil
}

func main() {
	//address truckI D: 0x6ebc14490020
	truckID := 12
	//address AnotherTruckID :0x6ebc14490020->truckID
	AnotherTruckID := &truckID
	fmt.Println("truck id pointer", &truckID)
	fmt.Println("Another id which point to truckid", *AnotherTruckID)

}
