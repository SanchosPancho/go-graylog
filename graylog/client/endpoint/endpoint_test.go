package endpoint_test

import (
	"testing"

	"github.com/SanchosPancho/go-graylog/v11/graylog/client/endpoint"
)

const (
	apiURL = "http://localhost:9000/api"
	ID     = "5a8c086fc006c600013ca6f5"
)

func TestNewEndpoints(t *testing.T) {
	if _, err := endpoint.NewEndpoints(""); err == nil {
		t.Fatal("invalid argument")
	}
	if _, err := endpoint.NewEndpoints("http://localhost:9000/api"); err != nil {
		t.Fatal(err)
	}
}
