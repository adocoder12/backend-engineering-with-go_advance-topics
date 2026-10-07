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
	//the power  of interfaces is that apply composition so we wouldnt be able to change truck battery cuz this processTruck only access to its method
	// if i woudl like to change i rather would do in LoadCargo() which access to its attributes
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

	truck := &Truck{ID: "Truck 1"}
	Etruck := &ElectrictTruck{ID: "ETruck 1", Cargo: 0, Battery: 10}

	fmt.Printf("Truck %s arrived.\n", truck.ID)
	if err := processTruck(truck); err != nil {
		log.Fatalf("error processing truck: %s", err)
	}
	if err := processTruck(Etruck); err != nil {
		log.Fatalf("error processing truck: %s", err)
	}
	fmt.Printf("truck status %+v\n", truck)
	fmt.Printf("Etruck status %+v\n", Etruck)
	fmt.Printf("truck %s Departed!\n", truck.ID)

}

func UnmarshallExampleInMapInterface() {
	//Common use of empty interface:
	// For example when working with json and want to unmarshall if dont know what the payload is
	// var person = interface{}

	// another way would be this to achieve same here ,this is a map of key: string , which can accept any value for example : struct , string,int or even another interface
	//another use of this map[string]interface{} is map[string]any
	person := make(map[string]any, 0)
	person["name"] = "ado"
	person["age"] = 33
	//here we can see we can set any value the problem is that we are not able to access to for example person.age
	// so we achieve that with this (we have to expesify the value : .(int) or .(string))
	age, exist := person["age"].(int)
	if exist {
		fmt.Printf("Person age is : %d", age)
	}
}
