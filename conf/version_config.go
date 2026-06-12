package conf

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// VersionConfig 版本配置结构
type VersionConfig struct {
	Version             string `yaml:"version"`
	EnableUpgradeCheck  bool   `yaml:"enable_upgrade_check"`
	LatestVersion       string `yaml:"latest_version"`
}

var versionConfig *VersionConfig

// LoadVersionConfig 加载版本配置文件
func LoadVersionConfig(configPath string) error {
	if configPath == "" {
		configPath = "etc/version.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		// 如果配置文件不存在，使用默认配置
		if os.IsNotExist(err) {
			versionConfig = &VersionConfig{
				Version:             "",
				EnableUpgradeCheck:  true,
				LatestVersion:       "v0.1.0",
			}
			return nil
		}
		return err
	}

	var config VersionConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	versionConfig = &config
	return nil
}

// GetVersionConfig 获取版本配置
func GetVersionConfig() *VersionConfig {
	if versionConfig == nil {
		// 尝试从默认路径加载
		_ = LoadVersionConfig("")
	}
	if versionConfig == nil {
		// 如果还是 nil，返回默认配置
		return &VersionConfig{
			Version:             "",
			EnableUpgradeCheck:  true,
			LatestVersion:       "v0.1.0",
		}
	}
	return versionConfig
}

// InitVersionConfig 初始化版本配置
func InitVersionConfig() error {
	// 尝试从多个可能的路径加载
	paths := []string{
		"etc/version.yaml",
		"./etc/version.yaml",
		"../etc/version.yaml",
	}

	for _, path := range paths {
		absPath, _ := filepath.Abs(path)
		if _, err := os.Stat(absPath); err == nil {
			return LoadVersionConfig(absPath)
		}
	}

	// 如果都不存在，使用默认配置
	return LoadVersionConfig("")
}
