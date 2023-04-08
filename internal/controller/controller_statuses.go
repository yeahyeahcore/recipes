package controller

import (
	"net/http"

	"github.com/labstack/echo"

	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/models"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/repository"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/service"
)

var controllerStatusesMap = map[error]int{
	repository.ErrCreateRecord:         http.StatusInternalServerError,
	repository.ErrDeleteRecord:         http.StatusInternalServerError,
	repository.ErrEmptyArgs:            http.StatusBadRequest,
	repository.ErrGetRecord:            http.StatusInternalServerError,
	repository.ErrIterateRows:          http.StatusInternalServerError,
	repository.ErrMethodNotImplemented: http.StatusInternalServerError,
	repository.ErrNoRecord:             http.StatusNotFound,
	repository.ErrScanRecord:           http.StatusInternalServerError,
	repository.ErrUpdateRecord:         http.StatusInternalServerError,
	ErrInvalidParam:                    http.StatusBadRequest,
	service.ErrEmptyArgs:               http.StatusBadRequest,
	service.ErrInternal:                http.StatusInternalServerError,
}

func response(ctx echo.Context, message string, err error) error {
	if err == nil {
		return ctx.JSON(http.StatusOK, models.HTTPErrorResponse{
			Message:    message,
			StatusCode: http.StatusOK,
		})
	}

	status, ok := controllerStatusesMap[err]
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{
			Message:    message,
			StatusCode: http.StatusInternalServerError,
		})
	}

	return ctx.JSON(status, models.HTTPErrorResponse{
		Message:    message,
		StatusCode: status,
	})
}
