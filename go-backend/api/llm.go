package api

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
	"log"
	"os"
)

func Llm(question string) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

	log.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content
}
