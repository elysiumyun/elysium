package server

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/elysiumyun/elysium/internal/app/server/router"
	"github.com/elysiumyun/elysium/internal/pkg/handlers"
	"github.com/elysiumyun/elysium/internal/pkg/system/prepare"
	"github.com/elysiumyun/elysium/pkg/logger"
	"github.com/elysiumyun/elysium/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Usage() string {
	var buffer strings.Builder

	return buffer.String()
}

func Flags() (bool, error) {
	var err error
	var argv = os.Args[1:]
	var isBreak bool = false

	switch argv[0] {
	default:
		err = errors.New("please check usage")
	}
	return isBreak, err
}

func Execute() error {
	// service prepare
	prepare.Environment()

	// system initialize
	prepare.Configure()

	// configure web server
	engine := gin.New()
	handlers.SetHandlers(engine)
	router.SetRouter(engine, "/elysium")

	var addr string = utils.Resolver()

	server := &http.Server{
		Addr:           addr,
		Handler:        engine,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 1 << 20,
	}

	// start web server
	logger.Printf("Server Listening %s Success...\n", addr)
	err := server.ListenAndServe()
	return err
}
