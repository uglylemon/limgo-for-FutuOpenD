package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Qot_GetStaticInfo"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(3202), "QotGetStaticInfo")

	var err error
	err = limgo.On("send.QotGetStaticInfo", QotGetStaticInfoSend)
	err = limgo.On("recv.QotGetStaticInfo", QotGetStaticInfoRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// QotGetStaticInfoSend handler
func QotGetStaticInfoSend(conn *limgo.Request, marketCode string, secType int32) error {
	ftpack := &limgo.FutuPack{}
	ftpack.SetProtoID(uint32(3202))

	market := transMarket(marketCode)

	reqData := &Qot_GetStaticInfo.Request{
		C2S: &Qot_GetStaticInfo.C2S{
			Market:  &market,
			SecType: &secType,
		},
	}
	pbData, err := proto.Marshal(reqData)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	ftpack.SetBody(pbData)
	err = conn.Send(ftpack)

	return err
}

// QotGetStaticInfoRecv handler
func QotGetStaticInfoRecv(data []byte) error {
	fut := &Qot_GetStaticInfo.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	// fmt.Println(fut)
	for _, info := range fut.S2C.StaticInfoList {
		fmt.Println(*info.Basic.Name, *info.Basic.Security.Code, *info.Basic.Security.Market)
	}

	return nil
}
