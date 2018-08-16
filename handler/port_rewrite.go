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
	"github.com/go-chassis/go-chassis/core/handler"
	"github.com/go-chassis/go-chassis/core/invocation"
	"strings"
)

//PortMapForPilot is a constant
const PortMapForPilot = "port-selector"

//PortSelectionHandler ..
type PortSelectionHandler struct {
}

//Handle function replace the provider port to mesher port so that traffic goes through mesher
func (ps *PortSelectionHandler) Handle(chain *handler.Chain, inv *invocation.Invocation, cb invocation.ResponseCallBack) {
	inv.Endpoint = replacePort(inv.Endpoint)

	if inv.Endpoint == "" {
		r := &invocation.Response{
			Err: fmt.Errorf("invalid endpoint"),
		}
		cb(r)
		return
	}

	chain.Next(inv, func(r *invocation.Response) error {
		return cb(r)
	})
}

//replacePort will replace the provider port with mesher port.
func replacePort(endpoint string) string {
	eps := strings.Split(endpoint, ":")
	if len(eps) != 2 {
		return ""
	}

	eps[1] = "30101"

	return strings.Join(eps, ":")
}

//Name returns name
func (ps *PortSelectionHandler) Name() string {
	return PortMapForPilot
}

//New create new port for pilot handler and retuns
func New() handler.Handler {
	return &PortSelectionHandler{}
}

func init() {
	handler.RegisterHandler(PortMapForPilot, New)
}
