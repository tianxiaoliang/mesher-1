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
 * Unless required by applicable law or agreeed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	_ "github.com/ServiceComb/cse-collector"            // use cse monitoring
	"github.com/ServiceComb/go-chassis"                 //Use go chassis EE
	_ "github.com/ServiceComb/go-chassis/config-center" //use config center
	"github.com/ServiceComb/go-chassis/core/lager"
	"github.com/go-chassis/mesher/adminapi/version"
	"github.com/go-chassis/mesher/bootstrap"
	"github.com/go-chassis/mesher/cmd"
	"github.com/go-chassis/mesher/config"
	_ "github.com/go-chassis/mesher/handler"
	_ "github.com/go-chassis/mesher/protocol/dubbo/client/chassis"
	_ "github.com/go-chassis/mesher/protocol/dubbo/server"
	_ "github.com/go-chassis/mesher/protocol/dubbo/simpleRegistry"
	_ "github.com/go-chassis/mesher/protocol/http"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// server init
	if err := cmd.Init(); err != nil {
		panic(err)
	}
	if err := cmd.Configs.GeneratePortsMap(); err != nil {
		panic(err)
	}
	bootstrap.RegisterFramework()
	bootstrap.SetHandlers()
	if err := chassis.Init(); err != nil {
		lager.Logger.Error("Go chassis init failed, Mesher is not available", err)
	}
	if err := bootstrap.Start(); err != nil {
		lager.Logger.Error("Bootstrap failed ", err)
		panic(err)
	}
	lager.Logger.Infof("Version is %s", version.Ver().Version)
	if config.GetConfig().PProf != nil {
		if config.GetConfig().PProf.Enable {
			go func() {
				if config.GetConfig().PProf.Listen == "" {
					config.GetConfig().PProf.Listen = "127.0.0.1:6060"
				}
				lager.Logger.Warn("Enable pprof on "+config.GetConfig().PProf.Listen, nil)
				if err := http.ListenAndServe(config.GetConfig().PProf.Listen, nil); err != nil {
					lager.Logger.Error("Can not enable pprof", err)
				}
			}()
		}
	}
	chassis.Run()
}
