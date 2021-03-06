package models

type TripsRecords struct {
	Id             string  `json:"trip_id"`
	Driver         string  `json:"name"`
	DriverCareer   string  `json:"career"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	InitTripTime   string  `json:"init_trip_time"`
	TotalTip       float64 `json:"total_tips"`
	TotalPassenger int     `json:"total_passenger"`
	Address        string  `json:"address"`
	DriverID       int     `json:"driver_id"`
	UUID           string  `json:"uuid_fb,omitempty"`
}

type Tip struct {
	Id    string `json:"trip_id"`
	Total string `json:"tips"`
}

type NewTripsRecords struct {
	DriverId  int     `json:"user_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Time      string  `json:"start_time"`
	Address   string  `json:"address"`
}

type TripRequest struct {
	UserName     string  `json:"user_name,omitempty"`
	UserCareer   string  `json:"user_career,omitempty"`
	UserID       int     `json:"user_id,omitempty"`
	TripID       int     `json:"trip_id,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Location     string  `json:"location,omitempty"`
	Contribution int     `json:"contribution,omitempty"`
	UUID         string  `json:"uuid_fb,omitempty"`
}

type TripRequestDriver struct {
	Response        string `json:"status,omitempty"`
	Trip_id         int    `json:"trip_id,omitempty"`
	PassangerUserID int    `json:"user_id,omitempty"`
}

type TripRequestPassenger struct {
	Request TripRequest `json:"request,omitempty"`
	IfSend  bool
}

type TripResponseDriver struct {
	Response TripRequestDriver `json:"response,omitempty"`
}
type TripResponseTripsDriver struct {
	Trips   []TripsRecords `json:"trip"`
	HasData bool           `json:"has_data"`
}
