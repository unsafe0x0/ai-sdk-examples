# AI REST API

This is a simple REST API application that demonstrates how to use the `github.com/unsafe0x0/ai` package to interact with AI providers for language tasks.

## Features

-   RESTful endpoint for English-to-French translation using an AI model.
-   Uses GroqCloud (`openai/gpt-oss-20b`) as the provider.
-   Configurable via environment variable for API key.
-   Built with [Gin](https://github.com/gin-gonic/gin) web framework.


## Setup

See the [root readme](../readme.md) for prerequisites, installation, and configuration instructions.

## Usage

1.  Run the application:

    ```bash
    go run main.go
    ```

2.  The server will start (default: `localhost:8080`).

### Endpoints

-   `GET /ping` — Health check endpoint. Returns `{ "message": "pong" }`.
-   `POST /translate` — Translate English text to French.

    **Request Body:**
    ```json
    {
      "prompt": "Hello, how are you?"
    }
    ```

    **Response:**
    ```json
    {
      "role": "assistant",
      "content": "Bonjour, comment ça va ?"
    }
    ```

## How it Works

The API uses the `github.com/unsafe0x0/ai` package to create a GroqCloud client and sends a prompt to the model for translation. The `/translate` endpoint expects a JSON body with a `prompt` field (English text), and returns the translated French text.


## Documentation

See the [root readme](../readme.md) for SDK documentation links.