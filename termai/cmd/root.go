package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "termai",
	Short: "A Terminal CLI for interacting with GenAI Models",
	Long: `TermAi is a terminal cli/terminal based application for interacting with GenAI Models.
Currently the supported models are Gemini models and OpenAI models, they are also currently defaulted.
In the future there will be an option to choose the model that you want to use for response.`,
}

var setup = &cobra.Command{
	Use:   "setup",
	Short: "Setup Termai",
	Long: `Setup Termai. Choose the model you want to configure and provide the corresponding api key.
You won't be able to use Termai without this set.
All api keys are saved in the config file`,
	Run: func(cmd *cobra.Command, args []string) {
		aiProviderSelector := promptui.Select{
			Label: "Select AI Provider",
			Items: []string{"OpenAI", "Google"},
		}

		_, result, err := aiProviderSelector.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "OpenAI":
			var oaiKey string
			fmt.Println("OpenAI selected")
			fmt.Println("Enter OpenAI API Key:")
			fmt.Scanln(&oaiKey)

			// Set the key in viper
			viper.Set("TERMAI_OAI_KEY", oaiKey)

			// Save the config
			if err := viper.WriteConfig(); err != nil {
				fmt.Printf("Error saving OpenAI key: %v\n", err)
				return
			}
			fmt.Println("OpenAI API Key set successfully")

		case "Google":
			var googleKey string
			fmt.Println("Google selected")
			fmt.Println("Enter Google API Key:")
			fmt.Scanln(&googleKey)

			// Set the key in viper
			viper.Set("TERMAI_GOOGLE_KEY", googleKey)

			// Save the config
			if err := viper.WriteConfig(); err != nil {
				fmt.Printf("Error saving Google key: %v\n", err)
				return
			}
			fmt.Println("Google API Key set successfully")

		default:
			fmt.Println("No provider selected")
		}
	},
}

func initConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}

	viper.SetConfigName(".termai")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Fprintln(os.Stderr, "Config file not found. Creating a new one...")
			if err := createConfigFile(); err != nil {
				return fmt.Errorf("error creating config file: %w", err)
			}
			// Try to read the config again after creating it
			if err := viper.ReadInConfig(); err != nil {
				return fmt.Errorf("error reading newly created config file: %w", err)
			}
		} else {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	return nil
}

func createConfigFile() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}

	configPath := filepath.Join(home, ".termai.yaml")

	if err := os.MkdirAll(filepath.Dir(configPath), os.ModePerm); err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	// Create initial config content
	config := map[string]string{
		"TERMAI_GOOGLE_KEY": "",
		"TERMAI_OAI_KEY":    "",
	}

	// Set these values in viper
	for k, v := range config {
		viper.Set(k, v)
	}

	// Save the config file
	if err := viper.SafeWriteConfig(); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func Execute() {
	if err := initConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing config: %v\n", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(setup)
	rootCmd.AddCommand(invokeGemini)
	rootCmd.AddCommand(invokeOai)
}
