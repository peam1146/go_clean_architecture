package book_repo

import (
	"github.com/peam1146/go_clean_architecture/src/book/domain/entity"
)

type UpdateBookParam struct {
	ID     int64
	Name   string
	Author string
}

type BookRepository interface {
	GetBooks() (*[]*entity.BookEntity, error)
	GetBook(id int64) (*entity.BookEntity, error)
	CreateBook(name string, author string) (*entity.BookEntity, error)
	UpdateBook(param UpdateBookParam) (*entity.BookEntity, error)
	DeleteBook(id int64) error
}
