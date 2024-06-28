package generatechat

import (
	"context"

	chunkandembed "github.com/Naadborole/TextRAGApi/ChunkAndEmbed"
	"github.com/Naadborole/TextRAGApi/config"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/cohere"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/vectorstores"
)

var Llm llms.Model
var RAGChain chains.Chain

func init() {
	initializeCohereAPI()
}

func initializeGoogleAPI() {
	var err error
	Llm, err = googleai.New(context.Background(), googleai.WithAPIKey(config.GetConfigValue("GEMINI_API_KEY")))
	if err != nil {
		panic(err)
	}

	RAGChain = chains.NewConversationalRetrievalQAFromLLM(Llm,
		vectorstores.ToRetriever(chunkandembed.Store, 3), memory.NewConversationBuffer())
}

func initializeCohereAPI() {
	var err error
	Llm, err = cohere.New(cohere.WithToken(config.GetConfigValue("COHERE_API_KEY")))
	if err != nil {
		panic(err)
	}
	RAGChain = chains.NewConversationalRetrievalQAFromLLM(Llm,
		vectorstores.ToRetriever(chunkandembed.Store, 3), memory.NewConversationBuffer())
}

func GetResponse(text string) string {
	res, err := chains.Run(context.Background(), RAGChain, text)
	if err != nil {
		panic(err)
	}
	return res
}
