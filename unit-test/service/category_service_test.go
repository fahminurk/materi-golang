package service

import (
	"testing"
	"unit-test/entity"
	"unit-test/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

// func TestCategoryService_Get(t *testing.T){
// 	categoryRepository.Mock.On("FindById", "1").Return(nil)

// 	category, err := categoryService.Get("1")
// 	assert.Nil(t, category)
// 	assert.NotNil(t, err)
// }

func TestCategoryService_GetSuccess(t *testing.T){
	category := entity.Category{
		Id: "1",
		Name: "Laptop",
	}

	categoryRepository.Mock.On("FindById", "3").Return(category)

	res, err := categoryService.Get("3")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, category.Id, res.Id)
}