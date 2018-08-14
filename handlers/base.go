package handlers

import (
	"fmt"
	"limgo/futupb/Qot_Common"
	"strings"
)

func transMarket(code string) int32 {
	market := int32(0)

	switch code {
	case "HK":
		market = int32(1)
	case "US":
		market = int32(11)
	case "SH":
		market = int32(21)
	case "SZ":
		market = int32(22)
	}

	return market
}

func transStockCode(code string) *Qot_Common.Security {
	codes := strings.Split(code, ".")

	stock := &Qot_Common.Security{
		Market: new(int32),
		Code:   new(string),
	}

	*stock.Market = transMarket(codes[0])
	*stock.Code = codes[1]

	return stock
}

func transSubType(t string) int32 {
	k := "SubType_" + strings.Title(t)
	if v, ok := Qot_Common.SubType_value[k]; ok {
		return v
	}
	fmt.Println("transSubType error:", t)
	return int32(0)
}
