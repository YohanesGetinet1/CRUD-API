/*Package controllers implements the HTTP handlers for the CRUD API.
These handlers parse incoming requests,call the appropriate model 
functions to retrieve, update or delete data,and then return the 
data to the client in JSON format. The Gorilla Mux library is used 
to manage routing and URL parameters.
*/

package controllers
import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/YohanesGetinet1/CRUD-API/pkg/utils"
	"github.com/YohanesGetinet1/CRUD-API/pkg/models"
)

// NewBook is an instance of the Book model.
var NewBook models.Book

// GetBook gets all books in the database and returns them as a JSON response.
func GetBook(w http.ResponseWriter, r *http.Request){
	// Get all books in the database
	newBooks := models.GetAllBooks()
	
	// Convert books to JSON format
	res, _ := json.Marshal(newBooks)
	
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Send the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBookById gets a single book by its ID and returns it as a JSON response.
func GetBookById(w http.ResponseWriter, r *http.Request){
	// Get the book ID from the URL parameters
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Convert the book ID to an integer
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		// Handle errors while parsing the book ID
		fmt.Println("error while parsing")
	}

	// Get the book details from the database
	bookDetails, _ := models.GetBookById(ID)
	
	// Convert book details to JSON format
	res, _ := json.Marshal(bookDetails)

	// Set the response content type to JSON
	w.Header().Set("Content-Type","application/json")
	
	// Send the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook creates a new book and returns its details as a JSON response.
func CreateBook(w http.ResponseWriter, r *http.Request){
	// Create a new book instance
	CreateBook := &models.Book{}

	// Parse the request body into the new book instance
	utils.ParseBody(r, CreateBook)

	// Save the new book to the database
	b := CreateBook.CreateBook()

	// Convert the book details to JSON format
	res, _ := json.Marshal(b)

	// Send the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook deletes a book by its ID and returns a message as a JSON response.
func DeleteBook(w http.ResponseWriter, r*http.Request){
	// Get the book ID from the URL parameters
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Convert the book ID to an integer
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		// Handle errors while parsing the book ID
		fmt.Println("error while parsing")
	}

	// Delete the book from the database
	book := models.DeleteBook(ID)

	// Convert the delete message to JSON format
	res, _ := json.Marshal(book)
	
	// Send the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// UpdateBook updates a book by its ID and returns its details as a JSON response.
func UpdateBook(w http.ResponseWriter, r *http.Request){
	// Create a new book instance
	var updateBook = &models.Book{}

	// Parse the request body into the new book instance
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Convert the book ID string to an integer
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		// Print an error message if there was an error while parsing the book ID
		fmt.Println("error while parsing")
	}

	// Retrieve the book details from the database using its ID
	bookDetail, db:=models.GetBookById(ID)

	// Update the book details with the values from the request body, if present
	if updateBook.Name != ""{
		bookDetail.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetail.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetail.Publication = updateBook.Publication
	}

	// Save the updated book details back to the database
	db.Save(&bookDetail)

	// Marshal the updated book details to JSON
	res, _ := json.Marshal(bookDetail)

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the JSON response to the client
	w.Write(res)
}
