package messenger

type Messenger interface {
	SendMessage(title, hostname, body string) error
}
