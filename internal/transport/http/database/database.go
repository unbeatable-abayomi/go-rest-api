package database

import (
	"fmt"
	_"os"
    _"database/sql"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
)

// NewDatabase - returns a pointer to a database object

func NewDatabase()(*gorm.DB, error){
	fmt.Println("Setting up new database connection")
	// dbUsername := os.Getenv("DB_USERNAME")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbTable := os.Getenv("DB_TABLE")
	// dbPort := os.Getenv("DB_PORT")

	dbUsername := "postgres"
	dbPassword := "abayomi"
	//dbHost := "localhost"
	dbHost := "host.docker.internal"
	dbTable:= "Restapi"
	dbPort := 5432



	connectString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	db, err := gorm.Open("postgres", connectString)
if err != nil {
 return db, err
}

if err := db.DB().Ping(); err != nil{
	return db, err
}

	return db, nil
}



// func NewDatabase()(*sql.DB, error){
// 	fmt.Println("Setting up new database connection")


// 	dbUsername := "postgres"
// 	dbPassword := "abayomi"
// 	dbHost := "localhost"
// 	dbName := "Restapi"
// 	  dbTable:= "Comments"
// 	dbPort := 5432
// 	dbname := "CommentAPI"

	

// 	// connectString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
// 	// db, err := gorm.Open("postgres", connectString)


// 	url := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",dbHost, dbPort, dbUsername, dbPassword)
//      db, err := sql.Open("postgres", url)
// if err != nil {
//     panic(err)
// }
// defer db.Close()

// _, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbname))
// if err != nil {
//     panic(err)
// }

// return db, err

// //dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s",)
// //DB, _ := gorm.Open(postgres.Open(dsn))

// //createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", Config("DB_NAME"))
// //DB.Exec(createDatabaseCommand)

// }