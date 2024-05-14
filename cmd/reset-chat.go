/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// resetChatCmd represents the resetChat command
var resetChatCmd = &cobra.Command{
	Use:   "reset-chat",
	Short: "Reset the csv of chat history",
	Long:  `Reset the csv of chat history`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			// Handle error
			panic(err)
		}

		CSV_PATH := homeDir + "/.config/clai/chatData/chat.csv"

		// Check if the file exists
		if _, err := os.Stat(CSV_PATH); err == nil {
			// File exists, delete it
			if err := os.Remove(CSV_PATH); err != nil {
				fmt.Println("Error deleting file:", err)
				return
			}
		}

		// Create a new file
		newFile, err := os.Create(CSV_PATH)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer newFile.Close()

		fmt.Println("Chat history reset successfully.")
	},
}

func init() {
	rootCmd.AddCommand(resetChatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetChatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetChatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
