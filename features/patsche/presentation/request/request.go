package request

import "finalproject/features/patsche"

type Patsche struct {
	// DoctorID int `json:"doctorid"`
	AdminID int `json:"adminid"`
	Day  string `json:"day"`
	Time string `json:"time"`
}

func (req *Patsche) ToDomain() *patsche.Domain {
	return &patsche.Domain{
		// DoctorID: req.DoctorID,
		AdminID: req.AdminID,
		Day:  req.Day,
		Time: req.Time,
	}
}
