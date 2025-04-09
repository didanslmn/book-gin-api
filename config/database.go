package config

import (
	"fmt"
	"log"

	"github.com/didanslmn/gin-restful-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Membuka koneksi database menggunakan GORM dengan driver PostgreSQL.
	// &gorm.Config{} menggunakan konfigurasi default GORM.
	dsn := "host=localhost user=postgres password=godgame357 dbname=book_pustaka port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Fatal akan mencetak pesan error dan menghentikan eksekusi program.
		log.Fatal("db connection error", err)
	}
	// Menjalankan AutoMigrate: GORM akan memeriksa struct model.Book
	err = db.AutoMigrate(&model.Book{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
	// Pesan konfirmasi jika koneksi dan migrasi berhasil.
	fmt.Println("database conecteed and migrate succes")
	// Simpan instance koneksi GORM yang berhasil dibuat ke variabel global DB
	DB = db
}
