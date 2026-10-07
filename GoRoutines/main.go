package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// creating context key for passing values
type contextKey string

var UserIdKey contextKey = "userID"

// custom errors
var (
	ErrorNotFound     = errors.New("truck not found")
	ErrNotImplemented = errors.New("truck not processed")
)

type TruckInterface interface {
	LoadCargo() error
	UnloadCargo() error
}

type Truck struct {
	id    string
	cargo int
}

type DieselTruck struct {
	Truck
}

func (t *DieselTruck) LoadCargo() error {
	t.cargo += 1

	return nil
}

func (t *DieselTruck) UnloadCargo() error {
	t.cargo = 0

	return nil
}

type ElectrictTruck struct {
	Truck
	battery float64
}

func (t *ElectrictTruck) LoadCargo() error {
	t.cargo += 1
	t.battery -= 1

	return nil
}

func (t *ElectrictTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

func ProccessTruck(ctx context.Context, truck TruckInterface) error {
	//getting context value example
	// userID := ctx.Value(UserIdKey)

	// Simulate running long operarion
	// delay one second more than context would fail: Error processing fleet: context deadline exceeded
	delay := time.Second * 2
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

	//truck functionality
	fmt.Printf("processing Truck: %+v\n", truck)
	// time.Sleep(time.Second)
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	fmt.Printf("Finished processing Truck: %+v\n", truck)
	return nil
}

// errChan must be buffered with capacity >= len(trucks),
// otherwise senders block and wg.Wait() never returns.
func ProccessFleet(ctx context.Context, trucks []TruckInterface) error {
	var wg sync.WaitGroup
	var errChan = make(chan error, len(trucks))

	for _, truck := range trucks {
		wg.Add(1)
		go func(t TruckInterface) {
			defer wg.Done()
			if err := ProccessTruck(ctx, t); err != nil {
				errChan <- err
			}
		}(truck)
	}
	wg.Wait()
	close(errChan)

	//select way for err
	/*
		select {
		case err := <-errChan:
			return err
		default:
			return nil
		}
	*/

	//using for range to loop errors
	var errs []error
	for err := range errChan {
		log.Panicf("error processing truck: %v\n", err)
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return fmt.Errorf("fleet processing had %d erros", len(errs))
	}
	return nil

}

func main() {
	//sending value in context
	// ctx = context.WithValue(ctx, UserIdKey, 42)

	//this is basically a signal that if were we are passing the context what ever is doing last more than 3 second cancel it , so the user doesnt wait
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fleet := []TruckInterface{
		&DieselTruck{Truck{id: "Truck 1", cargo: 0}},
		&ElectrictTruck{Truck: Truck{id: "ETruck 1", cargo: 0}, battery: 10},
		&DieselTruck{Truck{id: "Truck 2", cargo: 0}},
		&ElectrictTruck{Truck: Truck{id: "ETruck 2", cargo: 0}, battery: 10},
	}

	if err := ProccessFleet(ctx, fleet); err != nil {
		log.Fatalf("error processing truck: %s", err)

	}

	fmt.Println("all Trucks processed successfully")
}
