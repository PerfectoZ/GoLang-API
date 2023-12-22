package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/SimpleBank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Cannot connect to Database", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
