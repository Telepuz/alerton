package alert

import (
	"fmt"
	"log/slog"

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
		slog.Debug(fmt.Sprintf("NewAlerts(): Got new alert: %s", alertConfig.Name))
		switch alertConfig.Type {
		default:
			alert, err := script.New(&alertConfig, scriptDir)
			if err != nil {
				return nil, err
			}
			alerts = append(alerts, alert)
		}
	}
	slog.Debug(fmt.Sprintf("NewAlerts(): Return alerts: %s", alerts))
	return alerts, nil
}
