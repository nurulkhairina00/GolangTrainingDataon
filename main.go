package main

import "fmt"

// Tire interfaze
type Tire interface {
	Type() string
}

// rubber tire struct
type RubberTire struct{}

// rubber tire method
func (r RubberTire) Type() string {
	return "Rubber Tire"
}

// wooden tire struct
type WoodenTire struct{}

// wooden tire method
func (w WoodenTire) Type() string {
	return "Wooden Tire"
}

// metal tire struct
type MetalTire struct{}

// metal tire method
func (m MetalTire) Type() string {
	return "Metal Tire"
}

// door struct
type Door struct {
	side string
}

// knock method  print a sound  based on the side of the door
func (d Door) Knock() {
	if d.side == "right" {
		fmt.Println("Knock")
	} else if d.side == "left" {
		fmt.Println("Beep")
	}
}

// Open method  print a sound  based on the side of the door
func (d Door) Open() {
	if d.side == "right" {
		fmt.Println("Beep")
	} else if d.side == "left" {
		fmt.Println("Knock")
	}
}

// car struct
type Car struct {
	tires     [4]Tire
	leftDoor  Door
	rightDoor Door
}

// NewCar creates a new Car instance with default doors
func NewCar() *Car {
	return &Car{
		leftDoor:  Door{side: "left"},
		rightDoor: Door{side: "right"},
	}
}

// change tire based on index
func (c *Car) ChangeTire(index int, tire Tire) {
	if index >= 0 && index < 4 {
		c.tires[index] = tire
	} else {
		fmt.Println("Invalid tire index")
	}
}

func main() {
	car := NewCar() // create new car

	// change tires on the car
	car.ChangeTire(0, RubberTire{})
	car.ChangeTire(1, WoodenTire{})
	car.ChangeTire(2, MetalTire{})
	car.ChangeTire(3, RubberTire{})

	// print the types of tires on the car
	for i, tire := range car.tires {
		if tire != nil {
			fmt.Printf("Tire %d: %s\n", i+1, tire.Type())
		}
	}
	fmt.Println("====================")

	// print right door
	fmt.Println("Right Door")
	car.rightDoor.Knock()
	car.rightDoor.Open()

	fmt.Println("====================")

	// print left door
	fmt.Println("Left Door")
	car.leftDoor.Knock()
	car.leftDoor.Open()

}
