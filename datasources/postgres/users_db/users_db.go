package users_db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "dev_bookstore"
	schema   = "users_db"
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=%s",
		host, port, user, password, dbname, schema)
	Client, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer Client.Close()
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
