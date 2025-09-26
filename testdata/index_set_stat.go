package testdata

import (
	"github.com/SanchosPancho/go-graylog"
)

func IndexSetStats() graylog.IndexSetStats {
	return graylog.IndexSetStats{
		Indices:   1,
		Documents: 0,
		Size:      1044,
	}
}
