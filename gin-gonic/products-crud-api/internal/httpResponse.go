package internal

import (
	"net/http"
)

type Product struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description" binding:"omitempty,max=250"`
	CreatedAt   string  `json:"created_at"`
}

type ProductResponse struct {
	GUID        string  `json:"guid"`
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description" binding:"omitempty,max=250"`
	CreatedAt   string  `json:"created_at"`
}

type HttpResponse struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewHttpResponse(status int, data interface{}) HttpResponse {
	switch status {
	case http.StatusBadRequest,
		http.StatusInternalServerError,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusRequestTimeout,
		http.StatusNotImplemented,
		http.StatusNotAcceptable,
		http.StatusMethodNotAllowed:

		if e, ok := data.(error); ok {
			return HttpResponse{
				Status:  status,
				Success: false,
				Message: e.Error(),
				Data:    nil,
			}
		}
		return HttpResponse{
			Status:  status,
			Success: false,
			Message: data.(string),
			Data:    nil,
		}
	default:
		return HttpResponse{
			Status:  status,
			Success: true,
			Message: "Success",
			Data:    data,
		}
	}
}
