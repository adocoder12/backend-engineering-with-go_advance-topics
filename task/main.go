package main

import (
	"errors"
	"fmt"
	"sync"
)

var ErrTruckNotFound = errors.New("truck not found")
var ErrTruckExist = errors.New("Truck already exist")
var ErrEmptyID = errors.New("id must not be empty")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
	//mutex
	sync.RWMutex
}

func NewTruckManager() *truckManager {
	return &truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (tm *truckManager) GetTruck(id string) (Truck, error) {
	if id == "" {
		return Truck{}, ErrEmptyID
	}
	tm.RLock()
	defer tm.RUnlock()
	truck, exist := tm.trucks[id]
	if !exist {
		return Truck{}, ErrTruckNotFound
	}
	// copy, so callers can't race with UpdateTruckCargo , we passed the value not its place in memory
	return *truck, nil
}

func (tm *truckManager) AddTruck(id string, cargo int) error {
	if id == "" {
		return ErrEmptyID
	}
	//its lock so we could work in our struct
	tm.Lock()
	defer tm.Unlock()
	if _, exist := tm.trucks[id]; exist {
		return ErrTruckExist
	}
	tm.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (tm *truckManager) UpdateTruckCargo(id string, cargo int) error {
	if id == "" {
		return errors.New("must insert the id")
	}
	tm.Lock()
	defer tm.Unlock()
	truck, exist := tm.trucks[id]
	if !exist {
		return ErrTruckNotFound
	}
	truck.Cargo = cargo
	return nil

}
func (tm *truckManager) RemoveTruck(id string) error {
	tm.Lock()
	defer tm.Unlock()
	if _, exist := tm.trucks[id]; !exist {
		return ErrTruckNotFound
	}
	delete(tm.trucks, id)
	return nil
}

func main() {
	var manager FleetManager = NewTruckManager()

	if err := manager.AddTruck("truck1", 100); err != nil {
		fmt.Println("add:", err)
	}
	if err := manager.AddTruck("truck2", 200); err != nil {
		fmt.Println("add:", err)
	}
	if err := manager.AddTruck("truck2", 300); err != nil {
		fmt.Println("add:", err) // truck already exists
	}

	if err := manager.UpdateTruckCargo("truck1", 10); err != nil {
		fmt.Println("update:", err)
	}

	truck, err := manager.GetTruck("truck1")
	if err != nil {
		fmt.Println("get:", err)
		return
	}
	fmt.Println(truck) // {truck1 10}

	if err := manager.RemoveTruck("nope"); err != nil {
		fmt.Println("remove:", err) // truck not found
	}
}
