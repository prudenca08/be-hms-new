package data

import (
	"finalproject/features/recipe"
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	PatientSessionID int
	ID int `gorm:"primary_key"`
	DoctorID int
	Title string
	DetailRecipe string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(rec Recipe) recipe.Domain{
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
		ID: domain.ID,
		DoctorID: domain.DoctorID,
		Title: domain.Title,
		DetailRecipe: domain.DetailRecipe,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func toDomainUpdate(rec Recipe) recipe.Domain{
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

func toDomainList(data []Recipe) []recipe.Domain{
	result := []recipe.Domain{}
	for _, rec := range data {
		result = append(result, toDomain(rec))
	}
	return result
}