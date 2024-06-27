package chunkandembed

import (
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
)

func SplitText(texts []string, metadatas []map[string]any) ([]schema.Document, error) {
	return textsplitter.CreateDocuments(textsplitter.NewRecursiveCharacter(textsplitter.WithChunkSize(500)), texts, metadatas)
}
