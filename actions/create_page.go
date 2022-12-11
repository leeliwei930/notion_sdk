package actions

import (
	"errors"

	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/config"
	"github.com/leeliwei930/notion_sdk/models"
)

type CreatePageOptions func(body *CreatePageOptionsBody)
type CreatePageOptionsBody struct {
	Parent     *models.PageParent          `json:"parent"`
	Properties map[string]*models.Property `json:"properties"`
	Children   *[]*models.Block            `json:"children,omitempty"`
	Icon       *models.Icon                `json:"icon,omitempty"`
	Cover      *models.File                `json:"cover,omitempty"`
}

func CreatePage(options ...CreatePageOptions) (notionPage *models.Page, err error) {
	requestBody := &CreatePageOptionsBody{}
	notionPage = &models.Page{}

	for _, opt := range options {
		opt(requestBody)
	}
	response, err := client.Notion().
		SetAccessToken("secret_j8c5tvxQM854jS32JLGTXNi2TKUFu5QeDKzI7PodtnC").
		InitializeConfig(&config.NotionConfig{
			NotionVersion: "2022-06-28",
		}).
		Request().
		SetBody(requestBody).
		SetResult(notionPage).
		Post("/pages")

	if response.IsError() {
		err = errors.New(response.Error().(*models.NotionError).Message)
	}
	return notionPage, err

}

func SetPageParent(parent *models.PageParent) CreatePageOptions {
	return func(option *CreatePageOptionsBody) {
		option.Parent = parent
	}
}

func SetPageProperties(properties map[string]*models.Property) CreatePageOptions {
	return func(option *CreatePageOptionsBody) {
		option.Properties = properties
	}
}

func SetPageChildren(children *[]*models.Block) CreatePageOptions {
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
