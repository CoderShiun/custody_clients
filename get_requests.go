package main

type CoinBase struct {
	coin    string
	address string
	side string
}

func (g *CoinBase) accountDetails() string {
	res := Request("GET", "/v1/custody/org_info/", map[string]string{})

	return res
}

func (g *CoinBase) coinDetails() string {
	res := Request("GET", "/v1/custody/coin_info/", map[string]string{
		"coin": g.coin,
	})

	return res
}

func (g *CoinBase) verifyDeopsitAddress() string {
	res := Request("GET", "/v1/custody/address_info/", map[string]string{
		"coin":    g.coin,
		"address": g.address,
	})

	return res
}

func (g *CoinBase) verifyValidAddress() string {
	return Request("GET", "/v1/custody/is_valid_address/", map[string]string{
		"coin": g.coin,
		"address": g.address,
	})
}

func (g *CoinBase) addressHistList() string {
	return Request("GET", "/v1/custody/address_history/", map[string]string{
		"coin": g.coin,
	})
}

func (g *CoinBase) loopAddressDetails() string {
	return Request("GET", "/v1/custody/internal_address_info/", map[string]string{
		"coin": g.coin,
		"address": g.address,
	})
}

func (g *CoinBase) transDetails() {
	return
}

func (g *CoinBase) transHistory() string {
	return Request("GET", "/v1/custody/transaction_history/", map[string]string{
		"coin": g.coin,
		"side": g.side,
	})
}

func (g *CoinBase) pendingTransaction() string {
	return Request("GET", "/v1/custody/pending_transactions/", map[string]string{
		"coin": g.coin,
		"side": g.side,
	})
}

func (g *CoinBase) pendingDepositDetails(id string) string {
	return Request("GET", "/v1/custody/deposit_info/", map[string]string{
		"id": id,
	})
}

func (g *CoinBase) withdrawalInformation(reqId string) string {
	return Request("GET", "/v1/custody/withdraw_info_by_request_id/", map[string]string{
		"request_id": reqId,
	})
}