package site

import (
	"strings"
	"net/http"
	"errors"
	"encoding/json"
	"fmt"
	"url"
)

const directoryKeyword = 1
const adverbKeyword int = 0

type Keyword struct {
	Class int
	Id    int
	Name  string
}

func (k Keyword) FindAll(cls int) []Keyword {
	items := make([]Keyword, 10)
	_, err := http.Get(url.KeywordIndex())
	if err == nil {
		body := `[{"Class":110,"Id":123,"Name":"hiscaler"},{"Class":111,"Id":456,"Name":"shuzi","dd":111}]`
		fmt.Println(body)
		//body, _ = ioutil.ReadAll(resp.Body)
		if err := json.Unmarshal([]byte(body), &items); err == nil {
			fmt.Printf("%#v\r\n", items)
			for _, item := range items {
				fmt.Println(item.Name)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		errors.New("HTTP ERROR: " + err.Error())
	}

	return items
}

func (k *Keyword) Save() (*Keyword, error) {
	k.Name = strings.Trim(k.Name, " ")
	if len(k.Name) == 0 || (k.Class != directoryKeyword && k.Class != adverbKeyword) {
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
