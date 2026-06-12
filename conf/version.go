package conf

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime"

	"github.com/fishcpy/frp-panel/pb"
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
	// 优先使用配置文件中的版本号，否则使用构建时注入的版本
	config := GetVersionConfig()
	version := gitVersion
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

// GetSystemVersion 获取系统配置的版本号（用于升级检查）
func GetSystemVersion() string {
	config := GetVersionConfig()
	if config == nil || config.Version == "" {
		return gitVersion
	}
	return config.Version
}

// NeedUpgrade 检查是否需要升级
func NeedUpgrade(currentVersion string) bool {
	config := GetVersionConfig()
	if config == nil || !config.EnableUpgradeCheck {
		return false
	}

	// dev-build 不提示升级
	if currentVersion == "" || currentVersion == "dev-build" {
		return false
	}

	systemVersion := GetSystemVersion()
	if systemVersion == "" || systemVersion == "dev-build" {
		return false
	}

	// 简单的版本比较：当前版本不等于系统版本时提示升级
	return currentVersion != systemVersion
}
