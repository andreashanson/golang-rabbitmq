package msg

type Message struct {
	DeliveryTag uint64
	Body        Body
	Exchange    string
}

type Body struct {
	Type      string `json:"type"`
	StartTime string `json:"start_time"`
}
