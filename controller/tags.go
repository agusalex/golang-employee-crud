package controller

import (
	. "github.com/agusalex/golang-employee-crud/controller/utils"
	. "github.com/agusalex/golang-employee-crud/services"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetTagsHandler Gets all tags
// Does not give member data on purpose, for searching members through tags there's the search endpoint
// Intent is for this endpoint to be used for populating a combo box for example or other types of UI
// @Summary Get all tags
// @Description Gets all tags
// @Tags tags
// @Produce json
// @Success 200 {array} models.Tag
// @Router /tags [get]
func GetTagsHandler(c *gin.Context) {
	tags, err := TagService.GetAllTags()
	if err != nil {
		Return500(c, err)
		return
	}
	Return200(c, tags)
}
