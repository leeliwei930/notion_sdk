package actions

import (
	"errors"

	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/database"
	"github.com/leeliwei930/notion_sdk/models"
)

type CreateDatabaseOptions func(body *CreateDatabaseBody)
type CreateDatabaseBody struct {
	Title      []models.RichText                     `json:"title,omitempty"`
	Parent     *models.PageParent                    `json:"parent,omitempty"`
	Icon       *models.Icon                          `json:"icon,omitempty"`
	Cover      *models.File                          `json:"cover,omitempty"`
	Properties map[string]*database.DatabaseProperty `json:"properties,omitempty"`
}

func CreateDatabase(options ...CreateDatabaseOptions) (*database.Database, error) {
	requestBody := &CreateDatabaseBody{}

	for _, opt := range options {
		opt(requestBody)
	}
	notionDatabase := &database.Database{}

	response, err := client.Notion().
		Request().
		SetBody(requestBody).
		SetResult(notionDatabase).
		Post("/databases")
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

	return notionDatabase, nil

}

func SetDatabaseTitle(title []models.RichText) CreateDatabaseOptions {
	return func(option *CreateDatabaseBody) {
		option.Title = title
	}
}

func SetDatabasePageParent(parent *models.PageParent) CreateDatabaseOptions {
	return func(option *CreateDatabaseBody) {
		option.Parent = parent
	}
}

func SetDatabasePageCover(cover *models.File) CreateDatabaseOptions {
	return func(option *CreateDatabaseBody) {
		option.Cover = cover
	}
}

func SetDatabasePageIcon(icon *models.Icon) CreateDatabaseOptions {
	return func(option *CreateDatabaseBody) {
		option.Icon = icon
	}
}

func SetDatabaseProperties(properties map[string]*database.DatabaseProperty) CreateDatabaseOptions {
	return func(option *CreateDatabaseBody) {
		option.Properties = properties
	}
}
