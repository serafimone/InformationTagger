package main

import (
	"github.com/serafimone/InformationTagger/app"
	"github.com/serafimone/InformationTagger/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
