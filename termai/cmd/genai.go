package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var invokeGemini = &cobra.Command{
	Use:   "gemini",
	Short: "Invoke Gemini",
	Long:  "Ask a question to gemini",
	Run: func(cmd *cobra.Command, args []string) {
		question := strings.Join(args, " ")
		api_key := viper.GetString("TERMAI_GOOGLE_KEY")
		ctx := context.Background()
		client, err := genai.NewClient(ctx, option.WithAPIKey(api_key))
		if err != nil {
			println("Error creating client:", err.Error())
			return // Exit the function if client creation fails
		}
		defer client.Close() // Safe to defer now

		model := client.GenerativeModel("gemini-1.5-flash")
		resp_stream := model.GenerateContentStream(ctx, genai.Text(question))
		for {
			resp, err := resp_stream.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			for _, resp := range resp.Candidates {
				if resp.Content != nil {
					for _, part := range resp.Content.Parts {
						fmt.Println(part)
					}
				}
			}
		}
	},
}

var invokeOai = &cobra.Command{
	Use:   "oai",
	Short: "Invoke OpenAI Models",
	Long:  "Ask a question to OpenAI Models",
	Run: func(cmd *cobra.Command, args []string) {
		question := strings.Join(args, " ")
		api_key := viper.GetString("TERMAI_OAI_KEY")
		url := "https://api.openai.com/v1/chat/completions"
		reqBody := map[string]interface{}{
			"model": "gpt-4o",
			"messages": []map[string]string{
				{"role": "user", "content": question},
			},
		}
		jsonBody, err := json.Marshal(reqBody)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+api_key)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		var responseBody map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		if err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}

		choices, ok := responseBody["choices"].([]interface{})
		if ok && len(choices) > 0 {
			choice, ok := choices[0].(map[string]interface{})
			if ok {
				message, ok := choice["message"].(map[string]interface{})
				if ok {
					content, ok := message["content"].(string)
					if ok {
						fmt.Println("AI Response:", content)
					}
				}
			}
		}
	},
}
