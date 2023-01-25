package client

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/leeliwei930/notion_sdk/config"
	"github.com/leeliwei930/notion_sdk/models"
)

var restyClient *resty.Client

func setupClient() *resty.Client {
	notionError := &models.NotionError{}
	if restyClient != nil {
		return restyClient
	}
	restyClient = resty.
		New().
		SetBaseURL("https://api.notion.com/v1/").
		SetHeader("Content-Type", "application/json").
		SetError(notionError)

	return restyClient

}

type NotionClient struct {
	config       *config.NotionConfig
	restyRequest *resty.Request
	restyClient  *resty.Client
}

func (client *NotionClient) GetHttpBaseClient() *http.Client {
	return client.restyClient.GetClient()
}

func (client *NotionClient) SetAccessToken(accessToken string) *NotionClient {

	client.restyRequest.SetAuthToken(accessToken)
	return client
}

func (client *NotionClient) InitializeConfig(config *config.NotionConfig) *NotionClient {

	client.config = config
	client.restyRequest.SetHeader("Notion-Version", client.config.NotionVersion)
	return client
}

func (client *NotionClient) Request() *resty.Request {
	return client.restyRequest
}
func Notion() *NotionClient {
	restyClient := setupClient()
	client := &NotionClient{
		restyClient:  restyClient,
		restyRequest: restyClient.R(),
	}
	return client
}
