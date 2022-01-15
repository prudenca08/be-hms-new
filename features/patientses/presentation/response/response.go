package response

import (
	"finalproject/features/patientses"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type CreatePatientsesResponse struct {
	ID      int    `json:"id"`
	DoctorID          int `json:"doctorid"`
	PatientID         int `json:"patientid"`
	PatientScheduleID int `json:"patientscheduleid"`
	Status  string `json:"status"`
	Date    string `json:"date"`
	Message string `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type BaseResponse struct {
	Meta struct {
		Status   int      `json:"rc"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, param interface{}) error{
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = param
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Something Error"
	response.Meta.Messages = []string{err.Error()}
	return c.JSON(status, response)
}

func FromDomainCreate(domain patientses.Domain) CreatePatientsesResponse{
	return CreatePatientsesResponse{
		ID: domain.ID,
		DoctorID: domain.DoctorID,
	PatientID: domain.PatientID ,
	PatientScheduleID: domain.PatientScheduleID,
		Status: domain.Status,
		Date: domain.Date,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		Message: "Create Patient Session Success",
	}
}
type PatientsesResponse struct{
	ID int `json:"id"`
	DoctorID          int `json:"doctorid"`
	PatientID         int `json:"patientid"`
	PatientScheduleID int `json:"patientscheduleid"`
	Status string `json:"status"`
	Date string `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}
func FromDomainAllPatientses(domain patientses.Domain) PatientsesResponse{
	return PatientsesResponse{
		ID: domain.ID,
		DoctorID: domain.DoctorID,
	PatientID: domain.PatientID,
	PatientScheduleID: domain.PatientScheduleID,
		Status: domain.Status,
		Date: domain.Date,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainUpdatePatientses(domain patientses.Domain) CreatePatientsesResponse{
	return CreatePatientsesResponse{
		Message : "Update Patient Session Success",
		ID: domain.ID,
		DoctorID: domain.DoctorID,
	PatientID: domain.PatientID,
	PatientScheduleID: domain.PatientScheduleID,
		Status: domain.Status,
		Date: domain.Date,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func FromPatientsesListDomain(domain []patientses.Domain) []PatientsesResponse{
	var response []PatientsesResponse
	for _, value := range domain {
		response = append(response, FromDomainAllPatientses(value))
	}
	return response
}