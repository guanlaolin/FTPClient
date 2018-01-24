package main

import (
	"fmt"
	"log"
)

const ERROR int = -1
const RESP_CODE_OPEN_DATA = 150
const RESP_CODE_SERVICE_READY int = 220
const RESP_CODE_CLOSE_DATA int = 226
const RESP_CODE_LOGGED_IN int = 230
const RESP_CODE_NEED_PASS int = 331

func CodeAnalyze(code int) {
	log.Println("code:", code)
	switch code {
	case RESP_CODE_OPEN_DATA:
		buf := make([]byte, 64)
		_, err := conn.Read(buf)
		if err != nil {
			log.Println("ReadALL:", err)
			return
		}
		fmt.Printf("%s\n", buf)

	case RESP_CODE_SERVICE_READY:
		var user string
		fmt.Print("Require username:")
		_, err := fmt.Scanln(&user)
		if err != nil {
			log.Println("Scanln:", err)
			return
		}
		User(user)
	case RESP_CODE_CLOSE_DATA:
		log.Println("close data connection")
	case RESP_CODE_LOGGED_IN:

	case RESP_CODE_NEED_PASS:
		var pass string
		fmt.Print("Require password:")
		_, err := fmt.Scanln(&pass)
		if err != nil {
			log.Println("Scanln:", err)
			return
		}
		Pass(pass)
	default:
		log.Println("Invalid code:", code)
	}

}
