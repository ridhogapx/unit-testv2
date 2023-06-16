package repository

import "learn_unit_test_v2/entity"

type PostRepository interface {
	FindById(string) *entity.Post
	Filter(string) *[]entity.Post
}
