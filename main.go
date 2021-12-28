package main

import (
	"movies/app"
	"movies/config"
)

func init() {
	config.LoadENV()
}

func main() {
	app.NewApp().Run(config.ENV.AppPort)
}
