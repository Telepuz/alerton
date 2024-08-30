package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Telegram struct {
	Token  string
	ChatID int64
}

type Message struct {
	ChatID              int64  `json:"chat_id"`
	ParseMode           string `json:"parse_mode"`
	DisablePreview      bool   `json:"disable_web_page_preview"`
	DisableNotification bool   `json:"disable_notification"`
	Text                string `json:"text"`
}

func NewTelegram(token string, chatID int64) (*Telegram, error) {
	return &Telegram{
			Token:  token,
			ChatID: chatID,
		},
		nil
}

func (t *Telegram) NewMessage(title, hostname, body string) *Message {
	return &Message{
		ChatID:              t.ChatID,
		ParseMode:           "Markdown",
		DisablePreview:      true,
		DisableNotification: false,
		Text: fmt.Sprintf(
			"*%s*\n\n*Host: %s*\n%s\n",
			title,
			hostname,
			body,
		),
	}
}

func (t *Telegram) SendMessage(title, hostname, body string) error {
	message := t.NewMessage(title, hostname, body)
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}
	response, err := http.Post(
		fmt.Sprintf(
			"https://api.telegram.org/bot%s/sendMessage",
			t.Token,
		),
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return err
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Println("failed to close response body")
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send successful request. Status was %q", response.Status)
	}
	log.Printf("SendMessage(): %s\n", title)
	return nil
}
