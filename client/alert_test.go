package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/SanchosPancho/go-graylog/client"
)

func TestClient_GetAlerts(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/alert/alerts.json")
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
								Path:         "/api/streams/alerts",
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

	_, _, _, err = cl.GetAlerts(ctx, 0, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetAlert(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/alert/alert.json")
	require.Nil(t, err)

	id := "5d84c1a92ab79c000d35d6c7"

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
								Path:         "/api/streams/alerts/" + id,
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

	if _, _, err := cl.GetAlert(ctx, ""); err == nil {
		t.Fatal("alert id is required")
	}

	if _, _, err := cl.GetAlert(ctx, id); err != nil {
		t.Fatal(err)
	}
}
