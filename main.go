package main

import (
	"github.com/meupacote/app"
	"github.com/meupacote/config"
)

func main() {
	config := config.GetConfig()
	app.Init(config)
	app.Run("10.202.2.62:3001")
}
