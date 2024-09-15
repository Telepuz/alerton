package app

import (
	"fmt"
	"log/slog"
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
				slog.Error(fmt.Sprintf(
					"Run() Alert Run: %s - %s",
					alertName,
					err,
				))
			}
			if isTriggered && c.Storage.IsCooldown(alertName) {
				err = c.Messenger.SendMessage(
					alertName,
					c.Config.Hostname,
					body,
				)
				if err != nil {
					slog.Error(fmt.Sprintf(
						"Run(): Messenger.SendMessage(): %s - %s",
						alertName,
						err,
					))
				}
			}
		}

		slog.Debug(fmt.Sprintf(
			"Run(): Sleeping for %s",
			c.Config.CheckInterval,
		))
		time.Sleep(c.Config.CheckInterval)
	}
}
