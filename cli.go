package main

/*
#include <stdio.h>

#ifdef WIN32
	#include <conio.h>
#endif

#ifdef linux
#include <string.h>
#include <termios.h>
int getch(void)
{
	int fd = 0;
	int ch = 0;
	struct termios tm, otm;

	memset(&tm, 0, sizeof(tm));
	memset(&otm, 0, sizeof(otm));

	if (tcgetattr(0, &tm) < 0) {
		return -1;
	}

	otm = tm;

	cfmakeraw(&tm);

	if (tcsetattr(0, TCSANOW, &tm) < 0) {
		return -1;
	}

	ch = getchar();

	if (tcgetattr(0, &tm) < 0) {
		return -1;
	}

	return ch;
}
#endif
*/
import "C"

import (
	"fmt"
	"log"
)

const KEY_CTRL_C int = 3
const KEY_UP int = 72
const KEY_DOWN int = 80
const KEY_LEFT int = 75
const KEY_RIGHT int = 77

func Cli() {
	var ch C.int

	for {
		fmt.Println("before")
		ch = C.getch()
		fmt.Println("after")

		switch int(ch) {
		case KEY_UP:
			fmt.Println("up")
		case KEY_DOWN:
			fmt.Println("down")
		case KEY_LEFT:
			fmt.Println("left")
		case KEY_RIGHT:
			fmt.Println("right")
		default:
			fmt.Printf("%d", ch)
			err := buffer.WriteByte(byte(ch))
			if err != nil {
				log.Println("WriteByte:", err)
				continue
			}
			//fmt.Fprint(os.Stdin, ch)
		}
	}
}
