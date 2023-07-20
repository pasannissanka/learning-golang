/*
Copyright © 2023 Pasan Nissanka pasannissanka@outlook.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	hnapi "github.com/pasannissanka/learning-golang/go-cli-hackernews/hn-api"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cli-hackernews",
	Short: "A brief description of your application",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		info := color.New(color.BgGreen)
		error := color.New(color.BgRed)

		info.Println("Fetching Latest HackerNews Posts...")

		postIds, err := hnapi.GetItems()
		if err != nil {
			error.Println(err)
			os.Exit(1)
		}

		prompt := promptui.Prompt{
			Label:     "Load more posts",
			IsConfirm: true,
			Default:   "y",
		}

		stories, err := hnapi.FetchStories(postIds[:5])
		if err != nil {
			error.Println(err)
			os.Exit(1)
		}
		fmt.Print(stories)

		for i := 1; i < len(postIds)/5; i++ {

			result, err := prompt.Run()
			if err != nil {
				info.Println("Exiting...")
				os.Exit(1)
			}
			if result == "n" {
				info.Println("Exiting...")
				os.Exit(1)
			}

			stories, err := hnapi.FetchStories(postIds[(i+1)*5 : (i+2)*5])
			if err != nil {
				error.Println(err)
				os.Exit(1)
			}

			fmt.Print(stories)
		}
		info.Println("Exiting...")
		os.Exit(1)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cli-hackernews.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
