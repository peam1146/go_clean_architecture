package book_router

import (
	book_usecase "github.com/peam1146/go_clean_architecture/src/book/domain/usecase"
	common_router "github.com/peam1146/go_clean_architecture/src/common/presentation/router"
)

type BookDependencies struct {
	GetBookUseCase  *book_usecase.GetBookUseCase
	GetBooksUseCase *book_usecase.GetBooksUseCase
	Router          common_router.IRouter
}

func NewBookRouter(param *BookDependencies) *BookRouter {
	return &BookRouter{
		getBookUseCase:  param.GetBookUseCase,
		getBooksUseCase: param.GetBooksUseCase,
		router:          param.Router,
	}
}

type BookRouter struct {
	getBookUseCase  *book_usecase.GetBookUseCase
	getBooksUseCase *book_usecase.GetBooksUseCase
	router          common_router.IRouter
}

func (br *BookRouter) Register() {
	br.router.Get("/book", br.getBookUseCase.Call)
	br.router.Get("/book/:id", br.getBooksUseCase.Call)
}
