package users_db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	PGHOST     = "PGHOST"
	PGPORT     = "PGPORT"
	PGUSER     = "PGUSER"
	PGPASSWORD = "PGPASSWORD"
	PGDATABASE = "PGDATABASE"
	PGSCHEMA   = "PGSCHEMA"
)

var (
	host     = os.Getenv(PGHOST)
	port     = os.Getenv(PGPORT)
	user     = os.Getenv(PGUSER)
	password = os.Getenv(PGPASSWORD)
	dbname   = os.Getenv(PGDATABASE)
	schema   = os.Getenv(PGSCHEMA)
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
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
