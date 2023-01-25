package actions

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/config"
	"github.com/leeliwei930/notion_sdk/database/filter"
	"github.com/leeliwei930/notion_sdk/models"
)

type QueryDatabaseOptions func(opt *QueryDatabaseRequest)
type QueryDatabaseRequest struct {
	Filter *filter.QueryProps `json:"filter,omitempty"`
}

func FilterWith(filter *filter.QueryProps) QueryDatabaseOptions {
	return func(opt *QueryDatabaseRequest) {
		opt.Filter = filter
	}
}

func QueryDatabase(databaseId uuid.UUID, options ...QueryDatabaseOptions) (*filter.Cursor, error) {
	requestBody := &QueryDatabaseRequest{}

	for _, opt := range options {
		opt(requestBody)
	}
	cursorResults := &filter.Cursor{}
	endpointPath := fmt.Sprintf("/databases/%s/query", databaseId.String())
	response, err := client.Notion().
		SetAccessToken("secret_j8c5tvxQM854jS32JLGTXNi2TKUFu5QeDKzI7PodtnC").
		InitializeConfig(&config.NotionConfig{
			NotionVersion: "2022-06-28",
		}).
		Request().
		SetBody(requestBody).
		SetResult(cursorResults).
		Post(endpointPath)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		respErr := response.Error().(*models.NotionError)
		err = errors.New(respErr.Message)
		return nil, err
	}
	return cursorResults, nil

}
