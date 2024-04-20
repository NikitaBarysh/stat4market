package repository

import (
	"context"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/NikitaBarysh/stat4market/internal/entity"
)

type Repository struct {
	DataRepository
}

func NewRepository(db driver.Conn) *Repository {
	return &Repository{
		DataRepository: NewDataRepository(db),
	}
}

type DataInterface interface {
	SetData(ctx context.Context, data entity.Event) error
	Get(ctx context.Context, eventType string, timeBefore, timeAfter time.Time) ([]entity.Event, error)
}
