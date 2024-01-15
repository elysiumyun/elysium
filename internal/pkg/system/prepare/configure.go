package prepare

import (
	"os"

	"github.com/elysiumyun/elysium/internal/pkg/config"
	"github.com/elysiumyun/elysium/pkg/logger"
	"github.com/elysiumyun/elysium/pkg/utils"
)

func Configure() {
	options := config.Cfg.Get()
	logger.Printf("system log path: %s\n", options.Logdir)
	systemPath(options.Logdir)
	logger.Printf("system data path: %s\n", options.Datadir)
	systemPath(options.Datadir)
}

func systemPath(dir string) {
	// 持久化系统日志以及数据
	if !utils.IsExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
