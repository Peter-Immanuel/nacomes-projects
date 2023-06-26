package models

type SEX string

const (
	Undefined SEX = SEX(iota)
	Male
	Femalte
)

type UserRequest struct {
	Name          string
	Age           int
	Gender        SEX
	MovieCategory string
}

type UserMovieResponse struct {
	MovieImage    string
	MovieTitle    string
	MovieDuration string
	MovieRating   uint
}

type DeviceRequest struct {
	ID       string
	Location string
	DateTime string
}
