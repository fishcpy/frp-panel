package platform

import (
	"fmt"
	"os"

	"github.com/fishcpy/frp-panel/common"
	"github.com/fishcpy/frp-panel/conf"
	"github.com/fishcpy/frp-panel/defs"
	"github.com/fishcpy/frp-panel/pb"
	"github.com/fishcpy/frp-panel/services/app"
	"github.com/fishcpy/frp-panel/services/dao"
	"github.com/gin-gonic/gin"
)

func GetPlatformInfo(appInstance app.Application) func(*gin.Context) {
	return func(c *gin.Context) {
		resp, err := getPlatformInfo(appInstance, c)
		if err != nil {
			common.ErrResp(c, &pb.CommonResponse{Status: &pb.Status{Code: pb.RespCode_RESP_CODE_INVALID, Message: err.Error()}}, err.Error())
			return
		}
		common.OKResp(c, resp)
	}
}

func getPlatformInfo(appInstance app.Application, c *gin.Context) (*pb.GetPlatformInfoResponse, error) {
	appCtx := app.NewContext(c, appInstance)
	userInfo := common.GetUserInfo(c)
	if !userInfo.Valid() {
		return &pb.GetPlatformInfoResponse{
			Status: &pb.Status{Code: pb.RespCode_RESP_CODE_INVALID, Message: "invalid user"},
		}, nil
	}
	totalServers, err := dao.NewQuery(appCtx).CountServers(userInfo)
	if err != nil {
		return nil, err
	}
	totalClients, err := dao.NewQuery(appCtx).CountClients(userInfo)
	if err != nil {
		return nil, err
	}

	configuredServers, err := dao.NewQuery(appCtx).CountConfiguredServers(userInfo)
	if err != nil {
		return nil, err
	}
	configuredClients, err := dao.NewQuery(appCtx).CountConfiguredClients(userInfo)
	if err != nil {
		return nil, err
	}

	unconfiguredServers := totalServers - configuredServers

	unconfiguredClients := totalClients - configuredClients

	// 从环境变量读取外部客户端 URL，而不是从 appInstance.GetConfig()
	// 避免被内部 server 的本地地址配置覆盖
	clientRPCUrl := os.Getenv(defs.EnvClientRPCUrl)
	clientAPIUrl := os.Getenv(defs.EnvClientAPIUrl)

	if len(clientRPCUrl) == 0 {
		clientRPCUrl = fmt.Sprintf("grpc://%s:%d", appInstance.GetConfig().Master.RPCHost, appInstance.GetConfig().Master.RPCPort)
	}

	if len(clientAPIUrl) == 0 {
		clientAPIUrl = fmt.Sprintf("%s://%s:%d", appInstance.GetConfig().Master.APIScheme, appInstance.GetConfig().Master.RPCHost, appInstance.GetConfig().Master.APIPort)
	}

	// 获取服务器版本信息
	versionInfo := conf.GetVersion()
	systemVersion := conf.GetSystemVersion()
	versionConfig := conf.GetVersionConfig()

	return &pb.GetPlatformInfoResponse{
		Status:                  &pb.Status{Code: pb.RespCode_RESP_CODE_SUCCESS, Message: "ok"},
		TotalClientCount:        int32(totalClients),
		TotalServerCount:        int32(totalServers),
		UnconfiguredClientCount: int32(unconfiguredClients),
		UnconfiguredServerCount: int32(unconfiguredServers),
		ConfiguredClientCount:   int32(configuredClients),
		ConfiguredServerCount:   int32(configuredServers),
		MasterRpcHost:           appInstance.GetConfig().Master.RPCHost,
		MasterRpcPort:           int32(appInstance.GetConfig().Master.RPCPort),
		MasterApiPort:           int32(appInstance.GetConfig().Master.APIPort),
		MasterApiScheme:         appInstance.GetConfig().Master.APIScheme,
		ClientRpcUrl:            clientRPCUrl,
		ClientApiUrl:            clientAPIUrl,
		GithubProxyUrl:          appInstance.GetConfig().App.GithubProxyUrl + "|" + versionInfo.GitVersion + "|" + systemVersion + "|" + fmt.Sprintf("%v", versionConfig.EnableUpgradeCheck),
	}, nil
}
