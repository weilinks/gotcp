package main

import (
	"github.com/sirupsen/logrus"
	"github.com/weilinks/gotcp"
	"github.com/weilinks/gotcp/examples/fixture"
	"github.com/weilinks/gotcp/examples/tcp/owner_packet/common"
	"os"
	"os/signal"
	"syscall"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	easytcp.Log = log

	s := easytcp.NewServer(&easytcp.ServerOption{
		// specify codec and packer
		Codec:  &easytcp.JsonCodec{},
		Packer: &common.CustomPacker{},
	})

	s.AddRoute(common.MsgLogin, handler, fixture.RecoverMiddleware(log), logMiddleware)

	go func() {
		if err := s.Serve(fixture.ServerAddr); err != nil {
			log.Errorf("serve err: %s", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	if err := s.Stop(); err != nil {
		log.Errorf("server stopped err: %s", err)
	}
}

func handler(ctx easytcp.Context) {
	var data common.Login
	_ = ctx.Bind(&data)

	err := ctx.SetResponse(common.MsgLoginResp, &common.LoginResp{
		Result: 1,
		Value: 1,
		TxnNo: data.TxnNo,
		MsgType: &common.MsgType{
			MsgType: common.MsgLoginResp,
		},
		//Data:    fmt.Sprintf("%s:%d:%t", data.Key1, data.Key2, data.Key3),
	})
	if err != nil {
		log.Errorf("set response failed: %s", err)
	}
}

func logMiddleware(next easytcp.HandlerFunc) easytcp.HandlerFunc {
	return func(ctx easytcp.Context) {
		//fullSize := ctx.Request().MustGet("fullSize")
		log.Infof("recv request  | id:(%v) dataSize(%d) data: %s", ctx.Request().ID, len(ctx.Request().Data), ctx.Request().Data)

		defer func() {
			resp := ctx.Response()
			if resp != nil {
				log.Infof("send response | dataSize:(%d) id:(%v) data: %s", len(resp.Data), resp.ID, resp.Data)
			} else {
				log.Infof("don't send response since nil")
			}
		}()
		next(ctx)
	}
}
