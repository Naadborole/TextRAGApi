package chunkandembed

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

var Store pgvector.Store

func init() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := GeminiEmbedder{os.Getenv("GEMINI_API_KEY")}

	Store, err := pgvector.New(
		context.Background(),
		pgvector.WithConnectionURL("postgres://testuser:testpass@localhost:5432/testdb?sslmode=disable"),
		pgvector.WithEmbedder(e),
	)
	if err != nil {
		log.Fatal("Cannot create vector store")
	}
	Store.DropTables(context.Background())
}
