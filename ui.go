package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func UI() {
	fmt.Println("Welcome to use dogod-ftp-client!")

	if len(os.Args) == 2 {
		Open(os.Args[1])
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("ftp:> ")

		line, _, err := reader.ReadLine()

		if err != nil {
			log.Println("ReadLine:", err)
			continue
		}
		if string(line) == "" {
			continue
		}

		log.Printf("command:%s\n", line)

		CMDAnalyze(string(line))
	}
}

func CMDAnalyze(_cmd string) {
	segs := strings.Split(_cmd, " ")
	if len(segs) == 1 {
		segs = append(segs, "")
	}

	switch segs[0] {
	case "open":
		Open(segs[1])
	case "user":
		User(segs[1])
	case "ls":
		fallthrough
	case "dir":
		List(segs[1])
	case "lls":
		fallthrough
	case "ldir":
		Lls()
	case "lcd":
		Lcd(segs[1])
	case "get":
		Get(segs[1])
	case "lpwd":
		Lpwd()
	case "help":
		Help(segs[1])
	case "exit":
		fallthrough
	case "quit":
		fallthrough
	case "q":
		os.Exit(0)
	default:
		fmt.Printf("%s:unrecognized command\n", _cmd)
	}
}
