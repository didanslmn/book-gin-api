package service

import (
	"github.com/didanslmn/gin-restful-api/dto"
	"github.com/didanslmn/gin-restful-api/model"
	"github.com/didanslmn/gin-restful-api/repository"
)

type BookService interface {
	Create(dto.CreateBookRequest) (*dto.BookResponse, error)
	FindAll() ([]*dto.BookResponse, error)
	FindByID(id int) (*dto.BookResponse, error)
	Update(id int, dto dto.UpdateBookRequest) (*dto.BookResponse, error)
	Delete(id int) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) Create(req dto.CreateBookRequest) (*dto.BookResponse, error) {
	book := model.Book{
		Title:       req.Title,
		Price:       req.Price,
		Description: req.Description,
		Rating:      req.Rating,
	}
	newBook, err := s.repo.Create(book)
	if err != nil {
		return nil, err
	}
	return dto.ToBookResponse(&newBook), nil
}

func (s *bookService) FindAll() ([]*dto.BookResponse, error) {
	books, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var responses []*dto.BookResponse
	for _, book := range books {
		responses = append(responses, dto.ToBookResponse(&book))
	}
	return responses, nil
}

func (s *bookService) FindByID(id int) (*dto.BookResponse, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return dto.ToBookResponse(&book), nil
}

func (s *bookService) Update(id int, req dto.UpdateBookRequest) (*dto.BookResponse, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	book.Title = req.Title
	book.Price = req.Price
	book.Description = req.Description
	book.Rating = req.Rating

	updatedBook, err := s.repo.Update(book)
	if err != nil {
		return nil, err
	}

	return dto.ToBookResponse(&updatedBook), nil
}

func (s *bookService) Delete(id int) error {
	return s.repo.Delete(id)
}
