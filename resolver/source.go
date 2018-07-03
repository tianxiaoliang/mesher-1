package resolver

import (
	"errors"
	"github.com/ServiceComb/go-chassis/core/lager"
	"github.com/ServiceComb/go-chassis/core/registry"
	"github.com/ServiceComb/go-sc-client/model"
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
