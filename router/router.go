package router

import (
	"github.com/didanslmn/gin-restful-api/handler"
	"github.com/didanslmn/gin-restful-api/repository"
	"github.com/didanslmn/gin-restful-api/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	bookRepo := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	v1 := r.Group("/v1")
	v1.POST("/books", bookHandler.CreateBook)
	v1.GET("/books", bookHandler.GetAllBooks)
	v1.GET("/books/:id", bookHandler.GetBookByID)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	return r
}
