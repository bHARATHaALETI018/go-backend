package db

import (
	"database/sql"
	"log"
	"os"
	"testing"


	_ "github.com/lib/pq"
)


const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:password@localhost:5432/opconnect?sslmode=disable"
)
var testQueries *Queries

func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver, dbSource)
	if err!=nil{
		log.Fatal("Cannnot connect to DB:", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}

