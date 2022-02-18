package book

import (
	"books-api-and-calculation/config"
	"books-api-and-calculation/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	table = "books"
)

// Read
func GetAll(ctx context.Context, filterbooks models.FilterBooks) ([]models.Books, error) {

	var books []models.Books

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	var filter string
	// title
	if filterbooks.Title != "" {
		filter = fmt.Sprintf("title = '%v'", filterbooks.Title)
	}
	// minYear
	if filter != "" {
		filter = filter + " and "
	}
	if filterbooks.MinYear != "" {
		filter = filter + fmt.Sprintf("release_year >= %s", filterbooks.MinYear)
	} else {
		filter = filter + "release_year >= 0"
	}
	// maxYear
	if filter != "" {
		filter = filter + " and "
	}
	if filterbooks.MaxYear != "" {
		filter = filter + fmt.Sprintf("release_year <= %s", filterbooks.MaxYear)
	} else {
		filter = filter + "release_year <= 100000"
	}
	// minpage
	if filter != "" {
		filter = filter + " and "
	}
	if filterbooks.MinPage != "" {
		filter = filter + fmt.Sprintf("total_page >= %s", filterbooks.MinPage)
	} else {
		filter = filter + "total_page >= 0"
	}
	// maxPage
	if filter != "" {
		filter = filter + " and "
	}
	if filterbooks.MaxPage != "" {
		filter = filter + fmt.Sprintf("total_page <= %s", filterbooks.MaxPage)
	} else {
		filter = filter + "total_page <= 100000"
	}
	// sort
	filter = filter + " Order By id "
	if filterbooks.Sort == "asc" {
		filter = filter + "asc"
	} else {
		filter = filter + "desc"
	}
	queryText := fmt.Sprintf("SELECT * FROM %v where %v", table, filter)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Books
		var createdAt, updatedAt string

		if err = rowQuery.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Kategori_ketebalan,
			&createdAt,
			&updatedAt); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// Create
func Insert(ctx context.Context, buku models.Books) error {

	db, err := config.MySQL()

	queryText := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, kategori_ketebalan, created_at, updated_at ) values('%v','%v','%v','%v','%v','%v','%v','%v','%v')",
		table,
		buku.Title,
		buku.Description,
		buku.Image_url,
		buku.ReleaseYear,
		buku.Price,
		buku.TotalPage,
		buku.Kategori_ketebalan,
		buku.CreatedAt,
		buku.UpdatedAt,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

// Update
func Update(ctx context.Context, buku models.Books) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set title = '%s', description ='%s', image_url = '%s', release_year = '%d', price = '%s', total_page = '%s', kategori_ketebalan = '%s', created_at = '%s', updated_at = '%s' where id = '%d'",
		table,
		buku.Title,
		buku.Description,
		buku.Image_url,
		buku.ReleaseYear,
		buku.Price,
		buku.TotalPage,
		buku.Kategori_ketebalan,
		buku.CreatedAt,
		buku.UpdatedAt,
		buku.ID,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

// Delete
func Delete(ctx context.Context, buku models.Books) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, buku.ID)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, _ := s.RowsAffected()

	if check == 0 {
		return errors.New("id tidak ada ")
	}

	return nil
}
