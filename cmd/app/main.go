package main

import (
	"flag"
	"log"

	"github.com/telepuz/alerton/internal/alert"
	"github.com/telepuz/alerton/internal/app"
	config "github.com/telepuz/alerton/internal/config"
	"github.com/telepuz/alerton/internal/messenger/telegram"
	"github.com/telepuz/alerton/internal/storage/memory"
)

var (
	scriptsDir *string
	configFile *string
)

func init() {
	scriptsDir = flag.String("scripts_dir", "/etc/alerton/scripts", "Scripts directory")
	configFile = flag.String("config_file", "/etc/alerton/alerton.yml", "Config filename")
}

func main() {
	flag.Parse()

	cfg, err := config.NewConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	alerts, err := alert.NewAlerts(&cfg.Alerts, *scriptsDir)
	if err != nil {
		log.Fatal(err)
	}

	msg, err := telegram.NewTelegram(cfg.TelegramToken, cfg.TelegramChatid)
	if err != nil {
		log.Fatal(err)
	}

	storage := memory.NewMemoryStorage(cfg.CooldownDuration)

	appContext := app.AppContext{
		Config:    cfg,
		Messenger: msg,
		Alerts:    alerts,
		Storage:   storage,
	}
	app.Run(&appContext)
}
