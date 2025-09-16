# AI CLI

This is a simple command-line interface (CLI) application that demonstrates how to use the `github.com/unsafe0x0/ai` package to interact with various AI providers.

## Features

-   Interactive provider selection using arrow keys.
-   Supports multiple AI providers:
    -   OpenRouter
    -   GroqCloud
    -   Mistral
-   Streaming responses from the AI provider.

## Prerequisites

-   Go 1.18 or later.
-   API keys for the AI providers you want to use.

## Installation

1.  Navigate to the `cli` directory:

    ```bash
    cd cli
    ```

2.  Install the dependencies:

    ```bash
    go mod tidy
    ```

## Configuration

1.  Create a `.env` file in this directory.

2.  Add your API keys to the `.env` file. You can obtain them from the following links:
    -   **OpenRouter:** [Get your API key](https://openrouter.ai/keys)
    -   **GroqCloud:** [Get your API key](https://console.groq.com/keys)
    -   **Mistral:** [Get your API key](https://console.mistral.ai/api-keys/)

    Your `.env` file should look like this:
    ```
    OPEN_ROUTER_API="your-open-router-api-key"
    GROQ_CLOUD_API="your-groq-cloud-api-key"
    MISTRAL_API="your-mistral-api-key"
    ```

## Usage

1.  Run the application:

    ```bash
    go run main.go
    ```

2.  Use the arrow keys to select an AI provider and press Enter.

3.  Enter your prompt and press Enter to get a response from the AI.

4.  Type `exit` to quit the application.

## How it Works

This application uses the `github.com/unsafe0x0/ai` package to create a client for the selected AI provider. The `ai.NewSDK` function is used to create a new client, and it takes a provider-specific configuration struct as an argument.

For example, to create a client for OpenRouter, we use the following code:

```go
client = ai.NewSDK(&ai.OpenRouterProvider{
    APIKey: openRouterApiKey,
    Model:  "openrouter/sonoma-sky-alpha",
})
```

Once the client is created, the `StreamComplete` method is used to send a prompt to the AI and receive a streaming response. The `StreamComplete` method takes a context, a slice of messages, and a callback function as arguments. The callback function is called for each chunk of the response.

```go
err := client.StreamComplete(ctx, messages, func(chunk string) error {
    fmt.Print(chunk)
    return nil
})
```
