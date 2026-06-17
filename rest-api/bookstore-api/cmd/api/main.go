package main
import (
	"fmt"
	"net/http"
	"bookstore-api/internal/config"
	"github.com/go-chi/chi/v5"
)


func main() {
	cfg, err := config.Load()

	if err != nil{
		fmt.Println("Error loading in config:", err)
	}

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	fmt.Printf("Server starting on port %s\n", cfg.Serverport)

	http.ListenAndServe(":"+cfg.Serverport, r)

}