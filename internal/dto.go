package dto

import (
	"github.com/telepuz/alerton/internal/alert"
	"github.com/telepuz/alerton/internal/config"
	"github.com/telepuz/alerton/internal/messenger"
)

type AppContext struct {
	Config    *config.Config
	Messenger messenger.Messenger
	Alerts    []alert.Alert
}
