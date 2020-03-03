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
//const API_KEY = "02d9e82b388fe783d0667596dca16e0cbf75f48071f5d610367d94f0d8583b3b0d"
const API_KEY = "03028cd5dad75c8434e19d94a5ef853f1088745dda9a701b763ef92582dce96df0"
//const API_SECRET = "8143a7d49eeebb2822980641429feac5a14ff1ddde22bb99aa67fec08b24e2f8"
const API_SECRET = "c4096a2326b97a9b6c7b690bc8b52f85fc30373282c0a685c15a346c3c8b8015"
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

func VerifyEcc(message string, signature string) bool {
	api_key, _ := hex.DecodeString(API_KEY)
	pubKey, _ := btcec.ParsePubKey(api_key, btcec.S256())

	sigBytes, _ := hex.DecodeString(signature)
	sigObj, _ := btcec.ParseSignature(sigBytes, btcec.S256())

	verified := sigObj.Verify([]byte(Hash256x2(message)), pubKey)
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
	return string(body)
}

func main() {
	var get CoinBase
	get.coin = "TETH"
	get.address = "0xced5c00ccf7ff9784e11f15206c6b841fe528ad5"
	get.side = "deposit"

	//GenerateRandomKeyPair()

	//fmt.Println(get.accountDetails())

	//fmt.Println(get.coinDetails())

	//res := post.newDepositAddress()

	//fmt.Println(get.verifyDeopsitAddress())

	//fmt.Println(get.verifyValidAddress())

	//fmt.Println(get.addressHistList())

	//fmt.Println(get.loopAddressDetails())

	//fmt.Println(get.transHistory())

	//fmt.Println(get.pendingTransaction())

	//fmt.Println(get.pendingDepositDetails())

	//fmt.Println(get.withdrawalInformation("123"))

	res := Request("GET", "/v1/custody/transaction_history/", map[string]string{
		"coin": "TETH",
		"side": "deposit",
		"begin_time": "1579091977070",
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
	}

	fmt.Println("Results: ", len(topUp.Result))


	//fmt.Println("Results: ", topUp.Result[0].Created_time)

	/*res := Request("GET", "/v1/custody/address_info/", map[string]string{
		"coin": "ETH",
		"address": "0x544094588811118b7701cf4a9dea056e775b4b4e",
	})
	fmt.Printf("res: %v", res)*/

	//GenerateRandomKeyPair()
}
