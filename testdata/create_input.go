package testdata

import (
	"github.com/SanchosPancho/go-graylog/v11/graylog/graylog"
)

func CreateInput() graylog.Input {
	return graylog.Input{
		Title: "gelf udp 2",
		Attrs: &graylog.InputGELFUDPAttrs{
			DecompressSizeLimit: 8388608,
			BindAddress:         "0.0.0.0",
			Port:                12201,
			RecvBufferSize:      262144,
		},
		Global: true,
	}
}
