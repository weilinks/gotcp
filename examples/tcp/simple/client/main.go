package main

import (
	"github.com/weilinks/gotcp"
	"github.com/weilinks/gotcp/examples/fixture"
	"github.com/weilinks/gotcp/examples/tcp/simple/common"
	"github.com/weilinks/gotcp/message"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", fixture.ServerAddr)
	if err != nil {
		panic(err)
	}
	log := logrus.New()
	packer := easytcp.NewDefaultPacker()
	go func() {
		// write loop
		for {
			time.Sleep(time.Second)
			rawData := []byte("ping, ping, ping")
			msg := &message.Entry{
				ID:   common.MsgIdPingReq,
				Data: rawData,
			}
			packedMsg, err := packer.Pack(msg)
			if err != nil {
				panic(err)
			}
			if _, err := conn.Write(packedMsg); err != nil {
				panic(err)
			}
			log.Infof("snd >>> | id:(%d) size:(%d) data: %s", msg.ID, len(rawData), rawData)
		}
	}()
	go func() {
		// read loop
		for {
			msg, err := packer.Unpack(conn)
			if err != nil {
				panic(err)
			}
			log.Infof("rec <<< | id:(%d) size:(%d) data: %s", msg.ID, len(msg.Data), msg.Data)
		}
	}()
	select {}
}
