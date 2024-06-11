package repositories

import (
	"api/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB      *gorm.DB
	Post    *PostRepository
	Comment *CommentRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB:      db,
		Post:    NewPostRepository(db),
		Comment: NewCommentRepository(db),
	}
}

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostRepository) FindAll(posts *[]models.Post) error {
	return r.DB.Preload("Comments").Find(posts).Error
}

func (r *PostRepository) FindByID(id string, post *models.Post) error {
	return r.DB.First(post, id).Error
}

func (r *PostRepository) Update(post *models.Post) error {
	return r.DB.Save(post).Error
}

func (r *PostRepository) Delete(id string) error {
	return r.DB.Delete(&models.Post{}, id).Error
}

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func (r *CommentRepository) Create(comment *models.Comment) error {
	return r.DB.Create(comment).Error
}

func (r *CommentRepository) FindAllByPostID(postID string, comments *[]models.Comment) error {
	return r.DB.Where("post_id = ?", postID).Find(comments).Error
}

func (r *CommentRepository) FindByID(id string, comment *models.Comment) error {
	return r.DB.First(comment, id).Error
}

func (r *CommentRepository) Update(comment *models.Comment) error {
	return r.DB.Save(comment).Error
}

func (r *CommentRepository) Delete(id string) error {
	return r.DB.Delete(&models.Comment{}, id).Error
}
