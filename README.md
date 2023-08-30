# Elevator Simulation in Go

This project demonstrates an elevator management simulation developed using the Go programming language. It aims to showcase the behavior of elevators and the interaction of people with these elevators through a simple simulation scenario.

## Project Structure

### Data Models

The code utilizes two primary data structures:

- **Elevator**: Represents the characteristics of elevators. It includes attributes such as capacity, current number of people, floor number, direction of movement, and cabin identifier.

- **ControlPanel**: Represents the control panel that manages all elevators. It includes attributes like a list of elevators and the total number of floors.

### Core Functions

- **GetControlPanel()**: Initializes and creates the control panel instance. Within this function, elevator instances and total floor count are defined.

- **CallElevator(currentFloor, personCount int, direction string) Elevator**: Simulates the process of calling an elevator. It aims to select the most suitable and nearest elevator.

- **ElevatorMoveManagement(elevator *Elevator)**: Simulates the movement of an elevator. It updates status information as the elevator moves between floors based on its direction.

- **SimulateElevatorTransaction()**: Simulates the process of people interacting with elevators, including waiting for an elevator, entering, moving, and exiting. Random floor numbers and passenger counts are generated.

### Main Function

- **main()**: The main function initiates the program by calling the necessary functions to simulate elevator and passenger behaviors.

## How to Use?

1. **Clone the Project**: Clone the source code and place it in a directory.

2. **Set Up Go Environment**: If Go is not installed, download and install it from the [Official Go Website](https://golang.org/).

3. **Run the Project**: Open a terminal, navigate to the project directory, and execute the following commands to run the project:

   ```bash
   go run main.go
