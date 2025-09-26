package testdata

import (
	"github.com/SanchosPancho/go-graylog/v11/graylog/graylog"
)

func RequestCreateEventNotification() *graylog.EventNotification {
	return &graylog.EventNotification{
		Title:       "http",
		Description: "",
		Config: map[string]interface{}{
			"type": "http-notification-v1",
			"url":  "http://example.com",
		},
	}
}
