package conf

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime"

	"github.com/VaalaCat/frp-panel/pb"
)

var (
	gitVersion = "dev-build"
	gitCommit  = ""
	gitBranch  = ""
	buildDate  = "1970-01-01T00:00:00Z"
)

type VersionInfo struct {
	GitVersion string `json:"gitVersion" yaml:"gitVersion"`
	GitCommit  string `json:"gitCommit" yaml:"gitCommit"`
	GitBranch  string `json:"gitBranch" yaml:"gitBranch"`
	BuildDate  string `json:"buildDate" yaml:"buildDate"`
	GoVersion  string `json:"goVersion" yaml:"goVersion"`
	Compiler   string `json:"compiler" yaml:"compiler"`
	Platform   string `json:"platform" yaml:"platform"`
}

func (v *VersionInfo) String() string {
	tempStr := "BinVersion: {{.GitVersion}}\nGitCommit: {{.GitCommit}}\nBuildDate: {{.BuildDate}}\nGoVersion: {{.GoVersion}}\nCompiler: {{.Compiler}}\nPlatform: {{.Platform}}"
	temp, err := template.New("version").Parse(tempStr)
	if err != nil {
		return ""
	}
	var result bytes.Buffer
	err = temp.Execute(&result, v)
	if err != nil {
		return ""
	}
	return result.String()
}

func (v *VersionInfo) ToProto() *pb.ClientVersion {
	return &pb.ClientVersion{
		GitVersion: v.GitVersion,
		GitCommit:  v.GitCommit,
		GitBranch:  v.GitBranch,
		BuildDate:  v.BuildDate,
		GoVersion:  v.GoVersion,
		Compiler:   v.Compiler,
		Platform:   v.Platform,
	}
}

func GetVersion() *VersionInfo {
	version := gitVersion

	// 如果配置文件中设置了自定义版本，则使用配置文件的版本
	config := GetVersionConfig()
	if config != nil && config.Version != "" {
		version = config.Version
	}

	return &VersionInfo{
		GitVersion: version,
		GitCommit:  gitCommit,
		GitBranch:  gitBranch,
		BuildDate:  buildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// NeedUpgrade 检查是否需要升级
func NeedUpgrade(currentVersion string) bool {
	config := GetVersionConfig()
	if config == nil || !config.EnableUpgradeCheck {
		return false
	}

	// 简单的版本比较：如果当前版本不等于最新版本，则提示升级
	// 这里可以使用更复杂的语义版本比较
	if currentVersion == "" || currentVersion == "dev-build" {
		return false
	}

	return currentVersion != config.LatestVersion
}

// GetLatestVersion 获取配置的最新版本
func GetLatestVersion() string {
	config := GetVersionConfig()
	if config == nil {
		return ""
	}
	return config.LatestVersion
}
