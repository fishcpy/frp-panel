package proxy

import (
	"context"
	"fmt"

	"github.com/fishcpy/frp-panel/common"
	"github.com/fishcpy/frp-panel/pb"
	"github.com/fishcpy/frp-panel/services/app"
	"github.com/fishcpy/frp-panel/services/dao"
	"github.com/fishcpy/frp-panel/utils/logger"
)

// GetProxyStatsByClientID get proxy info by client id
func GetProxyStatsByClientID(c *app.Context, req *pb.GetProxyStatsByClientIDRequest) (*pb.GetProxyStatsByClientIDResponse, error) {
	logger.Logger(c).Infof("get proxy by client id, req: [%+v]", req)
	var (
		clientID = req.GetClientId()
		userInfo = common.GetUserInfo(c)
	)

	if len(clientID) == 0 {
		return nil, fmt.Errorf("request invalid")
	}

	proxyStatsList, err := dao.NewQuery(c).GetProxyStatsByClientID(userInfo, clientID)
	if proxyStatsList == nil || err != nil {
		logger.Logger(context.Background()).WithError(err).Errorf("cannot get proxy, client id: [%s]", clientID)
		return nil, err
	}
	return &pb.GetProxyStatsByClientIDResponse{
		Status:     &pb.Status{Code: pb.RespCode_RESP_CODE_SUCCESS, Message: "ok"},
		ProxyInfos: convertProxyStatsList(proxyStatsList),
	}, nil
}
