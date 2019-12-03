package entity

type Event struct {
	Id          uint64 `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
