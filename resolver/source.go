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

package resolver

import (
	"errors"
	"github.com/go-chassis/go-chassis/core/lager"
	"github.com/go-chassis/go-chassis/core/registry"
	"github.com/go-chassis/go-sc-client/model"
	"github.com/go-chassis/mesher/common"
)

var (
	//ErrFoo is of type error
	ErrFoo = errors.New("resolved as a nil service")
)

//SourceResolver is a interface which has Resolve function
type SourceResolver interface {
	Resolve(source string) *registry.SourceInfo
}

var sr SourceResolver = &DefaultSourceResolver{}

//DefaultSourceResolver is a struct
type DefaultSourceResolver struct {
}

//Resolve is a method which resolves service endpoint
func (sr *DefaultSourceResolver) Resolve(source string) *registry.SourceInfo {
	if source == "127.0.0.1" {
		return nil
	}
	cacheDatum, ok := registry.IPIndexedCache.Get(source)
	if !ok {
		return nil
	}
	ms, ok := cacheDatum.(*model.MicroService)
	if !ok {
		return nil
	}

	if ms == nil {
		lager.Logger.Warnf("Service is nil for IP %s, err: %v", source, ErrFoo)
		return nil
	}
	sourceInfo := &registry.SourceInfo{}
	sourceInfo.Tags = make(map[string]string)
	sourceInfo.Name = ms.ServiceName
	sourceInfo.Tags[common.BuildInTagApp] = ms.AppID
	sourceInfo.Tags[common.BuildInTagVersion] = ms.Version
	for k, v := range ms.Properties {
		sourceInfo.Tags[k] = v
	}
	return sourceInfo
}

//GetSourceResolver returns interface object
func GetSourceResolver() SourceResolver {
	return sr
}
