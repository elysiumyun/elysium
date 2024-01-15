package config

// options
type Options struct {
	App    `yaml:",inline"` // App 运行设置
	System `yaml:",inline"` // App 系统级配置
}

// app env config
type App struct {
	Addr string `yaml:"addr"` // App 地址
	Port string `yaml:"port"` // App 端口
}

type System struct {
	Logdir   string `yaml:"logdir"`   // 日志路径
	Datadir  string `yaml:"datadir"`  // 数据路径
	LogLevel string `yaml:"loglevel"` // 日志等级
	TimeZone string `yaml:"timezone"` // 时区设置
}
