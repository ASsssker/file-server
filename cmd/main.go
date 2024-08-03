package main

import "file-server/internal/app"

func main() {
	app := app.NewApp()
	app.Run()
}