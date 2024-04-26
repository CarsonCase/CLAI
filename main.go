package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/CarsonCase/CLAI/cmd"

	"github.com/spf13/viper"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	cmd.Execute()
}

func propmt() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.config/clai")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	apiKey := viper.GetString("api_key")

	if apiKey == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("This must be your first time using clai! Paste your OpenAI API key here: ")
		apiKeyInput, _ := reader.ReadString('\n')
		// Store the API key securely
		apiKey = strings.TrimSpace(apiKeyInput)
		viper.Set("api_key", apiKey)

		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("Error writing config file:", err)
			return
		}
	}
	err = os.Setenv("OPENAI_API_KEY", apiKey)
	if err != nil {
		fmt.Println("Error storing API key:", err)
		return
	}

	ctx := context.Background()

	llm, err := openai.New(openai.WithModel("gpt-3.5-turbo-0125"))
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args[1:]
	prompt := ""

	for _, arg := range args {
		prompt += arg + " "
	}

	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}
