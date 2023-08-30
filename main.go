package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type (
	Elevator struct {
		Capacity      int
		InstantPeople int
		InstantFloor  int
		Direction     string // "up", "down", "none"
		CabineId      int
	}

	ControlPanel struct {
		Elevators  []Elevator
		TotalFloor int
	}
)

var ControlPanelInstance *ControlPanel
var controlPanelOnce sync.Once
var mu sync.Mutex // Mutex for protecting ControlPanelInstance

func GetControlPanel() *ControlPanel {
	controlPanelOnce.Do(func() {
		ControlPanelInstance = &ControlPanel{
			Elevators: []Elevator{
				Elevator{
					Capacity:      12,
					InstantPeople: 3,
					InstantFloor:  1,
					Direction:     "up",
					CabineId:      1,
				},
				Elevator{
					Capacity:      12,
					InstantPeople: 0,
					InstantFloor:  2,
					Direction:     "none",
					CabineId:      2,
				},
				Elevator{
					Capacity:      12,
					InstantPeople: 9,
					InstantFloor:  3,
					Direction:     "down",
					CabineId:      3,
				},
			},
			TotalFloor: 11,
		}
	})
	return ControlPanelInstance
}

func CallElevator(currentFloor, personCount int, direction string) *Elevator {
	mu.Lock()
	defer mu.Unlock()

	var closestElevator *Elevator
	shortestDistance := 1000 // A large initial value

	for i, elevator := range ControlPanelInstance.Elevators {
		if (elevator.Direction == "none" || elevator.Direction == direction) && elevator.InstantPeople+personCount <= elevator.Capacity {
			distance := abs(currentFloor - elevator.InstantFloor)
			if distance < shortestDistance {
				closestElevator = &ControlPanelInstance.Elevators[i]
				shortestDistance = distance
			}
		}
	}

	return closestElevator
}

func ElevatorMoveManagement(elevator *Elevator) {
	mu.Lock()
	defer mu.Unlock()

	if elevator.Direction == "none" {
		fmt.Printf("Elevator %d is not moving.\n", elevator.CabineId)
		return
	}

	fmt.Printf("Elevator %d is moving %s from floor %d\n", elevator.CabineId, elevator.Direction, elevator.InstantFloor)

	var startFloor, endFloor int
	if elevator.Direction == "up" {
		startFloor = elevator.InstantFloor + 1
		endFloor = ControlPanelInstance.TotalFloor
	} else {
		startFloor = elevator.InstantFloor - 1
		endFloor = 0
	}

	for floor := startFloor; floor != endFloor; {
		elevator.InstantFloor = floor
		fmt.Printf("Elevator %d is on floor %d with %d people\n", elevator.CabineId, elevator.InstantFloor, elevator.InstantPeople)
		time.Sleep(time.Second) // Simulate time passing

		if elevator.Direction == "up" {
			floor++
		} else {
			floor--
		}
	}

	fmt.Printf("Elevator %d has reached its destination.\n", elevator.CabineId)
	elevator.Direction = "none"
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SimulateElevatorTransaction() {
	for {
		currentFloor := rand.Intn(ControlPanelInstance.TotalFloor)
		personCount := rand.Intn(10) + 1
		directions := []string{"up", "down"}
		direction := directions[rand.Intn(len(directions))]

		closestElevator := CallElevator(currentFloor, personCount, direction)

		if closestElevator != nil {
			fmt.Printf("%d people are waiting on floor %d, calling Elevator %d\n", personCount, currentFloor, closestElevator.CabineId)
			closestElevator.InstantPeople += personCount
			ElevatorMoveManagement(closestElevator)

			time.Sleep(time.Second * 2)
			closestElevator.InstantPeople -= personCount
			fmt.Printf("%d people exited elevator %d\n", personCount, closestElevator.CabineId)
		}

		time.Sleep(time.Second * 2)
	}
}

func main() {
	// Initialize elevator instances
	GetControlPanel()

	// Run the ElevatorMoveManagement for each elevator to simulate their movement
	for i := range ControlPanelInstance.Elevators {
		go ElevatorMoveManagement(&ControlPanelInstance.Elevators[i])
	}

	go SimulateElevatorTransaction()

	// Keep the main thread running
	select {}
}
