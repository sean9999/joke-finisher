package pkg

import (
	"github.com/sean9999/harebrain"
)

// *fileStore implements the [Store] interface
var _ Store[string, *Joke] = (*fileStore)(nil)

type fileStore struct {
	harebrain.Table
}

func (f fileStore) Get(id string) (*Joke, error) {
	data, err := f.Table.Get(id)
	if err != nil {
		return nil, err
	}
	joke := new(Joke)
	err = joke.UnmarshalBinary(data)
	if err != nil {
		return nil, err
	}
	return joke, nil
}

func (f fileStore) Insert(joke *Joke) error {
	return f.Table.Insert(joke)
}

func (f fileStore) GetAll() map[string]*Joke {
	byteMap, err := f.Table.GetAll()
	if err != nil {
		panic(err)
	}
	jokeMap := make(map[string]*Joke, len(byteMap))
	for k, data := range byteMap {
		j := new(Joke)
		err := j.UnmarshalBinary(data)
		if err != nil {
			panic(err)
		}
		jokeMap[k] = j
	}
	return jokeMap
}
