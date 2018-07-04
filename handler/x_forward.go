package handler

import (
	"github.com/ServiceComb/go-chassis/client/rest"
	"github.com/ServiceComb/go-chassis/core/handler"
	"github.com/ServiceComb/go-chassis/core/invocation"
)

//XForward is a costant
const XForward = "x-forward"

//XForwardHandler ..
type XForwardHandler struct {
}

//Handle function
func (h *XForwardHandler) Handle(chain *handler.Chain, inv *invocation.Invocation, cb invocation.ResponseCallBack) {
	orgReq, ok := inv.Args.(*rest.Request)
	if ok && orgReq.Req.Header["X-Forwarded-Host"] == nil {
		orgHost := orgReq.Req.Header["Host"]
		orgReq.Req.Header["X-Forwarded-Host"] = orgHost
	}
	chain.Next(inv, func(r *invocation.Response) error {
		return cb(r)
	})
}

//Name returns name
func (h *XForwardHandler) Name() string { return XForward }

//NewHandler creates new handler and returns it
func NewHandler() handler.Handler { return &XForwardHandler{} }

func init() { handler.RegisterHandler(XForward, NewHandler) }
