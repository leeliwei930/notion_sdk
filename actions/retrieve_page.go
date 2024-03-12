package actions

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/models"
)

func RetrievePage(pageId uuid.UUID, filterProperties ...string) (*models.Page, error) {

	notionPage := &models.Page{}

	requestBody := make(map[string]interface{})

	if len(filterProperties) > 0 {
		requestBody["filter_properties"] = filterProperties
	}
	response, err := client.
		Notion().
		Request().
		SetBody(requestBody).
		SetResult(notionPage).
		Get(fmt.Sprintf("/pages/%s", pageId.String()))

	if err != nil {
		return nil, err
	}
	if response.IsError() {
		respErr, ok := response.Error().(*models.NotionError)
		if !ok {
			return nil, InvalidErrorResponse
		}
		err = errors.New(respErr.Message)
		return nil, err
	}
	return notionPage, nil
}
