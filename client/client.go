package client

import (
	"errors"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/leeliwei930/notion_sdk/config"
	"github.com/leeliwei930/notion_sdk/models"
)

var restyClient *resty.Client
var notionClient *NotionClient
var notionConfig *config.NotionConfig

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

func loadNotionClient() *NotionClient {
	if notionConfig == nil {
		panic(errors.New("Notion config isn't initialize, do you call client.LoadNotionConfig before accessing to Notion client instance?"))
	}

	if notionClient == nil {
		notionClient = &NotionClient{
			config:       notionConfig,
			restyClient:  restyClient,
			restyRequest: restyClient.R(),
		}
		notionClient.restyRequest.SetHeader("Notion-Version", notionClient.config.NotionVersion)
		notionClient.restyRequest.SetAuthToken(notionClient.config.AccessToken)
	}
	return notionClient
}

func (client *NotionClient) Request() *resty.Request {
	return client.restyRequest
}

func InitializeNotionConfig(config *config.NotionConfig) {
	notionConfig = config
}

func Notion() *NotionClient {
	restyClient = setupClient()
	notionClient = loadNotionClient()
	return notionClient
}
