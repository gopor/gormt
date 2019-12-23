package config

import (
	"fmt"
	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/tools"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// CfgBase base config struct
type CfgBase struct {
	SerialNumber       string `json:"serial_number" yaml:"serial_number"`             // version.版本号
	ServiceName        string `json:"service_name" yaml:"service_name"`               // service name .service名字
	ServiceDisplayName string `json:"service_displayname" yaml:"service_displayname"` // display name .显示名
	ServiceDesc        string `json:"service_desc" yaml:"service_desc"`               // service desc .service描述
	IsDev              bool   `json:"is_dev" yaml:"is_dev"`                           // Is it a development version?是否是开发版本
}

var _map = Config{}

func init() {
	onInit()
	dev.OnSetDev(_map.IsDev)
}

func onInit() {
	path := tools.GetModelPath()
	err := InitFile(path + "/config.yml")
	if err != nil {
		fmt.Println("InitFile: ", err.Error())
		return
	}
}

// InitFile default value from file .
func InitFile(filename string) error {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(bs, &_map); err != nil {
		fmt.Println("read toml error: ", err.Error())
		return err
	}

	return nil
}

// GetServiceConfig Get service configuration information
func GetServiceConfig() (name, displayName, desc string) {
	name = _map.ServiceName
	displayName = _map.ServiceDisplayName
	desc = _map.ServiceDesc
	return
}
