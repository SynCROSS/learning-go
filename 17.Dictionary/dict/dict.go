package dict

import "errors"

// Dictionary is an alias of the type
type Dictionary map[string]string

var (
	errNotFound     = errors.New("Not Found that Word")
	errAlreadyExist = errors.New("That word Already Exists")
	errCanNOTUpdate = errors.New("Can NOT Update non-existing word")
)

<<<<<<< HEAD
// * Dictionary is a hashmap.
=======
// * Dictonary is a hashmap.
>>>>>>> Changing the environment to Linux
// * And by default a hashmap already has the * included.

// SearchWord returns based on the presence of the word.
func (d Dictionary) SearchWord(word string) (string, error) {

	value, existence := d[word]
	if existence {
		return value, nil
	}
	return "", errNotFound
}

// AddWord is for adding new word
func (d Dictionary) AddWord(word, definition string) error {
	_, e := d.SearchWord(word)
	if e == errNotFound {
		d[word] = definition
	} else if e == nil {
		return errAlreadyExist
	}
	return nil
}

// UpdateWord is for Updating existing word
func (d Dictionary) UpdateWord(word, definition string) error {
	_, e := d.SearchWord(word)
	if e == nil {
		d[word] = definition
	} else if e == errNotFound {
		return errCanNOTUpdate
	}
	return nil
}

// DeleteWord is for deleting word
func (d Dictionary) DeleteWord(word string) {
	// * The reason why I didn't use the "error" type of return type is
	// * because the "delete" function has no return
	// * and does nothing in the event of an error.
	delete(d, word)
}
