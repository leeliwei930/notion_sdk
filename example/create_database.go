package main

import (
	"github.com/google/uuid"
	notionActions "github.com/leeliwei930/notion_sdk/actions"
	"github.com/leeliwei930/notion_sdk/database"
	"github.com/leeliwei930/notion_sdk/enums"
	notionModels "github.com/leeliwei930/notion_sdk/models"
)

func CreateDatabase(parentPageID uuid.UUID) (createdDatabase *database.Database, err error) {
	createdDatabase, err = notionActions.CreateDatabase(
		notionActions.SetDatabaseTitle(
			[]notionModels.RichText{
				{
					Type: enums.Text,
					Text: &notionModels.Text{
						Content: "Text",
						Link:    nil,
					},
				},
			},
		),
		notionActions.SetDatabasePageParent(&notionModels.PageParent{
			PageID: &parentPageID,
		}),
		notionActions.SetDatabaseProperties(map[string]*database.DatabaseProperty{
			"Title": {
				Title: &database.EmptyProperty{},
			},
			"Class": {
				Select: &database.SelectProperty{
					SelectOptions: []database.SelectOptions{
						{
							Name: "class A",
						},
						{
							Name: "class B",
						},
						{
							Name: "class C",
						},
					},
				},
			},
		}),
	)

	return
}
