package structs

// Weather Top-Level-Struct of the 7Timer response
type Weather struct {
	Product    string     `json:"product"`
	Init       string     `json:"init"`
	Dataseries []Datasery `json:"dataseries"`
}

// Datasery part of Weather struct
type Datasery struct {
	Timepoint    int64    `json:"timepoint"`
	Cloudcover   int64    `json:"cloudcover"`
	Seeing       int64    `json:"seeing"`
	Transparency int64    `json:"transparency"`
	LiftedIndex  int64    `json:"lifted_index"`
	Rh2M         int64    `json:"rh2m"`
	Wind10M      Wind10M  `json:"wind10m"`
	Temp2M       int64    `json:"temp2m"`
	PrecType     PrecType `json:"prec_type"`
}

// Wind10M contains wind speed information and is part of Datasery
type Wind10M struct {
	Direction string `json:"direction"`
	Speed     int64  `json:"speed"`
}

// PrecType a string that is part of Datasery
type PrecType string
const (
	// None ...
	None PrecType = "none"
)
