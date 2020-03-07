package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
)
import "time"
import "io"
import "strings"
import "crypto/hmac"
import "sort"
import "net/url"
import "io/ioutil"
import "encoding/hex"
import "crypto/sha256"
import "github.com/btcsuite/btcd/btcec"
import "net/http"

const HOST = "https://api.sandbox.cobo.com"
// Query Test
//const API_KEY = "03028cd5dad75c8434e19d94a5ef853f1088745dda9a701b763ef92582dce96df0"
//const API_SECRET = "c4096a2326b97a9b6c7b690bc8b52f85fc30373282c0a685c15a346c3c8b8015"

// Withdrawal
const API_KEY = "021d7ea6b0277ad20cb1c3638839ffefcf4eaee052dace2998f5ede063cf04b9e2"
const API_SECRET = "1f3a73c1aedb99585c14b92693e9d91111a18fec84e9c396d6be9a1bd8404ef2"


const SIG_TYPE = "ecdsa"

func GenerateRandomKeyPair() {
	apiSecret := make([]byte, 32)
	if _, err := rand.Read(apiSecret); err != nil {
		panic(err)
	}
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), apiSecret)
	apiKey := fmt.Sprintf("%x", privKey.PubKey().SerializeCompressed())
	apiSecretStr := fmt.Sprintf("%x", apiSecret)

	fmt.Printf("apiKey: %s, apiSecret: %s\n", apiKey, apiSecretStr)
}

func SortParams(params map[string]string) string {
	keys := make([]string, len(params))

	i := 0
	for k, _ := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	sorted := make([]string, len(params))
	i = 0
	for _, k := range keys {
		sorted[i] = k + "=" + url.QueryEscape(params[k])
		i++
	}

	return strings.Join(sorted, "&")
}

func Hash256(s string) string {
	hash_result := sha256.Sum256([]byte(s))
	hash_string := string(hash_result[:])
	return hash_string
}

func Hash256x2(s string) string {
	return Hash256(Hash256(s))
}

func SignHmac(message string) string {
	h := hmac.New(sha256.New, []byte(API_SECRET))
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func SignEcc(message string) string {
	api_secret, _ := hex.DecodeString(API_SECRET)
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), api_secret)
	sig, _ := privKey.Sign([]byte(Hash256x2(message)))
	return fmt.Sprintf("%x", sig.Serialize())
}

func VerifyEcc(body string, signature string, timestamp string, cobokey string) bool {
	api_key, _ := hex.DecodeString(cobokey)
	pubKey, _ := btcec.ParsePubKey(api_key, btcec.S256())

	sigBytes, _ := hex.DecodeString(signature)
	sigObj, _ := btcec.ParseSignature(sigBytes, btcec.S256())

	verified := sigObj.Verify([]byte(Hash256x2(body+"|"+timestamp)), pubKey)
	return verified
}

func Request(method string, path string, params map[string]string) string {
	client := &http.Client{}
	nonce := fmt.Sprintf("%d", time.Now().Unix()*1000)
	sorted := SortParams(params)
	var req *http.Request
	if method == "POST" {
		req, _ = http.NewRequest(method, HOST+path, strings.NewReader(sorted))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, _ = http.NewRequest(method, HOST+path+"?"+sorted, strings.NewReader(""))
	}
	content := strings.Join([]string{method, path, nonce, sorted}, "|")

	req.Header.Set("Biz-Api-Key", API_KEY)
	req.Header.Set("Biz-Api-Nonce", nonce)
	if SIG_TYPE == "hmac" {
		req.Header.Set("Biz-Api-Signature", SignHmac(content))
	} else if SIG_TYPE == "ecdsa" {
		req.Header.Set("Biz-Api-Signature", SignEcc(content))
	} else {
		fmt.Printf("Not supported signature type")
		return ""
	}

	resp, _ := client.Do(req)

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("Close Err: ", err)
		}
	}()

	body, _ := ioutil.ReadAll(resp.Body)

	if !VerifyEcc(string(body), resp.Header.Get("Biz_resp_signature"), resp.Header.Get("Biz-Timestamp"), "032f45930f652d72e0c90f71869dfe9af7d713b1f67dc2f7cb51f9572778b9c876") {
		return "Verify failed"
	}

	return string(body)
}

func main() {
	var get CoinBase
	get.coin = "TETH"
	get.address = "0xced5c00ccf7ff9784e11f15206c6b841fe528ad5"
	get.side = "deposit"

	var post CoinBase
	post.coin = "TETH"

	//GenerateRandomKeyPair()

	//fmt.Println(get.accountDetails())

	//fmt.Println(get.coinDetails())

	/*res := post.newDepositAddress()
	fmt.Println("New address: ", res)
	var newAddress NewAddress
	err := json.Unmarshal([]byte(res), &newAddress)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newAddress.Success)
	fmt.Println(newAddress.Result.Coin)
	fmt.Println(newAddress.Result.Address)*/

	//fmt.Println(get.verifyDeopsitAddress())

	//fmt.Println(get.verifyValidAddress())

	//fmt.Println(get.addressHistList())

	//fmt.Println(get.loopAddressDetails())

	//fmt.Println(get.transHistory())
	//get.transHistory()
	//fmt.Println(VerifyEcc(resp, "032f45930f652d72e0c90f71869dfe9af7d713b1f67dc2f7cb51f9572778b9c876"))

	//fmt.Println(get.pendingTransaction())

	//fmt.Println(get.pendingDepositDetails())

	res := get.withdrawalInformation("3")
	var withdrawalInfo WithdrawInfo
	err := json.Unmarshal([]byte(res), &withdrawalInfo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(withdrawalInfo.Result.Amount)
	fmt.Println(withdrawalInfo.Result.Status)
	fmt.Println(withdrawalInfo.Error_message)

	tm := time.Unix(1583620848763, 0)
	fmt.Println(tm)

	//fmt.Println(get.transDetails("20200115204400000323373000007145"))

	/*res := Request("GET", "/v1/custody/transaction_history/", map[string]string{
		"coin": "TETH",
		"side": "deposit",
		"begin_time": "0",
	})
	fmt.Printf("res: %v", res)

	var topUp TopUpHist
	err := json.Unmarshal([]byte(res), &topUp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	for i, _ := range topUp.Result {
		fmt.Println("Number: ", i)
		fmt.Println("Address: ", topUp.Result[i].Address)
		fmt.Println("Source Addr: ", topUp.Result[i].Source_address)
		fmt.Println("Begin Time: ", topUp.Result[i].Created_time)
		fmt.Println("AbsCoboFee: ", topUp.Result[i].Abs_cobo_fee)
	}

	fmt.Println("Results: ", len(topUp.Result))*/


	//fmt.Println("Results: ", topUp.Result[0].Created_time)

	/*res := Request("GET", "/v1/custody/address_info/", map[string]string{
		"coin": "ETH",
		"address": "0x544094588811118b7701cf4a9dea056e775b4b4e",
	})
	fmt.Printf("res: %v", res)*/

	//GenerateRandomKeyPair()


	/*res := post.submitWithdrawal("0xa1398b50E62Bf5512dd659198f1e225db7A8a41b","2500000000000000", "test9", time.Now().String())
	fmt.Println(res)
	var withdrawResult WithdrawResult
	err := json.Unmarshal([]byte(res), &withdrawResult)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(withdrawResult.Success)
	fmt.Println(withdrawResult.ErrorCode)
	fmt.Println(withdrawResult.ErrorMessage)*/


}
