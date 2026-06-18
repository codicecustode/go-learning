package service

import(
	"bookstore-api/internal/repository"
	"bookstore-api/internal/model"
)

type BookService struct{
	repo *repository.BookRepository
}
func NewBookService(repo *repository.BookRepository) *BookService{
	return &BookService{
		repo: repo,
	}
}


func (s *BookService) GetAll() []model.Book{
	return s.repo.GetAll()
}

func (s *BookService) GetByID(id int) (*model.Book, error){
	return s.repo.GetByID(id)
}

func(s *BookService) Create(book *model.Book) {
	 s.repo.Create(book)
}

func (s *BookService) Update(id int , book *model.Book) error{
	return s.repo.Update(id, book)
}

func (s *BookService) Delete(id int) error{
	return s.repo.Delete(id)
}