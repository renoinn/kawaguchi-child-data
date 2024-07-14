package entity

type Preschool struct {
	Code                   string
	Id                     string
	Name                   string
	KanaName               string
	EnglishName            string
	Kind                   string
	LocationGovernmentCode int
	TownId                 string
	LocationFull           string
	Prefectures            string
	Municipalities         string
	Town                   string
	StreetAddress          string
	BuildingName           string
	Latitude               float64
	Longitude              float64
}
