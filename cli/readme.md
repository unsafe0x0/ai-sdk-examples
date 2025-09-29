# AI CLI


This is a command line interface (CLI) application that demonstrates how to use the `github.com/unsafe0x0/ai` package to interact with various AI providers and models.

## Features

-   Interactive provider selection using arrow keys.
-   Supports multiple AI providers and models:
    -   OpenRouter (`openrouter/sonoma-dusk-alpha`)
    -   GroqCloud (`openai/gpt-oss-20b`)
    -   Mistral (`mistral-small-latest`)
    -   OpenAI (`gpt-3.5-turbo`)
    -   Perplexity (`sonar-pro`)
    -   Anthropic (`claude-3.5`)
    -   Gemini (`gemini-2.5-flash`)
    -   Annanas (`mistralai/mistral-small-3.2-24b-instruct:free`)
-   Streaming or non-streaming responses.
-   Optional configuration for max tokens, reasoning effort, temperature, and streaming.


## Setup

See the [root readme](../readme.md) for prerequisites, installation, and configuration instructions.

## Usage

1.  Run the application:

    ```bash
    go run main.go
    ```

2.  Use the arrow keys to select an AI provider and press Enter.
3.  Optionally, enter values for `Max tokens`, `Reasoning effort`, and `Temperature` when prompted. You can also choose whether to stream the response.
4.  Enter your prompt and press Enter to get a response from the AI.
5.  Type `exit` to quit the application.

## How it Works


This application uses the `github.com/unsafe0x0/ai` package to create a client for the selected AI provider and model. The provider and model are selected based on your choice at runtime. The application supports both streaming and non-streaming responses, and allows you to configure options such as max tokens, reasoning effort, and temperature.

Example usage for GroqCloud:

```go
client := ai.GroqCloud(groqCloudApiKey, "openai/gpt-oss-20b")
```

To generate a response (streaming or not):

```go
resp, err := client.Generate(ctx, messages, &opts, nil) // non-streaming
// or
_, err := client.Generate(ctx, messages, &opts, func(chunk string) error {
    fmt.Print(chunk)
    return nil
}) // streaming
```


## Documentation

See the [root readme](../readme.md) for SDK documentation links.

