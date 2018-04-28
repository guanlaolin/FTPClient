package main

import (
	"bytes"
)

const PORT_MIN int = 40000
const PORT_RANGE int = 5000

var exit = make(chan int)

var buffer bytes.Buffer

func main() {
	//go Cli()

	go UI()

	<-exit
}

func Debug() {

}
