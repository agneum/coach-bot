package scheduler

import (
	"github.com/agneum/scheduler-bot/pkg/storage"
)

type Scheduler struct {
	TplSvc storage.TemplateRepo
}

func NewScheduler(tplSvc storage.TemplateRepo) *Scheduler {
	return &Scheduler{
		TplSvc: tplSvc,
	}
}

func (s *Scheduler) Schedule() error {
	templates, err := s.TplSvc.GetTemplates()
	if err != nil {
		return err
	}

	for range templates {

	}

	return nil
}
