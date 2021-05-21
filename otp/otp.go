package otp

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"vaccine1/variable"
)

//************************************************************
// STRUCTURES
//************************************************************

//GetOTPResp is used to get the TxnId for GetOTP method
type GetOTPResp struct {
	TxnId string
}

// type ConfirmMyOTPReq struct {
// 	Otp   string
// 	TxnId string
// }

//ConfirmOTPRespToken is used to get response for ConfirmMyOTP method
type ConfirmOTPRespToken struct {
	Token        string
	IsNewAccount string
}

//************************************************************
// METHODS
//************************************************************

//GetOTP is used to initiate OTP process
func (re *GetOTPResp) GetOTP(mobile string) {
	ot := map[string]string{"mobile": mobile}
	resp, err := variable.Client.R().SetHeader("Accept", "application/json").SetBody(ot).Post(variable.GetOTP)

	if err != nil {
		fmt.Printf("Error fetching otp: %v", err)
	}

	json.Unmarshal(resp.Body(), re)

}

//ConfirmMyOTP method is used to confirm the OTP received
func (conf *ConfirmOTPRespToken) ConfirmMyOTP(txnId string, otpV string) string {

	//log.Println(otpV)
	//log.Println(txnId)
	h := sha256.New()
	h.Write([]byte(otpV))
	otpData := hex.EncodeToString(h.Sum(nil))
	// otpData := BytesToString(h.Sum(nil))
	//log.Printf("otp: %s", otpData)
	otp := map[string]interface{}{"otp": otpData, "txnId": txnId}
	//log.Println(otp)

	resp, err := variable.Client.R().SetHeader("Accept", "application/json").SetBody(otp).Post(variable.ConfirmOTP)
	if err != nil {
		log.Printf("Error fetching otp: %v", err)
	}
	//log.Println(resp)
	json.Unmarshal(resp.Body(), conf)
	//log.Printf("re: %v", conf.Token)
	variable.Token = conf.Token
	variable.Bearer = variable.Bearer + variable.Token
	//log.Printf("Bearer: %v", Bearer)
	return variable.Bearer
}

//************************************************************
// FUNCTIONS
//************************************************************

//CheckAuthentication function is used to get the authentication status of user
func CheckAuthentication() bool {

	btarr := strings.Split(variable.Bearer, " ")
	//log.Println(btarr[1])
	if len(btarr) > 1 && btarr[1] != "" {

		fmt.Println()
		fmt.Println()
		fmt.Println("*****Already Authenticated****")
		fmt.Println()

		return true
	} else {
		return false
	}
}

//GetAuthenticated function is used to initiate the authentication for user
func GetAuthenticated() string {

	fmt.Println()
	fmt.Println()
	log.Println("*********Authentication Started*****************")
	fmt.Println()
	fmt.Println()

	var otp GetOTPResp
	var mob string
	var otpG string
	fmt.Println("Enter mobile (ex: 9xxxxxxxxx)")
	fmt.Print(">>")
	fmt.Scanf("%s\n", &mob)
	otp.GetOTP(mob)
	var conf ConfirmOTPRespToken
	fmt.Println("Enter OTP (Check you mobile for OTP)")
	fmt.Print(">>")
	fmt.Scanf("%s\n", &otpG)
	bt := conf.ConfirmMyOTP(otp.TxnId, otpG)
	return bt
}
