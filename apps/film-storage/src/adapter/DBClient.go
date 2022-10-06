package adapter

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "<password>"
	dbname   = "<dbname>"
)

type DBClient struct {
	db *sql.DB
}

func DBNewConnection() (dbClient *DBClient) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Fatalf("error opening clover database: %s", err)
	}
	return &DBClient{db: db}
}

func (client *DBClient) Exec(query string) error {
	return client.Exec(query)
}

func (client *DBClient) Query(query string) (*sql.Rows, error) {
	rows, err := client.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rows, err
}

func (client *DBClient) Ping() error {
	return client.db.Ping()
}

func (client *DBClient) Close() {
	client.db.Close()
	log.Println("closing db connection")
}
