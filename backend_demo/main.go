// @title        Golang Service Template
// @version      0.1
// @description  Golang back-end service template, get started with back-end projects quickly
// @BasePath     /api

package main

import (
	"flybitch/app"
	"flybitch/model"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetReportCaller(true)
	model.Init()
	app.InitWebFramework()
	app.StartServer()
}
