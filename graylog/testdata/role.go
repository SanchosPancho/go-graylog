package testdata

import (
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/SanchosPancho/go-graylog/v11/graylog/graylog"
)

var (
	Role = graylog.Role{
		Name:        "foo",
		Description: "Allows reading and writing all views and extended searches (built-in)",
		Permissions: set.NewStrSet(
			"extendedsearch:create",
			"extendedsearch:use",
			"view:create",
			"view:edit",
			"view:read",
			"view:use",
		),
		ReadOnly: true,
	}
)
