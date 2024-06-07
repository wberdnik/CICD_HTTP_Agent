package main

import (
	"github.com/wberdnik/CICD_HTTP_Agent/internal/app/apiserver"
)

func main() {

	if err := apiserver.Start(); err != nil {
		panic("Halt of start HTTP server")
	}
}
