package testdata

import (
	"github.com/SanchosPancho/go-graylog"
)

func CreateStream() graylog.Stream {
	return graylog.Stream{
		Title:        "All system events",
		IndexSetID:   "5d84bfbe2ab79c000d35d4a9",
		Description:  "Stream containing all system events created by Graylog",
		MatchingType: "AND",
	}
}
