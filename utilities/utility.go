package utilities

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"vaccine1/center"
	"vaccine1/district"
	"vaccine1/pin"
	"vaccine1/states"
)

//GetStateData method gets the list of all the states
func GetStateData() {
	var states states.States
	states.GetStates()
	fmt.Println()
	fmt.Println("**************************************************")
	for _, state := range states.States {
		fmt.Printf("State: %s || StateId : %s\n", state.StateName, strconv.Itoa(state.StateID))
	}
	fmt.Println("**************************************************")
	fmt.Println()
	fmt.Println()
}

//GetDistrictData method gets the list of all the Districts
func GetDistrictData() {
	fmt.Println("Enter State Id (get state id from Opt.1)")
	fmt.Print(">>")
	var stId string
	fmt.Scanf("%s\n", &stId)
	var districts district.Districts
	districts.GetDistricts(stId)
	fmt.Println()
	fmt.Println("**************************************************")
	for _, district := range districts.Districts {

		fmt.Printf("District: %s\n", district.DistrictName)
	}
	fmt.Println("**************************************************")
	fmt.Println()
	fmt.Println()
}

//DisplayCenters displayes the center details
func DisplayCenters(cen *center.Centers) {
	if len(cen.Sessions) == 0 {
		log.Println("No Vaccines available for the location")
	} else {
		for _, session := range cen.Sessions {

			fmt.Println()
			fmt.Println()
			fmt.Println("**************************************************************************")
			fmt.Printf("Date: %v\n", session.Date)
			fmt.Printf("Address: %v\n", session.Address)
			fmt.Printf("Name: %v\n", session.Name)
			fmt.Printf("State: %v\n", session.StateName)
			fmt.Printf("District: %v\n", session.DistrictName)
			fmt.Printf("Block: %v\n", session.BlockName)
			fmt.Printf("Block: %v\n", session.BlockName)
			fmt.Printf("PIN: %v\n", session.Pincode)
			fmt.Printf("From: %v\n", session.From)
			fmt.Printf("To: %v\n", session.To)
			fmt.Printf("Fee Type: %v\n", session.FeeType)
			fmt.Printf("Fees: %v\n", session.Fee)
			fmt.Printf("Dosage 1: %v\n", session.AvailableCapacityDose1)
			fmt.Printf("Dosage 2: %v\n", session.AvailableCapacityDose2)
			fmt.Printf("Age Limit %v\n", session.MinAgeLimit)
			fmt.Printf("Vaccine %v\n", session.Vaccine)
			fmt.Printf("Slots %v\n", session.Slots)
			fmt.Println("**************************************************************************")
			fmt.Println()

		}
	}

}

//DisplayPinCenters displayes the center details based on pin
func DisplayPinCenters(pin *pin.CentersByPIN) {
	if len(pin.Sessions) == 0 {
		log.Println("No Vaccines available for the location")
	} else {
		for _, session := range pin.Sessions {

			fmt.Println()
			fmt.Println()
			fmt.Println("**************************************************************************")
			fmt.Printf("Date: %v\n", session.Date)
			fmt.Printf("Address: %v\n", session.Address)
			fmt.Printf("Name: %v\n", session.Name)
			fmt.Printf("State: %v\n", session.StateName)
			fmt.Printf("District: %v\n", session.DistrictName)
			fmt.Printf("Block: %v\n", session.BlockName)
			fmt.Printf("Block: %v\n", session.BlockName)
			fmt.Printf("PIN: %v\n", session.Pincode)
			fmt.Printf("From: %v\n", session.From)
			fmt.Printf("To: %v\n", session.To)
			fmt.Printf("Fee Type: %v\n", session.FeeType)
			fmt.Printf("Fees: %v\n", session.Fee)
			fmt.Printf("Dosage 1: %v\n", session.AvailableCapacityDose1)
			fmt.Printf("Dosage 2: %v\n", session.AvailableCapacityDose2)
			fmt.Printf("Age Limit %v\n", session.MinAgeLimit)
			fmt.Printf("Vaccine %v\n", session.Vaccine)
			fmt.Printf("Slots %v\n", session.Slots)
			fmt.Println("**************************************************************************")
			fmt.Println()
			fmt.Println()

		}
	}

}

