package client

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	ApiKey string
	Model  string
}

type Message struct {
	Role    string
	Content string
}

type ChatData struct {
	Choices []struct {
		Message struct {
			Content string
		}
	}
}

type Http interface {
	Post(url string, body []byte, headers map[string]string) (response []byte, err error)
}

const baseUrl = "https://api.openai.com/v1"

func Chat(httpClient Http, prompt string) (res string, err error) {
	response, err := makeRequest(httpClient, prompt)

	if err != nil {
		log.Fatalln(err)
	}

	var data ChatData
	err = json.Unmarshal(response, &data)

	if err != nil {
		log.Fatalln(err)
	}

	return data.Choices[0].Message.Content, err
}

func buildConfig() Config {
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		log.Fatalln("OPENAI_API_KEY is not set")
	}

	var model string
	modelPreference, _ := os.LookupEnv("OPENAI_MODEL")

	if modelPreference != "" {
		model = modelPreference
	} else {
		model = "gpt-3.5-turbo"
	}

	return Config{ApiKey: apiKey, Model: model}
}

func makeRequest(httpClient Http, prompt string) (response []byte, err error) {
	config := buildConfig()
	messages := []map[string]string{{"role": "user", "content": prompt}}

	body, _ := json.Marshal(map[string]interface{}{
		"model":       config.Model,
		"temperature": 0.9,
		"messages":    messages,
	})

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", config.ApiKey),
		"Content-Type":  "application/json",
	}

	url := fmt.Sprintf("%s/chat/completions", baseUrl)

	return httpClient.Post(url, body, headers)
}
