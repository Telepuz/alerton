package script

import (
	"os/exec"
	"strings"

	"github.com/telepuz/alerton/internal/config"
)

type Script struct {
	Name    string
	Command string
	Params  []string
}

func New(conf *config.Alert) (*Script, error) {
	return &Script{
			Name:    conf.Name,
			Command: conf.Command,
			Params:  conf.Params,
		},
		nil
}

func (s *Script) GetName() string {
	return s.Name
}

func (s *Script) Run() (bool, string, error) {
	isTriggered := false

	outByte, err := exec.Command(s.Command, s.Params...).Output()
	if err != nil {
		return isTriggered, "", err
	}
	out := string(outByte)
	if strings.ToUpper(out) != "OK\n" {
		isTriggered = true
	}

	return isTriggered, out, nil
}
