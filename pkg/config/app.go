/*The config package provides a way to establish a database 
connection to a MySQL database. It contains the Connect function 
which is responsible for creating a new connection to the MySQL 
database using the gorm ORM library and the MySQL driver. */
package config

import(
	"fmt"
	"github.com/jinzhu/gorm" // ORM library
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL driver
)

// DB_USERNAME is the username for the MySQL database.
const DB_USERNAME = "Yohanes"

// DB_PASSWORD is the password for the MySQL database.
const DB_PASSWORD = "alxCRUDAPIPROJECT1"

// DB_NAME is the name of the MySQL database.
const DB_NAME = "APIdb"

// DB_HOST is the IP address of the host where the MySQL database is located.
const DB_HOST = "127.0.0.1"

// DB_PORT is the port number of the MySQL database.
const DB_PORT = "3306"

// Database connection instance.
var db *gorm.DB

// Connect function establishes a connection to a MySQL database.
func Connect() {
	dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	fmt.Println("dsn : ", dsn)
	if err != nil {
		fmt.Println("Failed to connect to MySQL database:", err)
		panic(err)
	}
	db = d
}

// GetDB returns a reference to the database instance.
func GetDB() *gorm.DB{
	return db
}