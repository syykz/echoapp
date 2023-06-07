package usecase

import (
	"echoapp/domain/model"
	"echoapp/domain/repository"
)

type SampleUsecase interface {
	FindById(id string) (*model.Sample, error)
	FindAll() ([]*model.Sample, error)
}

type sampleUsecaseImpl struct {
	sr repository.SampleRepository
}

func NewSampleUsecase(sr repository.SampleRepository) SampleUsecase {
	return &sampleUsecaseImpl{
		sr: sr,
	}
}

type SampleInput struct {
	ID string `json:"id"`
}

func (s *sampleUsecaseImpl) FindById(id string) (*model.Sample, error) {
	sample, err := s.sr.FindById(id)
	return sample, err
}

func (s *sampleUsecaseImpl) FindAll() ([]*model.Sample, error) {
	sample, err := s.sr.FindAll()
	return sample, err
}

