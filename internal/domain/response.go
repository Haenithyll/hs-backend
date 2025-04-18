package domain

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func ToErrorResponse(c *gin.Context, err *DomainError) {
	switch err.Code {
	case ErrNotFound:
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Message})
	case ErrBadRequest:
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Message})
	case ErrForbidden:
		c.JSON(http.StatusForbidden, ErrorResponse{Error: err.Message})
	case ErrInternalServerError:
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Message})
	default:
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Message})
	}
}

func Ok(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, data)
}
