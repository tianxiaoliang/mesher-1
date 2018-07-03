package health

//StatusCode is type of string
type StatusCode string

const (
	//Red is a const
	Red StatusCode = "red"
	//Green is a const
	Green StatusCode = "green"
)

//Health has details about health of a service
type Health struct {
	ServiceName                 string     `json:"serviceName,omitempty"`
	Version                     string     `json:"version,omitempty"`
	Status                      StatusCode `json:"status,omitempty"`
	ConnectedConfigCenterClient bool       `json:"connectedConfigCenterClient"`
	ConnectedMonitoring         bool       `json:"connectedMonitoring"`
	Error                       string     `json:"error,omitempty"`
}
