package service

import (
	"github.com/didanslmn/gin-restful-api/dto"
	"github.com/didanslmn/gin-restful-api/model"
	"github.com/didanslmn/gin-restful-api/repository"
)

// BookService adalah interface yang mendefinisikan kontrak untuk logika bisnis terkait buku.
// Interface ini menjembatani antara handler dan repository.
type BookService interface {
	Create(dto.CreateBookRequest) (*dto.BookResponse, error)
	FindAll() ([]*dto.BookResponse, error)
	FindByID(id int) (*dto.BookResponse, error)
	Update(id int, dto dto.UpdateBookRequest) (*dto.BookResponse, error)
	Delete(id int) error
}

// bookService adalah struct yang mengimplementasikan interface BookService.
// Struct ini memiliki dependency pada BookRepository untuk berinteraksi dengan database.
type bookService struct {
	repo repository.BookRepository
}

// NewBookService membuat dan mengembalikan instance baru dari bookService.
// Membutuhkan instance dari BookRepository sebagai parameter.
func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo}
}

// Create mengimplementasikan method Create dari interface BookService.
// Menerima CreateBookRequest, membuat objek model.Book, dan menyimpan ke database melalui repository.
// Mengembalikan BookResponse dari buku yang baru dibuat.
func (s *bookService) Create(req dto.CreateBookRequest) (*dto.BookResponse, error) {
	// Membuat objek model.Book dari data DTO.
	book := model.Book{
		Title:       req.Title,
		Price:       req.Price,
		Description: req.Description,
		Rating:      req.Rating,
	}
	// Memanggil method Create pada repository untuk menyimpan buku ke database.
	newBook, err := s.repo.Create(book)
	if err != nil {
		return nil, err
	}
	// Mengkonversi model.Book yang baru dibuat menjadi BookResponse.
	return dto.ToBookResponse(&newBook), nil
}

func (s *bookService) FindAll() ([]*dto.BookResponse, error) {
	// Memanggil method FindAll pada repository untuk mendapatkan semua buku.
	books, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var responses []*dto.BookResponse
	// Iterasi melalui setiap buku dan mengkonversinya menjadi BookResponse.
	for _, book := range books {
		responses = append(responses, dto.ToBookResponse(&book))
	}
	return responses, nil
}

func (s *bookService) FindByID(id int) (*dto.BookResponse, error) {
	// Memanggil method FindByID pada repository untuk mendapatkan buku berdasarkan ID.
	book, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	// Mengkonversi model.Book yang ditemukan menjadi BookResponse.
	return dto.ToBookResponse(&book), nil
}

// Update mengimplementasikan method Update dari interface BookService.
// Mengambil buku berdasarkan ID, memperbarui fieldnya dengan data dari UpdateBookRequest,
// dan menyimpan perubahan ke database melalui repository. Mengembalikan BookResponse dari buku yang diperbarui.
func (s *bookService) Update(id int, req dto.UpdateBookRequest) (*dto.BookResponse, error) {
	// Mencari buku yang akan diupdate berdasarkan ID.
	book, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	// Memperbarui field buku dengan data dari UpdateBookRequest (jika ada).
	book.Title = req.Title
	book.Price = req.Price
	book.Description = req.Description
	book.Rating = req.Rating
	// Memanggil method Update pada repository untuk menyimpan perubahan ke database.
	updatedBook, err := s.repo.Update(book)
	if err != nil {
		return nil, err
	}
	// Mengkonversi model.Book yang diperbarui menjadi BookResponse.
	return dto.ToBookResponse(&updatedBook), nil
}

func (s *bookService) Delete(id int) error {
	// Memanggil method Delete pada repository untuk menghapus buku berdasarkan ID.
	return s.repo.Delete(id)
}
