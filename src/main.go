package main

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	book_datasource "github.com/peam1146/go_clean_architecture/src/book/data/datasource/remote_datasource"
	"github.com/peam1146/go_clean_architecture/src/book/data/repository"
	book_usecase "github.com/peam1146/go_clean_architecture/src/book/domain/usecase"
)


func main() {
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
    panic(err)
	}

  bookQueires := book_datasource.New(db)

  bookRepoImpl := book_repo_impl.NewBookRepositoryImpl(bookQueires)
  getBookUseCase := book_usecase.NewGetBookUseCase(bookRepoImpl)
  getBooksUseCase := book_usecase.NewGetBooksUseCase(bookRepoImpl)

  r := fiber.New()

  g := r.Group("/book")
}
