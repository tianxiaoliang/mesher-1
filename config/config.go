package config

import (
	"github.com/ServiceComb/go-chassis/core/archaius"
	"github.com/ServiceComb/go-chassis/core/config"
	"github.com/ServiceComb/go-chassis/core/config/model"
	"github.com/ServiceComb/go-chassis/core/lager"
	"github.com/ServiceComb/go-chassis/core/server"
	"github.com/ServiceComb/go-chassis/pkg/util/fileutil"
	"github.com/go-chassis/mesher/cmd"
	"github.com/go-chassis/mesher/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

//Constant for mesher conf file
const (
	ConfFile = "mesher.yaml"
)

//Mode is of type string which gives mode of mesher deployment
var Mode string
var mesherConfig *MesherConfig

//GetConfig returns mesher config
func GetConfig() *MesherConfig {
	return mesherConfig
}

//SetConfig sets new mesher config from input config
func SetConfig(nc *MesherConfig) {
	if mesherConfig == nil {
		mesherConfig = &MesherConfig{}
	}
	*mesherConfig = *nc
}

//GetConfigFilePath returns config file path
func GetConfigFilePath() (string, error) {
	if cmd.Configs.ConfigFile == "" {
		wd, err := fileutil.GetWorkDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(wd, "conf", ConfFile), nil
	}
	return cmd.Configs.ConfigFile, nil
}

//InitProtocols initiates protocols
func InitProtocols() error {
	// todo if sdk init failed, do not call the data
	if len(config.GlobalDefinition.Cse.Protocols) == 0 {
		config.GlobalDefinition.Cse.Protocols = map[string]model.Protocol{
			common.HTTPProtocol: {Listen: "127.0.0.1:30101"},
		}

		return server.Init()
	}
	return nil
}

//Init reads config and initiates
func Init() error {
	mesherConfig = &MesherConfig{}
	contents, err := GetConfigContents(ConfFile)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal([]byte(contents), mesherConfig); err != nil {
		return err
	}

	return nil
}

//GetConfigContents returns config contents
func GetConfigContents(key string) (string, error) {
	f, err := GetConfigFilePath()
	if err != nil {
		return "", err
	}
	var contents string
	//route rule yaml file's content is value of a key
	//So read from config center first,if it is empty, Try to set file content into memory key value
	contents = archaius.GetString(key, "")
	if contents == "" {
		contents = SetKeyValueByFile(key, f)
	}
	return contents, nil
}

//SetKeyValueByFile reads mesher.yaml and gets key and value
func SetKeyValueByFile(key, f string) string {
	var contents string
	if _, err := os.Stat(f); err != nil {
		lager.Logger.Warn(err.Error(), nil)
		return ""
	}
	b, err := ioutil.ReadFile(f)
	if err != nil {
		lager.Logger.Error("Can not read mesher.yaml", err)
		return ""
	}
	contents = string(b)
	archaius.AddKeyValue(key, contents)
	return contents
}
