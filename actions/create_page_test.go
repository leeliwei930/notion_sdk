package actions

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/models"
	"github.com/stretchr/testify/assert"
)

func TestCreatePageSuccess(t *testing.T) {
	httpmock.ActivateNonDefault(client.Notion().GetHttpBaseClient())
	mockResponseFile := httpmock.File("create_page_response.json")
	jsonResponder, _ := httpmock.NewJsonResponder(200, mockResponseFile)

	httpmock.RegisterResponder("POST", "https://api.notion.com/v1/pages", jsonResponder)
	defer httpmock.DeactivateAndReset()

	notionPage, _ := CreatePage()
	var mockCreatedPage models.Page
	json.Unmarshal(mockResponseFile.Bytes(), &mockCreatedPage)

	assert.Equal(t, mockCreatedPage, *notionPage)

}

func TestCreatePageFailure(t *testing.T) {
	httpmock.ActivateNonDefault(client.Notion().GetHttpBaseClient())
	mockErrResponseFile := httpmock.File("error_response.json")

	jsonResponder, _ := httpmock.NewJsonResponder(400, mockErrResponseFile)

	httpmock.RegisterResponder("POST", "https://api.notion.com/v1/pages", jsonResponder)
	defer httpmock.DeactivateAndReset()

	notionPage, err := CreatePage()
	assert.Nil(t, notionPage)
	assert.EqualError(t, err, errors.New("Unavailable").Error())

}
