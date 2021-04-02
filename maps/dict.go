package maps

import "errors"

type Dict map[string]string

var ErrNotFound = errors.New("could not find the word you are looking for")

func (d Dict) Search(word string) (string, error) {
	value, exists := d[word]
	if !exists {
		return "", ErrNotFound
	}

	return value, nil
}
