package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Hostname      string        `yaml:"hostname"`
	CheckInterval time.Duration `yaml:"check_interval"`
	Logger        Logger        `yaml:"logger"`
	Messenger     Messenger     `yaml:"messenger"`
	Storage       Storage       `yaml:"storage"`
	Alerts        []Alert       `yaml:"alerts"`
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

type Messenger struct {
	Type   string `yaml:"type"`
	Token  string `yaml:"token"`
	ChatID int64  `yaml:"chatid"`
}

type Storage struct {
	Type             string        `yaml:"type"`
	CooldownDuration time.Duration `yaml:"cooldown_duration"`
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
