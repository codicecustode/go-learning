package main
import (
	"fmt"
	"net/http"

	"bookstore-api/internal/config"
	"bookstore-api/internal/handler"
	"bookstore-api/internal/repository"
	"bookstore-api/internal/service"

	"github.com/go-chi/chi/v5"
)


func main() {
	cfg, err := config.Load()

	if err != nil{
		fmt.Println("Error loading in config:", err)
	}

	repo := repository.NewBookRepository()
	svc := service.NewBookService(repo)
	h := handler.NewBookHandler(svc)

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Get("/books", h.GetAll)
	r.Get("/books/{id}", h.GetByID)
	r.Post("/books", h.Create)
	r.Patch("/books/{id}", h.Update)
	r.Delete("/books/{id}", h.Delete)

	fmt.Printf("Server starting on port %s\n", cfg.Serverport)

	http.ListenAndServe(":"+cfg.Serverport, r)

}