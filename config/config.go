package config

// version refers to gonfleet package version
const version = "1.0.0"

// name refers to name of package
const name = "onfleet/gonfleet"

// defaults - can be overriden by user init params
const (
	defaultBaseUrl    = "https://onfleet.com"
	defaultPath       = "/api"
	defaultApiVersion = "/v2"
)
const DefaultUserTimeout int64 = 70000 // in milliseconds

// InitParams accepts user provided overrides to be set on Config
type InitParams struct {
	ApiKey      string
	UserTimeout int64
	BaseUrl     string
	Path        string
	ApiVersion  string
}

// Config provides necessary common data to all resources
type Config struct {
	ApiKey     string
	BaseUrl    string
	Path       string
	ApiVersion string
	Version    string
	Name       string
}

// InitConfig returns Config.
// Any default overrides provided by the user are applied
func InitConfig(params InitParams) Config {
	baseUrl := "https://onfleet.com"
	path := "/api"
	apiVersion := "/v2"

	if params.BaseUrl != "" {
		baseUrl = params.BaseUrl
	}
	if params.Path != "" {
		path = params.Path
	}
	if params.ApiVersion != "" {
		apiVersion = params.ApiVersion
	}

	return Config{
		ApiKey:     params.ApiKey,
		BaseUrl:    baseUrl,
		Path:       path,
		ApiVersion: apiVersion,
		Version:    version,
		Name:       name,
	}
}
