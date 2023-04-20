
# BookStore API

``` The BookStore API is a **RESTful API** built using Go programming language and the Gorilla Mux router. It provides CRUD (Create, Read, Update, Delete) operations for managing books in a database. The API allows users to create new books, retrieve all books, retrieve a specific book by ID, update a specific book, and delete a specific book.```

# Technologies Used

* #### Go

* #### Gorilla Mux(Router)

* ####  GORM(ORM)

* #### MySQL(Database)

* ####  Postman(API Testing)

## Installation and Setup
1. Clone the repository:
```bash
$ git clone https://github.com/YohanesGetinet1/CRUD-API.git

```
2. Change into the project directory:
```bash
$ cd CRUD-API
```
3. Install dependencies:
``` go
$ go mod download
```
4. Run the application:
```go
$ go run main.go
```

## Usage
The following endpoints are available:
| Method |Endpoint | Description |
| --- | --- | --- |
| POST| /book/ | Create a new book |
| GET | /book/ |Retrieve all books |
| GET | /book/{bookId}| Retrieve a specific book|
| PUT | /book/{bookId} | Update a specific book |
| DELETE | /book/{bookId}| Delete a specific book |

###  Create a new book
To create a new book, send a POST request to the `/book/` endpoint with the book data in JSON format in the request body.

Example request body:
```json
{
    "name": "The Alchemist",
    "author": "Paulo Coelho",
    "publication": "Penguin publishers"
}

```
###  Retrieve all books

To retrieve all books, send a GET request to the `/book/` endpoint.

###  Retrieve a specific book

To retrieve a specific book, send a GET request to the `/book/{bookId}` endpoint, where `bookId` is the ID of the book to retrieve.

###  Update a specific book

To update a specific book, send a PUT request to the `/book/{bookId}` endpoint with the updated book data in JSON format in the request body.

Example request body:
```json
{
    "name": "The Alchemist",
    "author": "Paulo Coelho",
    "publication": "HarperCollins"
}

```
###  Delete a specific book

To delete a specific book, send a DELETE request to the `/book/{bookId}` endpoint, where `bookId` is the ID of the book to delete.


## Contributing

Pull requests are welcome. 
## Contributors

-   Yohanes Getinet 

## License

This project is licensed under the MIT License - see the LICENSE file for details.