package topics

type Example struct {
	Value string `json:"value"`
}

func (c *Example) GetTopicName() string {
	return "example"
}

func (c *Example) GetPayloadMessage() (t TopicPayload) {
	t.Payload = c
	return
}
