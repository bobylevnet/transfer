package db

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var dbstr string

//Getdtbs интерфейс метод для возврата
type Getdtbs interface {
	GetDB() *sqlx.DB
}

type manager struct {
	db *sqlx.DB
}

var Mgr Getdtbs

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("dbuser")
	password := os.Getenv("dbpass")
	dbName := os.Getenv("dbname")
	dbHost := os.Getenv("dbhost")
	dbstr = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	db, err := sqlx.Connect("pgx", (dbstr))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//if Mgr == nil {
	Mgr = &manager{db: db}
	//}
}

//GetDB  возвращаем DB
func (gr *manager) GetDB() *sqlx.DB {
	return gr.db
}
