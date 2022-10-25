package book_controller

import (
	"net/http"
	"strconv"

	book_dto "github.com/peam1146/go_clean_architecture/src/book/domain/dto"
	book_repo "github.com/peam1146/go_clean_architecture/src/book/domain/repository"
	common_dto "github.com/peam1146/go_clean_architecture/src/common/dto"
	common_router "github.com/peam1146/go_clean_architecture/src/common/presentation/router"
)

type BookController struct {
	book_repo book_repo.BookRepository
}

func NewBookController(book_repo book_repo.BookRepository) *BookController {
	return &BookController{book_repo}
}

func (ctr *BookController) GetBooks(ctx common_router.IContext) {
	books, err := ctr.book_repo.GetBooks()

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

func (ctr *BookController) GetBook(ctx common_router.IContext) {
	stringId, err := ctx.Param("id")

	if err != nil {
		const status = http.StatusBadRequest
		ctx.JSON(status, common_dto.NewErrorResponse400(status, "No param id"))
		return
	}

	id, err := strconv.ParseInt(stringId, 10, 64)

	if err != nil {
		const status = http.StatusBadRequest
		ctx.JSON(status, common_dto.NewErrorResponse400(status, "Invalid id"))
		return
	}

	bookEntity, err := ctr.book_repo.GetBook(id)

	if err != nil {
		const status = http.StatusBadRequest
		ctx.JSON(status, common_dto.NewErrorResponse400(status, err.Error()))
	}

	bookDto := &book_dto.BookDto{ID: bookEntity.GetID(), Name: bookEntity.GetName(), Author: bookEntity.GetName()}

	const status = http.StatusOK
	ctx.JSON(status, bookDto)
}
