package repository

import (
	"github.com/didanslmn/gin-restful-api/config"
	"github.com/didanslmn/gin-restful-api/model"
)

// BookRepository adalah interface yang mendefinisikan kontrak untuk operasi terkait entitas Buku pada database.
type BookRepository interface {
	Create(book model.Book) (model.Book, error)
	FindAll() ([]model.Book, error)
	FindByID(id int) (model.Book, error)
	Update(book model.Book) (model.Book, error)
	Delete(id int) error
}

// bookRepository adalah struct yang mengimplementasikan interface BookRepository.
// Struct ini bertanggung jawab untuk berinteraksi langsung dengan database untuk operasi terkait buku.
type bookRepository struct{}

// NewBookRepository membuat dan mengembalikan instance baru dari bookRepository.
func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) Create(book model.Book) (model.Book, error) {
	err := config.DB.Create(&book).Error
	return book, err
}

func (r *bookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := config.DB.Find(&books).Error
	return books, err
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
