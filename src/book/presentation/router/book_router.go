package book_router

import (
	book_controller "github.com/peam1146/go_clean_architecture/src/book/domain/controllers"
	common_router "github.com/peam1146/go_clean_architecture/src/common/presentation/router"
)

func NewBookRouter(bookController *book_controller.BookController, router common_router.IRouter) *BookRouter {
	return &BookRouter{bookController, router}
}

type BookRouter struct {
	bookController *book_controller.BookController
	router         common_router.IRouter
}

func (br *BookRouter) Register() {
	br.router.Get("/", br.bookController.GetBooks)
	br.router.Get("/:id", br.bookController.GetBook)
}
