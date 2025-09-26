package graylog_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/SanchosPancho/go-graylog/v11/graylog/graylog"
	"github.com/SanchosPancho/go-graylog/v11/graylog/testdata"
	"github.com/suzuki-shunsuke/go-set/v2"
)

func TestUserNewUpdateParams(t *testing.T) {
	user := testdata.User()
	prms := user.NewUpdateParams()
	if user.Username != prms.Username {
		t.Fatalf(`prms.Username = "%s", wanted "%s"`, prms.Username, user.Username)
	}
}

func TestUserSetDefaultValues(t *testing.T) {
	user := &graylog.User{}
	user.SetDefaultValues()
	if user.SessionTimeoutMs == 0 {
		t.Fatal("user.SessionTimeoutMs must be set")
	}
	if user.Timezone == "" {
		t.Fatal("user.Timezone must be set")
	}
}

func TestUserJSON(t *testing.T) {
	u := &graylog.User{
		Username:  "jdoe",
		Email:     "john.doe@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "Secret123!",
		Roles:     set.NewStrSet("Reader"),
		External:  false,
	}
	b, err := json.Marshal(u)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(b)
	if !containsAll(s, `"first_name":"John"`, `"last_name":"Doe"`) {
		t.Fatalf("missing first/last in json: %s", s)
	}
}

func containsAll(s string, subs ...string) bool {
	for _, sub := range subs {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}
