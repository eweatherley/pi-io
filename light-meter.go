package main

import (
	"time"
	"fmt"
	"math"
)

func discharge() {
	fmt.Println("Discharging...")
	time.Sleep(time.Duration(100) * time.Millisecond)
	fmt.Println("Discharged...")
}

func lightFromR(R int64) float64 {
	return math.Log(float64(1000000)/float64(R)) * 10.0
}

func chargeTime() int64 {
	// GPIO setup
	t1 := time.Now();
	time.Sleep(time.Duration(100) * time.Millisecond)
	return int64(time.Since(t1)/time.Millisecond);

}

func analogRead() int64 {
	discharge()
	return int64(chargeTime())
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
	value := lightFromR(analogRead());
	fmt.Print(value);
}