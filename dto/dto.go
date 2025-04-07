package dto

type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Description string `json:"description"`
	Rating      int    `json:"rating" binding:"required,min=1,max=5"`
}

type UpdateBookRequest struct {
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}
