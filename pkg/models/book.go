/*The models package contains the code that interacts with the database to perform 
CRUD (Create, Read, Update, Delete) operations on the Book entity.*/
package models

import (
	"github.com/jinzhu/gorm"
	"github.com/YohanesGetinet1/CRUD-API/pkg/config"
	
)
var db *gorm.DB

// Book is a struct to represent a book entity.
type Book struct{
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

// init function initializes the database connection and auto-migrates the book table.
func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}


// CreateBook creates a new book and returns the created book instance.
func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}


// GetAllBooks returns all books from the database.
func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}


// GetBookById returns a book from the database based on its ID.
func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// DeleteBook deletes a book from the database based on its ID.
func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}