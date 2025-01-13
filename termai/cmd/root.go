package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*
rootCmd represents the base command when called without any subcommands
*/
var rootCmd = &cobra.Command{
	Use:   "termai",
	Short: "A Terminal CLI for interacting with GenAI Models",
	Long: `TermAi is a terminal cli/terminal based application for interacting with GenAI Models.
Currently the supported models are Gemini models and OpenAI models, they are also currently defaulted.
In the future there will be an option to choose the model that you want to use for response.`,
}

/*
setup represents the logic that needs to run when the setup command is called
*/
var setup = &cobra.Command{
	Use:   "setup",
	Short: "Setup Termai",
	Long: `Setup Termai. Choose the model you want to configure and provide the corresponding api key
You won't be able to use Termai without this step
All api keys are saved in the config file locally on device.`,
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

// initConfig reads in config file and environment variables if set.
// If the config file doesn't exist, it creates a new one with the default config values.
// If the config file exists but there's an error reading it, it returns an error.
// If the config file exists and there's no error reading it, it logs a message indicating that it's using the config file.
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

	fmt.Fprintln(os.Stderr, "Get API Key From Config File")
	return nil
}

// createConfigFile creates the config file if it doesn't exist, and sets the initial config
// values.
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

// Execute runs the root command and starts the application.
//
// It initializes the config with initConfig, and then runs the root command
// with Execute. If there are any errors, it will print an error message and
// exit with a code of 1.
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

// init registers the subcommands with the root command.
func init() {
	rootCmd.AddCommand(setup)
	rootCmd.AddCommand(invokeGemini)
	rootCmd.AddCommand(invokeOai)
}
