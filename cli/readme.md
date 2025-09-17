# AI CLI

This is a simple command line interface (CLI) application that demonstrates how to use the `github.com/unsafe0x0/ai` package to interact with various AI providers.

## Features

-   Interactive provider selection using arrow keys.
-   Supports multiple AI providers:
    -   OpenRouter
    -   GroqCloud
    -   Mistral
    -   OpenAI
    -   Perplexity
    -   Anthropic
    -   Gemini
-   Streaming responses from the AI provider.
-   Optional configuration for `MaxTokens` and `ReasoningEffort`.

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
    -   **OpenAI:** [Get your API key](https://platform.openai.com/api-keys)
    -   **Perplexity:** [Get your API key](https://www.perplexity.ai/settings/api)
    -   **Anthropic:** [Get your API key](https://console.anthropic.com/settings/keys)
    -   **Gemini:** [Get your API key](https://aistudio.google.com/app/apikey)

    Your `.env` file should look like this:
    ```
    OPEN_ROUTER_API_KEY="your-open-router-api-key"
    GROQ_API_KEY="your-groq-cloud-api-key"
    MISTRAL_API_KEY="your-mistral-api-key"
    OPENAI_API_KEY="your-openai-api-key"
    PERPLEXITY_API_KEY="your-perplexity-api-key"
    ANTHROPIC_API_KEY="your-anthropic-api-key"
    GEMINI_API_KEY="your-gemini-api-key"
    ```

## Usage

1.  Run the application:

    ```bash
    go run main.go
    ```

2.  Use the arrow keys to select an AI provider and press Enter.
3.  Optionally, enter values for `Max tokens` and `Reasoning effort` when prompted.
4.  Enter your prompt and press Enter to get a response from the AI.
5.  Type `exit` to quit the application.

## How it Works

This application uses the `github.com/unsafe0x0/ai` package to create a client for the selected AI provider. The `ai.NewSDK` function is used with a provider-specific constructor to create a new client.

For example, to create a client for GroqCloud, we use the following code:

```go
client = ai.NewSDK(ai.NewGroqCloudProvider(groqCloudApiKey, "llama3-8b-8192"))
```

Once the client is created, the `StreamCompleteWithOptions` method is used to send a prompt to the AI and receive a streaming response, along with any optional parameters.

```go
err := client.StreamCompleteWithOptions(ctx, messages, func(chunk string) error {
    fmt.Print(chunk)
    return nil
}, &opts)
```

## Documentation

For more detailed information about the `github.com/unsafe0x0/ai` SDK, please refer to the [main repository](https://github.com/unsafe0x0/ai).
