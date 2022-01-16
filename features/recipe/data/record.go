package data

import (
	doctorrecord "finalproject/features/doctor/data"
	"finalproject/features/recipe"
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	PatientSessionID int
	DoctorID int
	ID int
	Title string
	DetailRecipe string
	CreatedAt time.Time
	UpdatedAt time.Time
	Doctor doctorrecord.Doctor `gorm:"foreignKey:ID;references:DoctorID"`
}

func ToDomain(rec Recipe) recipe.Domain{
	return recipe.Domain{
		PatientSessionID: rec.PatientSessionID,
		ID: rec.ID,
		DoctorID: rec.DoctorID,
		Title: rec.Title,
		DetailRecipe: rec.DetailRecipe,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(domain recipe.Domain) Recipe{
	return Recipe{
		PatientSessionID: domain.PatientSessionID,
		DoctorID: domain.DoctorID,
		ID: domain.ID,
		Title: domain.Title,
		DetailRecipe: domain.DetailRecipe,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func toDomainUpdate(rec Recipe) recipe.Domain{
	return recipe.Domain{
		PatientSessionID: rec.PatientSessionID,
		DoctorID: rec.DoctorID,
		ID: rec.ID,
		Title: rec.Title,
		DetailRecipe: rec.DetailRecipe,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func toDomainList(data []Recipe) []recipe.Domain{
	result := []recipe.Domain{}
	for _, rec := range data {
		result = append(result, ToDomain(rec))
	}
	return result
}