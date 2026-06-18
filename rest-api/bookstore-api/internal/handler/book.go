package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bookstore-api/internal/model"
	"bookstore-api/internal/service"
	"bookstore-api/pkg/response"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	books := h.service.GetAll()
	response.JSON(w, http.StatusOK, books)
}

func (h *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	book, err := h.service.GetByID(id)
	if err != nil {
		response.JSONError(w, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, book)
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var book *model.Book
	err :=json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	h.service.Create(book)

	response.JSON(w, http.StatusCreated, book)

}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err!=nil{
		response.JSONError(w,http.StatusBadRequest, err.Error())
		return
	}
	var book *model.Book
	err =json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err=h.service.Update(id, book)
	if err != nil {
		response.JSONError(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, book)

}


func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		response.JSONError(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, "book deleted")
}


