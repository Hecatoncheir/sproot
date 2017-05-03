package data

import (
	"github.com/google/uuid"
	"strings"
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
	var itemFromDatabase, err = &ItemOfCompany{}, error

	if item.ID != "" {
		itemFromDatabase, err = GetItemById(item.ID)
		if err != nil {
			return err
		}
	}

	item = itemFromDatabase

	return nil
}
