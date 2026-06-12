package worker

import (
	"fmt"

	"github.com/fishcpy/frp-panel/common"
	"github.com/fishcpy/frp-panel/models"
	"github.com/fishcpy/frp-panel/pb"
	"github.com/fishcpy/frp-panel/services/app"
	"github.com/fishcpy/frp-panel/services/dao"
	"github.com/fishcpy/frp-panel/utils/logger"
	"github.com/samber/lo"
)

func GetWorker(ctx *app.Context, req *pb.GetWorkerRequest) (*pb.GetWorkerResponse, error) {
	logger.Logger(ctx).Infof("get worker req: %s", req.String())
	var (
		workerID = req.GetWorkerId()
		userInfo = common.GetUserInfo(ctx)
	)

	if len(workerID) == 0 {
		logger.Logger(ctx).Errorf("worker id is empty")
		return nil, fmt.Errorf("worker id is empty")
	}

	workerRecord, err := dao.NewQuery(ctx).GetWorkerByWorkerID(userInfo, workerID)
	if err != nil {
		logger.Logger(ctx).WithError(err).Errorf("get worker by id failed")
		return nil, err
	}

	return &pb.GetWorkerResponse{
		Status: &pb.Status{
			Code:    pb.RespCode_RESP_CODE_SUCCESS,
			Message: "ok",
		},
		Worker: workerRecord.ToPB(),
		Clients: lo.Map(workerRecord.Clients, func(client models.Client, index int) *pb.Client {
			c := client.ToPB()
			c.Config = nil
			c.Secret = nil
			return c
		}),
	}, nil
}
