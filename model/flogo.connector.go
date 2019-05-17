package model

// FlogoMarketplaceConnector type for flogo connectors
type FlogoMarketplaceConnector struct {
	Name             string              `json:"name"`
	Title            string              `json:"title"`
	Version          string              `json:"version"`
	Author           string              `json:"author"`
	LastModifiedDate string              `json:"lastModifiedDate"`
	Ref              string              `json:"ref"`
	Type             string              `json:"type"`
	Owner            string              `json:"owner"`
	Display          ContributionDisplay `json:"display"`
}

// ContributionDisplay to display the contribution details
type ContributionDisplay struct {
	/**
	 * display name for the contribution
	 */
	Name string `json:"name"`
	/**
	 * Specifies the category of the contribution e.g. Salesforce,Marketo , default is "Default"
	 */
	Category string `json:"category"`
	/**
	 * Specifies if a contribution is visible and enabled
	 */
	Visible bool `json:"visible"`
	/**
	 * Specifies the small icon path. The path is relative to the contribution folder root. i.e. icons/file.png
	 */
	SmallIcon string `json:"smallIcon"`
	/**
	 * Specifies the small icon path. The path is relative to the contribution folder root. i.e. icons/file.png
	 */
	LargeIcon string `json:"largeIcon"`
	/**
	 * Description
	 */
	Description string `json:"description"`
}
