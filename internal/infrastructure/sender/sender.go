package sender

// Sender Отправитель сообщений
type Sender interface {
	SendMessageToChats(string) error
}
