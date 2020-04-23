package main

import (
	"fmt"
	dex "github.com/coinexchain/dex-api-go/client"
	"github.com/coinexchain/dex-api-go/client/auth"
	"github.com/coinexchain/dex-api-go/client/bank"
	"github.com/coinexchain/dex-api-go/client/transactions"
	"github.com/coinexchain/dex-api-go/models"
	bear "github.com/coinexchain/polarbear"
	trans "github.com/go-openapi/runtime/client"
	"strconv"
	"time"
)

func main() {
	bear.BearInit("tmp")
	from := bear.GetAddress("lhr")
	fmt.Println(from)
	fmt.Println("coinex10hmcj9sp6gef5244wxkwt9jgweuwpp9fjcmwng")


	amount := "100000000"
	demom := "cet"
	unlockTime := "0"

	baseReq := models.BaseReq{
		AccountNumber:  *accountNum,
		ChainID:       "coinexdex-test1",
		Fees: []*models.Coin{
			{
				Amount: &amount,
				Denom:  &demom,
			},
		},
		From:          "coinex10hmcj9sp6gef5244wxkwt9jgweuwpp9fjcmwng",
		Gas:           "2000000",
		GasAdjustment: "1.1",
		Memo:          "",
		Sequence:      *seq,
		Simulate:      false,
	}

	param := &bank.SendCoinsParams{
		Account: bank.SendCoinsBody{
			Amount: []*models.Coin{
				&models.Coin{
					Amount: &amount,
					Denom:  &demom,
				},
			},
			BaseReq:    &baseReq,
			UnlockTime: &unlockTime,
		},
		Address:    "coinex1h6favnlytw3lgpy8cm6lcv530z0ctj6rplwt06",
		Context:    nil,
		HTTPClient: nil,
	}

	param.SetTimeout(time.Second)
	resp, err := cli.Bank.SendCoins(param)
	if err == nil {
		bz, _ := resp.GetPayload().MarshalBinary()
		an , _ := strconv.ParseUint(*accountNum,10,64)
		sq , _ := strconv.ParseUint(*seq,10,64)
		ret := bear.SignAndBuildBroadcast("lhr", "12345678", string(bz), "coinexdex-test1", "block", an, sq)
		fmt.Printf("%s\n", bz)
		fmt.Printf("%s\n", ret)
		var stdTx transactions.BroadcastTxBody
		_ = stdTx.UnmarshalBinary([]byte(ret))
		fmt.Println("hehe")
		fmt.Println(stdTx)
		param := transactions.BroadcastTxParams{
			TxBroadcast: stdTx,
			Context:     nil,
			HTTPClient:  nil,
		}
		param.SetTimeout(2*time.Second)
		ok, err := cli.Transactions.BroadcastTx(&param)
		fmt.Println(err)
		bz, err = ok.GetPayload().MarshalBinary()
		fmt.Printf("%s\n", bz)
	}

	fmt.Println(err)
}
