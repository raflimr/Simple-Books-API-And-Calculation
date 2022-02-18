package crud

import (
	book "books-api-and-calculation/books"
	"books-api-and-calculation/fungsi"
	"books-api-and-calculation/models"
	"books-api-and-calculation/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Getbook
func GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		var booksF models.FilterBooks
		queryValues := r.URL.Query()
		booksF.Title = queryValues.Get("title")
		booksF.MinYear = queryValues.Get("minYear")
		booksF.MaxYear = queryValues.Get("maxYear")
		booksF.MinPage = queryValues.Get("minPage")
		booksF.MaxPage = queryValues.Get("maxPage")
		booksF.Sort = queryValues.Get("sort")
		books, err := book.GetAll(ctx, booksF)

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, books, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// Postbook
func Postbook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var buku models.Books
		var bukuRequest models.BooksRequest

		if err := json.NewDecoder(r.Body).Decode(&bukuRequest); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		_, errorImageURL := url.ParseRequestURI(bukuRequest.Image_url)
		release, errorReleaseYear := fungsi.GetReleaseYear(bukuRequest.ReleaseYear)
		var errorMessage string
		if errorImageURL != nil {
			errorMessage = "Image URL tidak valid"
		}
		if errorReleaseYear != nil {
			errorMessage = errorMessage + "\n" + "Release Year tidak valid"
		}

		if errorImageURL != nil || errorReleaseYear != nil {
			utils.ResponseJSON(w, errors.New(errorMessage), http.StatusInternalServerError)
			return
		}

		var err error
		bukuRequest.ReleaseYear, err = strconv.Atoi(release)
		if err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}
		halaman := fungsi.GetTotalPage(bukuRequest.TotalPage)
		bukuRequest.Kategori_ketebalan = halaman
		price := fungsi.GetPrice(bukuRequest.Price)
		buku.ID = bukuRequest.ID
		buku.Title = bukuRequest.Title
		buku.Description = bukuRequest.Description
		buku.Image_url = bukuRequest.Image_url
		buku.ReleaseYear = bukuRequest.ReleaseYear
		buku.Price = price
		buku.TotalPage = bukuRequest.TotalPage
		buku.Kategori_ketebalan = bukuRequest.Kategori_ketebalan
		buku.CreatedAt = bukuRequest.CreatedAt
		buku.UpdatedAt = bukuRequest.UpdatedAt
		if err := book.Insert(ctx, buku); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// Updatebook
func Updatebook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var buku models.Books
		var bukuRequest models.BooksRequest

		if err := json.NewDecoder(r.Body).Decode(&bukuRequest); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		_, errorImageURL := url.ParseRequestURI(bukuRequest.Image_url)
		release, errorReleaseYear := fungsi.GetReleaseYear(bukuRequest.ReleaseYear)
		var errorMessage string
		if errorImageURL != nil {
			errorMessage = "Image URL tidak valid"
		}
		if errorReleaseYear != nil {
			errorMessage = errorMessage + "\n" + "Release Year tidak valid"
		}
		if errorImageURL != nil || errorReleaseYear != nil {
			utils.ResponseJSON(w, errors.New(errorMessage), http.StatusInternalServerError)
			return
		}
		var err error
		bukuRequest.ReleaseYear, err = strconv.Atoi(release)
		if err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}
		halaman := fungsi.GetTotalPage(bukuRequest.TotalPage)
		bukuRequest.Kategori_ketebalan = halaman
		price := fungsi.GetPrice(bukuRequest.Price)
		buku.ID = bukuRequest.ID
		buku.Title = bukuRequest.Title
		buku.Description = bukuRequest.Description
		buku.Image_url = bukuRequest.Image_url
		buku.ReleaseYear = bukuRequest.ReleaseYear
		buku.Price = price
		buku.TotalPage = bukuRequest.TotalPage
		buku.Kategori_ketebalan = bukuRequest.Kategori_ketebalan
		buku.CreatedAt = bukuRequest.CreatedAt
		buku.UpdatedAt = bukuRequest.UpdatedAt
		if err := book.Update(ctx, buku); err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// Deletebook
func Deletebook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Method == "DELETE" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var buku models.Books

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		buku.ID, _ = strconv.Atoi(id)
		if err := book.Delete(ctx, buku); err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}
