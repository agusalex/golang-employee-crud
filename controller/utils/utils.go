package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Return500(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func Return404(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
}

func Return400(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func Return200(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data})
}
