package msg

type Message struct {
	DeliveryTag uint64
	Body        Body
	Exchange    string
}

type Body struct {
	Type      string `json:"type"`
	Msg       string `json:"msg"`
	StartTime string `json:"start_time"`
}
