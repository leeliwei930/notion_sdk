package actions

import (
	"errors"

	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/models"
)

type CreatePageOptions func(body *CreatePageOptionsBody)
type CreatePageOptionsBody struct {
	Parent     *models.NotionPageParent          `json:"parent"`
	Properties map[string]*models.NotionProperty `json:"properties"`
	Children   *[]*models.NotionBlock            `json:"children,omitempty"`
	Icon       *models.NotionIcon                `json:"icon,omitempty"`
	Cover      *models.NotionFile                `json:"cover,omitempty"`
}

func CreatePage(options ...CreatePageOptions) (notionPage *models.NotionPage, err error) {
	requestBody := &CreatePageOptionsBody{}
	notionPage = &models.NotionPage{}

	for _, opt := range options {
		opt(requestBody)
	}
	response, err := client.Notion().
		SetBody(requestBody).
		SetResult(notionPage).
		Post("/pages")

	if response.IsError() {
		err = errors.New(response.Error().(*models.NotionError).Message)
	}
	return notionPage, err

}

func SetPageParent(parent *models.NotionPageParent) CreatePageOptions {
	return func(option *CreatePageOptionsBody) {
		option.Parent = parent
	}
}

func SetPageProperties(properties map[string]*models.NotionProperty) CreatePageOptions {
	return func(option *CreatePageOptionsBody) {
		option.Properties = properties
	}
}

func SetPageChildren(children *[]*models.NotionBlock) CreatePageOptions {
	return func(option *CreatePageOptionsBody) {
		option.Children = children
	}
}

func SetPageIcon(icon *models.NotionIcon) CreatePageOptions {
	return func(option *CreatePageOptionsBody) {
		option.Icon = icon
	}
}

func SetPageCover(cover *models.NotionFile) CreatePageOptions {
	return func(option *CreatePageOptionsBody) {
		option.Cover = cover
	}
}
