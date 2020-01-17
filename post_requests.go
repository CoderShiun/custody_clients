package main

func newDepositAddress(coin string) string {
	res := Request("POST", "/v1/custody/new_address/", map[string]string{
		"coin": coin,
	})

	return res
}