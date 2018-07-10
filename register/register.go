package register

import (
	chassisCommon "github.com/ServiceComb/go-chassis/core/common"
	"github.com/ServiceComb/go-chassis/core/config"
	chassisModel "github.com/ServiceComb/go-chassis/core/config/model"
	"github.com/ServiceComb/go-chassis/core/lager"
	"github.com/ServiceComb/go-chassis/core/registry"
	"github.com/ServiceComb/go-chassis/pkg/util/iputil"
	"github.com/go-chassis/mesher/common"
	"strings"
)

// AdaptEndpoints moves http endpoint to rest endpoint
func AdaptEndpoints() {
	// To be called by services based on CSE SDK,
	// mesher has to register endpoint with rest://ip:port
	oldProtoMap := config.GlobalDefinition.Cse.Protocols
	if _, ok := oldProtoMap[common.HTTPProtocol]; !ok {
		return
	}
	if _, ok := oldProtoMap[chassisCommon.ProtocolRest]; ok {
		return
	}

	newProtoMap := make(map[string]chassisModel.Protocol)
	for n, proto := range oldProtoMap {
		if n == common.HTTPProtocol {
			continue
		}
		newProtoMap[n] = proto
	}
	newProtoMap[chassisCommon.ProtocolRest] = oldProtoMap[common.HTTPProtocol]
	registry.InstanceEndpoints = registry.MakeEndpointMap(newProtoMap)
	for protocol, address := range registry.InstanceEndpoints {
		if address == "" {
			port := strings.Split(newProtoMap[protocol].Listen, ":")
			if len(port) == 2 { //check if port is not specified along with ip address, eventually in case port is not specified, server start will fail in subsequent processing.
				registry.InstanceEndpoints[protocol] = iputil.GetLocalIP() + ":" + port[1]
			}
		}
	}

	lager.Logger.Debug("Adapt endpoints success")
}
