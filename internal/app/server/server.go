package server

import (
	"errors"
	"os"
	"strings"

	"github.com/elysiumyun/elysium/internal/pkg/system/prepare"
	"github.com/elysiumyun/elysium/pkg/logger"
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

	logger.Println("Server init")
	return nil
}
