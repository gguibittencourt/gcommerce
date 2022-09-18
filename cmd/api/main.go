package main

import (
	"github.com/gguibittencourt/gcommerce/cmd/api/modules"
)

func main() {
	app := modules.NewApp()
	app.Run()
}
