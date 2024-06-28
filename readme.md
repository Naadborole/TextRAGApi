# API for RAG (Retrieval-Augmented Generation)

This project creates API endpoints which enables users to chat with or ask questions based on supplied text/knowledge.

The application makes use of LangChain [langchaingo](https://github.com/tmc/langchaingo) to communicate with LLMs and facilitate RAG.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
  - [Embed and Store Text](#embed-and-store-text)
  - [Query Document](#query-document)
  - [Query Chat](#query-chat)

## Installation

To install and run the project, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/Naadborole/TextRAGApi.git
    cd TextRAGApi
    ```

2. Start the server:

    ```bash
    go run .
    ```

## Usage

Once the server is running, you can interact with the API endpoints using tools like `curl`, Postman, or any HTTP client of your choice.

## API Endpoints

### Embed and Store Text

Endpoint: `/embedAndStore`

**Method:** POST

**Description:** Embed the uploaded text and store it for retrieval.

**Request Body:**

```json
{
    "text": "Your text data here."
}
```

**Response:**

```json
{
    "ID": "1097108e-49b0-4e65-a931-9b2a91f1d1da"
}
```

### Query Document

Endpoint: `/queryDoc`

**Method:** POST

**Description:** Query the stored documents based on the input text.

**Request Body:**

```json
{
    "text": "Your query here."
}
```

**Response:**

```json
{
    "results": [
        {
            "document": "Relevant document text.",
            "score": 0.95
        },
        ...
    ]
}
```

### Query Chat

Endpoint: `/queryChat`

**Method:** POST

**Description:** Chat with the system based on the input query, using RAG for generating responses.

**Request Body:**

```json
{
    "query": "Your query here."
}
```

**Response:**

```json
{
    "Generated response based on the input query."
}
```