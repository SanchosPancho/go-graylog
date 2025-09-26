package graylog

// PipelineRule represents a Graylog's Pipeline Rule.
// https://docs.graylog.org/en/latest/pages/pipelines/rules.html
type PipelineRule struct {
	// required
	Source string `json:"source,omitempty" v-create:"required" v-update:"required"`
	ID     string `json:"id,omitempty" v-create:"isdefault" v-update:"required"`
	// Note that title is ignored in create and update API.
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	// CreatedAt   string `json:"created_at,omitempty"`
	// ModifiedAt   string `json:"modified_at,omitempty"`
	// Errors string `json:"errors"`
}
