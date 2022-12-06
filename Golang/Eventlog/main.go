package main

import "golang.org/x/sys/windows/svc/eventlog"

func main() {
	eventlog.Open("")
}
