package main

import (
	"fmt"
	"github.com/kidoman/embd"
)

var blinker Blinker

func main() {
	var blinker Blinker
	err := embd.InitGPIO()
	if (err != nil) {
		blinker = new (MockBlinker)
	} else {
		blinker = new (GPIOBlinker)
	}
	fmt.Println(blinker.Blink())
}


type Blinker interface {
	Blink() string
}

type MockBlinker struct {

}

type GPIOBlinker struct {
	pinNumber int
}

func(mb MockBlinker) Blink() string {
	return "Mock Blink"
}

func(b GPIOBlinker) Blink() string {
	return fmt.Sprintf("Sending signal to GPIO pin %d", b.pinNumber);
}


