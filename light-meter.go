package main

import (
	"time"
	"github.com/kidoman/embd"
	"fmt"
	"math"
)

var pin_a = 18
var pin_b = 23

func discharge(pinA embd.DigitalPin, pinB embd.DigitalPin) {
	fmt.Println("Discharging...")

	pinA.SetDirection(embd.In)
	pinB.SetDirection(embd.Out)
	pinB.Write(embd.Low)

	time.Sleep(time.Duration(100) * time.Millisecond)
	fmt.Println("Discharged...")
}

func lightFromR(R int64) float64 {
	return math.Log(float64(1000000)/float64(R)) * 10.0
}

func chargeTime(pinA embd.DigitalPin, pinB embd.DigitalPin) int64 {
	pinA.SetDirection(embd.Out)
	pinB.SetDirection(embd.In)
	risen := make(chan interface{})
	t1 := time.Now();
	pinB.Watch(embd.EdgeRising, func(pin embd.DigitalPin) {
		risen <- pin
	})
	_ := <- risen
	return int64(time.Since(t1)/time.Millisecond)

}

func analogRead(pinA embd.DigitalPin, pinB embd.DigitalPin) int64 {
	discharge(pinA, pinB)
	return int64(chargeTime(pinA, pinB))
}

func readResistence() float64 {
	total := int64(0)
	for i := 0; i < 20; i++ {
		total = int64(total) + analogRead()
	}
	reading := float64(total) / float64(20)
	resistence := reading * 6.05 - 939
	return resistence
}

func main() {
	_ := embd.InitGPIO()
	pinA, _ := embd.NewDigitalPin(pin_a)
	pinB, _ := embd.NewDigitalPin(pin_b)
	defer embd.CloseGPIO()
	value := lightFromR(analogRead(pinA, pinB));
	fmt.Print(value);
}