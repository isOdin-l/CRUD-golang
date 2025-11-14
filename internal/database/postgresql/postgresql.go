package postgresql

import (
	"context"
	"fmt"

	"github.com/isOdin/RestApi/configs"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func NewPostgresDB(cfg *configs.Config) (*pgxpool.Pool, error) {
	conectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	conn, err := pgxpool.New(context.Background(), conectionString)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	logrus.Info("Database connected")
	return conn, nil
}
