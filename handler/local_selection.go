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
