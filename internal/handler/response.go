package handler

import (
	"net/http"

	"hs-backend/internal/error"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, data)
}

func BadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, error.ErrorResponse{Error: msg})
}

func NotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, error.ErrorResponse{Error: msg})
}

func InternalError(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, error.ErrorResponse{Error: msg})
}
