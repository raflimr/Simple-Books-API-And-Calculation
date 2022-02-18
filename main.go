package main

import (
	"books-api-and-calculation/crud"
	"books-api-and-calculation/fungsi"
	"books-api-and-calculation/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/bangun-datar/:bangundatar", fungsi.BangunDatar)
	router.GET("/books", middleware.Auth(crud.GetBooks))
	router.POST("/books/create", crud.Postbook)
	router.PUT("/books/update", middleware.Auth(crud.Updatebook))
	router.DELETE("/books/delete", middleware.Auth(crud.Deletebook))
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
