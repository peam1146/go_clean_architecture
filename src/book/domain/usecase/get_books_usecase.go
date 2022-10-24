package book_usecase

import (
	"net/http"

	book_dto "github.com/peam1146/go_clean_architecture/src/book/domain/dto"
	book_repo "github.com/peam1146/go_clean_architecture/src/book/domain/repository"
	common_dto "github.com/peam1146/go_clean_architecture/src/common/dto"
	common_router "github.com/peam1146/go_clean_architecture/src/common/presentation/router"
)

func NewGetBooksUseCase(book_repo book_repo.BookRepository) *GetBooksUseCase {
	return &GetBooksUseCase{book_repo}
}

type GetBooksUseCase struct {
	book_repo book_repo.BookRepository
}

func (u *GetBooksUseCase) Call(ctx common_router.IContext) {
	books, err := u.book_repo.GetBooks()

	if err != nil {
		const status = http.StatusBadRequest
		ctx.JSON(status, common_dto.NewErrorResponse400(status, err.Error()))
		return
	}

	var booksDto []book_dto.BookDto
	for _, v := range *books {
		booksDto = append(booksDto, book_dto.BookDto{ID: v.GetID(), Name: v.GetName(), Author: v.GetAuthor()})
	}

	const status = http.StatusOK
	ctx.JSON(status, booksDto)
}
