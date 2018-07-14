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
