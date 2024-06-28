package chunkandembed

import (
	"context"
	"log"

	"github.com/Naadborole/TextRAGApi/config"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

var Store pgvector.Store

func init() {

	e := GeminiEmbedder{config.GetConfigValue("GEMINI_API_KEY")}
	var err error
	Store, err = pgvector.New(
		context.Background(),
		pgvector.WithConnectionURL(config.GetConfigValue("POSTGRES_URL")),
		pgvector.WithEmbedder(e),
		pgvector.WithCollectionName("DocumentStore"),
	)
	if err != nil {
		log.Fatal("Cannot create vector store")
	}

}
