package models

import "time"

type (
	// Books
	Books struct {
		ID                 int       `json:"id"`
		Title              string    `json:"title"`
		Description        string    `json:"description"`
		Image_url          string    `json:"image_url"`
		ReleaseYear        int       `json:"release_year"`
		Price              string    `json:"price"`
		TotalPage          string    `json:"total_page"`
		Kategori_ketebalan string    `json:"kategori_ketebalan"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
	}
	BooksRequest struct {
		ID                 int       `json:"id"`
		Title              string    `json:"title"`
		Description        string    `json:"description"`
		Image_url          string    `json:"image_url"`
		ReleaseYear        int       `json:"release_year"`
		Price              int       `json:"price"`
		TotalPage          string    `json:"total_page"`
		Kategori_ketebalan string    `json:"kategori_ketebalan"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
	}
	FilterBooks struct {
		Title   string
		MinYear string
		MaxYear string
		MinPage string
		MaxPage string
		Sort    string
	}
)
