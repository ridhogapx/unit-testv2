package repository

import (
	"learn_unit_test_v2/entity"

	"github.com/stretchr/testify/mock"
)

type PostRepositoryMock struct {
	Mock mock.Mock
}

// Fix this
func (repository *PostRepositoryMock) FindById(id string) entity.Post {
	args := repository.Mock.Called(author)
	if args.Get(0) == nil {
		return nil
	} else {
		return
	}
}
