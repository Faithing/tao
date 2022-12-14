package chat

import (
	"context"

	"github.com/Faithing/tao"
	"github.com/leesper/holmes"
)

const (
	// ChatMessage is the message number of chat message.
	ChatMessage     uint64 = 1
	ChatMessageName        = "ChatMessage"
)

// Message defines the chat message.
type Message struct {
	Content string
}

// RequestCommand returns the request command.
func (cm Message) RequestCommand() uint64 {
	return ChatMessage
}

// ResponseCommand returns the response command.
func (cm Message) ResponseCommand() uint64 {
	return ChatMessage
}

// RequestName returns the request name.
func (cm Message) RequestName() string {
	return ChatMessageName
}

// ResponseName returns the response name.
func (cm Message) ResponseName() string {
	return ChatMessageName
}

// Serialize Serializes Message into bytes.
func (cm Message) Serialize() ([]byte, error) {
	return []byte(cm.Content), nil
}

func (cm Message) SetCustom(interface{}) {}

// DeserializeMessage deserializes bytes into Message.
func DeserializeMessage(data []byte) (message tao.Message, err error) {
	if data == nil {
		return nil, tao.ErrNilData
	}
	content := string(data)
	msg := Message{
		Content: content,
	}
	return msg, nil
}

// ProcessMessage handles the Message logic.
func ProcessMessage(ctx context.Context, conn tao.WriteCloser) {
	holmes.Infof("ProcessMessage")
	s, ok := tao.ServerFromContext(ctx)
	if ok {
		msg := tao.MessageFromContext(ctx)
		s.Broadcast(msg)
	}
}
