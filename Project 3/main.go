package main

import (
	"echo-example/http_router"
)

const portEcho int = 8080

func main() {
	finished := make(chan bool)

	go http_router.EchoNew(portEcho)

	<-finished
}
