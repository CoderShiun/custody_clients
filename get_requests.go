package main

type GetFunc struct {
	coin    string
	address string
	side string
}

func (g *GetFunc) accountDetails() string {
	res := Request("GET", "/v1/custody/org_info/", map[string]string{})

	return res
}

func (g *GetFunc) coinDetails() string {
	res := Request("GET", "/v1/custody/coin_info/", map[string]string{
		"coin": g.coin,
	})

	return res
}

func (g *GetFunc) verifyDeopsitAddress() string {
	res := Request("GET", "/v1/custody/address_info/", map[string]string{
		"coin":    g.coin,
		"address": g.address,
	})

	return res
}

func (g *GetFunc) verifyValidAddress() string {
	return Request("GET", "/v1/custody/is_valid_address/", map[string]string{
		"coin": g.coin,
		"address": g.address,
	})
}

func (g *GetFunc) addressHistList() string {
	return Request("GET", "/v1/custody/address_history/", map[string]string{
		"coin": g.coin,
	})
}

func (g *GetFunc) loopAddressDetails() string {
	return Request("GET", "/v1/custody/internal_address_info/", map[string]string{
		"coin": g.coin,
		"address": g.address,
	})
}

func (g *GetFunc) transDetails() {
	return
}

func (g *GetFunc) transHistory() string {
	return Request("GET", "/v1/custody/transaction_history/", map[string]string{
		"coin": g.coin,
		"side": g.side,
	})
}

func (g *GetFunc) pendingTransaction() string {
	return Request("GET", "/v1/custody/pending_transactions/", map[string]string{
		"coin": g.coin,
		"side": g.side,
	})
}

func (g *GetFunc) pendingDepositDetails(id string) string {
	return Request("GET", "/v1/custody/deposit_info/", map[string]string{
		"id": id,
	})
}

func (g *GetFunc) withdrawalInformation(reqId string) string {
	return Request("GET", "/v1/custody/withdraw_info_by_request_id/", map[string]string{
		"request_id": reqId,
	})
}