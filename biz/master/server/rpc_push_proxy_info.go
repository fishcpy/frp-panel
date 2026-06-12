package server

import (
	"github.com/fishcpy/frp-panel/models"
	"github.com/fishcpy/frp-panel/pb"
	"github.com/fishcpy/frp-panel/services/app"
	"github.com/fishcpy/frp-panel/services/dao"
)

func PushProxyInfo(ctx *app.Context, req *pb.PushProxyInfoReq) (*pb.PushProxyInfoResp, error) {
	var srv *models.ServerEntity
	var err error

	if srv, err = ValidateServerRequest(ctx, req.GetBase()); err != nil {
		return nil, err
	}

	if err = dao.NewMutation(ctx).AdminUpdateProxyStats(srv, req.GetProxyInfos()); err != nil {
		return nil, err
	}
	return &pb.PushProxyInfoResp{
		Status: &pb.Status{Code: pb.RespCode_RESP_CODE_SUCCESS, Message: "ok"},
	}, nil
}
