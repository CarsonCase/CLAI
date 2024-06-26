/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// askCmd represents the ask command
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask a stand alone question to your LLM.",
	Long: `ask asks a standard text completion to the llm with no context provided

ulike with say chatGPT where there is context of previous messages`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := ConvertAllArgsToPrompt(args)
		var prompts []string
		prompts = append(prompts, prompt)

		Propmt(prompts, func(ctx context.Context, llm *openai.LLM, prompts []string) {
			completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompts[0])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(completion)
		})

	},
}

type GenerateAndPrintLLM func(ctx context.Context, llm *openai.LLM, prompts []string)

func init() {
	rootCmd.AddCommand(askCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// askCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// askCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ConvertAllArgsToPrompt(args []string) string {
	prompt := ""

	for _, arg := range args {
		prompt += arg + " "
	}

	return prompt
}

func Propmt(prompts []string, llmResponseAndPrint GenerateAndPrintLLM) {
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

	model := viper.GetString("model")

	llm, err := openai.New(openai.WithModel(model))
	if err != nil {
		log.Fatal(err)
	}
	llmResponseAndPrint(ctx, llm, prompts)
}
