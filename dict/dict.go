package dict

import "errors"

// Dictionary is an alias of the type
type Dictionary map[string]string

// Search returns based on the presence of the word.
func (d Dictionary) Search(word string) (string, error) {
	var errNoValue = errors.New("Not Found that Word")

	value, existence := d[word]
	if existence {
		return value, nil
	}
	return "", errNoValue
}
