package textsplitter

import (
	"context"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func Test1(t *testing.T) {
	var ge = GeminiEmbedder{os.Getenv("GEMINI_API_KEY")}
	texts := make([]string, 1)
	texts[0] = "Hello World"
	ge.CreateEmbedding(context.Background(), texts)
}
