package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/joho/godotenv"
	"github.com/unsafe0x0/ai"
)

func main() {
	_ = godotenv.Load()

	openRouterApiKey := os.Getenv("OPEN_ROUTER_API")
	groqCloudApiKey := os.Getenv("GROQ_CLOUD_API")
	mistralApiKey := os.Getenv("MISTRAL_API")

	provider := ""
	prompt := &survey.Select{
		Message: "Choose a provider:",
		Options: []string{"OpenRouter", "GroqCloud", "Mistral"},
	}
	survey.AskOne(prompt, &provider)

	var client *ai.SDK
	switch provider {
	case "OpenRouter":
		client = ai.NewSDK(&ai.OpenRouterProvider{
			APIKey: openRouterApiKey,
			Model:  "openrouter/sonoma-sky-alpha",
		})
	case "GroqCloud":
		client = ai.NewSDK(&ai.GroqCloudProvider{
			APIKey: groqCloudApiKey,
			Model:  "openai/gpt-oss-20b",
		})
	case "Mistral":
		client = ai.NewSDK(&ai.MistralProvider{
			APIKey: mistralApiKey,
			Model:  "mistral-small-latest",
		})
	default:
		fmt.Println("No provider selected")
		os.Exit(1)
	}

	systemMessage := ai.Message{
		Role:    "system",
		Content: "you are an helpful coding assistant",
	}

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nPrompt: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		messages := []ai.Message{
			systemMessage,
			{
				Role:    "user",
				Content: input,
			},
		}

		fmt.Print("\nResponse: ")
		err := client.StreamComplete(ctx, messages, func(chunk string) error {
			fmt.Print(chunk)
			return nil
		})

		if err != nil {
			fmt.Println("\nError", err)
		}
	}
}
