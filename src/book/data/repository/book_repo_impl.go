package book_repo_impl

import (
	"context"

	book_datasource "github.com/peam1146/go_clean_architecture/src/book/data/datasource/remote_datasource"
	book_entity "github.com/peam1146/go_clean_architecture/src/book/domain/entity"
	book_repo "github.com/peam1146/go_clean_architecture/src/book/domain/repository"
)

func NewBookRepositoryImpl(bookDB *book_datasource.Queries) *bookRepositoryImpl {
	return &bookRepositoryImpl{bookDB}
}

type bookRepositoryImpl struct {
	bookDB *book_datasource.Queries
}

func (br *bookRepositoryImpl) GetBooks() (*[]*book_entity.BookEntity, error) {
	books, err := br.bookDB.ListBook(context.Background())
	if err != nil {
		return nil, err
	}
	var booksEntity []*book_entity.BookEntity
	for _, v := range books {
		booksEntity = append(booksEntity, book_entity.NewBookEntity(v.ID, v.Name, v.Author))
	}
	return &booksEntity, nil
}

func (br *bookRepositoryImpl) GetBook(id int64) (*book_entity.BookEntity, error) {
	book, err := br.bookDB.GetBook(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return book_entity.NewBookEntity(book.ID, book.Name, book.Author), nil
}

func (br *bookRepositoryImpl) CreateBook(name, author string) (*book_entity.BookEntity, error) {
	createBook := book_datasource.CreateBookParams{
		Name:   name,
		Author: author,
	}
	book, err := br.bookDB.CreateBook(context.Background(), createBook)
	if err != nil {
		return nil, err
	}
	return book_entity.NewBookEntity(book.ID, book.Name, book.Author), nil
}

func (br *bookRepositoryImpl) UpdateBook(param book_repo.UpdateBookParam) (*book_entity.BookEntity, error) {
	p := book_datasource.UpdateBookParams{ID: param.ID, Name: param.Name, Author: param.Author}
	newBook, err := br.bookDB.UpdateBook(context.Background(), p)
	if err != nil {
		return nil, nil
	}
	return book_entity.NewBookEntity(newBook.ID, newBook.Name, newBook.Author), nil
}

func (br *bookRepositoryImpl) DeleteBook(id int64) error {
	return br.bookDB.DeleteBook(context.Background(), id)
}
