package request

import (
	"finalproject/features/doctor"
)

type Doctor struct {
	DoctorSessionID int    `json:"doctorsessionid"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Name            string `json:"name"`
	NIP             string `json:"nip"`
	Experience      string `json:"experience"`
	Specialist      string `json:"specialist"`
	Room            string `json:"room"`
	Phone_Number    string `json:"phone_number"`
	Status          string `json:"status"`
}

func (req *Doctor) ToDomain() *doctor.Domain {
	return &doctor.Domain{
		DoctorSessionID: req.DoctorSessionID,
		Username:        req.Username,
		Password:        req.Password,
		Name:            req.Name,
		NIP:             req.NIP,
		Experience:      req.Experience,
		Specialist:      req.Specialist,
		Room:            req.Room,
		Phone_Number:    req.Phone_Number,
		Status:          req.Status,
	}
}

type DoctorLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
