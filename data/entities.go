package data

import (
	"github.com/google/uuid"
	"strings"
	"errors"
)

func CreateNewItemOfCompany(name string) *ItemOfCompany {
	id := uuid.New().String()
	tags := strings.Split(name, " ")
	return &ItemOfCompany{ID: id, Tags: tags}
}

type ItemOfCompany struct {
	ID   string
	Tags []string
}

func (item *ItemOfCompany) Create() error {
	return InsertItemToDatabase(item)
}

func (item *ItemOfCompany) Read() error {
	var itemFromDatabase *ItemOfCompany
	var err error

	if item.ID != "" {
		itemFromDatabase, err = GetItemById(item.ID)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Item must have id")
	}

	item.Tags = itemFromDatabase.Tags

	return nil
}
