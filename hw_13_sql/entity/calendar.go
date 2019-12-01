package entity

type Calendar struct {
	Dates []*Date `json:"dates,omitempty"`
}

type Date struct {
	Id            uint32 `json:"id"`
	Day           string `json:"day,omitempty"`
	Description   string
	IsCelebration bool
	Records       []Record `json:"records,omitempty"`
}