//GetCenterByDistrict method gets the list of centers used in GetCentersListByDistrict
func GetCenterByDistrict(dStateId string, dDistrictName string, dDate string, dVac string) {
	fmt.Println("*********Getting Center Information*********")
	fmt.Println(dStateId)
	fmt.Println(dDistrictName)
	fmt.Println(dDate)
	fmt.Println(dVac)

	// *dStateName = flag.String("s", "Karnataka", "The required state..Default: Karnataka")
	// *dDistrictName = flag.String("d", "BBMP", "The desired district..Default: BBMP")
	// *dDate = flag.String("dt", "24-05-2021", "The desired date..Default: 24-05-2021")
	// *dVac = flag.String("v", "COVISHIELD", "The desired vaccine..Default: COVISHIELD")
	// flag.Parse()

	var states states.States
	//get the states data
	states.GetStates()

	if len(states.States) == 0 {
		fmt.Println("No Vaccines available for the location")
	} else {
		//provide the value of the states
		for _, state := range states.States {

			//filter for state
			if state.StateName == dStateId {
				//fmt.Printf("State: %s", state.StateName)
				var districts district.Districts
				//fmt.Printf("State id : %s", strconv.Itoa(state.StateID))

				//get the list of districts
				districts.GetDistricts(strconv.Itoa(state.StateID))
				for _, district := range districts.Districts {

					// filter for district
					if district.DistrictName == dDistrictName {
						//fmt.Printf("District ID: %d", district.DistrictID)

						var centers center.Centers
						//get the centers for district , date, vaccine
						centers.GetCenters(strconv.Itoa(district.DistrictID), dDate, dVac)
						DisplayCenters(&centers)

					}

				}
			}

		}
	}

}

//GetCentersListByDistrict gets the list of centers based on State and District
func GetCentersListByDistrict() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter State (ex: Karnataka)")
	fmt.Print(">>")
	//var st string
	//fmt.Scanf("%s\n", &st)
	st, err := reader.ReadString('\r')
	if err != nil {
		log.Printf("Error getting the input: %v", err)
	}
	st = strings.Replace(st, "\r", "", -1)
	fmt.Println("Enter District (ex: BBMP)")
	fmt.Print(">>")
	// var dt string
	// fmt.Scanf("%s\n", &dt)
	dt, err := reader.ReadString('\r')
	if err != nil {
		log.Printf("Error getting teh input: %v", err)
	}
	dt = strings.Replace(dt, "\r", "", -1)
	fmt.Println("Enter Vaccine (ex: D -> COVISHIELD, X -> COVAXIN)")
	fmt.Print(">>")
	var vc string
	fmt.Scanf("%s\n", &vc)
	// vc, _ := reader.ReadString('\n')
	fmt.Println("Enter Date (ex: 31-05-2021)")
	fmt.Print(">>")
	var dDate string
	fmt.Scanf("%s\n", &dDate)
	fmt.Println("******************************************")
	fmt.Println()
	fmt.Println()
	GetCenterByDistrict(st, dt, dDate, vc)
	fmt.Println()
	fmt.Println()
}

//CallOTPBasedPin gets the list of centers based on PIN info
func CallOTPBasedPin() {
	var mypin string
	var date string
	fmt.Println("Enter PIN for desired location (ex: 560017)")
	fmt.Print(">>")
	fmt.Scanf("%s\n", &mypin)
	fmt.Println("Enter Date (ex: 31-05-2021)")
	fmt.Print(">>")
	fmt.Scanf("%s\n", &date)
	var pinCenters pin.CentersByPIN
	pinCenters.GetCentersByPIN(mypin, date)
	DisplayPinCenters(&pinCenters)
}
