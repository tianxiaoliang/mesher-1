package config

//MesherConfig has all mesher config
type MesherConfig struct {
	PProf      *PProf  `yaml:"pprof"`
	Plugin     *Plugin `yaml:"plugin"`
	Admin      Admin   `yaml:"admin"`
	ProxyedPro string  `yaml:"proxyedProtocol"`
}

//PProf has enable and listen attribute for pprof
type PProf struct {
	Enable bool   `yaml:"enable"`
	Listen string `yaml:"listen"`
}

//Policy has attributes for destination, tags and loadbalance
type Policy struct {
	Destination   string            `yaml:"destination"`
	Tags          map[string]string `yaml:"tags"`
	LoadBalancing map[string]string `yaml:"loadBalancing"`
}

//Plugin has attributes for destination and source resolver
type Plugin struct {
	DestinationResolver string `yaml:"destinationResolver"`
	SourceResolver      string `yaml:"sourceResolver"`
}

//Admin has attributes for enabling, serverURI and metrics for admin data
type Admin struct {
	Enable           *bool  `yaml:"enable"`
	ServerURI        string `yaml:"serverUri"`
	GoRuntimeMetrics bool   `yaml:"goRuntimeMetrics"`
}
