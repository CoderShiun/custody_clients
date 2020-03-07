package main

type TopUpHist struct {
	Success bool              `json:"success"`
	Result  []TopUpHistResult `json:"result"`
}

type TopUpHistResult struct {
	ID             string `json:"id"`
	Coin           string `json:"coin"`
	Display_code   string `json:"display_code"`
	Description    string `json:"description"`
	Decimal        int    `json:"decimal"`
	Address        string `json:"address"`
	Source_address string `json:"source_address"`
	Side           string `json:"side"`
	Amount         string `json:"amount"`
	Abs_amount     string `json:"abs_amount"`
	Abs_cobo_fee   string `json:"abs_cobo_fee"`
	Txid           string `json:"txid"`
	Vout_n         int    `json:"vout_n"`
	Resquest_id    string `json:"resquest_id"`
	Status         string `json:"status"`
	Created_time   int64  `json:"created_time"`
	Last_time      int64  `json:"last_time"`
	Confirmed_num  int    `json:"confirmed_num"`
}

type NewAddress struct {
	Success bool             `json:"success"`
	Result  NewAddressResult `json:"result"`
}

type NewAddressResult struct {
	Coin    string `json:"coin"`
	Address string `json:"address"`
}

type WithdrawResult struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Result       string `json:"result"`
}

type WithdrawInfo struct {
	Success bool `json:"success"`
	Result WithdrawInfoResult `json:"result"`
	Error_code int `json:"error_code"`
	Error_message string `json:"error_message"`
	Error_description string `json:"error_description"`
}

type WithdrawInfoResult struct {
	ID                   string `json:"id"`
	Coin                 string `json:"coin"`
	Display_code         string `json:"display_code"`
	Description          string `json:"description"`
	Address              string `json:"address"`
	Source_address       string `json:"source_address"`
	Side                 string `json:"side"`
	Amount               string `json:"amount"`
	Decimal              int    `json:"decimal"`
	Abs_amount           string `json:"abs_amount"`
	Txid                 string `json:"txid"`
	Vout_n               int64 `json:"vout_n"`
	Resquest_id          string `json:"resquest_id"`
	Status               string `json:"status"`
	Created_time         int64  `json:"created_time"`
	Last_time            int64  `json:"last_time"`
	Memo                 string `json:"memo"`
	Confirming_threshold int64  `json:"confirming_threshold"`
	Confirmed_num        int64  `json:"confirmed_num"`
	Type                 string `json:"type"`
}