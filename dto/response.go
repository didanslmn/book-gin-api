package dto

import (
	"time"

	"github.com/didanslmn/gin-restful-api/model"
)

// Struct ini berisi informasi detail buku yang akan dikirimkan sebagai response API.
type BookResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	Rating      int       `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time ` json:"updated_at"`
}

// ToBookResponse adalah fungsi yang mengkonversi objek model.Book menjadi objek BookResponse.
// Fungsi ini berguna untuk memformat data buku dari layer model menjadi format response API.
func ToBookResponse(b *model.Book) *BookResponse {
	return &BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}
