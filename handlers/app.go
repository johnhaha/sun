package handlers

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type AppInfo struct {
	Name   string `json:"name" yaml:"name"`
	Avatar string `json:"avatar" yaml:"avatar"`
	Des    string `json:"des" yaml:"des"`
}

func GetAppInfo() ([]AppInfo, error) {
	data, err := ioutil.ReadFile("assets/apps.yaml")
	if err != nil {
		return nil, err
	}
	var apps []AppInfo
	err = yaml.Unmarshal(data, &apps)
	if err != nil {
		return nil, err
	}
	return apps, nil
}
