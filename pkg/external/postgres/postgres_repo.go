package postgres

import (
	"fmt"

	"github.com/andreashanson/golang-rabbitmq/pkg/msg"
)

type PostgresRepo struct {
	Conn PostgresConnection
}

func NewPostgresRepo(c PostgresConnection) *PostgresRepo {
	return &PostgresRepo{
		Conn: c,
	}
}

func (pr PostgresRepo) Get(m msg.Message) {
	fmt.Println("Get data for type: ", m.Body.Type)
}
func (pr PostgresRepo) Post(m msg.Message) {
	fmt.Println("Post data for type: ", m.Body.Type)
}
func (pr PostgresRepo) Update(m msg.Message) {
	fmt.Println("Update data for type: ", m.Body.Type)
}

func (pr PostgresRepo) GetAll(m msg.Message) {
	fmt.Println("Get all for type: ", m.Body.Type)
}
