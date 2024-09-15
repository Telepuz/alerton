package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Hostname         string        `yaml:"hostname"`
	CheckInterval    time.Duration `yaml:"check_interval"`
	Logger           Logger        `yaml:"logger"`
	CooldownDuration time.Duration `yaml:"cooldown_duration"`
	AlertTimeout     time.Duration `yaml:"alert_timeout"`
	TelegramToken    string        `yaml:"telegram_token"`
	TelegramChatid   int64         `yaml:"telegram_chatid"`
	Alerts           []Alert       `yaml:"alerts"`
}

type Alert struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Command string   `yaml:"command"`
	Params  []string `yaml:"params"`
}

type Logger struct {
	Format string `yaml:"format"`
	Level  string `yaml:"level"`
}

func NewConfig(configFile string) (*Config, error) {
	host, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("NewConfig(): hostname error: %w", err)
	}
	cfg := &Config{
		Hostname: host,
	}
	err = cleanenv.ReadConfig(configFile, cfg)
	if err != nil {
		return nil, fmt.Errorf("NewConfig(): config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
