package dictionary

// Notes:
// Map is reference type, do not need to pass by pointer
// so Add function works with (d Dictionary) instead of (d *Dictionary)

import "errors"

// Dictionary is just a map with some methods
type Dictionary map[string]string

// ErrNotFound is returned when searching for a non-existant word
var ErrNotFound = errors.New("could not find the word you were looking for")

// ErrWordExists is returned if attempting to add an existing word
var ErrWordExists = errors.New("cannot add existing word to dictionary")

// ErrWordDoesNotExist is returned if attempting to update a non-existant word
var ErrWordDoesNotExist = errors.New("cannot update non-existant word")

// Search searches for a word in a map
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add adds a word to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update updates the definition of an existing word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}
