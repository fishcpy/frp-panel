package proxy

import (
	"github.com/fishcpy/frp-panel/biz/master/client"
	"github.com/fishcpy/frp-panel/common"
	"github.com/fishcpy/frp-panel/models"
	"github.com/fishcpy/frp-panel/pb"
	"github.com/fishcpy/frp-panel/services/app"
	"github.com/fishcpy/frp-panel/services/dao"
	"github.com/fishcpy/frp-panel/utils/logger"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/samber/lo"
)

func StartProxy(ctx *app.Context, req *pb.StartProxyRequest) (*pb.StartProxyResponse, error) {

	var (
		userInfo  = common.GetUserInfo(ctx)
		clientID  = req.GetClientId()
		serverID  = req.GetServerId()
		proxyName = req.GetName()
	)

	clientEntity, err := GetClientWithMakeShadow(ctx, clientID, serverID)
	if err != nil {
		logger.Logger(ctx).WithError(err).Errorf("cannot get client, id: [%s]", clientID)
		return nil, err
	}

	_, err = dao.NewQuery(ctx).GetServerByServerID(userInfo, serverID)
	if err != nil {
		logger.Logger(ctx).WithError(err).Errorf("cannot get server, id: [%s]", serverID)
		return nil, err
	}

	proxyConfig, err := dao.NewQuery(ctx).GetProxyConfigByFilter(userInfo, &models.ProxyConfigEntity{
		ClientID: clientID,
		ServerID: serverID,
		Name:     proxyName,
	})
	if err != nil {
		logger.Logger(ctx).WithError(err).Errorf("cannot get proxy config, client: [%s], server: [%s], proxy name: [%s]", clientID, serverID, proxyName)
		return nil, err
	}

	// 1. 更新proxy状态
	proxyConfig.Stopped = false
	err = dao.NewMutation(ctx).UpdateProxyConfig(userInfo, proxyConfig)
	if err != nil {
		logger.Logger(ctx).WithError(err).Errorf("cannot update proxy config, client: [%s], server: [%s], proxy name: [%s]", clientID, serverID, proxyName)
		return nil, err
	}

	typedProxyConfig, err := proxyConfig.GetTypedProxyConfig()
	if err != nil {
		logger.Logger(ctx).WithError(err).Errorf("cannot get typed proxy config, client: [%s], server: [%s], proxy name: [%s]", clientID, serverID, proxyName)
		return nil, err
	}

	// 2. 添加 proxy到client
	if oldCfg, err := clientEntity.GetConfigContent(); err != nil {
		logger.Logger(ctx).WithError(err).Errorf("cannot get client config, id: [%s]", clientID)
		return nil, err
	} else {
		oldCfg.Proxies = lo.Filter(oldCfg.Proxies, func(proxy v1.TypedProxyConfig, _ int) bool {
			return proxy.GetBaseConfig().Name != proxyName
		})
		oldCfg.Proxies = append(oldCfg.Proxies, typedProxyConfig)

		if err := clientEntity.SetConfigContent(*oldCfg); err != nil {
			logger.Logger(ctx).WithError(err).Errorf("cannot set client config, id: [%s]", clientID)
			return nil, err
		}
	}

	// 3. 更新client的配置
	rawCfg, err := clientEntity.MarshalJSONConfig()
	if err != nil {
		logger.Logger(ctx).WithError(err).Errorf("cannot marshal client config, id: [%s]", clientID)
		return nil, err
	}

	_, err = client.UpdateFrpcHander(ctx, &pb.UpdateFRPCRequest{
		ClientId: &clientEntity.ClientID,
		ServerId: &serverID,
		Config:   rawCfg,
		Comment:  &clientEntity.Comment,
		FrpsUrl:  &clientEntity.FrpsUrl,
	})
	if err != nil {
		logger.Logger(ctx).WithError(err).Warnf("cannot update frpc, id: [%s]", clientID)
	}

	return &pb.StartProxyResponse{
		Status: &pb.Status{
			Code:    pb.RespCode_RESP_CODE_SUCCESS,
			Message: "start proxy success",
		},
	}, nil
}
