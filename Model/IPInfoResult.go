package Model

// IPInfoResult IP信息
type IPInfoResult struct {
	IP       string `json:"ip"`
	Region   string `json:"region"`
	Duration string `json:"duration"`
	Error    string `json:"error,omitempty"`
}
