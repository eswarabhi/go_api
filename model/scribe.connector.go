package model

// ScribeOrgConnector type for scribe installed connectors
type ScribeOrgConnector struct {
	ID               string          `json:"id"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	VendorName       string          `json:"vendorName"`
	IsCloudSupported bool            `json:"isCloudSupported"`
	UName            string          `json:"uName"`
	UVersion         string          `json:"uVersion"`
	SolutionRoles    []SolutionRoles `json:"solutionRoles"`
}

// ScribeMarketplaceConnector type for scribe all connectors
type ScribeMarketplaceConnector struct {
	AdapterTypeID           string                   `json:"adapterTypeId"`
	AgentInstallStatuses    []AgentInstallStatus     `json:"agentInstallStatuses"`
	CertificationDate       string                   `json:"certificationDate"`
	ContactURL              string                   `json:"contactUrl"`
	Description             string                   `json:"description"`
	Installed               bool                     `json:"installed"`
	IsAssociated            bool                     `json:"isAssociated"`
	IsFree                  bool                     `json:"isFree"`
	IsISSource              bool                     `json:"isISSource"`
	IsISTarget              bool                     `json:"isISTarget"`
	IsMSSource              bool                     `json:"isMSSource"`
	IsMSTarget              bool                     `json:"isMSTarget"`
	IsPublic                bool                     `json:"isPublic"`
	IsRSSource              bool                     `json:"isRSSource"`
	IsRSTarget              bool                     `json:"isRSTarget"`
	Link                    string                   `json:"link"`
	Logo                    []string                 `json:"logo"`
	LogoURL                 string                   `json:"logoUrl"`
	Name                    string                   `json:"name"`
	ProviderName            string                   `json:"providerName"`
	RequiresJava            bool                     `json:"requiresJava"`
	SupportsCloudAgents     bool                     `json:"supportsCloudAgents"`
	ConnectorUpdateVersions []ConnectorUpdateVersion `json:"connectorUpdateVersions"`
}

type AgentInstallStatus struct {
	agentId   string
	agentName string
	status    string
}
type ConnectorUpdateVersion struct {
	lastModificationDate string
	version              string
	isHotFix             bool
	isLocked             bool
}

type ScribeUserInfosso struct {
	Id                  string `json:"id"`
	CurrentTenantId     string `json:"currentTenantId"`
	CurrentOrgNumber    int    `json:"currentOrgNumber"`
	CurrentOrgName      string `json:"currentOrgName"`
	CurrentOrgRole      string `json:"currentOrgRole"`
	CurrentOrgType      string `json:"currentOrgType"`
	Culture             string `json:"culture"`
	PasswordMustChange  bool   `json:"passwordMustChange"`
	CurrentUserName     string `json:"currentUserName"`
	CurrentTenantStatus string `json:"currentTenantStatus"`
	AcceptedEula        string `json:"acceptedEula"`
	IsTibcoOrg          bool   `json:"isTibcoOrg"`
	SessionId           string `json:"sessionId"`
}

type SolutionRoles struct {
	name string
}
