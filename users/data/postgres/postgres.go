package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var (
	counts   int64
	ClientDB *sql.DB
)

func init() {
	_ = godotenv.Load("C:/Projects/go-bookstore/users/.env")

	connectToDB()
}

func openDB(dsn string) (*sql.DB, error) {
	var err error
	ClientDB, err = sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	err = ClientDB.Ping()
	if err != nil {
		panic(err)
	}
	return ClientDB, nil
}

func connectToDB() *sql.DB {
	host := os.Getenv("userdb_host")
	port := os.Getenv("userdb_port")
	user := os.Getenv("userdb_user")
	password := os.Getenv("userdb_password")
	dbname := os.Getenv("userdb_dbname")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC", host, port, user, password, dbname)
	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return conn
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}
		log.Println("Waiting for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
