package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func GetMessageFromOpenAI() *string {
	response := ""
	client := openai.NewClient("api key")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return &response
	}

	return &resp.Choices[0].Message.Content
}
