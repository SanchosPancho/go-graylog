package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/SanchosPancho/go-graylog/v11/graylog/client"
)

func TestClient_GetAlertConditions(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/alert_condition/alert_conditions.json")
	require.Nil(t, err)

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method:       "GET",
								Path:         "/api/alerts/conditions",
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: string(buf),
							},
						},
					},
				},
			},
		},
	})

	_, _, _, err = cl.GetAlertConditions(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
