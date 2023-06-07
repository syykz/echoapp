package database

import (
	"echoapp/domain/model"
	"echoapp/domain/repository"
)

func NewSampleRepository() repository.SampleRepository {
	return &sampleRepositoryImpl{}
}

type sampleRepositoryImpl struct {}

func (s *sampleRepositoryImpl) FindById(id string) (*model.Sample, error) {
	return &model.Sample{
		ID: id,
		Name: "sample",
	}, nil
}

func (s *sampleRepositoryImpl) FindAll() ([]*model.Sample, error) {
	return []*model.Sample{
		&model.Sample{
			ID: "1",
			Name: "sample1",
		},
		&model.Sample{
			ID: "2",
			Name: "sample2",
		},
	}, nil
}