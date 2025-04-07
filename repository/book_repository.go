package repository

import (
	"github.com/didanslmn/gin-restful-api/config"
	"github.com/didanslmn/gin-restful-api/model"
)

type BookRepository interface {
	Save(book model.Book) (model.Book, error)
	FindAll() []model.Book
	FindByID(id int) (model.Book, error)
	Update(book model.Book) (model.Book, error)
	Delete(id int) error
}

type bookRepository struct{}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) Save(book model.Book) (model.Book, error) {
	err := config.DB.Create(&book).Error
	return book, err
}

func (r *bookRepository) FindAll() []model.Book {
	var books []model.Book
	config.DB.Find(&books)
	return books
}

func (r *bookRepository) FindByID(id int) (model.Book, error) {
	var book model.Book
	err := config.DB.First(&book, id).Error
	return book, err
}

func (r *bookRepository) Update(book model.Book) (model.Book, error) {
	err := config.DB.Save(&book).Error
	return book, err
}

func (r *bookRepository) Delete(id int) error {
	return config.DB.Delete(&model.Book{}, id).Error
}
