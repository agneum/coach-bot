package storage

import (
	"gopkg.in/reform.v1"

	"github.com/agneum/scheduler-bot/pkg/models"
)

type UserRepo struct {
	db *reform.Querier
}

func NewUserRepo(db *reform.Querier) *UserRepo {
	return &UserRepo{db: db}
}

func (e UserRepo) AddUser(user *models.User) error {
	return e.db.Insert(user)
}
