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
	CooldownDuration time.Duration `yaml:"cooldown_duration"`
	TelegramToken    string        `yaml:"telegram_token"`
	TelegramChatid   string        `yaml:"telegram_chatid"`
	Alerts           []Alert       `yaml:"alerts"`
}

type Alert struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Command string   `yaml:"command"`
	Params  []string `yaml:"params"`
}

func NewConfig() (*Config, error) {
	host, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("hostname error: %w", err)
	}
	cfg := &Config{
		Hostname: host,
	}
	err = cleanenv.ReadConfig("./config/alerton.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
