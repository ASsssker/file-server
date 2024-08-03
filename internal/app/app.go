package app

import "file-server/internal/server"

type App struct {
	server *server.Server
}

func NewApp() *App {
	conf := GetConfig()
	srv := server.NewServer(conf.Host, conf.DirPath)
	return &App{
		server: srv,
	}
}

func (a *App) Run() {
	a.server.Run()
}
