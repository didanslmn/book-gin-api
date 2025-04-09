package router

import (
	"github.com/didanslmn/gin-restful-api/handler"
	"github.com/didanslmn/gin-restful-api/repository"
	"github.com/didanslmn/gin-restful-api/service"
	"github.com/gin-gonic/gin"
)

// SetupRouter mengkonfigurasi dan mengembalikan instance dari Gin router.
// Fungsi ini bertanggung jawab untuk mendefinisikan semua endpoint API dan menghubungkannya ke handler yang sesuai.
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Membuat instance dari BookRepository.
	bookRepo := repository.NewBookRepository()
	// Membuat instance dari BookService, dengan menyuntikkan BookRepository sebagai dependency.
	bookService := service.NewBookService(bookRepo)
	// Membuat instance dari BookHandler, dengan menyuntikkan BookService sebagai dependency.
	bookHandler := handler.NewBookHandler(bookService)
	// Membuat grup route dengan prefix "/v1" untuk API versi 1.
	v1 := r.Group("/v1")
	v1.POST("/books", bookHandler.CreateBook)
	v1.GET("/books", bookHandler.GetAllBooks)
	v1.GET("/books/:id", bookHandler.GetBookByID)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	return r
}
