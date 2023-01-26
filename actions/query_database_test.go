package actions_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"
	"github.com/leeliwei930/notion_sdk/actions"
	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/database/filter"
	"github.com/stretchr/testify/assert"
)

func TestQueryDatabaseWithFilter(t *testing.T) {
	databaseUUID, _ := uuid.NewRandom()
	httpmock.ActivateNonDefault(client.Notion().GetHttpBaseClient())
	defer httpmock.DeactivateAndReset()
	mockQueryDatabaseResponse := httpmock.File("tests/response_sample/query_database_response.json")
	httpmock.RegisterResponder("POST", fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", databaseUUID.String()), func(req *http.Request) (*http.Response, error) {
		var parsedReqBody actions.QueryDatabaseRequest
		json.NewDecoder(req.Body).Decode(&parsedReqBody)
		assert.Equal(t, actions.QueryDatabaseRequest{
			Filter: &filter.QueryProps{
				Property: "Title",
				RichText: &filter.Text{
					Equals: "Title A",
				},
				And: []filter.QueryProps{
					{
						Property: "Page",
						Number: &filter.Number{
							Gt: 10,
						},
					},
				},
			},
		}, parsedReqBody)
		return httpmock.NewJsonResponse(200, mockQueryDatabaseResponse)

	})

	queryResult, _ := actions.QueryDatabase(databaseUUID, []actions.QueryDatabaseOptions{
		actions.FilterWith(
			&filter.QueryProps{
				Property: "Title",
				RichText: &filter.Text{
					Equals: "Title A",
				},
				And: []filter.QueryProps{
					{
						Property: "Page",
						Number: &filter.Number{
							Gt: 10,
						},
					},
				},
			},
		),
	}...)

	var mockQueryResultCursor filter.Cursor

	unmarshalErr := json.Unmarshal(mockQueryDatabaseResponse.Bytes(), &mockQueryResultCursor)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
	}
	assert.Equal(t, mockQueryResultCursor, *queryResult)
}
