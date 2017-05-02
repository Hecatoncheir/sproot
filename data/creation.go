package data

import "log"
import rethink "gopkg.in/gorethink/gorethink.v3"

var historyDataBaseSession *rethink.Session

func init() {
	var err error

	historyDataBaseSession, err = rethink.Connect(rethink.ConnectOpts{
		Address: "0.0.0.0:28015",
	})

	err = checkDataBase("History")
	if err != nil {
		log.Fatalln(err.Error())
	}

	historyDataBaseSession, err = rethink.Connect(rethink.ConnectOpts{
		Address:  "0.0.0.0:28015",
		Database: "History",
	})

	err = checkTable("Items")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func checkDataBase(databaseNameForCheck string) error {
	dbList := make([]string, 128)

	dbListCursor, err := rethink.DBList().Run(historyDataBaseSession)
	defer dbListCursor.Close()

	err = dbListCursor.All(&dbList)
	if err != nil {
		return err
	}

	dataBaseNotExist := true
	for _, dbName := range dbList {
		if dbName == databaseNameForCheck {
			dataBaseNotExist = false
		}
	}

	if dataBaseNotExist == true {
		_, err = rethink.DBCreate(databaseNameForCheck).Run(historyDataBaseSession)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkTable(tableNameForCheck string) error {
	var err error

	tableListCursor, err := rethink.TableList().Run(historyDataBaseSession)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer tableListCursor.Close()

	tableList := make([]string, 128)
	err = tableListCursor.All(&tableList)

	historyTableNotExist := true
	for _, tableName := range tableList {
		if tableName == tableNameForCheck {
			historyTableNotExist = false
		}
	}

	if historyTableNotExist == true {
		_, err = rethink.TableCreate(tableNameForCheck, rethink.TableCreateOpts{PrimaryKey: "ID"}).Run(historyDataBaseSession)
		if err != nil {
			return err
		}
	}

	return nil
}
