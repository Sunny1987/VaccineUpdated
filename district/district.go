package district

import (
	"encoding/json"
	"fmt"
	"log"
	"vaccine1/variable"
)

//************************************************************
// STRUCTURES
//************************************************************

//Districts is the structure for districts response
type Districts struct {
	Districts []District `json:"districts"`
	TTL       int        `json:"ttl"`
}

//District is the individual district structure
type District struct {
	DistrictID   int    `json:"district_id"`
	DistrictName string `json:"district_name"`
}

//************************************************************
// METHODS
//************************************************************

//GetDistricts method is used to get the list of available Districts
func (dis *Districts) GetDistricts(stateId string) {
	variable.GetDist = variable.GetDist + stateId
	fmt.Printf("Token: %v", variable.Bearer)
	resp, err := variable.Client.R().SetHeader("Accept", "application/json").SetAuthToken(variable.Bearer).Get(variable.GetDist)
	if err != nil {
		log.Printf("ERROR in getting response: %v", err)
	}
	json.Unmarshal(resp.Body(), dis)

}
