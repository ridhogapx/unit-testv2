package service

import (
	"learn_unit_test_v2/entity"
	"learn_unit_test_v2/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var postRepository = &repository.PostRepositoryMock{Mock: mock.Mock{}}
var postService = PostService{Repository: postRepository}

func TestPost_Get_1(t *testing.T) {
	postRepository.Mock.On("FindById", "3").Return(nil)
	post, error := postService.Get("3")
	assert.Nil(t, post)
	assert.NotNil(t, error)

}

func TestPost_Get_2(t *testing.T) {
	single_post := entity.Post{
		Id:     "5",
		Title:  "Kerusuhan 1998",
		Author: "Anon",
	}

	postRepository.Mock.On("FindById", "5").Return(single_post)
	post, error := postService.Get("5")

	assert.Nil(t, error)
	assert.NotNil(t, post)

	assert.Equal(t, single_post.Id, post.Id, "Result is not expected")
	assert.Equal(t, single_post.Title, post.Title)
}

func TestFilterPost(t *testing.T) {
	posts := []entity.Post{
		{
			Id:     "7",
			Title:  "Tips menjaga tubuh sehat",
			Author: "Dr. Johny",
		},
		{
			Id:     "10",
			Title:  "Tips melatih AIM",
			Author: "Fearless",
		},
	}

	postRepository.Mock.On("Filter", "Tips").Return(posts)

	result, error := postService.Query("Tips")

	assert.Nil(t, error)
	assert.NotNil(t, result)
	// Fix this!!!
	assert.Equal(t, posts, result)
}
