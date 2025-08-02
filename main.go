package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/sean9999/harebrain"
	"github.com/sean9999/joke-finisher/pkg"
	"github.com/tmc/langchaingo/llms/openai"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	jokeSetUp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	DoJoke(strings.TrimSpace(jokeSetUp))
}

func CreateRepertoire() (*pkg.Repertoire, error) {
	db := harebrain.NewDatabase()
	err := db.Open("jokes")
	if err != nil {
		log.Fatal(err)
	}
	rep := pkg.NewRepertoire("jokes")
	err = rep.Load()
	if err != nil {
		return nil, err
	}
	return rep, nil
}

func DoJoke(setup string) {

	rep, err := CreateRepertoire()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}

	joke, err := rep.Create(ctx, strings.TrimSpace(setup), llm)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(joke.String())

}
