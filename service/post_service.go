package service

import (
	"errors"
	"learn_unit_test_v2/entity"
	"learn_unit_test_v2/repository"
)

type PostService struct {
	Repository repository.PostRepository
}

func (service PostService) Get(id string) (*entity.Post, error) {
	post := service.Repository.FindById(id)
	if post == nil {
		return post, errors.New("Post is not found!")
	} else {
		return post, nil
	}
}
