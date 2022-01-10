package presentation

import (
	"finalproject/features/doctor"
	"finalproject/features/doctor/presentation/request"
	"finalproject/features/doctor/presentation/response"
	"finalproject/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DoctorHandler struct {
	doctorHand doctor.Service
}

func NewHandlerDoctor(handler doctor.Service) *DoctorHandler {
	return &DoctorHandler{
		doctorHand: handler,
	}
}

func (ctrl *DoctorHandler) Register(c echo.Context) error {

	registerReq := request.Doctor{}

	if err := c.Bind(&registerReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.doctorHand.Register(registerReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *DoctorHandler) Login(c echo.Context) error {

	loginReq := request.DoctorLogin{}

	if err := c.Bind(&loginReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result, err := ctrl.doctorHand.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainLogin(result))
}
func (ctrl *DoctorHandler) Update(c echo.Context) error {

	updateReq := request.Doctor{}

	if err := c.Bind(&updateReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	getData, _ := ctrl.doctorHand.DoctorByID(id)
	result, err := ctrl.doctorHand.Update(id, updateReq.ToDomain())
	result.ID = getData.ID

	result.Name = getData.Name

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainUpdateDoctor(result))

}
func (ctrl *DoctorHandler) DoctorByID(c echo.Context) error {

	doctorID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.doctorHand.DoctorByID(doctorID)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return response.NewSuccessResponse(c, response.FromDomainAllDoctor(result))
}
func (ctrl *DoctorHandler) Delete(c echo.Context) error {

	orgzID := middleware.GetUser(c)
	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.doctorHand.Delete(orgzID.ID, deletedId)

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, result)

}
func (ctrl *DoctorHandler) AllDoctor(c echo.Context) error {

	result, err := ctrl.doctorHand.AllDoctor()

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDoctorListDomain(result))

}
