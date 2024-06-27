package chunkandembed

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GeminiEmbedder struct {
	apiKey string
}

type content struct {
	Parts []map[string]string `json:"parts"`
}

type GeminiEmbeddingRequestBody struct {
	Model   string  `json:"model"`
	Content content `json:"content"`
}

type GeminiResponse struct {
	Embedding struct {
		Values []float32 `json:"values"`
	} `json:"embedding"`
}

func newGeminiEmbeddingRequestBody(text string) GeminiEmbeddingRequestBody {
	var postBody GeminiEmbeddingRequestBody
	postBody.Model = "models/embedding-001"
	postBody.Content = content{Parts: make([]map[string]string, 1)}
	postBody.Content.Parts[0] = make(map[string]string)
	postBody.Content.Parts[0]["text"] = text
	return postBody
}

func (e GeminiEmbedder) EmbedDocuments(ctx context.Context, texts []string) ([][]float32, error) {

	var embeddings [][]float32
	for _, text := range texts {
		postBody := newGeminiEmbeddingRequestBody(text)

		jsonRequestBody, err := json.Marshal(postBody)
		if err != nil {
			return nil, err
		}

		url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/embedding-001:embedContent?key=%s", e.apiKey)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonRequestBody))

		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var response GeminiResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		embeddings = append(embeddings, response.Embedding.Values)
	}
	return embeddings, nil
}

func (e GeminiEmbedder) EmbedQuery(ctx context.Context, text string) ([]float32, error) {
	postBody := newGeminiEmbeddingRequestBody(text)

	jsonRequestBody, err := json.Marshal(postBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/embedding-001:embedContent?key=%s", e.apiKey)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonRequestBody))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response GeminiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.Embedding.Values, nil
}
