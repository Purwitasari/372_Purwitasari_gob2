package repository

import "hello-mock/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
