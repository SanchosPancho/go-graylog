package client

import (
	"context"

	"github.com/SanchosPancho/go-graylog/graylog"
)

// GetAlertConditions returns all alert conditions.
func (client *Client) GetAlertConditions(ctx context.Context) (
	[]graylog.AlertCondition, int, *ErrorInfo, error,
) {
	conditions := &graylog.AlertConditionsBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints().AlertConditions(), nil, conditions)
	return conditions.AlertConditions, conditions.Total, ei, err
}
