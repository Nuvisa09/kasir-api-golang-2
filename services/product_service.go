package services

import (
	"kasir-api-golang-2/models"
	"kasir-api-golang-2/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService{
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Categories, error){
	return s.repo.GetAll()
}

func (s *CategoryService) Create(data *models.Categories) error {
return  s.repo.Create(data)
}

func (s *CategoryService) GetByID(id int) (*models.Categories, error){
return  s.repo.GetByID(id)
}

func (s *CategoryService) Update(category *models.Categories) error{
return  s.repo.Update(category)
}
func (s *CategoryService) Delete(id int) error {
return  s.repo.Delete(id)
}