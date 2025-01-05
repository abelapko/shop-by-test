package models

type APIRate struct {
	Code    string  `json:"Cur_Abbreviation"`
	Name    string  `json:"Cur_Name"`
	Rate    float64 `json:"Cur_OfficialRate"`
	Date    string  `json:"Date"`
	Nominal int     `json:"Cur_Scale"`
}
