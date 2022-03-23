package main

import "fmt"

func main() {
	mallardDuck := MallardDuck{}
	TestDuck(mallardDuck)

	drone := SuperDrone{}
	droneAdapter := DroneAdapter{drone}
	TestDuck(droneAdapter)
}

type Duck interface {
	Quack()
	Fly()
}

func TestDuck(duck Duck) {
	duck.Quack()
	duck.Fly()
}

// MallardDuck

type MallardDuck struct{}

func (d MallardDuck) Quack() {
	fmt.Println("mallard duck quacking...")
}

func (d MallardDuck) Fly() {
	fmt.Println("mallard duck flying...")
}

// Drone

type Drone interface {
	Beep()
	SpinRotors()
	TakeOff()
}

// SuperDrone

type SuperDrone struct{}

func (d SuperDrone) Beep() {
	fmt.Println("beep beep beep")
}

func (d SuperDrone) SpinRotors() {
	fmt.Println("rotors are spinning")
}

func (d SuperDrone) TakeOff() {
	fmt.Println("taking off")
}

// DroneAdapter

type DroneAdapter struct {
	d Drone
}

func (a DroneAdapter) Quack() {
	a.d.Beep()
}

func (a DroneAdapter) Fly() {
	a.d.SpinRotors()
	a.d.TakeOff()
}
