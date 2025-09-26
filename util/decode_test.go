package util_test

import (
	"testing"

	"github.com/SanchosPancho/go-graylog/util"
	"github.com/github.com/SanchosPancho/go-graylog/graylog"
)

func TestMSDecode(t *testing.T) {
	data := map[string]interface{}{
		"bind_address": "0.0.0.0",
	}
	var attrs graylog.InputAttrs = &graylog.InputBeatsAttrs{}
	if err := util.MSDecode(data, attrs); err != nil {
		t.Fatal(err)
	}
	a, _ := attrs.(*graylog.InputBeatsAttrs)
	if a.BindAddress != "0.0.0.0" {
		t.Fatalf(`bind_address = "%s", wanted "%s"`, a.BindAddress, "0.0.0.0")
	}
}
