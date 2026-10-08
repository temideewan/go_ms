package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a truck cargo", func(t *testing.T) {
			// arrange
			normalTruck := &NormalTruck{id: "Truck-1"}
			electricTruck := &ElectricTruck{id: "Electric-Truck-1"}
			// act
			err := processTruck(normalTruck)
			if err != nil {
				t.Fatalf("Error processing truck %s \n", err)
			}
			err = processTruck(electricTruck)
			if err != nil {
				t.Fatalf("Error processing truck %s \n", err)
			}

			// assert
			want := 0
			got := normalTruck.cargo
			if want != got {
				t.Fatalf("expected %d got %d", want, got)
			}

			want = -2
			got = int(electricTruck.battery)

			if want != got {
				t.Fatalf("expected %d got %d", want, got)
			}
		})
	})
}
