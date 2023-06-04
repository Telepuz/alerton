package telegram

import "fmt"

type Telegram struct {
	Name string
}

func NewTelegram(name string) (*Telegram, error) {
	return &Telegram{
			Name: name,
		},
		nil
}

func (t *Telegram) SendMessage(title, hostname, body string) error {
	fmt.Printf("Title: %s\nHostname: %s\nBody: %s\n", title, hostname, body)
	return nil
}
