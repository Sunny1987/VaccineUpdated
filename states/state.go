package states

import (
	"encoding/json"
	"log"
	"vaccine1/variable"
)

//************************************************************
// STRUCTURES
//************************************************************

//States is the structure for states response
type States struct {
	States []State `json:"states"`
	TTL    int     `json:"ttl"`
}

//State is the individual state structure
type State struct {
	StateID   int    `json:"state_id"`
	StateName string `json:"state_name"`
}

//************************************************************
// METHODS
//************************************************************

//GetStates method is used to get the list of available states
func (sts *States) GetStates() {
	//resp, err := Client.R().SetHeader("Accept", "application/json").SetAuthToken(Bearer).Get(GETSTATES)
	resp, err := variable.Client.R().SetHeader("Accept", "application/json").SetAuthToken(variable.Bearer).Get(variable.GETSTATES)
	if err != nil {
		log.Printf("ERROR in getting response: %v", err)
	}
	json.Unmarshal(resp.Body(), sts)

}
