package main

import (
	"fmt"
	"github.com/coinexchain/dex-api-go/client/bank"
	"github.com/coinexchain/dex-api-go/context"
	"github.com/coinexchain/dex-api-go/models"
	bear "github.com/coinexchain/polarbear"
)

func main() {

	//init keybase
	bear.BearInit("tmp")
	from := bear.GetAddress("cat")

	//init api context
	ctx := context.DefaultApiContext()
	ctx.SetChainID("coinexdex-test1")
	ctx.SetFromAddress(from)
	ctx.SetName("cat")
	ctx.SetPassword("12345678")

	//prepare params
	unlockTime := "0"
	to := "coinex10c22dwn7hxps77tnkpj5pzu9zcpq5zf76xms55"
	transferredAmount := "1000000"
	err := ctx.RefreshAccNumAndSeq()
	if err != nil {
		panic(err)
	}
	param := &bank.SendCoinsParams{
		Account: bank.SendCoinsBody{
			Amount: []*models.Coin{
				{
					Amount: &transferredAmount,
					Denom:  &context.DefaultDenom,
				},
			},
			BaseReq:    &ctx.BaseReq,
			UnlockTime: &unlockTime,
		},
		Address: to,
	}
	param.SetTimeout(context.DefaultTimeOut)

	//transfer coins
	resp, err := ctx.Client.Bank.SendCoins(param)
	if err == nil {
		bz, _ := resp.GetPayload().MarshalBinary()
		resp, err := ctx.SignAndBroadcast(bz)
		fmt.Printf("response:%s error:%v\n", resp, err)
	}
}
