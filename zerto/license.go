//
// Zerto API Interface Wrapper
//
// Author: Martin Weber <martin.weber@de.clara.net>
//
package zerto

import "encoding/json"

//
// /v1/license
//
type ZertoLicense struct {
	Details						ZertoLicenseDetails
	Usage							ZertoLicenseUsage
}

type ZertoLicenseDetails struct {
	ExpiryTime				string
	LicenseKey				string
	LicenseType				string
	MaxVms						string
}

type ZertoLicenseUsage struct {
	SitesUsage				*ZertoLicenseSiteUsage
	TotalVmsCount			float64
}

type ZertoLicenseSiteUsage struct {
	ProtectedVmsCount	float64
	SiteIdentifier		string
	SiteName					string
}
//
// Action: /v1/license
//
// Fetch Information about local Zerto instance
//
func (z *Zerto) LicenseInformations() ZertoLicense {
	resp, _ := z.makeRequest("GET", "/license", RequestParams{})
	data := json.NewDecoder(resp.Body)

	var d ZertoLicense
	data.Decode(&d)

	return d
}
