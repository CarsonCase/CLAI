/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Ask a question with all the context of your past interactions, and log this one as well",
	Long:  `Unlike with ask, chat will save each prompt and response much like with the Chat GPT interface. Allowing prompts to build on and reference previous prompts and responses`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := ConvertAllArgsToPrompt(args)
		var prompts []string
		prompts = append(prompts, prompt)

		// Get the home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			// Handle error
			panic(err)
		}

		CSV_PATH := homeDir + "/.config/clai/chatData/chat.csv"

		// save prompt to .csv, in ./chatData/chat.csv
		WriteToCSV(CSV_PATH, llms.ChatMessageTypeHuman, prompt)

		// get response as a chat
		Propmt(prompts, func(ctx context.Context, llm *openai.LLM, prompts []string) {

			// create system prompt
			content := []llms.MessageContent{
				llms.TextParts(llms.ChatMessageTypeSystem, "You are a helpful assistant based in a linux CLAI"),
			}

			// read in all past prompts
			contentPointer, err := ReadCSVToContent(CSV_PATH, content)

			content = *contentPointer

			if err != nil {
				log.Fatal(err)
			}

			// append the newest prompt to the content
			content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, prompts[0]))

			var completeResponse string
			if _, err := llm.GenerateContent(ctx, content,
				llms.WithMaxTokens(1024),
				llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
					completeResponse += string(chunk)
					fmt.Print(string(chunk))
					return nil
				})); err != nil {
				log.Fatal(err)
			}

			// append the complete response to the csv
			WriteToCSV(CSV_PATH, llms.ChatMessageTypeAI, completeResponse)
		})
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func WriteToCSV(path string, messageType llms.ChatMessageType, message string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	timestamp := time.Now().Format(time.RFC3339)
	record := []string{timestamp, string(messageType), message}

	if err := writer.Write(record); err != nil {
		return err
	}

	return nil
}

func ReadCSVToContent(path string, content []llms.MessageContent) (*[]llms.MessageContent, error) {
	file, err := os.Open(path)
	if err != nil {
		return &[]llms.MessageContent{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()

	if err != nil {
		return &[]llms.MessageContent{}, err
	}

	for _, line := range data {
		var messageType llms.ChatMessageType
		var text string

		for j, field := range line {
			switch j {
			case 1:
				{
					messageType = llms.ChatMessageType(field)
				}
			case 2:
				{
					text = field
				}
			}
		}
		content = append(content, llms.TextParts(messageType, text))
	}
	return &content, nil

}
