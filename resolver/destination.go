package resolver

import (
	"errors"
	"log"
	"net/url"

	"github.com/ServiceComb/go-chassis/core/lager"
	"github.com/ServiceComb/go-chassis/core/util/string"
	"github.com/go-chassis/mesher/config"
)

var dr DestinationResolver

//DestinationResolverPlugins is a map
var DestinationResolverPlugins map[string]func() DestinationResolver

//SelfEndpoint is a string
var SelfEndpoint = "#To be init#"

//DefaultPlugin is a contant which stores default plugin name
const DefaultPlugin = "host"

//ErrUnknownResolver is of type error
var ErrUnknownResolver = errors.New("unknown Destination Resolver")

//DestinationResolver is a interface with Resolve method
type DestinationResolver interface {
	Resolve(sourceAddr string, header map[string]string, rawURI string, destinationName *string) error
}

//DefaultDestinationResolver is a struct
type DefaultDestinationResolver struct {
}

//Resolve resolves service's endpoint
func (dr *DefaultDestinationResolver) Resolve(sourceAddr string, header map[string]string, rawURI string, destinationName *string) error {
	u, err := url.Parse(rawURI)
	if err != nil {
		lager.Logger.Error("Can not parse url", err)
		return err
	}

	if u.Host == "" {
		return errors.New(`Invalid uri, please check:
1, For provider, mesher listens on external ip
2, Set http_proxy as mesher address, before sending request`)
	}

	if u.Host == SelfEndpoint {
		return errors.New(`uri format must be: http://serviceName/api`)
	}

	if h := stringutil.SplitFirstSep(u.Host, ":"); h != "" {
		*destinationName = h
		return nil
	}

	*destinationName = u.Host
	return nil
}

//New function returns new DefaultDestinationResolver struct object
func New() DestinationResolver {
	return &DefaultDestinationResolver{}
}

//GetDestinationResolver returns destinationResolver object
func GetDestinationResolver() DestinationResolver {
	return dr
}

//InstallDestinationResolver function installs new plugin
func InstallDestinationResolver(name string, newFunc func() DestinationResolver) {
	DestinationResolverPlugins[name] = newFunc
	log.Printf("Installed DestinationResolver Plugin, name=%s", name)
}
func init() {
	DestinationResolverPlugins = make(map[string]func() DestinationResolver)
	dr = &DefaultDestinationResolver{}
	InstallDestinationResolver(DefaultPlugin, New)
}

//Init function reads config and initiates it
func Init() error {
	var name string
	if config.GetConfig().Plugin != nil {
		name = config.GetConfig().Plugin.DestinationResolver
	}
	if name == "" {
		name = DefaultPlugin
	}
	df, ok := DestinationResolverPlugins[name]
	if !ok {
		return ErrUnknownResolver
	}
	dr = df()
	return nil
}
