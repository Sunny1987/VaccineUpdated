package center

import (
	"encoding/json"
	"fmt"
	"log"
	"vaccine1/variable"
)

//************************************************************
// STRUCTURES
//************************************************************

//Centers is the structure for Centers response
type Centers struct {
	Sessions []Session `json:"sessions"`
}

//Session is the individual Session structure
type Session struct {
	CenterID               int      `json:"center_id"`
	Name                   string   `json:"name"`
	Address                string   `json:"address"`
	StateName              string   `json:"state_name"`
	DistrictName           string   `json:"district_name"`
	BlockName              string   `json:"block_name"`
	Pincode                int      `json:"pincode"`
	From                   string   `json:"from"`
	To                     string   `json:"to"`
	Lat                    int      `json:"lat"`
	Long                   int      `json:"long"`
	FeeType                string   `json:"fee_type"`
	SessionID              string   `json:"session_id"`
	Date                   string   `json:"date"`
	AvailableCapacityDose1 int      `json:"available_capacity_dose1"`
	AvailableCapacityDose2 int      `json:"available_capacity_dose2"`
	AvailableCapacity      int      `json:"available_capacity"`
	Fee                    string   `json:"fee"`
	MinAgeLimit            int      `json:"min_age_limit"`
	Vaccine                string   `json:"vaccine"`
	Slots                  []string `json:"slots"`
}

//************************************************************
// METHODS
//************************************************************

//GetCenters method is used to get the list of available Centers
func (c *Centers) GetCenters(district_id string, date string, vaccine string) {

	fmt.Println()

	resp, err := variable.Client.R().SetQueryParams(map[string]string{
		"district_id": district_id,
		"date":        date,
		"vaccine":     vaccine,
	}).SetHeader("Accept", "application/json").SetAuthToken(variable.Bearer).Get(variable.GetCenters)
	if err != nil {
		log.Printf("ERROR in getting response: %v", err)
	}
	json.Unmarshal(resp.Body(), c)
}
