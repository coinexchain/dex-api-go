# dex-api-go
dex链go语言api库，提供了dex链上所有的rest接口的封装，结合polarbear库可以方便的实现rest接口的组装，签名和调用。钱包开发者可以更加便捷的使用dex提供的功能。

#### 使用举例

**初始化签名用的keybase，获取交易发送者地址`from**`

```
bear.BearInit("keybase_path")
from := bear.GetAddress("key_name")
```

**初始化全局api上下文**

```
ctx := context.DefaultApiContext()
ctx.SetChainID("coinexdex2")
ctx.SetFromAddress(from)
ctx.SetName("key_name")
ctx.SetPassword("12345678")
```

在context中设置链ID，sender地址，sender的key_name， sender的密码等

另外还可以设置交易fees和gas

```
func (a *ApiContext) SetFees(amount string)
func (a *ApiContext) SetGas(amount string)
```

**组装rest接口参数**

以一笔普通转账为例，提供转账目标地址，转账数量，转账币种，解锁时间

```
unlockTime := "0"
to := "coinex1h6favnlytw3lgpy8cm6lcv530z0ctj6rplwt06"
transferredAmount := "400000000"
//refresh new account number and sequence first
ctx.RefreshAccNumAndSeq()
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
```

**发起签名和转账**

```
unsignedResp, err := ctx.Client.Bank.SendCoins(param)
if err == nil {
   bz, _ := unsignedResp.GetPayload().MarshalBinary()
   resp, err := ctx.SignAndBroadcast(bz)
}
```

