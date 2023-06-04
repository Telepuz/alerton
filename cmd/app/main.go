package main

import (
	"log"

	dto "github.com/telepuz/alerton/internal"
	"github.com/telepuz/alerton/internal/alert"
	"github.com/telepuz/alerton/internal/app"
	config "github.com/telepuz/alerton/internal/config"
	"github.com/telepuz/alerton/internal/messenger/telegram"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	alerts, err := alert.NewAlerts(&cfg.Alerts)
	if err != nil {
		log.Fatal(err)
	}

	msg, err := telegram.NewTelegram("telegramName")
	if err != nil {
		log.Fatal(err)
	}

	appContext := dto.AppContext{
		Config:    cfg,
		Messenger: msg,
		Alerts:    alerts,
	}
	app.Run(&appContext)
}
