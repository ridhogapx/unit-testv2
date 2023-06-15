package repository

import "learn_unit_test_v2/entity"

type PostRepository interface {
	FindByAuthor(string) entity.Post
	FindById(string) entity.Post
}
