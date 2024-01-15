package prepare

import (
	"os"

	"github.com/elysiumyun/elysium/internal/pkg/system/service"
	"github.com/elysiumyun/elysium/pkg/info"
	"github.com/elysiumyun/elysium/pkg/logger"
	"github.com/elysiumyun/elysium/pkg/timezone"
)

type AppMode string

const (
	Test    AppMode = "test"
	Debug   AppMode = "debug"
	Release AppMode = "release"
)

var ENV_MODE_KEY string = "GCA_APP_MODE"

func Environment() {
	appInfo()
	appMode()
	appTimeZone()
}

func appInfo() {
	logger.Printf("micro service: %s\n", info.MicroService)
}

func appMode() {
	var mode string
	mode = os.Getenv(ENV_MODE_KEY)
	if len(mode) == 0 {
		mode = string(Test)
	}

	if (mode != string(Test)) && (mode != string(Debug)) && (mode != string(Release)) {
		logger.Printf("Invalid Mode: %s\n", mode)
		mode = string(Test)
	}

	service.MODE = mode
	logger.Printf("System Running Mode: %s\n", service.MODE)
}

func appTimeZone() {
	tz := timezone.TZ()
	os.Setenv("TZ", tz)
	logger.Printf("System Running TimeZone: %s\n", tz)
}
