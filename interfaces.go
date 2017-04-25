package main

import (
	"fmt"
	"github.com/kidoman/embd"
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

var blinker Blinker


func output(c chan int) {
	seconds := <-c
	fmt.Printf("Outputting Hi for %d s\n", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println("Finished outputting Hi")
}

func main() {
	var blinker Blinker
	c := make(chan int)
	go output(c)
	err := embd.InitGPIO()
	if err != nil {
		blinker = new (MockBlinker)
	} else {
		blinker = new (GPIOBlinker)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter number: ")
		text, _ := reader.ReadString('\n')

		intValue, err := strconv.Atoi(strings.TrimRight(text, "\n"))
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(intValue)
		c <- intValue
		fmt.Println(blinker.Blink(intValue))
	}
}


type Blinker interface {
	Blink(number int) string
}

type MockBlinker struct {

}

type GPIOBlinker struct {

}

func(mb MockBlinker) Blink(number int) string {
	return fmt.Sprintf("Mock Blink (%v)", number)
}

func(b GPIOBlinker) Blink(number int) string {
	return fmt.Sprintf("Sending signal to GPIO pin %d", number);
}


