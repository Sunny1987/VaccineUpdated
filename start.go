package main

import (
	"fmt"
	"vaccine1/otp"
	"vaccine1/utilities"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewFigure("Cowin Client", "graffiti", true)
	myFigure.Print()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	//Initiate Authentication
	otp.GetAuthenticated()

	//Check Authentication
	if otp.CheckAuthentication() {

		var ex bool
		for !ex {
			ch := GetOptions()

			switch ch {
			case "1":
				utilities.GetStateData()
				ex = false
			case "2":
				utilities.GetDistrictData()
				ex = false
			case "3":
				utilities.GetCentersListByDistrict()
				ex = false
			case "4":
				utilities.CallOTPBasedPin()
				ex = false
			case "5":
				fmt.Println("********Exiting the Cowin Client****************")
				ex = true

			}
		}

	} else {
		//If the Authetication check is failed
		fmt.Println("*********Authentication not successful*******")
	}

}

//GetOptions will return the options available
func GetOptions() string {

	var ch string
	fmt.Println("Choose options")
	fmt.Println("*********************************************")
	fmt.Println("1. Get States names List")
	fmt.Println("2. Get States based district names List (Please get State Id from Opt.1)")
	fmt.Println("3. State-Dist-Center list !!Help from #Opts {1 & 2} ")
	fmt.Println("4. PIN based Center list")
	fmt.Println("5. Exit")
	fmt.Println("*********************************************")
	fmt.Print(">>")
	fmt.Scanf("%s\n", &ch)
	return ch
}
