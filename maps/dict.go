package maps

import "errors"

type Dict map[string]string

func (d Dict) Search(word string) (string, error) {
	value, exists := d[word]
	if !exists {
		return "", errors.New("could not find the word you are looking for")
	}

	return value, nil
}
