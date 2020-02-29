package main

type Targetconfig struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Firewallip string `json:"firewallip"`
	Username   string `json:"username"`
	Apikey     string `json:"apikey"`
	Enabled    string `json:"enabled"`
	Devtype    string `json:"devtype"`
	Status     string `json:"status"`
}

type Backendconfig struct {
	Targets []Targetconfig `json:"targets"`
}

type Frontendconfig struct {
	Hostname     string `json:"hostname"`
	Port         string `json:"port"`
	Headertext   string `json:"headertext"`
	Logofile     string `json:"logofile"`
	Showuserinfo bool   `json:"showuserinfo"`
	Headertitle  string `json:"headertitle"`
}

type Config struct {
	BE Backendconfig  `json:"backend"`
	FE Frontendconfig `json:"frontend"`
}

type FOSDetectedDeviceResults struct {
	Results []FOSDetectedDevice `json:"results"`
}

type FOSDetectedDevice struct {
	Mac                             string   `json:"mac"`
	MasterMac                       string   `json:"master_mac"`
	IsMasterDevice                  bool     `json:"is_master_device"`
	HardwareVendor                  string   `json:"hardware_vendor"`
	HardwareType                    string   `json:"hardware_type"`
	HardwareFamily                  string   `json:"hardware_family"`
	OsName                          string   `json:"os_name"`
	OsVersion                       string   `json:"os_version"`
	Hostname                        string   `json:"hostname"`
	Ipv4Address                     string   `json:"ipv4_address"`
	Ipv6Address                     string   `json:"ipv6_address"`
	DetectedInterface               string   `json:"detected_interface"`
	DetectedInterfaceFortitelemetry bool     `json:"detected_interface_fortitelemetry"`
	IsDetectedInterfaceRoleWan      bool     `json:"is_detected_interface_role_wan"`
	LastSeen                        int      `json:"last_seen"`
	IsOnline                        bool     `json:"is_online"`
	OnlineInterfaces                []string `json:"online_interfaces"`
	Sourcefg                        string   `json:"sourcefg"`
}

type FOSCMDBSystemVdom struct {
	Name       string `json:"name"`
	QOriginKey string `json:"q_origin_key"`
	ShortName  string `json:"short-name"`
	VclusterID int    `json:"vcluster-id"`
	Flag       int    `json:"flag"`
}

type FOSCMDBSystemVdomResults struct {
	Results []FOSCMDBSystemVdom `json:"results"`
	FOSResults
}

type FOSResults struct {
	HTTPMethod string `json:"http_method"`
	Revision   string `json:"revision"`
	Vdom       string `json:"vdom"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	HTTPStatus int    `json:"http_status"`
	Serial     string `json:"serial"`
	Version    string `json:"version"`
	Build      int    `json:"build"`
}

type TargetApiCall struct {
	Target string `json:"target"`
	Path   string `json:"path"`
	Method string `json:"method"`
	Data   []byte `json:"data"`
}

type DbObj struct {
	Id   []byte `json:"id"`
	Data []byte `json:"data"`
}
