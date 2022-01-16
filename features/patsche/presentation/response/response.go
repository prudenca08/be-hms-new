package response

import (
	doctorresponse "finalproject/features/doctor/presentation/response"
	patientresponse "finalproject/features/patientses/presentation/response"
	"finalproject/features/patsche"
	"time"

	"net/http"

	echo "github.com/labstack/echo/v4"
)

type CreatePatscheResponse struct {
	Message   string    `json:"message"`
	AdminID int `json:"adminid"`
	DoctorID int `json:"doctorid"`
	ID        int       `json:"id:"`
	Day       string    `json:"day"`
	Time      string    `json:"time"`
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

func NewSuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = param

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Something not right"
	response.Meta.Messages = []string{err.Error()}

	return c.JSON(status, response)
}

func FromDomainCreate(domain patsche.Domain) CreatePatscheResponse {
	return CreatePatscheResponse{
		Message:   "Create  Patient Schedule Success",
		AdminID: domain.AdminID,
		DoctorID: domain.DoctorID,
		ID:        domain.ID,
		Day:       domain.Day,
		Time:      domain.Time,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type PatscheResponse struct {
	Message   string    `json:"message"`
	AdminID int `json:"adminid"`
	DoctorID int `json:"doctorid"`
	ID        int       `json:"id:"`
	Day       string    `json:"day"`
	Time      string    `json:"time"`
	PatientSession patientresponse.PatientsesResponse `json:"patientsession"`
	Doctor doctorresponse.DoctorResponse `json:"doctor"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainAllPatsche(domain patsche.Domain) PatscheResponse {
	return PatscheResponse{
		AdminID: domain.AdminID,
		DoctorID: domain.DoctorID,
		ID:        domain.ID,
		Day:       domain.Day,
		Time:      domain.Time,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		Doctor: doctorresponse.FromDomainAllDoctor(domain.Doctor),
		PatientSession: patientresponse.FromDomainAllPatientses(domain.PatientSession),
	}
}

func FromDomainUpdatePatsche(domain patsche.Domain) CreatePatscheResponse {
	return CreatePatscheResponse{
		Message:   "Update  Patient Schedule Success",
		AdminID: domain.AdminID,
		DoctorID: domain.DoctorID,
		ID:        domain.ID,
		Day:       domain.Day,
		Time:      domain.Time,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromPatscheListDomain(domain []patsche.Domain) []PatscheResponse {
	var response []PatscheResponse
	for _, value := range domain {
		response = append(response, FromDomainAllPatsche(value))
	}
	return response
}
