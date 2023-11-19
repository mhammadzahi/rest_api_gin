package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "muhammad"
//	password = "k9999"
//	dbname   = "testdb"
//)

func main6() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")

}
