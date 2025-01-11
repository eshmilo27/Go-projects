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

func main() {
	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	reader := bufio.NewReader(os.Stdin)
	// userInput stores the input provided by the user as a string.
	var userInput = ""
	for userInput == "" {
		fmt.Print("Enter your question: ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		completion, err := llm.Call(ctx, "Human: "+userInput+"Assistant:",
			llms.WithTemperature(0.8),
			llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				fmt.Print(string(chunk))
				return nil
			}),
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("")

		if completion != "" {
			fmt.Print("Do you want to continue? (y/n): ")
			userInputExit, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			if userInputExit[0] == 'n' {
				break
			}
		}

		userInput = ""
	}
}
