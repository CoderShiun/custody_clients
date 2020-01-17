package main

func (p *CoinBase)newDepositAddress() string {
	res := Request("POST", "/v1/custody/new_address/", map[string]string{
		"coin": p.coin,
	})

	return res
}

func (p *CoinBase)submitWithdrawal(amount, reqId, memo string) string {
	return Request("POST", "/v1/custody/new_withdraw_request/", map[string]string{
		"coin": p.coin,
		"address": p.address,
		"amount": amount,
		"request_id": reqId,
		"memo": memo,
	})
}

