package simpleregistry

import (
//	"github.com/go-chassis/mesher/plugin/dubbo/dubbo"
//"github.com/go-chassis/mesher/plugin/dubbo/utils"
)

//RegistryURL is a struct which has attributes of a URL
type RegistryURL struct {
	Protocol   string
	Username   string
	Password   string
	Host       string
	Port       int
	Path       string
	Parameters map[string]string
}
