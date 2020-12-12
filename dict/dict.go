package dict

import "errors"

// Dictionary is an alias of the type
type Dictionary map[string]string

var (
	errNotFound     = errors.New("Not Found that Word")
	errAlreadyExist = errors.New("That word Already Exists")
	errCanNOTUpdate = errors.New("Can NOT Update non-existing word")
)

// SearchWord returns based on the presence of the word.
func (d Dictionary) SearchWord(word string) (string, error) {

	value, existence := d[word]
	if existence {
		return value, nil
	}
	return "", errNotFound
}

// AddWord for adding new word
func (d Dictionary) AddWord(word, definition string) error {
	_, e := d.SearchWord(word)
	if e == errNotFound {
		d[word] = definition
	} else if e == nil {
		return errAlreadyExist
	}
	return nil
}

// UpdateWord for Updating existing word
func (d Dictionary) UpdateWord(word, definition string) error {
	_, e := d.SearchWord(word)
	if e == nil {
		d[word] = definition
	} else if e == errNotFound {
		return errCanNOTUpdate
	}
	return nil
}
