package context

import (
	"fmt"
	dex "github.com/coinexchain/dex-api-go/client"
	"github.com/coinexchain/dex-api-go/client/auth"
	"github.com/coinexchain/dex-api-go/models"
	trans "github.com/go-openapi/runtime/client"
	"time"
)

var DefaultApiContext ApiContext = NewApiContext()
var DefaultTimeOut = 10 * time.Second
var DefaultFeesAmount string = "5000000"
var DefaultDenom string = "cet"
var DefaultGas string = "2000000"

type ApiContext struct {
	Name     string
	Password string
	BaseReq  models.BaseReq
	Client   *dex.CETLiteForCoinExChain
}

func (a *ApiContext) RefreshAccNumAndSeq() error {
	getAccountParam := auth.GetAccountParams{
		Address:    a.BaseReq.From,
		Context:    nil,
		HTTPClient: nil,
	}
	getAccountParam.SetTimeout(DefaultTimeOut)
	ok, _, err := a.Client.Auth.GetAccount(&getAccountParam)
	if err != nil {
		return err
	}
	seq := ok.Payload.Result.Sequence
	accountNum := ok.Payload.Result.AccountNumber
	a.BaseReq.AccountNumber = *accountNum
	a.BaseReq.Sequence = *seq
	fmt.Printf("seq:%s, accNum:%s\n", *seq, *accountNum)
	return nil
}

func NewApiContext() ApiContext {
	transport := trans.New("127.0.0.1:1317", "/", []string{"http"})
	cli := dex.Default
	cli.SetTransport(transport)
	req := NewBaseReq()
	return ApiContext{
		BaseReq: req,
		Client:  cli,
	}
}

func NewBaseReq() models.BaseReq {
	fees := []*models.Coin{
		{
			Amount: &DefaultFeesAmount,
			Denom:  &DefaultDenom,
		},
	}
	return models.BaseReq{
		AccountNumber: "",
		ChainID:       "coinexdex2",
		Fees:          fees,
		From:          "",
		Gas:           DefaultGas,
		GasAdjustment: "1.1",
		Memo:          "",
		Sequence:      "",
		Simulate:      false,
	}
}
