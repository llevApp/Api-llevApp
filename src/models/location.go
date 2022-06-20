package models

type location struct {
	TripID    int     `json:"trip_id,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type LocationResponse struct {
	User_type string   `json:"type,omitempty"`
	UserID    int      `json:"user_id,omitempty"`
	Location  location `json:"location,omitempty"`
}
