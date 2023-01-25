package actions_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"
	"github.com/leeliwei930/notion_sdk/actions"
	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/enums"
	"github.com/leeliwei930/notion_sdk/models"
	"github.com/stretchr/testify/assert"
)

func TestCreatePageSuccess(t *testing.T) {
	httpmock.ActivateNonDefault(client.Notion().GetHttpBaseClient())
	mockResponseFile := httpmock.File("tests/response_sample/create_page_response.json")
	var testUUID uuid.UUID
	testUUID, _ = uuid.NewRandom()

	pageOptions := []actions.CreatePageOptions{
		actions.SetPageParent(&models.PageParent{
			PageID: &testUUID,
		}),
		actions.SetPageProperties(
			map[string]models.Property{
				"title": {
					Title: []models.RichText{
						{
							Text: &models.Text{
								Content: "Hello from Go",
							},
						},
					},
				},
			},
		),
		actions.SetPageChildren([]models.Block{
			{
				Type: enums.H1BlockType,
				Heading1: &models.Heading1Block{
					RichText: []models.RichText{
						{
							Text: &models.Text{
								Content: "This is the content",
							},
						},
					},
				},
			},
		}),
		actions.SetPageIcon(
			&models.Icon{
				Type:  "emoji",
				Emoji: "🤖",
			},
		),
		actions.SetPageCover(&models.File{
			External: &models.ExternalFile{
				Url: "",
			},
		}),
	}
	httpmock.RegisterResponder("POST", "https://api.notion.com/v1/pages", func(req *http.Request) (*http.Response, error) {
		var parsedReqBody actions.CreatePageOptionsBody
		json.NewDecoder(req.Body).Decode(&parsedReqBody)
		assert.Equal(t, actions.CreatePageOptionsBody{
			Parent: &models.PageParent{
				PageID: &testUUID,
			},
			Properties: map[string]models.Property{
				"title": {
					Title: []models.RichText{
						{
							Text: &models.Text{
								Content: "Hello from Go",
							},
						},
					},
				},
			},
			Children: []models.Block{
				{
					Type: enums.H1BlockType,
					Heading1: &models.Heading1Block{
						RichText: []models.RichText{
							{
								Text: &models.Text{
									Content: "This is the content",
								},
							},
						},
					},
				},
			},
			Icon: &models.Icon{
				Type:  "emoji",
				Emoji: "🤖",
			},
			Cover: &models.File{
				External: &models.ExternalFile{
					Url: "",
				},
			},
		}, parsedReqBody)
		return httpmock.NewJsonResponse(200, mockResponseFile)
	})
	defer httpmock.DeactivateAndReset()

	notionPage, _ := actions.CreatePage(pageOptions...)
	var mockCreatedPage models.Page
	json.Unmarshal(mockResponseFile.Bytes(), &mockCreatedPage)

	assert.Equal(t, mockCreatedPage, *notionPage)

}

func TestCreatePageFailure(t *testing.T) {
	httpmock.ActivateNonDefault(client.Notion().GetHttpBaseClient())
	mockErrResponseFile := httpmock.File("tests/response_sample/error_response.json")

	jsonResponder, _ := httpmock.NewJsonResponder(400, mockErrResponseFile)

	httpmock.RegisterResponder("POST", "https://api.notion.com/v1/pages", jsonResponder)
	defer httpmock.DeactivateAndReset()

	notionPage, err := actions.CreatePage()
	assert.Nil(t, notionPage)
	assert.EqualError(t, err, errors.New("Unavailable").Error())

}
