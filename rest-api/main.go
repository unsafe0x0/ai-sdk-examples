package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/unsafe0x0/ai"
)

func main() {

	_ = godotenv.Load()

	server := gin.Default()
	server.Use(gin.Logger())

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	server.POST("/translate", func(c *gin.Context) {
		type TranslateRequest struct {
			Prompt string `json:"prompt"`
		}
		var req TranslateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
			return
		}

		key := os.Getenv("GROQ_API_KEY")
		if key == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "GROQ_API_KEY not set"})
			return
		}
		client := ai.GroqCloud(key, "openai/gpt-oss-20b")

		systemMsg := ai.Message{
			Role:    "system",
			Content: "You are a helpful assistant that translates English to French.",
		}

		opts := ai.Options{
			MaxCompletionTokens: 100,
			ReasoningEffort:     "low",
			Temperature:         0.7,
			Stream:              false,
		}

		messages := []ai.Message{
			systemMsg,
			{Role: "user", Content: req.Prompt},
		}
		resp, err := client.Generate(c.Request.Context(), messages, &opts, nil)
		if err != nil {
			println("AI SDK error:", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	server.Run()
}
