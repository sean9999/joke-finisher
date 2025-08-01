package pkg

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
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
	return json.Marshal(joke)
}

// Hash hashes just the set-up, so "why did the chicken cross the road" can only exist once
func (joke *Joke) Hash() string {
	id := md5.Sum([]byte(joke.Setup))
	return fmt.Sprintf("%x", id[:])
}

func (joke *Joke) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, joke)
}

func GeneratePunchLine(setup string) (string, error) {
	panic("implement me")
}
