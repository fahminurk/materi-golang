package repository

import "unit-test/entity"

type CategotyRepository interface {
	FindById(id string) *entity.Category
}