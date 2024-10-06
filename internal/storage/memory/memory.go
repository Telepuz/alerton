package memory

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/telepuz/alerton/internal/config"
)

type Alert struct {
	Name      string
	Timestamp time.Time
}

type Memory struct {
	CooldownDuration time.Duration
	Alerts           []Alert
}

func NewMemoryStorage(cfg *config.Storage) *Memory {
	slog.Debug(fmt.Sprintf(
		"NewMemoryStorage(): Created new memory storage. Default cooldown - %v seconds",
		cfg.CooldownDuration.Seconds(),
	))
	return &Memory{
		CooldownDuration: cfg.CooldownDuration,
	}
}

func (m *Memory) IsCooldown(name string) bool {
	for _, alert := range m.Alerts {
		if alert.Name == name {
			slog.Debug(fmt.Sprintf(
				"IsCooldown(): Found alert - %s",
				name,
			))
			return false
		}
	}
	m.Alerts = append(
		m.Alerts,
		Alert{
			Name:      name,
			Timestamp: time.Now(),
		},
	)
	slog.Debug(fmt.Sprintf(
		"IsCooldown(): Appended alert to storage - %s",
		name,
	))
	return true
}

func (m *Memory) ClearByTTL() {
	temp := []Alert{}
	for _, alert := range m.Alerts {
		delta := time.Since(alert.Timestamp).Seconds()
		if delta <= m.CooldownDuration.Seconds() {
			temp = append(temp, alert)
		}
	}
	slog.Debug(fmt.Sprintf(
		"ClearByTTL(): Return alerts - %v",
		temp,
	))
	m.Alerts = temp
}
