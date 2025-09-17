package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/joho/godotenv"
	"github.com/unsafe0x0/ai"
)

func main() {
	_ = godotenv.Load()

	openRouterApiKey := os.Getenv("OPEN_ROUTER_API_KEY")
	groqCloudApiKey := os.Getenv("GROQ_API_KEY")
	mistralApiKey := os.Getenv("MISTRAL_API_KEY")
	openAiApiKey := os.Getenv("OPENAI_API_KEY")
	perplexityApiKey := os.Getenv("PERPLEXITY_API_KEY")
	anthropicApiKey := os.Getenv("ANTHROPIC_API_KEY")
	geminiApiKey := os.Getenv("GEMINI_API_KEY")

	provider := ""
	prompt := &survey.Select{
		Message: "Choose a provider:",
		Options: []string{"OpenRouter", "GroqCloud", "Mistral", "OpenAI", "Perplexity", "Anthropic", "Gemini"},
	}
	survey.AskOne(prompt, &provider)

	var client *ai.SDK
	switch provider {
	case "OpenRouter":
		if openRouterApiKey == "" {
			fmt.Println("OPEN_ROUTER_API_KEY not set")
			return
		}
		client = ai.NewSDK(ai.NewOpenRouterProvider(openRouterApiKey, "openrouter/sonoma-dusk-alpha"))
	case "GroqCloud":
		if groqCloudApiKey == "" {
			fmt.Println("GROQ_API_KEY not set")
			return
		}
		client = ai.NewSDK(ai.NewGroqCloudProvider(groqCloudApiKey, "llama3-8b-8192"))
	case "Mistral":
		if mistralApiKey == "" {
			fmt.Println("MISTRAL_API_KEY not set")
			return
		}
		client = ai.NewSDK(ai.NewMistralProvider(mistralApiKey, "mistral-small-latest"))
	case "OpenAI":
		if openAiApiKey == "" {
			fmt.Println("OPENAI_API_KEY not set")
			return
		}
		client = ai.NewSDK(ai.NewOpenAiProvider(openAiApiKey, "gpt-3.5-turbo"))
	case "Perplexity":
		if perplexityApiKey == "" {
			fmt.Println("PERPLEXITY_API_KEY not set")
			return
		}

		client = ai.NewSDK(ai.NewPerplexityProvider(perplexityApiKey, "sonar-pro"))
	case "Anthropic":
		if anthropicApiKey == "" {
			fmt.Println("ANTHROPIC_API_KEY not set")
			return
		}
		client = ai.NewSDK(ai.NewAnthropicProvider(anthropicApiKey, "claude-3.5"))
	case "Gemini":
		if geminiApiKey == "" {
			fmt.Println("GEMINI_API_KEY not set")
			return
		}
		client = ai.NewSDK(ai.NewGeminiProvider(geminiApiKey, "gemini-2.5-flash"))
	default:
		fmt.Println("No provider selected")
		os.Exit(1)
	}

	systemMessage := ai.Message{
		Role:    "system",
		Content: "your custom system prompt",
	}

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	var maxTokensStr string
	var reasoningEffortStr string
	survey.AskOne(&survey.Input{Message: "Max tokens (optional, press enter to skip):"}, &maxTokensStr)
	survey.AskOne(&survey.Input{Message: "Reasoning effort (optional, press enter to skip, e.g. 1=low, 2=medium, 3=high):"}, &reasoningEffortStr)

	var opts ai.Options
	if maxTokensStr != "" {
		if v, err := strconv.Atoi(maxTokensStr); err == nil {
			opts.MaxTokens = &v
		}
	}
	if reasoningEffortStr != "" {
		if v, err := strconv.Atoi(reasoningEffortStr); err == nil {
			opts.ReasoningEffort = &v
		}
	}

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
		err := client.StreamCompleteWithOptions(ctx, messages, func(chunk string) error {
			fmt.Print(chunk)
			return nil
		}, &opts)

		if err != nil {
			fmt.Println("\nError", err)
		}
	}
}
