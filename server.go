package main

import (
	"github.com/bagastri07/be-test-kumparan/services/application"
	"github.com/bagastri07/be-test-kumparan/services/config"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	conf := config.GetConfig()

	figure.NewColorFigure(conf.AppName, "", "green", true).Print()

	app := application.New(&conf)

	app.Start()
}
