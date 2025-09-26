package graylog

const (
	// InputTypeBeats is one of input types.
	InputTypeBeats string = "org.graylog.plugins.beats.BeatsInput"
)

// NewInputBeatsAttrs is the constructor of InputBeatsAttrs.
func NewInputBeatsAttrs() InputAttrs {
	return &InputBeatsAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputBeatsAttrs) InputType() string {
	return InputTypeBeats
}

// InputBeatsAttrs represents Beats Input's attributes.
type InputBeatsAttrs struct {
	BindAddress           string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	OverrideSource        string `json:"override_source,omitempty"`
	TLSKeyFile            string `json:"tls_key_file,omitempty"`
	TLSKeyPassword        string `json:"tls_key_password,omitempty"`
	TLSClientAuthCertFile string `json:"tls_client_auth_cert_file,omitempty"`
	TLSClientAuth         string `json:"tls_client_auth,omitempty"`
	TLSCertFile           string `json:"tls_cert_file,omitempty"`
	TLSEnable             bool   `json:"tls_enable"`
	TCPKeepAlive          bool   `json:"tcp_keepalive"`
	Port                  int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize        int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
}
