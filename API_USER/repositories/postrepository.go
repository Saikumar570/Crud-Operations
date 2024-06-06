package repositories

import (
	"api/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func FunctionOfRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Create(post *models.Post) error {
	return r.DB.Create(post).Error
}

func (r *Repository) FindAll(posts *[]models.Post) error {
	return r.DB.Find(posts).Error
}

func (r *Repository) FindByID(id string, post *models.Post) error {
	return r.DB.First(post, id).Error
}

func (r *Repository) Update(post *models.Post) error {
	return r.DB.Save(post).Error
}

func (r *Repository) Delete(id string) error {
	return r.DB.Delete(&models.Post{}, id).Error
}
