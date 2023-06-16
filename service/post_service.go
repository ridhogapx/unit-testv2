package service

import (
	"errors"
	"learn_unit_test_v2/entity"
	"learn_unit_test_v2/repository"
)

type PostService struct {
	Repository repository.PostRepository
}

// Todo: Add more service

func (service PostService) Get(id string) (*entity.Post, error) {
	post := service.Repository.FindById(id)
	if post == nil {
		return post, errors.New("Post is not found!")
	} else {
		return post, nil
	}
}

func (service PostService) Query(query string) (*[]entity.Post, error) {
	posts := service.Repository.Filter(query)
	if posts == nil {
		return posts, errors.New("Post is not found")
	} else {
		return posts, nil
	}
}
