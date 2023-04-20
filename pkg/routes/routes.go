/*The routes package defines the HTTP routes for the BookStore API using the Gorilla Mux router.
It sets up the routes for creating, retrieving, updating, and deleting book resources, and maps 
each route to its corresponding controller function. The RegisterBookStoreRoutes variable is a
function that takes a pointer to a Mux router as its argument, and sets up the HTTP routes for 
the BookStore API on that router.*/
package routes

import (
	"github.com/YohanesGetinet1/CRUD-API/pkg/controllers"
	"github.com/gorilla/mux"
)
// RegisterBookStoreRoutes sets up the HTTP routes for the BookStore API.
var RegisterBookStoreRoutes = func (router *mux.Router)  {
	// POST /book/ creates a new book resource
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	
	// GET /book/ retrieves all books
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")	
	
	// GET /book/{bookId} retrieves a specific book resource by ID
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	
	// PUT /book/{bookId} updates a specific book resource by ID
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	
	// DELETE /book/{bookId} deletes a specific book resource by ID
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
