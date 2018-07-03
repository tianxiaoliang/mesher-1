package handler

import (
	"fmt"
	"github.com/ServiceComb/go-chassis/core/handler"
	"github.com/ServiceComb/go-chassis/core/invocation"
	"github.com/go-chassis/mesher/cmd"
	"github.com/go-chassis/mesher/common"
)

//LocalSelection is a constant
const LocalSelection = "local-selection"

//LocalSelectionHandler ..
type LocalSelectionHandler struct {
}

//Handle function gets locally defined handler
func (ls *LocalSelectionHandler) Handle(chain *handler.Chain, inv *invocation.Invocation, cb invocation.ResponseCallBack) {
	// if work as sidecar and handler request from remote,then endpoint should be localhost:port
	inv.Endpoint = cmd.Configs.PortsMap[inv.Protocol]
	if inv.Endpoint == "" {
		r := &invocation.Response{
			//			Err: errors.New(
			//				fmt.Sprintf("[%s] is not supported, [%s] didn't set env [%s] or cmd parameter --service-ports before mesher start",
			//					inv.Protocol, inv.MicroServiceName, common.EnvServicePorts)),
			Err: fmt.Errorf("[%s] is not supported, [%s] didn't set env [%s] or cmd parameter --service-ports before mesher start",
				inv.Protocol, inv.MicroServiceName, common.EnvServicePorts),
		}
		cb(r)
		return
	}
	chain.Next(inv, func(r *invocation.Response) error {
		return cb(r)
	})
}

//Name returns name
func (ls *LocalSelectionHandler) Name() string {
	return LocalSelection
}

//New create new local selection handler and retuns
func New() handler.Handler {
	return &LocalSelectionHandler{}
}
func init() {
	handler.RegisterHandler(LocalSelection, New)
}
