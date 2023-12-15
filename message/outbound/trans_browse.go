package outbound

// getBrowse -> fetchData -> sendProxyMessage -> sendRequest -> Java
type TransactionBrowseBody struct {
	RequestId   string `json:"requestId"`
	StartDate   int64  `json:"startDate"`
	EndDate     int64  `json:"endDate"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	AppCode     string `json:"appCode"`
	AppId       string `json:"appId"`
	AppVer      int    `json:"appVer"`
	BuildNumber int    `json:"buildNumber"`
	Channel     string `json:"channel"`
	DeviceOs    string `json:"deviceOS"`
}
