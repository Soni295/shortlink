package entity

type Link struct {
	ID      uint64 `json:"id"`
	Url     string `json:"url"`
	Shorted string `json:"shorted"`
}
