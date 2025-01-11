package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

// main is the entry point of the program.
func main() {
	// Initialize the LLM with the "llama2" model.
	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	// Create a new reader to read input from the standard input (console).
	reader := bufio.NewReader(os.Stdin)
	// userInput stores the input provided by the user as a string.
	var userInput = ""
	for userInput == "" {
		// Prompt the user to enter a question.
		fmt.Print("Enter your question: ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Call the LLM with the user's input and get the completion.
		completion, err := llm.Call(ctx, "Human: "+userInput+"Assistant:",
			llms.WithTemperature(0.8),
			llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				// Print each chunk of the response as it is received.
				fmt.Print(string(chunk))
				return nil
			}),
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("")

		// Ask the user if they want to continue.
		if completion != "" {
			fmt.Print("Do you want to continue? (y/n): ")
			userInputExit, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			// Exit the loop if the user enters 'n'.
			if userInputExit[0] == 'n' {
				break
			}
		}

		// Reset userInput to prompt for another question.
		userInput = ""
	}
}
