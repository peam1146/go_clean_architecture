package entity

func NewBookEntity(id int64, name, author string) *BookEntity {
	return &BookEntity{id, name, author}
}

type BookEntity struct {
	id     int64
	name   string
	author string
}

func (b *BookEntity) GetID() int64 {
	return b.id
}

func (b *BookEntity) GetName() string {
	return b.name
}

func (b *BookEntity) GetAuthor() string {
	return b.author
}
