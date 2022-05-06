package models

type TripsRecords struct {
	Driver       string  `json:"name"`
	DriverCareer string  `json:"career"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	InitTripTime string  `json:"init_trip_time"`
}

type NewTripsRecords struct {
	DriverId  int     `json:"user_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Time      string  `json:"start_time"`
}

type TripRequest struct {
	UserName     string  `json:"user_name,omitempty"`
	UserCareer   string  `json:"user_career,omitempty"`
	UserID       int     `json:"user_id,omitempty"`
	TripID       int     `json:"trip_id,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Contribution int     `json:"contribution,omitempty"`
}
