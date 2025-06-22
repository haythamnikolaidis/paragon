package ollama

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

// GetDefinition queries the Ollama API to get the definition, an example of use, and the connotation of a word.
func GetDefinition(word string) (string, error) {
	// Create the Ollama API client from the environment.
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return "", fmt.Errorf("failed to create Ollama client: %w", err)
	}

	// Construct the messages for the chat request.
	messages := []api.Message{
		{
			Role:    "system",
			Content: "You are a helpful assistant that provides definitions, examples, and connotations of words.",
		},
		{
			Role:    "user",
			Content: fmt.Sprintf("Define the word '%s'. Provide a definition, an example of its use, and its connotation.", word),
		},
	}

	// Create the chat request.
	req := &api.ChatRequest{
		Model:    "llama3.2", // Replace with the appropriate model name.
		Messages: messages,
	}

	// Use a context for the API call.
	ctx := context.Background()

	// Variable to store the response content.
	var result string

	// Define the response handler function.
	respFunc := func(resp api.ChatResponse) error {
		// Append the response content to the result.
		result += resp.Message.Content
		return nil
	}

	// Make the API call.
	err = client.Chat(ctx, req, respFunc)
	if err != nil {
		return "", fmt.Errorf("failed to query Ollama API: %w", err)
	}

	// Return the result.
	return result, nil
}
