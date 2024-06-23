package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var database *sql.DB

func SetupDB() {
	db := connectToDB()
	database = db
	createMissingTables()
}

func connectToDB() *sql.DB {
	user, password, dbName := getEnv()
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbName)
	db, err := sql.Open("postgres", connStr)

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func getEnv() (string, string, string) {
	user := os.Getenv("user")
	if user == "" {
		log.Fatal("Database username wasn't set")
	}
	password := os.Getenv("password")
	if user == "" {
		log.Fatal("Database password wasn't set")
	}
	dbName := os.Getenv("databaseName")
	if user == "" {
		log.Fatal("Database name wasn't set")
	}
	return user, password, dbName
}

func createMissingTables() {
	query := "CREATE TABLE IF NOT EXISTS User (id VARCHAR(50),email VARCHAR(50),password VARCHAR(50),vatNumber VARCHAR(50),name VARCHAR(50));"
	ExecUpdate(query)
}

func ExecQuery(query string) *sql.Rows {
	res, err := database.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func ExecUpdate(query string) {
	_, err := database.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
