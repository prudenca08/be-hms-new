package patientses

import (
	doctorentity "finalproject/features/doctor"
	recipeentity "finalproject/features/recipe"
	"time"
)

type Domain struct {
	ID                int
	AdminID           int
	DoctorID          int
	PatientID         int
	PatientScheduleID int
	Date              string
	Status            string
	Doctor doctorentity.Domain
	Recipe recipeentity.Domain
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	AllPatientses() ([]Domain, error)
	Create(pssID int, domain *Domain) (Domain, error)
	Update(admID int, pssID int, domain *Domain) (Domain, error)
	Delete(pssID, id int) (string, error)
	PatientsesByID(id int) (Domain, error)
}
type Repository interface {
	AllPatientses() ([]Domain, error)
	Create(pssID int, domain *Domain) (Domain, error)
	Update(admID int, pssID int, domain *Domain) (Domain, error)
	Delete(pssID, id int) (string, error)
	PatientsesByID(id int) (Domain, error)
}