# API Book CRUD dengan Framework Gin

Aplikasi sederhana untuk manajemen data buku dengan operasi CRUD (Create, Read, Update, Delete) menggunakan:
- Bahasa: Go
- Database: Postgres sql
- Framework: Gin
- Arsitektur: Layered (Handler-Service-Repository-Model)

## Fitur
- Membuat data buku baru
- Melihat daftar semua buku
- Mengupdate data buku
- Menghapus data buku
- Validasi form input

## Prasyarat
- Go 1.18+
- MySQL 5.7+
- Go Modules
- Gin library/framework
- GORM

## Instalasi

1. Clone repository:
```bash
git clone https://github.com/didanslmn/book-gin-api
cd crud-go
```
2. Gin framework:
```bash
go get -u github.com/gin-gonic/gin
```
3. GORM library:
```bash
go get -u gorm.io/gorm
```

## Endpoint

HTTP Method,	Endpoint	,Handler Function,	Deskripsi

- POST	  | /v1/books	    | CreateBook	| Membuat data buku baru
- GET	    | /v1/books	    | GetAllBooks	| Mendapatkan semua data buku
- GET	    | /v1/books/:id	| GetBookByID	| Mendapatkan buku berdasarkan ID
- PUT	    | /v1/books/:id	| UpdateBook	| Memperbarui data buku berdasarkan ID
- DELETE	| /v1/books/:id	| DeleteBook	| Menghapus data buku berdasarkan ID


Base URL :
```bash
POST   http://localhost:8080/v1/books
GET    http://localhost:8080/v1/books
GET    http://localhost:8080/v1/books/123
PUT    http://localhost:8080/v1/books/123
DELETE http://localhost:8080/v1/books/123
```
