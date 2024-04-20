package repository

import (
	"context"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/NikitaBarysh/stat4market/internal/entity"
)

type DataRepository struct {
	db driver.Conn
}

func NewDataRepository(NewDB driver.Conn) DataRepository {
	return DataRepository{db: NewDB}
}

func (d *DataRepository) SetData(ctx context.Context, data entity.Event) error {
	sql := `INSERT INTO events (eventID, eventType, userID, eventTime, payload) 
			VALUES ($1, $2, $3, $4, $5)`

	err := d.db.Exec(ctx, sql, data.EventID, data.EventType, data.UserID, data.EventTime, data.Payload)
	if err != nil {
		log.Printf("Error inserting event: %v", err)
		return err
	}

	return nil
}

func (d *DataRepository) Get(ctx context.Context, eventType string, timeBefore, timeAfter time.Time) ([]entity.Event, error) {
	sql := `SELECT * FROM events WHERE eventType=$1 AND eventTime>$2 AND eventTime<$3`

	rows, err := d.db.Query(ctx, sql, eventType, timeBefore, timeAfter)

	if err != nil {
		log.Printf("Error getting events: %v", err)
		return nil, err
	}

	defer rows.Close()

	eventSlice := make([]entity.Event, 0)

	for rows.Next() {
		var event entity.Event
		err = rows.Scan(&event.EventID, &event.EventType, &event.UserID, &event.EventTime, &event.Payload)
		if err != nil {
			log.Printf("Error getting event: %v", err)
			return nil, err
		}
		eventSlice = append(eventSlice, event)
	}

	return eventSlice, nil
}
