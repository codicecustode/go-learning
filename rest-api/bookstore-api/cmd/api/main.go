package main
import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()

	port := os.Getenv("SERVER_PORT")

	if port == ""{
		port = "8000"
	}

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	fmt.Printf("Server starting on port %s\n", port)

	http.ListenAndServe(":"+port, r)

}