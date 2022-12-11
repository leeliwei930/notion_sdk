package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/leeliwei930/notion_sdk/config"
	"github.com/leeliwei930/notion_sdk/models"
)

func newNotionClient(cfg *config.NotionConfig) *resty.Client {
	notionError := &models.NotionError{}
	client := resty.
		New().
		SetBaseURL("https://api.notion.com/v1/").
		SetAuthToken(cfg.AccessToken).
		SetHeader("Notion-Version", "2021-08-16").
		SetHeader("Content-Type", "application/json").
		SetError(notionError)

	return client

}

func Notion() *resty.Request {
	return newNotionClient(
		&config.NotionConfig{
			AccessToken: "secret_j8c5tvxQM854jS32JLGTXNi2TKUFu5QeDKzI7PodtnC",
		}).R()
}
