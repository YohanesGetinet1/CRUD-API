/*The main package is the entry point of a Go program.
It contains the main() function, which is executed when 
the program is run. In this specific code, the main() 
function sets up a Gorilla Mux router, registers the 
routes for a book store API using the routes package,
sets the router as the default handler for all incoming 
requests, and starts the server at port 9010*/

package main

// Import the necessary packages.
import(
	"log"
	"net/http"


	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/YohanesGetinet1/CRUD-API/pkg/routes"
)

// Define the main function that starts the server.
func main(){
	// Create a new Gorilla Mux router.
	r := mux.NewRouter()

	// Register the routes for the book store API using the router.
	routes.RegisterBookStoreRoutes(r)

	// Set the router as the default handler for all incoming requests.
	http.Handle("/", r)

	// Start the server at port 9010 and log any errors that occur.
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}