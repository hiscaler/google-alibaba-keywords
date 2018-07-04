package site

import "strings"

const directoryKeyword = 1
const adverbKeyword int = 0

type Keyword struct {
	Class int
	Id    int
	Name  string
}

func (k *Keyword) save() (*Keyword, error) {
	k.Name = strings.Trim(k.Name, " ")
	if len(k.Name) == 0 || k.Class != directoryKeyword || k.Class != adverbKeyword {
		return k, nil
	}
	if k.Id == 0 {
		// Insert
		k.Id = 1
	} else {
		// Update
	}

	return k, nil
}
