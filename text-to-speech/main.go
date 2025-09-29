package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/unsafe0x0/ai"
)

func speakWithElevenLabs(apiKey, text string) error {
	url := "https://api.elevenlabs.io/v1/text-to-speech/JBFqnCBsd6RMkjVDRZzb?output_format=mp3_44100_128"
	payload := strings.NewReader(fmt.Sprintf(`{"text": %q, "model_id": "eleven_multilingual_v2"}`, text))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("xi-api-key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("TTS API error: %s\n%s", res.Status, string(body))
	}

	outFile, err := os.Create("output.mp3")
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Unsafe0x0/ai demo with elevenlabs")

	_ = godotenv.Load()

	elevenLabsAPIKey := os.Getenv("ELEVENLABS_API_KEY")
	groqApiKey := os.Getenv("GROQ_API_KEY")
	if elevenLabsAPIKey == "" || groqApiKey == "" {
		fmt.Println("Please set the ELEVENLABS_API_KEY and GROQ_API_KEY environment variables.")
		return
	}

	aiClient := ai.GroqCloud(groqApiKey, "openai/gpt-oss-20b")

	systemMessage := ai.Message{
		Role:    "system",
		Content: "be clear and concise in your answers use plain text as output",
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your question (or type 'exit' to quit): ")
		prompt, _ := reader.ReadString('\n')
		prompt = strings.TrimSpace(prompt)
		if prompt == "exit" {
			break
		}

		messages := []ai.Message{systemMessage, {Role: "user", Content: prompt}}

		opts := &ai.Options{
			MaxCompletionTokens: 512,
			ReasoningEffort:     "low",
			Temperature:         0.1,
			Stream:              false,
		}

		response, err := aiClient.Generate(context.Background(), messages, opts, nil)
		if err != nil {
			fmt.Printf("Chat error: %v\n", err)
			continue
		}
		fmt.Printf("Answer: %s\n", response)
		if err := speakWithElevenLabs(elevenLabsAPIKey, response); err != nil {
			fmt.Printf("[TTS error] %v\n", err)
		}
	}
}
