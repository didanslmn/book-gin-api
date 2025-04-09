package handler

import (
	"net/http"
	"strconv"

	"github.com/didanslmn/gin-restful-api/dto"
	"github.com/didanslmn/gin-restful-api/service"
	"github.com/gin-gonic/gin"
)

// BookHandler adalah struct yang menangani request terkait buku.
// Struct ini memiliki dependency pada BookService untuk logika bisnis.
type BookHandler struct {
	Service service.BookService
}

// NewBookHandler membuat instance baru dari BookHandler.
// Membutuhkan instance dari BookService sebagai parameter.
func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{Service: service}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var bookDTO dto.CreateBookRequest
	// BindJSON digunakan untuk mengikat data JSON dari body request ke struct CreateBookRequest.
	if err := c.ShouldBindJSON(&bookDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Panggil method Create pada service untuk membuat buku baru.
	book, err := h.Service.Create(bookDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Memanggil method FindAll pada service untuk mengambil daftar semua buku.
func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.Service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get book"})
	}
	c.JSON(http.StatusOK, books)
}

// Mengambil ID dari parameter route dan memanggil method FindByID pada service
func (h *BookHandler) GetBookByID(c *gin.Context) {
	// Ambil nilai parameter "id" dari URL dan konversikan ke integer.
	id, _ := strconv.Atoi(c.Param("id"))
	// Panggil method FindByID pada service untuk mencari buku berdasarkan ID.
	book, err := h.Service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Mengambil ID dari parameter route dan data pembaruan dari body request.
func (h *BookHandler) UpdateBook(c *gin.Context) {
	// Ambil nilai parameter "id" dari URL dan konversikan ke integer.
	id, _ := strconv.Atoi(c.Param("id"))
	var bookDTO dto.UpdateBookRequest
	// BindJSON digunakan untuk mengikat data JSON dari body request ke struct UpdateBookRequest.
	if err := c.ShouldBindJSON(&bookDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Panggil method Update pada service untuk memperbarui buku berdasarkan ID dan data yang diberikan.
	updatedBook, err := h.Service.Update(id, bookDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}

// Mengambil ID dari parameter route dan memanggil method Delete pada service.
func (h *BookHandler) DeleteBook(c *gin.Context) {
	// Ambil nilai parameter "id" dari URL dan konversikan ke integer.
	id, _ := strconv.Atoi(c.Param("id"))
	// Panggil method Delete pada service untuk menghapus buku berdasarkan ID.
	err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
