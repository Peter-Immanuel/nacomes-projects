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
	DeviceID string `json:"deviceID"`
	Location string `json:"location"`
	DateTime string `json:"dateTime"`
}

func (dr *DeviceRequest) Validate() bool {
	if dr.DeviceID == "" || dr.Location == "" || dr.DateTime == "" {
		return false
	}
	return true
}
