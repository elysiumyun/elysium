/*
Copyright © 2023 Mardan https://www.mardan.wiki
*/
package app

import (
	"github.com/elysiumyun/elysium/internal/app/server"
	"github.com/elysiumyun/elysium/pkg/info"
)

func init() {
	info.ShowBanner()
}

type service string

const (
	Server service = "server"
)

func (service service) Usage() string {
	switch service {
	case Server:
		return server.Usage()
	}
	return ""
}

func (service service) Flags() func() (bool, error) {
	switch service {
	case Server:
		return server.Flags
	}
	return func() (bool, error) {
		return true, nil
	}
}

func (service service) Service() func() error {
	switch service {
	case Server:
		return server.Execute
	}
	return nil
}
