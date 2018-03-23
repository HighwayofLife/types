package client

const (
	AuthnConfigType          = "authnConfig"
	AuthnConfigFieldOptions  = "options"
	AuthnConfigFieldSans     = "sans"
	AuthnConfigFieldStrategy = "strategy"
)

type AuthnConfig struct {
	Options  map[string]string `json:"options,omitempty" yaml:"options,omitempty"`
	Sans     []string          `json:"sans,omitempty" yaml:"sans,omitempty"`
	Strategy string            `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}
