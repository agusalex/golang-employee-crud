package controller

import (
	"fmt"
	. "github.com/agusalex/golang-employee-crud/controller/utils"
	"github.com/agusalex/golang-employee-crud/db"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// PingHandler Checks the availability of the server
// @Summary Ping endpoint
// @Description Checks the availability of the server
// @Tags ping
// @Produce plain
// @Success 200 {string} string "Ok"
// @Router /ping [get]
func PingHandler(c *gin.Context) {
	Return200(c, "Ok")
}

// HealthHandler Checks the health status of the server and database connection
// @Summary Health endpoint
// @Description Checks the health status of the server and database connection
// @Tags health
// @Produce plain
// @Success 200 {string} string "Ok"
// @Failure 500 {string} string "Could not establish a connection to the database"
// @Router /health [get]
func HealthHandler(c *gin.Context) {
	if err := db.DB.DB().Ping(); err != nil {
		Return500(c, fmt.Errorf("could not establish a connection to the database"))
		return
	}
	Return200(c, "Ok")
}
