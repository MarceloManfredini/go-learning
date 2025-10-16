package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Carrega o .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	r := gin.Default()
	log.Default().Println(os.Getenv("OPENAI_API_KEY"))

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	r.POST("/chat", func(c *gin.Context) {
		var req struct {
			Message string `json:"message"`
		}

		//POST http://localhost:8080/chat \
		//"Content-Type: application/json" \
		//'{"message": "Olá, tudo bem?"}'

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida"})
			return
		}

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT4oMini,
				Messages: []openai.ChatCompletionMessage{
					{Role: "user", Content: req.Message},
				},
			},
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"response": resp.Choices[0].Message.Content,
		})
	})

	r.Run(":8080")
}
