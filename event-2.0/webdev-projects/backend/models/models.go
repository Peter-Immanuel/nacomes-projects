package models

type SEX string

const (
	Undefined SEX = SEX(iota)
	Male
	Female
)

type UserRequest struct {
	Name          string `json:"name"`
	Age           int    `json:"age"`
	Gender        SEX    `json:"gender"`
	MovieCategory string `json:"movieCategory"`
}

type UserMovieResponse struct {
	Title  string
	Poster string
	Year   string
	Type   string
}

type DeviceRequest struct {
	ID       string
	Location string
	DateTime string
}
