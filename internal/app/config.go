package app

import "flag"

const (
	defaultHTTPHost = "127.0.0.1:8888"
	defaultDirPath  = "."
)

type AppConfig struct {
	Host    string
	DirPath string
}

func GetConfig() *AppConfig {
	conf := AppConfig{}

	flag.StringVar(&conf.Host, "addr", defaultHTTPHost, "Хост на котором запустится сервер")
	flag.StringVar(&conf.DirPath, "path", defaultDirPath, "Путь до нужной директории")

	flag.Parse()

	return &conf
}
