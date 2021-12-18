package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("RABBIT_HOST", "rabbithost")
	os.Setenv("RABBIT_USER", "rabbituser")
	os.Setenv("RABBIT_PW", "testpw")
	os.Setenv("POSTGRES_HOST", "postgreshost")
	os.Setenv("POSTGRES_USER", "postgresuser")
	os.Setenv("POSTGRES_PW", "postgrespw")
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "test new config",
			want: &Config{
				Postgres: &PostgresConfig{
					Host:     "postgreshost",
					User:     "postgresuser",
					Password: "postgrespw",
				},
				RabbitMQ: &RabbitMQConfig{
					Host:     "rabbithost",
					User:     "rabbituser",
					Password: "testpw",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewConfig()
			assert.Equal(t, *got.Postgres, *tt.want.Postgres)
			assert.Equal(t, *got.RabbitMQ, *tt.want.RabbitMQ)
			//if got := NewConfig(); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			//}
		})
	}
}
