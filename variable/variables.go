package variable

import "github.com/go-resty/resty/v2"

var (
	Client     *resty.Client = resty.New()
	ConfirmOTP string        = "https://cdn-api.co-vin.in/api/v2/auth/public/confirmOTP"
	GetOTP     string        = "https://cdn-api.co-vin.in/api/v2/auth/public/generateOTP"
	GETSTATES  string        = "https://cdn-api.co-vin.in/api/v2/admin/location/states"
	GetDist    string        = "https://cdn-api.co-vin.in/api/v2/admin/location/districts/"
	GetCenters string        = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByDistrict"
	Token      string        = ""
	Bearer     string        = "Bearer "
	FindByPin  string        = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByPin"
)
