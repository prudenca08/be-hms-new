package response

import (
	"finalproject/features/doctor"
	"time"

	"net/http"

	echo "github.com/labstack/echo/v4"
)

type DoctorRegisterResponse struct {
	Message      string    `json:"message"`
	ID           int       `json:"id:"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Name         string    `json:"name"`
	NIP          string    `json:"nip"`
	Experience   string    `json:"experience"`
	Spesialist   string    `json:"specialist"`
	Room         string    `json:"room"`
	Phone_Number string    `json:"phone_number"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type DoctorResponse struct {
	ID           int       `json:"id:"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Name         string    `json:"name"`
	NIP          string    `json:"nip"`
	Experience   string    `json:"experience"`
	Spesialist   string    `json:"specialist"`
	Room         string    `json:"room"`
	Phone_Number string    `json:"phone_number"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
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

func FromDomainRegister(domain doctor.Domain) DoctorRegisterResponse {
	return DoctorRegisterResponse{
		Message:      "Register Doctor Success",
		ID:           domain.ID,
		Username:     domain.Username,
		Password:     domain.Password,
		Name:         domain.Name,
		NIP:          domain.NIP,
		Experience:   domain.Experience,
		Spesialist:   domain.Specialist,
		Room:         domain.Room,
		Phone_Number: domain.Phone_Number,
		Status:       domain.Status,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

type DoctorLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func FromDomainLogin(domain doctor.Domain) DoctorLoginResponse {
	return DoctorLoginResponse{
		Message: "Doctor Login Success",
		Token:   domain.Token,
	}
}
func FromDomainUpdateDoctor(domain doctor.Domain) DoctorRegisterResponse {
	return DoctorRegisterResponse{
		Message:      "Update Doctor Success",
		ID:           domain.ID,
		Username:     domain.Username,
		Password:     domain.Password,
		Name:         domain.Name,
		NIP:          domain.NIP,
		Experience:   domain.Experience,
		Spesialist:   domain.Specialist,
		Room:         domain.Room,
		Phone_Number: domain.Phone_Number,
		Status:       domain.Status,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func FromDomainAllDoctor(domain doctor.Domain) DoctorResponse {
	return DoctorResponse{
		ID:           domain.ID,
		Username:     domain.Username,
		Password:     domain.Password,
		Name:         domain.Name,
		NIP:          domain.NIP,
		Experience:   domain.Experience,
		Spesialist:   domain.Specialist,
		Room:         domain.Room,
		Phone_Number: domain.Phone_Number,
		Status:       domain.Status,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func FromDoctorListDomain(domain []doctor.Domain) []DoctorResponse {
	var response []DoctorResponse
	for _, value := range domain {
		response = append(response, FromDomainAllDoctor(value))
	}
	return response
}
