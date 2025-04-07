package service

import (
	"github.com/didanslmn/gin-restful-api/dto"
	"github.com/didanslmn/gin-restful-api/model"
	"github.com/didanslmn/gin-restful-api/repository"
)

type BookService interface {
	Create(dto.CreateBookRequest) (model.Book, error)
	FindAll() []model.Book
	FindByID(id int) (model.Book, error)
	Update(id int, dto dto.UpdateBookRequest) (model.Book, error)
	Delete(id int) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) Create(req dto.CreateBookRequest) (model.Book, error) {
	book := model.Book{
		Title:       req.Title,
		Price:       req.Price,
		Description: req.Description,
		Rating:      req.Rating,
	}
	return s.repo.Save(book)
}

func (s *bookService) FindAll() []model.Book {
	return s.repo.FindAll()
}

func (s *bookService) FindByID(id int) (model.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) Update(id int, req dto.UpdateBookRequest) (model.Book, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return book, err
	}

	book.Title = req.Title
	book.Price = req.Price
	book.Description = req.Description
	book.Rating = req.Rating

	return s.repo.Update(book)
}

func (s *bookService) Delete(id int) error {
	return s.repo.Delete(id)
}
