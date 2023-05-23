package controller

import (
	. "github.com/agusalex/golang-employee-crud/controller/utils"
	. "github.com/agusalex/golang-employee-crud/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @BasePath /api/v1

// SearchMembersHandler Search members
// @Summary Search members
// @Description Provides search capabilities on the members, allows for querying by type and by tag
// @Tags members
// @Produce json
// @Param tags query []string false "Tags to search for"
// @Param type query string false "Member type to search for"
// @Success 200 {array} models.Member
// @Router /members/search [get]

type SearchMembersRequest struct {
	Tags       []string `form:"tags"`
	MemberType string   `form:"type" validate:"omitempty,oneof=CONTRACTOR EMPLOYEE"`
}

func SearchMembersHandler(c *gin.Context) {
	var request SearchMembersRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		Return400(c, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		Return400(c, err)
		return
	}

	members, err := SearchService.SearchMembers(request.Tags, request.MemberType)
	if err != nil {
		Return500(c, err)
		return
	}

	c.JSON(200, members)
}
