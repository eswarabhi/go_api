package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../config"
	"../mock"
	"../model"
)

// HealthCheck will respond with the status whether the server is up or not
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Health Check Ping!!", r.URL.Path)
	encoder := json.NewEncoder(w)
	encoder.Encode(model.Health{
		Title: "Server is up!",
	})
}

// GetProviders will fetch all the providers
func GetProviders(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	encoder.Encode(mock.Providers)
}

// GetConnectors will fetch the connectors from all the providers
func GetConnectors(w http.ResponseWriter, r *http.Request) {
	fmt.Print(model.SCRIBE)
	var flogoConnectors []model.FlogoMarketplaceConnector
	res := model.Connector{
		Connectors: []model.ConnectorList{},
		Errors:     []model.ConnectorError{},
	}
	encoder := json.NewEncoder(w)

	// Creating a Http Client for Scribe API
	scribeClient := &http.Client{
		CheckRedirect: scribeRedirectPolicy,
	}

	// Get scribe Org Info
	scribeInfo := getScribeInfo(scribeClient)

	// Get scribe connectors
	scribeAllConnectors := getScribeAllConnectors(scribeClient, scribeInfo.CurrentOrgNumber)
	scribeInstalledConnectors := getScribeInstalledConnectors(scribeClient, scribeInfo.CurrentOrgNumber)

	// mark installed
	scribeConnectors := markInstalled(scribeAllConnectors, scribeInstalledConnectors)

	// Get flogo connectors
	if err := json.Unmarshal([]byte(mock.FlogoConnectors), &flogoConnectors); err != nil {
		log.Fatal("parsing flogo connectors JSON failed!", err)
	}

	//Format connectors and push to response connectors array
	formatScribeConnectors(&res, scribeConnectors, scribeInfo.CurrentOrgNumber)
	formatFlogoConnectors(&res, flogoConnectors)
	// Sending all connectors as response
	encoder.Encode(res)
}

func formatScribeConnectors(res *model.Connector, connectors []model.ScribeMarketplaceConnector, scribeOrgID int) {
	for _, c := range connectors {
		res.Connectors = append(res.Connectors, model.ConnectorList{
			ID:                c.AdapterTypeID,
			Name:              c.Name,
			Description:       c.Description,
			Provider:          fmt.Sprint(model.SCRIBE),
			Icon:              fmt.Sprintf("%s/orgs/%d/connectors/%s/logo", config.Global.ScribeBaseUrl, scribeOrgID, c.AdapterTypeID),
			CertificationDate: c.CertificationDate,
			Installed:         c.Installed,
			Uninstallable:     true,
			IsFree:            c.IsFree,
			ContactURL:        c.ContactURL,
		})
	}
}

func formatFlogoConnectors(res *model.Connector, connectors []model.FlogoMarketplaceConnector) {
	for _, c := range connectors {
		res.Connectors = append(res.Connectors, model.ConnectorList{
			ID:                c.Name,
			Name:              c.Title,
			Description:       c.Display.Description,
			Provider:          fmt.Sprint(model.FLOGO),
			Icon:              fmt.Sprintf("/wistudio/v1/contributions/Tibco/%s/%s/%s", c.Display.Category, c.Name, c.Display.SmallIcon),
			CertificationDate: c.LastModifiedDate,
			Installed:         true,
			Uninstallable:     false,
		})
	}
}

func getScribeInfo(client *http.Client) model.ScribeUserInfosso {
	var info model.ScribeUserInfosso

	if req, err := http.NewRequest("GET", config.Global.ScribeBaseUrl+"/users/infosso", nil); err != nil {
		log.Fatal("getScribeInfo: Create Request Error: ", err)
	} else {
		// Add auth headers
		req.SetBasicAuth("tpeng@d-link.ml", "Owen.2019")

		// Make the request
		if resp, err := client.Do(req); err != nil {
			log.Fatal("getScribeInfo: Response Error: ", err)
		} else {
			rawInfo, _ := ioutil.ReadAll(resp.Body)
			if json.Unmarshal(rawInfo, &info) != nil {
				log.Fatal("Unmarshalling Scribe Info failed!")
			}
		}
	}
	return info
}

func getScribeInstalledConnectors(client *http.Client, orgID int) []model.ScribeOrgConnector {
	var connectors []model.ScribeOrgConnector
	URL := fmt.Sprintf("%s/orgs/%d/connectors", config.Global.ScribeBaseUrl, orgID)
	log.Print(URL)
	if req, err := http.NewRequest("GET", URL, nil); err != nil {
		log.Fatal("Create Request Error: ", err)
	} else {
		// Add auth headers
		req.SetBasicAuth("tpeng@d-link.ml", "Owen.2019")

		// Make the request
		if resp, err := client.Do(req); err != nil {
			log.Fatal("Scribe Response Error: ", err)
		} else {
			rawConnectors, _ := ioutil.ReadAll(resp.Body)
			if json.Unmarshal(rawConnectors, &connectors) != nil {
				log.Fatal("Unmarshalling Scribe connectors failed!")
			}
		}
	}
	return connectors
}

func getScribeAllConnectors(client *http.Client, orgID int) []model.ScribeMarketplaceConnector {
	var connectors []model.ScribeMarketplaceConnector
	var tail string
	if orgID > 0 {
		tail = fmt.Sprintf("?orgId=%d", orgID)
	} else {
		tail = ""
	}

	URL := fmt.Sprintf("%s/marketplace/connectors%s", config.Global.ScribeBaseUrl, tail)
	log.Print(URL)
	if req, err := http.NewRequest("GET", URL, nil); err != nil {
		log.Fatal("Create Request Error: ", err)
	} else {
		// Add auth headers
		req.SetBasicAuth("tpeng@d-link.ml", "Owen.2019")

		// Make the request
		if resp, err := client.Do(req); err != nil {
			log.Fatal("Scribe Response Error: ", err)
		} else {
			rawConnectors, _ := ioutil.ReadAll(resp.Body)
			if json.Unmarshal(rawConnectors, &connectors) != nil {
				log.Fatal("Unmarshalling Scribe All connectors failed!")
			}
		}
	}
	return connectors
}

func markInstalled(allConnectors []model.ScribeMarketplaceConnector, installedConnectors []model.ScribeOrgConnector) []model.ScribeMarketplaceConnector {

	for _, ic := range installedConnectors {
		for _, ac := range allConnectors {
			if ac.AdapterTypeID == ic.ID {
				ac.Installed = true
				break
			}
		}
	}
	return allConnectors

}
func scribeRedirectPolicy(req *http.Request, via []*http.Request) error {
	req.SetBasicAuth("tpeng@d-link.ml", "Owen.2019")
	return nil
}
