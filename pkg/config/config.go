package config

type Config struct {
	RabbitMQ *RabbitMQConfig
}

type RabbitMQConfig struct {
	Host     string
	User     string
	Password string
}
