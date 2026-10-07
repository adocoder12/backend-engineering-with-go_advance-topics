package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("ProccessTruck", func(t *testing.T) {
		t.Run("Should load and unload a truck cargo", func(t *testing.T) {
			truck := &Truck{ID: "Truck 1", Cargo: 42}
			Etruck := &ElectrictTruck{ID: "ETruck 1", Cargo: 0, Battery: 10}

			if err := processTruck(truck); err != nil {
				t.Fatalf("error processing truck: %s", err)
			}
			if err := processTruck(Etruck); err != nil {
				t.Fatalf("error processing truck: %s", err)
			}

			//assert
			if truck.Cargo != 0 {
				t.Fatalf("Normal Truck cargo should be 0, got : %d", truck.Cargo)
			}
			if Etruck.Battery != 9 {
				t.Fatalf("Electrict batteru should be 9 , got: %g", Etruck.Battery)
			}

		})
	})

}

/* to test go test -v *.go  */
