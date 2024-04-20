package service

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/NikitaBarysh/stat4market/internal/entity"
	"github.com/NikitaBarysh/stat4market/internal/model"
	"github.com/NikitaBarysh/stat4market/internal/repository"
	"github.com/bxcodec/faker/v3"
)

type DataService struct {
	rep *repository.Repository
}

func NewDataService(rep *repository.Repository) DataService {
	return DataService{rep}
}

func (s *DataService) Set(ctx context.Context, data model.Event) error {
	bodyTime, err := time.Parse("2006-01-02 15:04:05", data.EventTime)
	if err != nil {
		log.Printf("Error parsing time: %v", err)
		return err
	}

	event := entity.Event{
		EventID:   rand.Int63n(100),
		EventType: data.EventType,
		UserID:    data.UserID,
		EventTime: bodyTime,
		Payload:   data.Payload,
	}
	err = s.rep.SetData(ctx, event)
	if err != nil {
		log.Printf("Error setting data: %v", err)
		return err
	}
	return nil
}

func (s *DataService) Get(ctx context.Context, eventType string, timeBefore, timeAfter time.Time) ([]entity.Event, error) {
	res, err := s.rep.Get(ctx, eventType, timeBefore, timeAfter)
	if err != nil {
		log.Printf("Error getting data: %v", err)
		return nil, err
	}
	return res, nil
}

func (s *DataService) Scheduler(ctx context.Context) {
	ticker := time.NewTicker(time.Minute * 5)
	go func(ticker *time.Ticker) {
		for {
			select {
			case <-ticker.C:
				s.fakeData(ctx)
			}
		}
	}(ticker)
}

func (s *DataService) fakeData(ctx context.Context) {
	var randomEvent entity.Event

	err := faker.FakeData(&randomEvent)

	if err != nil {
		log.Printf("Error faking data: %v", err)
	}

	randomEvent.EventID = rand.Int63n(110)
	randomEvent.EventType = faker.UUIDDigit()
	randomEvent.UserID = rand.Int63n(10)
	randomEvent.EventTime = time.Unix(rand.Int63n(2342232323), 0)
	randomEvent.Payload = faker.UUIDDigit()

	err = s.rep.SetData(ctx, randomEvent)
	if err != nil {
		log.Printf("Error setting fake data: %v", err)
	}
}
