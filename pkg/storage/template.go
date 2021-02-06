package storage

import (
	"gopkg.in/reform.v1"

	"github.com/agneum/scheduler-bot/pkg/models"
)

type TemplateRepo struct {
	db *reform.Querier
}

func NewTemplateRepo(db *reform.Querier) *TemplateRepo {
	return &TemplateRepo{db: db}
}

func (t *TemplateRepo) GetTemplates() ([]*models.Template, error) {
	return []*models.Template{}, nil
}
