package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/actions"
	"github.com/leeliwei930/notion_sdk/database/filter"
)

func ListDatabases() {

}

func QueryDatabase() (results *filter.Cursor, err error) {
	databaseID, _ := uuid.Parse("736d33a72c4141888e54d13f44d7245c")

	results, err = actions.QueryDatabase(databaseID, actions.FilterWith(
		&filter.QueryProps{
			Property: "Title",
			RichText: &filter.Text{
				Contains: "John",
			},
		},
	))

	if err != nil {
		fmt.Print(err.Error())
	}

	for _, record := range results.Results {
		props := record.Properties
		for _, title := range props["Title"].Title {
			fmt.Println(title.PlainText)
		}
	}
	return
}
