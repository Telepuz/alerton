package alert

import (
	"github.com/telepuz/alerton/internal/alert/script"
	"github.com/telepuz/alerton/internal/config"
)

type Alert interface {
	GetName() string
	Run() (bool, string, error)
}

func NewAlerts(configs *[]config.Alert, scriptDir string) ([]Alert, error) {
	alerts := []Alert{}
	for _, alertConfig := range *configs {
		switch alertConfig.Type {
		default:
			alert, err := script.New(&alertConfig, scriptDir)
			if err != nil {
				return nil, err
			}
			alerts = append(alerts, alert)
		}
	}
	return alerts, nil
}
