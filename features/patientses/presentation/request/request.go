package request

import "finalproject/features/patientses"

type Patientses struct {
	AdminID           int `json:"adminid"`
	DoctorID          int `json:"doctorid"`
	PatientID         int `json:"patientid"`
	PatientScheduleID int `json:"patientscheduleid"`
	Date              string `json:"date"`
	Status            string `json:"status"`
}

func (req *Patientses) ToDomain() *patientses.Domain {
	return &patientses.Domain{
		AdminID: req.AdminID,
		DoctorID: req.DoctorID,
		PatientID: req.PatientID,
		PatientScheduleID: req.PatientScheduleID,
		Date:              req.Date,
		Status:            req.Status,
	}
}
