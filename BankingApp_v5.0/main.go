package main

import (
	"github.com/SundarBank/app"
	"github.com/SundarBank/logger"
)

func main(){
	logger.Info("Application started...")
	app.Start()
}
