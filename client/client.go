package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/leeliwei930/notion_sdk/config"
	"github.com/leeliwei930/notion_sdk/models"
)

func setupRequest() *resty.Request {
	notionError := &models.NotionError{}
	client := resty.
		New().
		SetBaseURL("https://api.notion.com/v1/").
		SetHeader("Content-Type", "application/json").
		SetError(notionError)

	return client.R()

}

type NotionClient struct {
	config  *config.NotionConfig
	request *resty.Request
}

func (client *NotionClient) SetAccessToken(accessToken string) *NotionClient {

	client.request.SetAuthToken(accessToken)
	return client
}

func (client *NotionClient) InitializeConfig(config *config.NotionConfig) *NotionClient {

	client.config = config
	client.request.SetHeader("Notion-Version", client.config.NotionVersion)

	return client
}

func (client *NotionClient) Request() *resty.Request {
	return client.request
}
func Notion() *NotionClient {
	req := setupRequest()
	client := &NotionClient{
		request: req,
	}
	return client
}
