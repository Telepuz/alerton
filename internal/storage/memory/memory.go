package memory

import (
	"time"
)

type Alert struct {
	Name      string
	Timestamp time.Time
}

type Memory struct {
	CooldownDuration time.Duration
	Alerts           []Alert
}

func NewMemoryStorage(cooldownDuration time.Duration) *Memory {
	return &Memory{
		CooldownDuration: cooldownDuration,
	}
}

func (m *Memory) IsCooldown(name string) bool {
	for _, alert := range m.Alerts {
		if alert.Name == name {
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
	m.Alerts = temp
}
