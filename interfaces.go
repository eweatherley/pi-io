package main

import (
	"fmt"
	"github.com/kidoman/embd"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var blinker Blinker

func main() {
	var blinker Blinker
	err := embd.InitGPIO()
	if err != nil {
		blinker = new (MockBlinker)
	} else {
		blinker = new (GPIOBlinker)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	text, _ := reader.ReadString('\n')

	intValue, err := strconv.Atoi(strings.TrimRight(text, "\n"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(intValue)
	fmt.Println(blinker.Blink(intValue))
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


