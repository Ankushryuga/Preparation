package main

import (
	"fmt"
	"time"
)

type Direction int

const (
	Idle Direction = iota
	Up
	Down
)

type Request struct {
	Floor     int
	Direction Direction
}

// ElevatorCar interface
type ElevatorCar interface {
	ID() int
	CurrentFloor() int
	Direction() Direction
	AssignRequest(req Request)
	Step()
}

// Car struct implementing ElevatorCar
type Car struct {
	id           int
	currentFloor int
	direction    Direction
	requests     []Request
}

func NewCar(id int) *Car {
	return &Car{id: id, direction: Idle}
}

func (c *Car) ID() int             { return c.id }
func (c *Car) CurrentFloor() int   { return c.currentFloor }
func (c *Car) Direction() Direction { return c.direction }

func (c *Car) AssignRequest(req Request) {
	c.requests = append(c.requests, req)
	if c.direction == Idle {
		if req.Floor > c.currentFloor {
			c.direction = Up
		} else if req.Floor < c.currentFloor {
			c.direction = Down
		}
	}
}

func (c *Car) Step() {
	if len(c.requests) == 0 {
		c.direction = Idle
		return
	}

	// Move one floor
	if c.direction == Up {
		c.currentFloor++
	} else if c.direction == Down {
		c.currentFloor--
	}

	// Handle requests at current floor
	var remaining []Request
	for _, r := range c.requests {
		if r.Floor == c.currentFloor {
			fmt.Printf("🚪 Car %d stopped at floor %d\n", c.id, c.currentFloor)
		} else {
			remaining = append(remaining, r)
		}
	}
	c.requests = remaining

	// Decide next direction
	if len(c.requests) == 0 {
		c.direction = Idle
	} else {
		next := c.requests[0]
		if next.Floor > c.currentFloor {
			c.direction = Up
		} else {
			c.direction = Down
		}
	}
}

// Controller to manage elevators
type Controller struct {
	cars []ElevatorCar
}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) RegisterCar(e ElevatorCar) {
	ctrl.cars = append(ctrl.cars, e)
}

func (ctrl *Controller) SubmitRequest(req Request) {
	// Assign to the first idle or least loaded car
	best := ctrl.cars[0]
	for _, car := range ctrl.cars {
		if car.Direction() == Idle {
			best = car
			break
		}
	}
	best.AssignRequest(req)
}

func (ctrl *Controller) Step() {
	for _, car := range ctrl.cars {
		car.Step()
	}
}

// Simulation
func main() {
	ctrl := NewController()

	// Create 3 elevators
	for i := 0; i < 3; i++ {
		car := NewCar(i)
		ctrl.RegisterCar(car)
	}

	// Submit elevator requests
	ctrl.SubmitRequest(Request{Floor: 5, Direction: Up})
	ctrl.SubmitRequest(Request{Floor: 2, Direction: Down})
	ctrl.SubmitRequest(Request{Floor: 8, Direction: Down})
	ctrl.SubmitRequest(Request{Floor: 1, Direction: Up})

	// Simulate steps
	for step := 0; step < 15; step++ {
		fmt.Printf("\n=== ⏱️ Step %d ===\n", step)
		ctrl.Step()
		time.Sleep(500 * time.Millisecond)
	}
}
