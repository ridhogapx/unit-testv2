package service

import (
	"learn_unit_test_v2/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var postRepository = &repository.PostRepositoryMock{Mock: mock.Mock{}}
var postService = PostService{Repository: postRepository}

func TestPost_Get(t *testing.T) {
	postRepository.Mock.On("FindById", "3").Return(nil)
	post, error := postService.Get("3")
	assert.Nil(t, post)
	assert.NotNil(t, error)

}
