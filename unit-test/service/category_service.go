package service

import (
	"errors"
	"unit-test/entity"
	"unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategotyRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error){
category := service.Repository.FindById(id)
if category == nil {
	return nil, errors.New("Category not found")
} else {
	return category, nil
}
}