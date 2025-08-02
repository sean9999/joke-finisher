package pkg

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

var JokeAlreadyExists = errors.New("joke already exists")

// A Joke is a unique combination of set-up and punchline.
type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func (joke *Joke) String() string {
	return fmt.Sprintf("%s\n%s\n", joke.Setup, joke.Punchline)
}

func (joke *Joke) MarshalBinary() ([]byte, error) {
	return json.MarshalIndent(joke, "", "\t")
}

// Hash hashes just the set-up, so "why did the chicken cross the road" can only exist once
func (joke *Joke) Hash() string {
	id := md5.Sum([]byte(joke.Setup))
	return fmt.Sprintf("%x.json", id[:])
}

func (joke *Joke) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, joke)
}

const systemPrompt = "create a novel punchline for the joke deliminated by three backticks"

func GeneratePunchLine(ctx context.Context, s string, llm *openai.LLM) (string, error) {
	prompt := fmt.Sprintf("%s.\n```%s```\n", systemPrompt, s)
	return llms.GenerateFromSinglePrompt(ctx, llm, prompt)
}
