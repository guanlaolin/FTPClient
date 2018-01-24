package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const FMT string = "%-15s%s\n"

var conn net.Conn

var ok chan int

func ExecCMD(cmd string) {
	checkConn()

	buf := make([]byte, 512)

	log.Println(cmd)

	_, err := conn.Write([]byte(cmd))
	if err != nil {
		log.Println("Write:", err)
		return
	}

	_, err = conn.Read(buf)
	if err != nil {
		log.Println("Read:", err)
		return
	}

	fmt.Printf("%s", buf)

	CodeAnalyze(getResponCode(buf))
}

func Open(addr string) {
	if addr == "" {
		fmt.Println("Usage:open host[:port]")
		return
	}

	if conn != nil {
		fmt.Println("You are conncted to ", addr)
		return
	}

	buf := make([]byte, 512)

	segs := strings.Split(addr, ":")
	//validate ip
	//

	if len(segs) > 2 {
		fmt.Println("Usage:open host[:port]")
		return
	}
	//if port == ""
	if len(segs) == 1 {
		addr = addr + ":21"
	}

	du, err := time.ParseDuration("10s")
	if err != nil {
		log.Println("ParseDuration:", err)
		return
	}

	conn, err = net.DialTimeout("tcp", addr, du)
	if err != nil {
		log.Println("Dial:", err)
		return
	}

	_, err = conn.Read(buf)
	if err != nil {
		log.Println("Read:", err)
		return
	}

	fmt.Printf("%s", buf)

	CodeAnalyze(getResponCode(buf))
}

func User(user string) {
	if user == "" {
		fmt.Println("Usage:user username")
		return
	}

	cmd := "USER " + user + "\r\n"

	ExecCMD(cmd)
}

func Pass(pass string) {
	cmd := "PASS " + pass + "\r\n"
	ExecCMD(cmd)
}

func List(dir string) {
	Port()

	if dir == "" {
		dir = "/"
	}

	cmd := "LIST " + dir + "\r\n"
	ExecCMD(cmd)
}

func Port() {
	addr := conn.LocalAddr().String()
	segs := strings.Split(addr, ":")
	log.Println("addr:", segs[0])
	n5, n6 := PORT(segs[0])

	cmd := strings.Replace(addr, ",", ".", -1)
	cmd = "PORT " + cmd + "." + strconv.Itoa(n5) + "." + strconv.Itoa(n6) + "\r\n"
	log.Println(cmd)

	ExecCMD(cmd)
}

var listener net.Listener

func PORT(addr string) (n5 int, n6 int) {
	var err error

	rand.Seed(time.Now().Unix())
	port := PORT_MIN + rand.Intn(PORT_RANGE)
	log.Println("port:", port)

	for i := 0; i < PORT_RANGE; i++ {
		listener, err = net.Listen("tcp", addr+":"+strconv.Itoa(port))
		if err != nil {
			log.Println("listen:", port, " ERROR:", err)
			port++
			time.Sleep(1 * time.Millisecond)
			continue
		}
		break
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept:", err)
			return
		}

		defer listener.Close()
		defer conn.Close()

		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Println("ReadAll:", err)
			return
		}

		fmt.Printf("\n%s\n", buf)
	}()

	return port / 256, port % 256

}

func Help() {
	fmt.Printf(FMT, "open", "connect to remote ftp server.")
	fmt.Printf(FMT, "user", "")
	fmt.Printf(FMT, "help", "show help.")
	fmt.Printf(FMT, "quit", "exit program.")
}

func getResponCode(buf []byte) int {
	temp := string(buf[0:3])
	code, _ := strconv.Atoi(temp)
	return code
}

func checkConn() {
	if conn == nil {
		fmt.Println("Not connected")
		return
	}
}
