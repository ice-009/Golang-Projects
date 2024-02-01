package routes

import (
	"github.com/gorilla/mux"
	"github.com/ice-009/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{BookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{BookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{BookId}", controllers.DeleteBook).Methods("DELETE")
}