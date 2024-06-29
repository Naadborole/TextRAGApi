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
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/vectorstores"
)

var Llm llms.Model
var RAGChain chains.Chain

func init() {
	var err error
	Llm, err = cohere.New(cohere.WithToken(config.GetConfigValue("COHERE_API_KEY")))
	if err != nil {
		panic(err)
	}
	initializeAlternateFlow(Llm)
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

//A different prompt based flow with history

func initializeAlternateFlow(llm llms.Model) {
	template := `Use the following pieces of context to answer the question at the end.
	If you don't know the answer, just say that you don't know, don't try to make up an answer.
	Keep the answer as concise as possible.

	{{.context}}

	Question: {{.question}}

	Helpful Answer:`

	CustomRAGTemplateChain := chains.NewLLMChain(llm, prompts.NewPromptTemplate(template, []string{"context", "question"}))
	// CustomRAGTemplateChain.OutputKey = "answer"
	stuffDocChain := chains.NewStuffDocuments(CustomRAGTemplateChain)
	chatHist := memory.NewChatMessageHistory()
	RAGChain = chains.NewConversationalRetrievalQA(stuffDocChain, chains.LoadCondenseQuestionGenerator(llm),
		vectorstores.ToRetriever(chunkandembed.Store, 3), memory.NewConversationBuffer(memory.WithReturnMessages(true), memory.WithChatHistory(chatHist)))

}
