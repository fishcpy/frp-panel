package defs

import "github.com/fishcpy/frp-panel/pb"

type Connector struct {
	CliID   string
	Conn    pb.Master_ServerSendServer
	CliType string
}
