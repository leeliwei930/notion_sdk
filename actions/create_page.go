package actions

import (
	"errors"

	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/models"
)

type CreatePageOptions func(body *CreatePageOptionsBody)
type CreatePageOptionsBody struct {
	Children   []models.Block             `json:"children,omitempty"`
	Parent     *models.PageParent         `json:"parent"`
	Properties map[string]models.Property `json:"properties"`
	Icon       *models.Icon               `json:"icon,omitempty"`
	Cover      *models.File               `json:"cover,omitempty"`
}

func CreatePage(options ...CreatePageOptions) (*models.Page, error) {
	requestBody := &CreatePageOptionsBody{}

	for _, opt := range options {
		opt(requestBody)
	}
	notionPage := &models.Page{}

	response, err := client.Notion().
		Request().
		SetBody(requestBody).
		SetResult(notionPage).
		Post("/pages")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		respErr := response.Error().(*models.NotionError)

		err = errors.New(respErr.Message)
		return nil, err
	}
	return notionPage, nil
}

func SetPageParent(parent *models.PageParent) CreatePageOptions {
	return func(option *CreatePageOptionsBody) {
		option.Parent = parent
	}
}

func SetPageProperties(properties map[string]models.Property) CreatePageOptions {
	return func(option *CreatePageOptionsBody) {
		option.Properties = properties
	}
}

func SetPageChildren(children []models.Block) CreatePageOptions {
	return func(option *CreatePageOptionsBody) {
		option.Children = children
	}
}

func SetPageIcon(icon *models.Icon) CreatePageOptions {
	return func(option *CreatePageOptionsBody) {
		option.Icon = icon
	}
}

func SetPageCover(cover *models.File) CreatePageOptions {
	return func(option *CreatePageOptionsBody) {
		option.Cover = cover
	}
}
