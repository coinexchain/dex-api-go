package context

import (
	"fmt"
	dex "github.com/coinexchain/dex-api-go/client"
	"github.com/coinexchain/dex-api-go/client/auth"
	"github.com/coinexchain/dex-api-go/client/transactions"
	"github.com/coinexchain/dex-api-go/models"
	bear "github.com/coinexchain/polarbear"
	trans "github.com/go-openapi/runtime/client"
	"strconv"
	"time"
)

var DefaultContext ApiContext = DefaultApiContext()
var DefaultTimeOut = 10 * time.Second
var DefaultFeesAmount string = "50000000"
var DefaultDenom string = "cet"
var DefaultGas string = "2000000"

type ApiContext struct {
	Name     string
	password string
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

func (a *ApiContext) SignAndBroadcast(signedTxBytes []byte) (string, error) {
	an, _ := strconv.ParseUint(a.BaseReq.AccountNumber, 10, 64)
	sq, _ := strconv.ParseUint(a.BaseReq.Sequence, 10, 64)
	ret := bear.SignAndBuildBroadcast(a.Name, a.password, string(signedTxBytes), a.BaseReq.ChainID, "block", an, sq)
	var stdTx transactions.BroadcastTxBody
	_ = stdTx.UnmarshalBinary([]byte(ret))
	param := transactions.BroadcastTxParams{
		TxBroadcast: stdTx,
		Context:     nil,
		HTTPClient:  nil,
	}
	param.SetTimeout(DefaultTimeOut)
	ok, err := a.Client.Transactions.BroadcastTx(&param)
	if err != nil {
		return "", err
	}
	bz, err := ok.GetPayload().MarshalBinary()
	fmt.Printf("%s\n", bz)
	return string(bz), nil
}

func (a *ApiContext) SetFees(amount string) {
	a.BaseReq.Fees[0].Amount = &amount
}

func (a *ApiContext) SetGas(gas string) {
	a.BaseReq.Gas = gas
}

func (a *ApiContext) SetFromAddress(from string) {
	a.BaseReq.From = from
}

func (a *ApiContext) SetName(name string) {
	a.Name = name
}

func (a *ApiContext) SetPassword(password string) {
	a.password = password
}

func (a *ApiContext) SetChainID(chainID string) {
	a.BaseReq.ChainID = chainID
}

func (a *ApiContext) SetNodeUrl(url string) {
	transport := trans.New(url, "/", []string{"http"})
	a.Client.SetTransport(transport)
}

func DefaultApiContext() ApiContext {
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
