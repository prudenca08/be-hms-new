package recipe

import "time"

type Domain struct {
	ID               int
	DoctorID         int
	PatientSessionID int
	Title        string
	DetailRecipe string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Repository interface {
	AllRecipe() ([]Domain, error)
	Create(recID int, domain *Domain) (Domain, error)
	Update(docID int, recID int, domain *Domain) (Domain, error)
	Delete(recID, id int) (string, error)
	RecipeByID(id int) (Domain, error)
}
type Service interface {
	AllRecipe() ([]Domain, error)
	Create(recID int, domain *Domain) (Domain, error)
	Update(docID int, recID int, domain *Domain) (Domain, error)
	Delete(recID, id int) (string, error)
	RecipeByID(id int) (Domain, error)
}