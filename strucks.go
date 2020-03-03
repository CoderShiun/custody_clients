package main

type TopUpHist struct {
	Success bool `json:"success"`
	Result []TopUpHistResult `json:"result"`
}

type TopUpHistResult struct {
	ID string `json:"id"`
	Coin string `json:"coin"`
	Display_code string `json:"display_code"`
	Description string `json:"description"`
	Decimal int `json:"decimal"`
	Address string `json:"address"`
	Source_address string `json:"source_address"`
	Side string `json:"side"`
	Amount string `json:"amount"`
	Abs_amount string `json:"abs_amount"`
	Txid string `json:"txid"`
	Vout_n int `json:"vout_n"`
	Resquest_id string `json:"resquest_id"`
	Status string `json:"status"`
	Created_time int64 `json:"created_time"`
	Last_time int64 `json:"last_time"`
	Confirmed_num int `json:"confirmed_num"`
}

