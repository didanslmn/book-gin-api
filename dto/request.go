package dto

// struct yang merepresentasikan data yang dibutuhkan untuk membuat buku baru.
type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Description string `json:"description"`
	Rating      int    `json:"rating" binding:"required,min=1,max=5"`
}

//struct yang merepresentasikan data yang dapat diubah untuk memperbarui informasi buku yang sudah ada.
type UpdateBookRequest struct {
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}
