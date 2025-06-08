package privateapi

import (
	"github.com/yusologia/go-core/v2/api"
	"github.com/yusologia/go-core/v2/response"
	"net/url"
	"service/internal/pkg/config"
)

type TestingAPI interface {
	Get() logiares.ResponseSuccessWithPagination                      // Response bebas, sesuaikan dengan kebutuhan
	Store(payload interface{}) logiares.ResponseSuccessWithPagination // Response bebas, sesuaikan dengan kebutuhan
	RollBack(data interface{})                                        // Response bebas, sesuaikan dengan kebutuhan
}

func NewTestingAPI() TestingAPI {
	conf := config.PrivateAPIClient["testing"]

	client := logiaapi.NewLogiaAPI(logiaapi.LogiaAPIOption{
		Headers: map[string]string{
			"Client-ID":     conf["client-id"].(string),
			"Client-Name":   conf["client-name"].(string),
			"Client-Secret": conf["client-secret"].(string),
		},
	})

	api := testingAPI{
		baseURL: conf["host"].(string),
		client:  client,
	}

	return &api
}

type testingAPI struct {
	baseURL string
	client  logiaapi.LogiaAPI
}

func (api *testingAPI) Get() logiares.ResponseSuccessWithPagination {
	return api.client.Get(api.baseURL+"/testings", url.Values{})
}

func (api *testingAPI) Store(payload interface{}) logiares.ResponseSuccessWithPagination {
	return api.client.Post(api.baseURL+"/testings", payload)
}

func (api *testingAPI) RollBack(data interface{}) {
	api.client.Post(api.baseURL+"/testings/roll-back", data)
}
