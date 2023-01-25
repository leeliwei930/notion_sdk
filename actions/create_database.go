package actions

import (
	"errors"

	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/config"
	"github.com/leeliwei930/notion_sdk/database"
	"github.com/leeliwei930/notion_sdk/models"
)

type CreateDatabaseOptions func(body *CreateDatabaseBody)
type CreateDatabaseBody struct {
	Parent     *models.PageParent                    `json:"parent,omitempty"`
	Icon       *models.Icon                          `json:"icon,omitempty"`
	Title      []*models.RichText                    `json:"title,omitempty"`
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
		SetAccessToken("secret_j8c5tvxQM854jS32JLGTXNi2TKUFu5QeDKzI7PodtnC").
		InitializeConfig(&config.NotionConfig{
			NotionVersion: "2022-06-28",
		}).
		Request().
		SetBody(requestBody).
		SetResult(notionDatabase).
		Post("/databases")
	if err != nil {
		return nil, err
	}

	if response.IsError() {
		respErr := response.Error().(*models.NotionError)
		err = errors.New(respErr.Message)
		return nil, err
	}

	return notionDatabase, nil

}

func SetDatabaseTitle(title []*models.RichText) CreateDatabaseOptions {
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
