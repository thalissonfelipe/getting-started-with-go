package maps

import "errors"

type Dict map[string]string

var (
	ErrNotFound          = errors.New("could not find the word you are looking for")
	ErrWordAlreadyExists = errors.New("cannot add the word because it already exists")
)

func (d Dict) Search(word string) (string, error) {
	value, exists := d[word]
	if !exists {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dict) Add(word, value string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = value
	case nil:
		return ErrWordAlreadyExists
	default:
		return err
	}

	return nil
}
