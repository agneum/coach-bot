package scheduler

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/agneum/scheduler-bot/pkg/models"
	"github.com/agneum/scheduler-bot/pkg/storage"
)

type Scheduler struct {
	TplSvc   *storage.TemplateRepo
	EventSvc *storage.EventRepo
}

func NewScheduler(tplSvc *storage.TemplateRepo, eventSvc *storage.EventRepo) *Scheduler {
	return &Scheduler{
		TplSvc:   tplSvc,
		EventSvc: eventSvc,
	}
}

func (s *Scheduler) Schedule() error {
	templates, err := s.TplSvc.GetTemplates()
	if err != nil {
		return err
	}

	now := time.Now()
	for _, tpl := range templates {
		dayDiff := tpl.Weekday - int(now.Weekday())
		if dayDiff < 0 {
			dayDiff += 7
		}

		stTime, err := time.Parse("15:04", tpl.StartTime)
		if err != nil {
			return err
		}

		startDate := time.Date(now.Year(), now.Month(), now.Day(), stTime.Hour(), stTime.Minute(), 0, 0, now.Location())
		startDate = startDate.Add(time.Duration(dayDiff*24) * time.Hour)

		event := &models.Event{
			Title:       tpl.Title,
			Description: tpl.Description,
			Type:        tpl.Type,
			Note:        tpl.Note,
			CoachID:     tpl.CoachID,
			PlaceID:     tpl.PlaceID,
			StartDate:   startDate,
			Duration:    tpl.Duration,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := s.EventSvc.AddEvent(event); err != nil {
			return err
		}
	}

	return nil
}

func (s *Scheduler) ShowTemplates() string {
	str := strings.Builder{}

	tpls, err := s.TplSvc.GetTemplates()
	if err != nil {
		log.Fatal(err)
	}

	for _, ev := range tpls {
		str.WriteString(fmt.Sprintln(ev))
	}

	return str.String()
}

const EventTpl = `
%s (%s) %s (%s)
Тип: %s%s
Зал: %d
Тренер: %d
---`

func (s Scheduler) ShowEvents() string {
	str := strings.Builder{}

	events, err := s.EventSvc.GetEvents()
	if err != nil {
		log.Fatal(err)
	}

	for _, ev := range events {
		note := ""
		if ev.Note != "" {
			note = fmt.Sprintf(" (%s)", ev.Note)
		}

		str.WriteString(
			fmt.Sprintf(EventTpl, ev.StartDate.Weekday().String(), ev.StartDate.Format("02-01-2006"),
				ev.StartDate.Format("15:04"), ev.Duration, ev.Type, note, ev.PlaceID, ev.CoachID),
		)
	}

	return str.String()
}
