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
	sts, err := t.db.SelectAllFrom(models.TemplateTable, "")
	if err != nil {
		return nil, err
	}

	templates := make([]*models.Template, 0, len(sts))

	for _, st := range sts {
		template := st.(*models.Template)
		templates = append(templates, template)
	}

	return templates, nil
}
