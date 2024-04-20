package service

import (
	"context"
	"time"

	"github.com/NikitaBarysh/stat4market/internal/entity"
	"github.com/NikitaBarysh/stat4market/internal/repository"
)

type Service struct {
	DataService
}

type DataInterface interface {
	Set(ctx context.Context, data entity.Event) error
	Get(ctx context.Context, eventType string, timeBefore, timeAfter time.Time) ([]entity.Event, error)
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		DataService: NewDataService(rep),
	}
}
