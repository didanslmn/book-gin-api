package main

import (
	"github.com/didanslmn/gin-restful-api/config"
	"github.com/didanslmn/gin-restful-api/router"
)

func main() {
	config.ConnectDB()
	r := router.SetupRouter()
	r.Run()
}

//CRUD
// create
// book := model.Book{}
// book.Title = "tiger"
// book.Price = 100
// book.Description = "dwwqq"
// book.Rating = 5

// err = db.Create(&book).Error
// if err != nil {
// 	fmt.Println("error creating book")
// }
// read
// var book model.Book

// read
// err = db.First(&book, 1).Error
// if err != nil {
// 	fmt.Println("error finding book")
// }
// fmt.Println("Title", book.Title)
// fmt.Println("book object %v", book)
//
// get by id
// var book model.Book

// err = db.Debug().Where("id=?", 1).First(&book).Error
// if err != nil {
// 	fmt.Println("error finding book")
// }
// update
// book.Title = "man tiger (edited)"
// err = db.Save(&book).Error
// if err != nil {
// 	fmt.Println("error update book")
// }
//delete
// err = db.Delete(&book).Error
// if err != nil {
// 	fmt.Println("error update book")
// }
