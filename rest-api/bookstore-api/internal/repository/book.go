package repository

import (
	"errors"
	"sync"
	"time"

	"bookstore-api/internal/model"
)


type BookRepository struct{
	books  	[]model.Book
	nextID 	int
	mu   	sync.Mutex   
}

func NewBookRepository() *BookRepository{
	return &BookRepository{
		books: []model.Book{},
		nextID: 1,
	}
}


func (r *BookRepository) GetAll() []model.Book{
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.books
}

func (r *BookRepository) GetByID(id int)  (*model.Book, error){
	for i:= range len(r.books){
		if r.books[i].ID == id{
			book := &r.books[i]
			return book, nil
		}
	}

	return nil, errors.New("Book not found")
}


func (r *BookRepository)Create(book *model.Book){
	r.mu.Lock()
	defer r.mu.Unlock()

	book.ID = r.nextID
	r.nextID++
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	r.books = append(r.books, *book)

}

func (r *BookRepository) Update (id int, book *model.Book) error{
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.books{
		if r.books[i].ID == id{
			r.books[i].Title = book.Title
			r.books[i].Author = book.Author
			r.books[i].Price = book.Price
			r.books[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return errors.New("Book not found")
}

func (r *BookRepository) Delete(id int) error{
	r.mu.Lock()
	defer r.mu.Unlock()

	for i:= range r.books{
		if r.books[i].ID==id{
			r.books = append(r.books[:i],r.books[i+1:]... )
			return nil
		}
	}
	return errors.New("Book not found")
}