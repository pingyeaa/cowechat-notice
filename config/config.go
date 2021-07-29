package config

import (
	"cowechat-notice/types"
	"encoding/json"
	"os"
)

// GetConfig 获取配置
func GetConfig() (*types.Config, error) {
	var config *types.Config
	f, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
