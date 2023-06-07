package repository

import (
	"echoapp/domain/model"
)

type SampleRepository interface {
	FindById(id string) (*model.Sample, error)
	FindAll() ([]*model.Sample, error)
}