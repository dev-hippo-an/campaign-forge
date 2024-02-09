package main

import (
	"context"
	"encoding/json"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"log"
)

func GetCampaignList(category CampaignCategory) []Campaign {
	marshal, err := json.Marshal(category)

	if err != nil {
		return make([]Campaign, 0)
	}
	response := CampaignsResponse{}
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(`%s; request: %s`, initialScript, marshal),
				},
			},
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return response.Campaigns
	}

	responseContent := &resp.Choices[0].Message.Content
	err = json.Unmarshal([]byte(*responseContent), &response)
	if err != nil {
		return make([]Campaign, 0)
	}

	return response.Campaigns
}

func GetCampaign(campaign Campaign) CampaignDetail {
	marshal, err := json.Marshal(campaign)

	if err != nil {
		return CampaignDetail{}
	}

	response := CampaignDetailResponse{}
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(`%s; request: %s`, initialScript, marshal),
				},
			},
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return CampaignDetail{}
	}

	responseContent := &resp.Choices[0].Message.Content
	err = json.Unmarshal([]byte(*responseContent), &response)
	if err != nil {
		return CampaignDetail{}
	}

	return response.CampaignDetail
}
