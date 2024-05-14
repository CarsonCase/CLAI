/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setModelCmd represents the setModel command
var setModelCmd = &cobra.Command{
	Use:   "set-model",
	Short: "Set the AI model to use | currently only openai",
	Long: `Sets the AI model to be used.

At the moment you can only use the following models from OpenAI:
gpt-4o
	DESCRIPTION: Our most advanced, multimodal flagship model that’s cheaper and faster than GPT-4 Turbo. Currently points to gpt-4o-2024-05-13.
	CONTEXT WINDOW: 128,000 tokens
	TRAINING DATA: Up to Oct 2023
gpt-4-turbo
	DESCRIPTION: New GPT-4 Turbo with Vision
	CONTEXT WINDOW: 128,000 tokens
	TRAINING DATA: Up to Dec 2023
gpt-4-turbo-2024-04-09
	DESCRIPTION: GPT-4 Turbo with Vision model.
	CONTEXT WINDOW: 128,000 tokens
	TRAINING DATA: Up to Dec 2023
gpt-4-turbo-preview
	DESCRIPTION: GPT-4 Turbo preview model.
	CONTEXT WINDOW: 128,000 tokens
	TRAINING DATA: Up to Dec 2023
gpt-4-0125-preview
	DESCRIPTION: GPT-4 Turbo preview model intended to reduce cases of “laziness” where the model doesn’t complete a task.
	CONTEXT WINDOW: 128,000 tokens
	TRAINING DATA: Up to Dec 2023
gpt-4-1106-preview
	DESCRIPTION: GPT-4 Turbo preview model featuring improved instruction following, JSON mode, reproducible outputs, parallel function calling, and more.
	CONTEXT WINDOW: 128,000 tokens
	TRAINING DATA: Up to Apr 2023
gpt-4-vision-preview
	DESCRIPTION: GPT-4 model with the ability to understand images, in addition to all other GPT-4 Turbo capabilities.
	CONTEXT WINDOW: 128,000 tokens
	TRAINING DATA: Up to Apr 2023
gpt-4-1106-vision-preview
	DESCRIPTION: GPT-4 model with the ability to understand images, in addition to all other GPT-4 Turbo capabilities.
	CONTEXT WINDOW: 128,000 tokens
	TRAINING DATA: Up to Apr 2023
gpt-4
	DESCRIPTION: Currently points to gpt-4-0613.
	CONTEXT WINDOW: 8,192 tokens
	TRAINING DATA: Up to Sep 2021
gpt-4-0613
	DESCRIPTION: Snapshot of gpt-4 from June 13th 2023 with improved function calling support.
	CONTEXT WINDOW: 8,192 tokens
	TRAINING DATA: Up to Sep 2021
gpt-4-32k
	DESCRIPTION: Currently points to gpt-4-32k-0613. This model was never rolled out widely in favor of GPT-4 Turbo.
	CONTEXT WINDOW: 32,768 tokens
	TRAINING DATA: Up to Sep 2021
gpt-4-32k-0613
	DESCRIPTION: Snapshot of gpt-4-32k from June 13th 2023 with improved function calling support. This model was never rolled out widely in favor of GPT-4 Turbo.
	CONTEXT WINDOW: 32,768 tokens
	TRAINING DATA: Up to Sep 2021
gpt-3.5-turbo-0125
	DESCRIPTION: New Updated GPT 3.5 Turbo
	CONTEXT WINDOW: 16,385 tokens
	TRAINING DATA: Up to Sep 2021
gpt-3.5-turbo
	DESCRIPTION: Currently points to gpt-3.5-turbo-0125.
	CONTEXT WINDOW: 16,385 tokens
	TRAINING DATA: Up to Sep 2021
gpt-3.5-turbo-1106
	DESCRIPTION: GPT-3.5 Turbo model with improved instruction following, JSON mode, reproducible outputs, parallel function calling, and more.
	CONTEXT WINDOW: 16,385 tokens
	TRAINING DATA: Up to Sep 2021
gpt-3.5-turbo-instruct
	DESCRIPTION: Similar capabilities as GPT-3 era models. Compatible with legacy Completions endpoint and not Chat Completions.
	CONTEXT WINDOW: 4,096 tokens
	TRAINING DATA: Up to Sep 2021
gpt-3.5-turbo-16k
	DESCRIPTION: Legacy Currently points to gpt-3.5-turbo-16k-0613.
	CONTEXT WINDOW: 16,385 tokens
	TRAINING DATA: Up to Sep 2021
gpt-3.5-turbo-0613
	DESCRIPTION: Legacy Snapshot of gpt-3.5-turbo from June 13th 2023. Will be deprecated on June 13, 2024.
	CONTEXT WINDOW: 4,096 tokens
	TRAINING DATA: Up to Sep 2021
gpt-3.5-turbo-16k-0613
	DESCRIPTION: Legacy Snapshot of gpt-3.5-16k-turbo from June 13th 2023. Will be deprecated on June 13, 2024.
	CONTEXT WINDOW: 16,385 tokens
	TRAINING DATA: Up to Sep 2021
.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Need to pass exactly 1 model name")
		}
		SetModel(args[0])
	},
}

func init() {
	rootCmd.AddCommand(setModelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setModelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setModelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func SetModel(model string) {
	models := [...]string{
		"gpt-4o",
		"gpt-4-turbo",
		"gpt-4-turbo-2024-04-09",
		"gpt-4-turbo-preview",
		"gpt-4-0125-preview",
		"gpt-4-1106-preview",
		"gpt-4-vision-preview",
		"gpt-4-1106-vision-preview",
		"gpt-4",
		"gpt-4-0613",
		"gpt-4-32k",
		"gpt-4-32k-0613",
		"gpt-3.5-turbo-0125",
		"gpt-3.5-turbo",
		"gpt-3.5-turbo-1106",
		"gpt-3.5-turbo-instruct",
		"gpt-3.5-turbo-16k",
		"gpt-3.5-turbo-0613",
		"gpt-3.5-turbo-16k-0613",
	}

	selectedModel := ""
	for _, current := range models {
		if model == current {
			selectedModel = current
		}
	}

	if selectedModel == "" {
		fmt.Println("Invalid model selected")
		return
	}
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.config/clai")
	viper.Set("model", selectedModel)

	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("Error writing config file:", err)
		return
	}

}
