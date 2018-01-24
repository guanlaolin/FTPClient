package main

import (
	"os"
	"os/signal"
)

const PORT_MIN int = 40000
const PORT_RANGE int = 5000

var exit chan int

func main() {
	signal.Ignore(os.Interrupt, os.Kill)

	UI()
	<-exit
}
