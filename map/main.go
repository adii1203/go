package main

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrNotFound         = errors.New("could not find the word you were looking for")
	ErrWordExist        = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search a word in dictionary.
func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

// Add a word in dictionary.
func (d Dictionary) Add(word, value string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = value
	case nil:
		return ErrWordExist
	default:
		return err
	}

	return nil
}

// Update a word in dictionary.
func (d Dictionary) Update(word, newValue string) error {
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

// Delete a word from dictionary.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func main() {
	fmt.Println("Hello World")
}
