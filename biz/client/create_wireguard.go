package client

import (
	"github.com/fishcpy/frp-panel/defs"
	"github.com/fishcpy/frp-panel/pb"
	"github.com/fishcpy/frp-panel/services/app"
)

func CreateWireGuard(ctx *app.Context, req *pb.CreateWireGuardRequest) (*pb.CreateWireGuardResponse, error) {

	log := ctx.Logger().WithField("op", "CreateWireGuard")
	cfg := defs.WireGuardConfig{WireGuardConfig: req.GetWireguardConfig()}

	log.Debugf("create wireguard service, cfg: %s", cfg.String())

	wgSvc, err := ctx.GetApp().GetWireGuardManager().CreateService(&cfg)
	if err != nil {
		log.WithError(err).Errorf("create wireguard service failed")
		return nil, err
	}

	err = wgSvc.Start()
	if err != nil {
		log.WithError(err).Errorf("start wireguard service failed")
		// Remove the failed service from the manager to prevent it from being used
		if removeErr := ctx.GetApp().GetWireGuardManager().RemoveService(cfg.GetInterfaceName()); removeErr != nil {
			log.WithError(removeErr).Warnf("failed to remove failed wireguard service")
		}
		return nil, err
	}

	log.Debugf("start wireguard service success")

	return &pb.CreateWireGuardResponse{Status: &pb.Status{Code: pb.RespCode_RESP_CODE_SUCCESS, Message: "success"}}, nil
}
