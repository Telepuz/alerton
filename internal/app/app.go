package app

import (
	"log"
	"time"

	dto "github.com/telepuz/alerton/internal"
)

func Run(c *dto.AppContext) {
	for {
		for _, alert := range c.Alerts {
			isTriggered, body, err := alert.Run()
			if err != nil {
				log.Printf("Alert.Run(): %s - %s", alert.GetName(), err)
			}
			if isTriggered {
				err = c.Messenger.SendMessage(
					alert.GetName(),
					c.Config.Hostname,
					body,
				)
				if err != nil {
					log.Printf("Messenger.SendMessage(): %s - %s", alert.GetName(), err)
				}
			}
		}
		time.Sleep(c.Config.CheckInterval)
	}
}
