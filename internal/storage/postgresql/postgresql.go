package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	UsersTable      string = "users"
	todoListsTable  string = "todo_lists"
	todoItemsTable  string = "todo_items"
	usersListsTable string = "users_lists"
	listsItemsTable string = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewPostgresDB(cfg *Config) (*pgxpool.Pool, error) {
	conectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	conn, err := pgxpool.New(context.Background(), conectionString)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conn, nil
}
