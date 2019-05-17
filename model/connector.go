package model

// Health response type
type Health struct {
	Title string `json:"title"`
}

type Connector struct {
	Connectors []ConnectorList  `json:"connectors"`
	Errors     []ConnectorError `json:"errors"`
}
type ConnectorList struct {
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	Provider          string      `json:"provider"`
	Icon              interface{} `json:"icon"`
	CertificationDate interface{} `json:"certificationDate"`
	Installed         bool        `json:"installed"`
	Uninstallable     bool        `json:"uninstallable"`
	IsFree            bool        `json:"isFree"`
	ContactURL        string      `json:"contactUrl"`
}
type ConnectorError struct {
	Type   Provider
	status int
}
type ConnectorProviders struct {
	providers []ConnectorProvider
}
type ConnectorProvider struct {
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
}

type Provider int

const (
	FLOGO Provider = iota
	SCRIBE
)

func (p Provider) String() string {
	return [...]string{"flogo", "scribe"}[p]
}

type ProviderDisplayName int

const (
	DFLOGO ProviderDisplayName = iota
	DSCRIBE
)

func (p ProviderDisplayName) String() string {
	return [...]string{"Flogo", "Scribe"}[p]
}
