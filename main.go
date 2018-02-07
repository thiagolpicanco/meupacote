package main

import (
	"github.com/thiagolpicanco/meupacote/app"
	"github.com/thiagolpicanco/meupacote/config"
)

func main() {
	config := config.GetConfig()
	app.Init(config)
	app.Run("10.202.2.62:3001")
}
