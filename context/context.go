package context

import "github.com/coinexchain/dex-api-go/models"

var DefaultApiContext ApiContext

type ApiContext struct {
	BaseReq models.BaseReq
}

func NewApiContext() ApiContext {
	return ApiContext{}
}

func NewBaseReq(from string) *models.BaseReq {
	fees := []*models.Coin{
		{
			Amount: 4000000,
			Denom:  nil,
		},
	}
	return models.BaseReq{
		AccountNumber: "",
		ChainID:       "coinexdex2",
		Fees:          nil,
		From:          "",
		Gas:           "1000000",
		GasAdjustment: "1.1",
		Memo:          "",
		Sequence:      "",
		Simulate:      false,
	}
}
