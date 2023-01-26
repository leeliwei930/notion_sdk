package main

import (
	"github.com/google/uuid"
	notionClient "github.com/leeliwei930/notion_sdk/actions"
	notionEnums "github.com/leeliwei930/notion_sdk/enums"
	notionModels "github.com/leeliwei930/notion_sdk/models"
)

func CreatePage(parentPageID uuid.UUID) (createdPage *notionModels.Page, err error) {
	createdPage, err = notionClient.CreatePage(
		notionClient.SetPageParent(
			&notionModels.PageParent{
				PageID: &parentPageID,
			},
		),
		notionClient.SetPageProperties(
			map[string]notionModels.Property{
				"title": {
					Title: []notionModels.RichText{
						{
							Text: &notionModels.Text{
								Content: "Hello from Go",
							},
						},
					},
				},
			},
		),
		notionClient.SetPageChildren([]notionModels.Block{
			{
				Type: notionEnums.H1BlockType,
				Heading1: &notionModels.Heading1Block{
					RichText: []notionModels.RichText{
						{
							Text: &notionModels.Text{
								Content: "This is the content",
							},
						},
					},
				},
			},
			{
				Type: notionEnums.CodeBlockType,
				Code: &notionModels.CodeBlock{
					Language: "go",
					RichText: []notionModels.RichText{
						{
							Text: &notionModels.Text{
								Content: `package main
import "fmt"
func main(){
	fmt.Println("Hello Go")
}
`,
							},
						},
					},
				},
			},
		}),
		notionClient.SetPageIcon(
			&notionModels.Icon{
				Type:  "emoji",
				Emoji: "🤖",
			},
		),
	)
	return

}
