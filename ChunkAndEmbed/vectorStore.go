package chunkandembed

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

var Store pgvector.Store

func init() {
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	ProjectRootPath := filepath.Join(filepath.Dir(b), "../")
	err := godotenv.Load(ProjectRootPath + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := GeminiEmbedder{os.Getenv("GEMINI_API_KEY")}

	Store, err = pgvector.New(
		context.Background(),
		pgvector.WithConnectionURL(os.Getenv("POSTGRES_URL")),
		pgvector.WithEmbedder(e),
		pgvector.WithCollectionName("DocumentStore"),
	)
	if err != nil {
		log.Fatal("Cannot create vector store")
	}

}
