package outbound

// fetchDataDetail -> fetchData -> sendProxyMessage -> sendRequest -> Java
type TransactionDetailBody struct {
	RequestId   string `json:"requestId"`
	TransId     string `json:"transId"`
	ServiceId   string `json:"serviceId"`
	AppCode     string `json:"appCode"`
	AppId       string `json:"appId"`
	AppVer      int    `json:"appVer"`
	BuildNumber int    `json:"buildNumber"`
	Channel     string `json:"channel"`
	DeviceOs    string `json:"deviceOS"`
}
