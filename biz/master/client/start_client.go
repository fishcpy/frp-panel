package client

import (
	"context"

	"github.com/fishcpy/frp-panel/common"
	"github.com/fishcpy/frp-panel/pb"
	"github.com/fishcpy/frp-panel/services/app"
	"github.com/fishcpy/frp-panel/services/dao"
	"github.com/fishcpy/frp-panel/services/rpc"
	"github.com/fishcpy/frp-panel/utils/logger"
)

func StartFRPCHandler(ctx *app.Context, req *pb.StartFRPCRequest) (*pb.StartFRPCResponse, error) {
	logger.Logger(ctx).Infof("master get a start client request, origin is: [%+v]", req)

	userInfo := common.GetUserInfo(ctx)
	clientID := req.GetClientId()

	if !userInfo.Valid() {
		return &pb.StartFRPCResponse{
			Status: &pb.Status{Code: pb.RespCode_RESP_CODE_INVALID, Message: "invalid user"},
		}, nil
	}

	if len(clientID) == 0 {
		return &pb.StartFRPCResponse{
			Status: &pb.Status{Code: pb.RespCode_RESP_CODE_INVALID, Message: "invalid client id"},
		}, nil
	}

	cli, err := dao.NewQuery(ctx).GetClientByClientID(userInfo, clientID)
	if err != nil {
		return nil, err
	}

	client := cli.ClientEntity

	client.Stopped = false

	if err = dao.NewMutation(ctx).UpdateClient(userInfo, client); err != nil {
		return nil, err
	}

	go func() {
		resp, err := rpc.CallClient(app.NewContext(context.Background(), ctx.GetApp()), req.GetClientId(), pb.Event_EVENT_START_FRPC, req)
		if err != nil {
			logger.Logger(context.Background()).WithError(err).Errorf("start client event send to client error, client id: [%s]", req.GetClientId())
		}

		if resp == nil {
			logger.Logger(ctx).Errorf("cannot get response, client id: [%s]", req.GetClientId())
		}
	}()

	return &pb.StartFRPCResponse{
		Status: &pb.Status{Code: pb.RespCode_RESP_CODE_SUCCESS, Message: "ok"},
	}, nil
}
