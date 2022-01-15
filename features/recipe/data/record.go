package data

import (
	"finalproject/features/recipe"
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	PatientSessionID int
	DoctorID int
	ID int
	// Day string
	// Time string
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
		// Day : rec.Day,
		// Time: rec.Time,
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
		// Day: domain.Day,
		// Time: domain.Time,
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
		// Day: rec.Day,
		// Time: rec.Time,
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