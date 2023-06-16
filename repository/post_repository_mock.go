package repository

import (
	"learn_unit_test_v2/entity"

	"github.com/stretchr/testify/mock"
)

type PostRepositoryMock struct {
	Mock mock.Mock
}

// Todo: add more test
func (repository *PostRepositoryMock) FindById(id string) *entity.Post {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		post := args.Get(0).(entity.Post)
		return &post
	}
}

func (repository *PostRepositoryMock) Filter(query string) *[]entity.Post {
	args := repository.Mock.Called(query)
	if args.Get(0) == nil {
		return nil
	} else {
		// This is fixed
		posts := args.Get(0).([]entity.Post)
		return &posts

	}
}
