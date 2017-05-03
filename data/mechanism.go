package data

import (
	"errors"
)

import rethink "gopkg.in/gorethink/gorethink.v3"

var ItemAlreadyExistError = errors.New("Item already exist")

func InsertItemToDatabase(item *ItemOfCompany) error {
	_, err := rethink.Table("Items").Insert(item).RunWrite(historyDataBaseSession)
	if err != nil {
		return ItemAlreadyExistError
	}

	return nil
}

func GetItemById(itemId string) (*ItemOfCompany, error) {
	cursor, err := rethink.Table("Items").Get(itemId).Run(historyDataBaseSession)
	defer cursor.Close()

	if err != nil {
		return nil, err
	}

	item := &ItemOfCompany{}
	err = cursor.One(item)

	if err == rethink.ErrEmptyResult {
		// row not found
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return item, nil
}

func GetItemByName(itemName string) (*ItemOfCompany, error) {
	companyItem := &ItemOfCompany{}

	return companyItem, nil
}
