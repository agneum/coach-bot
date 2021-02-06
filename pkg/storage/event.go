package storage

import (
	"time"

	"gopkg.in/reform.v1"

	"github.com/agneum/scheduler-bot/pkg/models"
)

type EventRepo struct {
	db *reform.Querier
}

func NewEventRepo(db *reform.Querier) *EventRepo {
	return &EventRepo{db: db}
}

func (e *EventRepo) GetEvents() ([]*models.Event, error) {
	sts, err := e.db.SelectAllFrom(models.EventTable, "WHERE start_date > ?", time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	events := make([]*models.Event, 0, len(sts))

	for _, st := range sts {
		ev := st.(*models.Event)
		events = append(events, ev)
	}

	return events, nil
}

func (e EventRepo) AddEvent(ev *models.Event) error {
	return e.db.Insert(ev)
}
