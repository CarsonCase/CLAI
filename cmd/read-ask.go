/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// readAskCmd represents the readAsk command
var readAskCmd = &cobra.Command{
	Use:   "read-ask",
	Short: "Read a file to postpend to your prompt",
	Long: `
	example usage:
	clai read-ask myfile.js rewrite this js code into a single python file, code only: 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		prompt := ConvertAllArgsToPrompt(args[0:])

		var prompts []string

		file, err := os.ReadFile(fileName)

		if err != nil {
			log.Fatal(err)
		}

		prompts = append(prompts, prompt)
		prompts = append(prompts, string(file))

		fmt.Println(prompts)

		Propmt(prompts, func(ctx context.Context, llm *openai.LLM, prompts []string) {
			completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompts[0])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(completion)
		})
	},
}

func init() {
	rootCmd.AddCommand(readAskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readAskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readAskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
