package pkg

import (
	"context"
	"github.com/tmc/langchaingo/llms/openai"
)

// A Repertoire is a set of jokes with the ability to persist them somewhere
type Repertoire struct {
	Jokes map[string]*Joke
	Store Store[string, *Joke]
}

func NewRepertoire(folder string) *Repertoire {
	return &Repertoire{
		Jokes: make(map[string]*Joke),
		Store: NewFileStore(folder),
	}
}

func (r *Repertoire) save(joke *Joke) error {
	err := r.Store.Insert(joke)
	if err != nil {
		return err
	}
	r.Jokes[joke.Hash()] = joke
	return nil
}

func (r *Repertoire) Create(ctx context.Context, setup string, llm *openai.LLM) (*Joke, error) {

	// TODO: is it necessary to ask the db here? why not just check the map?
	_, err := r.Store.Get(setup)
	if err == nil {
		return nil, JokeAlreadyExists
	}
	punchline, err := GeneratePunchLine(ctx, setup, llm)
	if err != nil {
		return nil, err
	}
	joke := new(Joke)
	joke.Setup = setup
	joke.Punchline = punchline
	err = r.save(joke)
	if err != nil {
		return nil, err
	}
	return joke, nil
}

func (r *Repertoire) Delete(id string) error {
	err := r.Store.Delete(id)
	if err != nil {
		return err
	}
	delete(r.Jokes, id)
	return nil
}

// Load loads all jokes from the database into memory
func (r *Repertoire) Load() error {
	for id, record := range r.Store.GetAll() {
		r.Jokes[id] = record
	}
	return nil
}
