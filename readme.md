# AI SDK Examples

This repository contains examples demonstrating how to use the `github.com/unsafe0x0/ai` Go SDK to build applications that interact with various AI providers.

## Core SDK

All examples in this repository are built using the [`github.com/unsafe0x0/ai`](https://github.com/unsafe0x0/ai) Go SDK. This SDK provides a unified interface for interacting with a variety of AI models and providers.

## Prerequisites

- Go 1.18 or later
- API keys for the AI providers you want to use (see below)

## Installation

1. Clone this repository and navigate to the example you want to run:
   ```bash
   git clone <repo-url>
   cd examples/<example-dir>
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Configuration

Create a `.env` file in the example directory and add your API keys accordingly. Example:

```env
OPEN_ROUTER_API_KEY="your-open-router-api-key"
GROQ_API_KEY="your-groq-cloud-api-key"
MISTRAL_API_KEY="your-mistral-api-key"
OPENAI_API_KEY="your-openai-api-key"
PERPLEXITY_API_KEY="your-perplexity-api-key"
ANTHROPIC_API_KEY="your-anthropic-api-key"
GEMINI_API_KEY="your-gemini-api-key"
ANANNAS_API_KEY="your-annanas-api-key"
```

You can obtain API keys from:

- [OpenRouter](https://openrouter.ai/keys)
- [GroqCloud](https://console.groq.com/keys)
- [Mistral](https://console.mistral.ai/api-keys/)
- [OpenAI](https://platform.openai.com/api-keys)
- [Perplexity](https://www.perplexity.ai/settings/api)
- [Anthropic](https://console.anthropic.com/settings/keys)
- [Gemini](https://aistudio.google.com/app/apikey)
- [Annanas](https://anannas.ai/)

## Examples

Each directory contains a self-contained example:

- [`cli/`](./cli): Command-line interface (CLI) chat with different AI models.
- [`rest-api/`](./rest-api): REST API for English-to-French translation using an AI model.
- [`text-to-speech/`](./text-to-speech): Generate text with an AI model and convert it to speech using ElevenLabs.

See the `readme.md` file within each example's directory for usage and features specific to that example.

## Documentation

For more detailed information about the `github.com/unsafe0x0/ai` SDK, please refer to the [main repository](https://github.com/unsafe0x0/ai).

## Contributing

Contributions are welcome! If you have an idea for a new example or want to improve an existing one, please feel free to open an issue or submit a pull request.

1.  Fork the repository.
2.  Create a new branch (`git checkout -b feature/your-feature`).
3.  Commit your changes (`git commit -m 'Add some feature'`).
4.  Push to the branch (`git push origin feature/your-feature`).
5.  Open a pull request.
