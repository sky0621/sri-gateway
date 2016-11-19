package server

import (
	"log"
	"net"

	gw "github.com/sky0621/sri-gateway"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Run ...
func Run() error {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	gw.RegisterMessageServiceServer(s, &RegisterMessageServiceServerImpl{})

	s.Serve(l)
	return nil
}

// RegisterMessageServiceServerImpl ...
type RegisterMessageServiceServerImpl struct{}

// SendMessage ...
func (i *RegisterMessageServiceServerImpl) SendMessage(ctx context.Context, msg *gw.SriMessage) (*gw.SriResponse, error) {
	log.Println(msg.MsgId)
	log.Println(msg.TextMsg)
	res := &gw.SriResponse{
		Ok:         true,
		ErrMessage: "",
		MsgId:      0,
	}
	return res, nil
}
