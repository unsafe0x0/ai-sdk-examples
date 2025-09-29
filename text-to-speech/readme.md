# AI Text-to-Speech

This is a small example that demonstrates using the `github.com/unsafe0x0/ai` package to generate text with a language model (GroqCloud) and convert the result to speech with ElevenLabs.

## Features

- Sends a user prompt to a GroqCloud model (`openai/gpt-oss-20b`) to generate a concise text response.
- Calls the ElevenLabs Text-to-Speech API to synthesize the AI response to an MP3 file (`output.mp3`).
- Interactive loop: enter prompts repeatedly, type `exit` to quit.

## Setup

See the [root readme](../readme.md) for prerequisites and installing Go.

This example expects two environment variables to be set:

- `GROQ_API_KEY` GroqCloud API key used by the `ai` client.
- `ELEVENLABS_API_KEY` ElevenLabs API key used to call the TTS endpoint.

You can create a `.env` file in the example folder (the example uses `godotenv` if present):

```
GROQ_API_KEY=your_groq_api_key_here
ELEVENLABS_API_KEY=your_elevenlabs_api_key_here
```

## Usage

Run the example from the `examples/text-to-speech` directory:

```bash
go run main.go
```

When the program runs you will be prompted to enter a question. The program will:

1. Send the prompt to GroqCloud to generate an answer.
2. Print the answer to stdout.
3. Send the answer text to ElevenLabs and save the resulting audio as `output.mp3` in the current directory.

Type `exit` to quit the interactive loop.

## How it works

The example creates a GroqCloud client via:

```go
client := ai.GroqCloud(groqApiKey, "openai/gpt-oss-20b")
```

It then builds a small chat payload and calls `client.Generate(...)` to get a text response. The `speakWithElevenLabs` helper performs an HTTP POST to the ElevenLabs TTS endpoint and writes the binary MP3 response to `output.mp3`.

Important implementation details:

- The example uses a `system` message to request clear, concise plain-text answers.
- The AI generation options are set to conservative values (512 tokens, low reasoning effort, temperature 0.1, non-streaming).

## Notes

- The example saves synthesized audio to `output.mp3` in the current working directory. Overwriting will occur if the file already exists.
- Ensure both API keys are valid and have sufficient quota.

## Documentation

See the [root readme](../readme.md) for SDK documentation and provider configuration details.
