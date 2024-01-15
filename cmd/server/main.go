package main

import (
	"os"

	"github.com/elysiumyun/elysium/internal/app"
)

func main() {
	// bootstrap
	service := app.App
	service.Usage = app.Server.Usage()
	service.Flags = app.Server.Flags()
	service.Service = app.Server.Service()
	code, err := service.Run()
	defer os.Exit(code)
	if err != nil {
		panic(err)
	}
}
