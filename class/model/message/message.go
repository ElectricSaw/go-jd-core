package message

func NewMessage() *Message {
	return &Message{}
}

func NewMessageWithBody(body interface{}) *Message {
	return &Message{
		Body: body,
	}
}

type Message struct {
	Headers map[string]interface{}
	Body    interface{}
}

func (m *Message) RemoveHeader(name string) {
	delete(m.Headers, name)
}
