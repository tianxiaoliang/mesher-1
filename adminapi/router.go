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

package adminapi

import (
	"crypto/tls"
	gorestful "github.com/emicklei/go-restful"
	chassisCom "github.com/go-chassis/go-chassis/core/common"
	"github.com/go-chassis/go-chassis/core/lager"
	chassisTLS "github.com/go-chassis/go-chassis/core/tls"
	"github.com/go-chassis/mesher/common"
	"github.com/go-chassis/mesher/config"
	"github.com/go-chassis/mesher/metrics"
	"net"
	"net/http"
	"strings"
	"time"
)

//Init function initiates admin server config and runs it
func Init() (err error) {
	var isAdminEnable *bool = config.GetConfig().Admin.Enable

	if isAdminEnable != nil && *isAdminEnable == false {
		lager.Logger.Infof("admin api are not enable")
		return nil
	}

	errCh := make(chan error)
	metrics.Init()

	adminServerURI := config.GetConfig().Admin.ServerURI

	if adminServerURI == "" {
		adminServerURI = "0.0.0.0:30102"
	}
	ln, err := net.Listen("tcp", adminServerURI)
	if err != nil {
		return
	}
	tlsConfig, err := getTLSConfig()
	if err != nil {
		return
	}
	if tlsConfig != nil {
		lager.Logger.Infof("admin server is using ssl")
		ln = tls.NewListener(ln, tlsConfig)
	} else {
		lager.Logger.Infof("admin server is not using ssl")
	}

	go func() {
		lager.Logger.Infof("admin server listening on %s", ln.Addr().String())
		restfulWebService := GetWebService()
		gorestful.Add(&restfulWebService)
		if err := http.Serve(ln, nil); err != nil {
			errCh <- err
			return
		}
	}()

	select {
	case err = <-errCh:
		lager.Logger.Warnf("got Admin Server Error, err: %v", err)
	case <-time.After(time.Second):
		lager.Logger.Infof("admin server start success")
		err = nil
	}
	return
}

func getTLSConfig() (*tls.Config, error) {
	var tlsConfig *tls.Config
	sslTag := genTag(common.ComponentName, chassisCom.Provider)
	tmpTLSConfig, sslConfig, err := chassisTLS.GetTLSConfigByService(common.ComponentName, "", chassisCom.Provider)
	if err != nil {
		if !chassisTLS.IsSSLConfigNotExist(err) {
			return nil, err
		}
	} else {
		lager.Logger.Warnf("%s TLS mode, verify peer: %t, cipher plugin: %s.",
			sslTag, sslConfig.VerifyPeer, sslConfig.CipherPlugin)
		tlsConfig = tmpTLSConfig
	}
	return tlsConfig, nil
}

func genTag(s ...string) string {
	return strings.Join(s, ".")
}
