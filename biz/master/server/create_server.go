package server

import (
	"github.com/fishcpy/frp-panel/common"
	"github.com/fishcpy/frp-panel/models"
	"github.com/fishcpy/frp-panel/pb"
	"github.com/fishcpy/frp-panel/services/app"
	"github.com/fishcpy/frp-panel/services/dao"
	"github.com/fishcpy/frp-panel/utils"
	"github.com/google/uuid"
)

func InitServerHandler(c *app.Context, req *pb.InitServerRequest) (*pb.InitServerResponse, error) {
	var (
		userServerID = req.GetServerId()
		serverIP     = req.GetServerIp()
		userInfo     = common.GetUserInfo(c)
	)

	if !userInfo.Valid() {
		return &pb.InitServerResponse{
			Status: &pb.Status{Code: pb.RespCode_RESP_CODE_INVALID, Message: "invalid user"},
		}, nil
	}

	if len(userServerID) == 0 || len(serverIP) == 0 || !utils.IsClientIDPermited(userServerID) {
		return &pb.InitServerResponse{
			Status: &pb.Status{Code: pb.RespCode_RESP_CODE_INVALID, Message: "request invalid"},
		}, nil
	}

	globalServerID := app.GlobalClientID(userInfo.GetUserName(), "s", userServerID)

	if err := dao.NewMutation(c).CreateServer(userInfo,
		&models.ServerEntity{
			ServerID:      globalServerID,
			TenantID:      userInfo.GetTenantID(),
			UserID:        userInfo.GetUserID(),
			ConnectSecret: uuid.New().String(),
			ServerIP:      serverIP,
		}); err != nil {
		return nil, err
	}

	return &pb.InitServerResponse{
		Status:   &pb.Status{Code: pb.RespCode_RESP_CODE_SUCCESS, Message: "ok"},
		ServerId: &globalServerID,
	}, nil
}
