package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Naadborole/TextRAGApi/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ConnPool *pgxpool.Pool

func init() {
	ConnPool = initDB()
	initializeSchema()
}

func initDB() *pgxpool.Pool {
	ConnPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection succeeded!")
	return ConnPool
}

func initializeSchema() {
	_, err := ConnPool.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS index (id VARCHAR(255), name VARCHAR(255), nDocuments INT)")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create index table")
	}
	// _, err := ConnPool.Exec(context.Background(), "")
}

func GetIndexList() []models.Index {
	rows, _ := ConnPool.Query(context.Background(), "SELECT * FROM index")
	indexes, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Index])
	if err != nil {
		log.Fatal(err)
	}
	return indexes
}

func DoesIndexExists(id string) bool {
	var answer bool
	err := ConnPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT id FROM index where id = $1)", id).Scan(&answer)
	if err != nil {
		log.Fatal((err))
	}
	return answer
}

func AddIndex(i models.Index) (bool, error) {
	_, err := ConnPool.Exec(context.Background(), "INSERT INTO index VALUES($1, $2, $3)", i.Id, i.Name, i.NDocuments)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetIndex(id string) (models.Index, error) {
	rows, err := ConnPool.Query(context.Background(), "SELECT * FROM index WHERE id = $1", id)
	if err != nil {
		return models.Index{}, err
	}
	ind, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Index])
	if err != nil {
		return models.Index{}, err
	}
	return ind, nil
}
