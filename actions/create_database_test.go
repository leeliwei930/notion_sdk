package actions_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"
	"github.com/leeliwei930/notion_sdk/actions"
	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/config"
	"github.com/leeliwei930/notion_sdk/database"
	"github.com/leeliwei930/notion_sdk/enums"
	"github.com/leeliwei930/notion_sdk/models"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	client.InitializeNotionConfig(&config.NotionConfig{})

	exitCode := m.Run()

	os.Exit(exitCode)
}
func TestCreateDatabaseSuccess(t *testing.T) {
	var parentPageID uuid.UUID
	parentPageID, _ = uuid.NewRandom()
	httpmock.ActivateNonDefault(client.Notion().GetHttpBaseClient())
	mockCreatedDatabaseResponse := httpmock.File("tests/response_sample/create_database_response.json")

	httpmock.RegisterResponder("POST", "https://api.notion.com/v1/databases", func(req *http.Request) (*http.Response, error) {
		var parsedReqBody actions.CreateDatabaseBody
		json.NewDecoder(req.Body).Decode(&parsedReqBody)
		assert.Equal(t, actions.CreateDatabaseBody{
			Cover: &models.File{
				External: &models.ExternalFile{
					Url: "some file",
				},
			},
			Icon: &models.Icon{
				Type:  "emoji",
				Emoji: "🤖",
			},
			Title: []models.RichText{
				{
					Type: enums.Text,
					Text: &models.Text{
						Content: "Text",
						Link:    nil,
					},
				},
			},
			Parent: &models.PageParent{
				PageID: &parentPageID,
			},
			Properties: map[string]*database.DatabaseProperty{
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
			},
		}, parsedReqBody)
		return httpmock.NewJsonResponse(200, mockCreatedDatabaseResponse)
	})
	defer httpmock.DeactivateAndReset()

	notionDatabase, _ := actions.CreateDatabase(
		actions.SetDatabaseTitle(
			[]models.RichText{
				{
					Type: enums.Text,
					Text: &models.Text{
						Content: "Text",
						Link:    nil,
					},
				},
			},
		), actions.SetDatabasePageCover(&models.File{
			External: &models.ExternalFile{
				Url: "some file",
			},
		}),
		actions.SetDatabasePageParent(&models.PageParent{
			PageID: &parentPageID,
		}),
		actions.SetDatabasePageIcon(&models.Icon{
			Type:  "emoji",
			Emoji: "🤖",
		}),
		actions.SetDatabaseProperties(map[string]*database.DatabaseProperty{
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
		}))
	var mockCreatedDatabase database.Database

	unmarshalErr := json.Unmarshal(mockCreatedDatabaseResponse.Bytes(), &mockCreatedDatabase)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
	}
	assert.Equal(t, mockCreatedDatabase, *notionDatabase)

}

func TestCreateDatabaseFail(t *testing.T) {

	httpmock.ActivateNonDefault(client.Notion().GetHttpBaseClient())
	mockErrorResponse := httpmock.File("tests/response_sample/error_response.json")
	jsonResponder, _ := httpmock.NewJsonResponder(400, mockErrorResponse)
	httpmock.RegisterResponder("POST", "https://api.notion.com/v1/databases", jsonResponder)
	defer httpmock.DeactivateAndReset()

	notionDatabase, err := actions.CreateDatabase()
	assert.Nil(t, notionDatabase)

	assert.EqualError(t, err, errors.New("Unavailable").Error())
}
