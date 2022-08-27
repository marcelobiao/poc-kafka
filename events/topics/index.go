package topics

type ITopic interface {
	GetTopicName() string
	GetPayloadMessage() TopicPayload
}

type TopicPayload struct {
	Payload ITopic `json:"payload"`
}

type TopicHeader struct {
	DbName    string `json:"dbName"`
	AccountID string `json:"accountId"`
}
