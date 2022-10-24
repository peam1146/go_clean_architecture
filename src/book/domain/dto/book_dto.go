package book_dto

type BookDto struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type BooksDto struct {
	Total uint       `json:"total"`
	Books []*BookDto `json:"books"`
}

type CreateBookDto struct {
	Name   string `json:"name"`
	Author string `json:"aothor"`
}

type UpdateBookDto struct {
	Name   string `json:"name"`
	Author string `json:"aothor"`
}
