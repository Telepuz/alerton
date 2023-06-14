package app

import (
	"log"
	"time"

	"github.com/telepuz/alerton/internal/alert"
	"github.com/telepuz/alerton/internal/config"
	"github.com/telepuz/alerton/internal/messenger"
	"github.com/telepuz/alerton/internal/storage"
)

type AppContext struct {
	Config    *config.Config
	Messenger messenger.Messenger
	Alerts    []alert.Alert
	Storage   storage.Storage
}

func Run(c *AppContext) {
	for {
		c.Storage.ClearByTTL()
		for _, alert := range c.Alerts {
			alertName := alert.GetName()
			isTriggered, body, err := alert.Run()
			if err != nil {
				log.Printf("Alert.Run(): %s - %s", alertName, err)
			}
			if isTriggered && c.Storage.IsCooldown(alertName) {
				err = c.Messenger.SendMessage(
					alertName,
					c.Config.Hostname,
					body,
				)
				if err != nil {
					log.Printf("Messenger.SendMessage(): %s - %s", alertName, err)
				}
			}
		}
		time.Sleep(c.Config.CheckInterval)
	}
}
