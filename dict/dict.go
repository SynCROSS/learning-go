package dict

import "errors"

// Dictionary is an alias of the type
type Dictionary map[string]string

var errNoValue = errors.New("Not Found that Word")
var errAlreadyExist = errors.New("That word Already Exists")

// Search returns based on the presence of the word.
func (d Dictionary) Search(word string) (string, error) {

	value, existence := d[word]
	if existence {
		return value, nil
	}
	return "", errNoValue
}

func (d Dictionary) AddWord(word string, definition string) error {
	_, e := d.Search(word)
	if e == errNoValue {
		d[word] = definition
	} else if e == nil {
		return errAlreadyExist
	}
	return nil
}
