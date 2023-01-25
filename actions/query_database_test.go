package actions_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"
	"github.com/leeliwei930/notion_sdk/client"
)

func TestQueryDatabase(t *testing.T) {
	var parentPageID uuid.UUID
	parentPageID, _ = uuid.NewRandom()
	httpmock.ActivateNonDefault(client.Notion().GetHttpBaseClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", parentPageID.String()), func(r *http.Request) (*http.Response, error) {

	})

}
