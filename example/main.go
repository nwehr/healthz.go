package main

import (
	"time"

	"github.com/nwehr/healthz.go"
)

func main() {
	ready := false

	healthz.Start(":8192", &ready)

	time.Sleep(10 * time.Second)

	ready = true

	time.Sleep(10 * time.Second)
}
