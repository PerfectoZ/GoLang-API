package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"testing"
	"time"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/SimpleBank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to Database", err)
	}
	testQueries = New(conn)
	defer conn.Close(ctx)
	os.Exit(m.Run())
}
