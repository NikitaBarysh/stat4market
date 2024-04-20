package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func NewDB(ctx context.Context, addr, port, db, user, pass string) (driver.Conn, error) {
	address := fmt.Sprintf("%s:%s", addr, port)
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{address},
		Auth: clickhouse.Auth{
			Database: db,
			Username: user,
			Password: pass,
		},
		Debugf: func(msg string, args ...interface{}) {
			log.Printf(msg, args)
		},
	})

	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Printf("[DB][EXCEPTION][%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		log.Printf("Error pinging database connection: %v", err)
		return nil, err
	}
	return conn, nil
}
