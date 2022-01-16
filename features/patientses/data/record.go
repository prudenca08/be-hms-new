package data

import (
	doctorrecord "finalproject/features/doctor/data"
	"finalproject/features/patientses"
	datarecipe "finalproject/features/recipe/data"
	"time"

	"gorm.io/gorm"
)

type Patientses struct {
	gorm.Model
	ID                int `gorm:"primary_key"`
	AdminID           int
	DoctorID          int 
	PatientID         int
	PatientScheduleID int 
	Date              string
	Status            string
	Doctor doctorrecord.Doctor `gorm:"foreignKey:ID;references:DoctorID"`
	Recipe datarecipe.Recipe `gorm:"foreignKey:PatientSessionID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
func ToDomain(pss Patientses) patientses.Domain {
	return patientses.Domain{
		ID:                pss.ID,
		AdminID:           pss.AdminID,
		DoctorID:          pss.DoctorID,
		PatientID:         pss.PatientID,
		PatientScheduleID: pss.PatientScheduleID,
		Date:              pss.Date,
		Status:            pss.Status,
		Doctor: doctorrecord.ToDomain(pss.Doctor),
		Recipe: datarecipe.ToDomain(pss.Recipe),
		CreatedAt: pss.CreatedAt,
		UpdatedAt: pss.UpdatedAt,
	}
}
func fromDomain(domain patientses.Domain) Patientses {
	return Patientses{
		ID:                domain.ID,
		AdminID:           domain.AdminID,
		DoctorID:          domain.DoctorID,
		PatientID:         domain.PatientID,
		PatientScheduleID: domain.PatientScheduleID,
		Date:              domain.Date,
		Status:            domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func toDomainUpdate(pss Patientses) patientses.Domain{
	return patientses.Domain{
		ID: pss.ID,
		AdminID: pss.AdminID,
		DoctorID: pss.DoctorID,
		PatientID: pss.PatientID,
		PatientScheduleID: pss.PatientScheduleID,
		Date: pss.Date,
		Status: pss.Status,
		CreatedAt: pss.CreatedAt,
		UpdatedAt: pss.UpdatedAt,
	}
}
func toDomainList(data []Patientses) []patientses.Domain {
	result := []patientses.Domain{}
	for _, pss := range data{
		result = append(result, ToDomain(pss))
	}
	return result
}