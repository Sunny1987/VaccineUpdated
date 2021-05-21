package pin

import (
	"encoding/json"
	"fmt"
	"log"
	"vaccine1/variable"
)

//************************************************************
// STRUCTURES
//************************************************************

//CentersByPIN is the structure for CentersByPIN response
type CentersByPIN struct {
	Sessions []CenterByPIN `json:"sessions"`
}

//CenterByPIN is the individual CenterByPIN structure
type CenterByPIN struct {
	CenterID               int      `json:"center_id"`
	Name                   string   `json:"name"`
	NameL                  string   `json:"name_l"`
	Address                string   `json:"address"`
	AddressL               string   `json:"address_l"`
	StateName              string   `json:"state_name"`
	StateNameL             string   `json:"state_name_l"`
	DistrictName           string   `json:"district_name"`
	DistrictNameL          string   `json:"district_name_l"`
	BlockName              string   `json:"block_name"`
	BlockNameL             string   `json:"block_name_l"`
	Pincode                string   `json:"pincode"`
	Lat                    float64  `json:"lat"`
	Long                   float64  `json:"long"`
	From                   string   `json:"from"`
	To                     string   `json:"to"`
	FeeType                string   `json:"fee_type"`
	Fee                    string   `json:"fee"`
	SessionID              string   `json:"session_id"`
	Date                   string   `json:"date"`
	AvailableCapacity      int      `json:"available_capacity"`
	AvailableCapacityDose1 int      `json:"available_capacity_dose1"`
	AvailableCapacityDose2 int      `json:"available_capacity_dose2"`
	MinAgeLimit            int      `json:"min_age_limit"`
	Vaccine                string   `json:"vaccine"`
	Slots                  []string `json:"slots"`
}

//************************************************************
// METHODS
//************************************************************

//GetCentersByPIN method is used to get the list of available Centers
func (pin *CentersByPIN) GetCentersByPIN(pinCode string, date string) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	// log.Printf("pincode : %s\n", pinCode)
	// log.Printf("date : %s\n", date)
	resp, err := variable.Client.R().SetQueryParams(map[string]string{
		"pincode": pinCode,
		"date":    date,
	}).SetHeader("Accept", "application/json").SetAuthToken(variable.Bearer).Get(variable.FindByPin)
	if err != nil {
		log.Printf("ERROR in getting response: %v", err)
	}
	//log.Println(resp.Body())
	json.Unmarshal(resp.Body(), pin)
}
