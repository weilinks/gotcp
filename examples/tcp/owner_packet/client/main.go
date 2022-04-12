package main

import (
	"github.com/weilinks/gotcp"
	"github.com/weilinks/gotcp/examples/fixture"
	"github.com/weilinks/gotcp/examples/tcp/owner_packet/common"
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
	codec := &easytcp.JsonCodec{}
	packer := &common.CustomPacker{}
	go func() {
		// write loop
		for {
			time.Sleep(time.Second)
			req := common.Login{
				DevType: common.EXCHANGE_CABINET,
				DevId: "CHZD12XHSO190923001",
				HardVersion: "HDTTV234",
				Imsi: "460034200879463",
				Ccid: "89860412101870601575",
				Imei: "863517023000770",
				ProtocolVersion: "V2",
				SoftVersion: "3.0.3_ONE",
				CabSta: common.RETURN,
				TxnNo: "1591646735703",
				MsgType: &common.MsgType{
					MsgType: common.MsgLogin,
				},
			}
			data, err := codec.Encode(req)
			if err != nil {
				panic(err)
			}
			msg := &message.Entry{
				ID:   "json01-req",
				Data: data,
			}
			packedMsg, err := packer.Pack(msg)
			if err != nil {
				panic(err)
			}
			if _, err := conn.Write(packedMsg); err != nil {
				panic(err)
			}
		}
	}()
	go func() {
		// read loop
		for {
			msg, err := packer.Unpack(conn)

			if err != nil {
				log.Errorf("read resp err: %s", err)
				//panic(err)
			} else {
				log.Infof("rec <<< | id:(%d) size:(%d) data: %s", msg.ID, len(msg.Data), msg.Data)
			}
		}
	}()
	select {}
}
