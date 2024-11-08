package message

func NewMessage() *Message {
	return &Message{
		Headers: make(map[string]interface{}),
	}
}

func NewMessageWithBody(body interface{}) *Message {
	return &Message{
		Headers: make(map[string]interface{}),
		Body:    body,
	}
}

type Message struct {
	Headers map[string]interface{}
	Body    interface{}
}

func (m *Message) RemoveHeader(name string) {
	delete(m.Headers, name)
}
