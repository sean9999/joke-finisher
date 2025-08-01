package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	jokeSetUp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	DoJoke(jokeSetUp)
}

func DoJoke(setup string) {
	ctx := context.Background()
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}
	prompt := fmt.Sprintf("Create a novel punchline for the following joke set-up: %s", setup)
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}
