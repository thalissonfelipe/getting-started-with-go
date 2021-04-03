package maps

const (
	ErrNotFound          = ErrDict("could not find the word you are looking for")
	ErrWordAlreadyExists = ErrDict("cannot add the word because it already exists")
	ErrWordDoesNotExist  = ErrDict("the word could not be updated because it does not exist")
)

type ErrDict string

func (e ErrDict) Error() string {
	return string(e)
}

type Dict map[string]string

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

func (d Dict) Update(word, newValue string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newValue
	default:
		return err
	}

	return nil
}

func (d Dict) Delete(word string) {
	delete(d, word)
}
