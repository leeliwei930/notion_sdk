package actions_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/leeliwei930/notion_sdk/actions"
	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/config"
	"github.com/leeliwei930/notion_sdk/database"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	client.InitializeNotionConfig(&config.NotionConfig{})

	exitCode := m.Run()

	os.Exit(exitCode)
}
func TestCreateDatabaseSuccess(t *testing.T) {

	httpmock.ActivateNonDefault(client.Notion().GetHttpBaseClient())
	mockCreatedDatabaseResponse := httpmock.File("tests/response_sample/create_database_response.json")
	jsonResponder, _ := httpmock.NewJsonResponder(200, mockCreatedDatabaseResponse)
	httpmock.RegisterResponder("POST", "https://api.notion.com/v1/databases", jsonResponder)
	defer httpmock.DeactivateAndReset()

	notionDatabase, _ := actions.CreateDatabase()
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
