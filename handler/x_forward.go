/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handler

import (
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core/handler"
	"github.com/go-chassis/go-chassis/core/invocation"
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
