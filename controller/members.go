package controller

import (
	"fmt"
	. "github.com/agusalex/golang-employee-crud/controller/utils"
	. "github.com/agusalex/golang-employee-crud/controller/validators"
	"github.com/agusalex/golang-employee-crud/models"
	. "github.com/agusalex/golang-employee-crud/services"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetMembersHandler Get all members with tags
// @Summary Get all members with tags
// @Description Get all members with tags
// @Tags members
// @Produce json
// @Success 200 {array} models.Member
// @Router /members [get]
func GetMembersHandler(c *gin.Context) {
	members, err := MemberService.GetAllMembers()
	if err != nil {
		Return500(c, err)
		return
	}
	Return200(c, members)
}

// GetMemberHandler Get a member by ID with tags
// @Summary Get a member by ID with tags
// @Description Get a member by ID with tags
// @Tags members
// @Param id path string true "Member ID"
// @Produce json
// @Success 200 {object} models.Member
// @Router /members/{id} [get]
func GetMemberHandler(c *gin.Context) {
	id := c.Param("id")
	member, err := MemberService.GetMemberByID(id)
	if err != nil {
		Return404(c, fmt.Errorf("member not found"))
		return
	}
	Return200(c, member)
}

// PostMembersHandler Add a new member
// @Summary Add a new member
// @Description Add a new member
// @Tags members
// @Accept json
// @Produce json
// @Param member body models.Member true "Member object that needs to be added"
// @Success 200 {object} models.Member
// @Router /members [post]
func PostMembersHandler(c *gin.Context) {
	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		Return400(c, err)
		return
	}

	if err := Validate.Struct(member); err != nil {
		Return400(c, err)
		return
	}

	member, err := MemberService.CreateMember(member)
	if err != nil {
		Return500(c, err)
		return
	}

	Return200(c, member)
}

// DeleteMemberHandler Delete a member
// @Summary Delete a member
// @Description Delete a member by ID
// @Tags members
// @Param id path string true "Member ID"
// @Success 200 {string} string "Member deleted successfully"
// @Router /members/{id} [delete]
func DeleteMemberHandler(c *gin.Context) {
	id := c.Param("id")
	err := MemberService.DeleteMember(id)
	if err != nil {
		Return500(c, err)
		return
	}

	Return200(c, fmt.Sprintf("member %s deleted successfully", id))
}

// PutMemberHandler Update a member by ID
// @Summary Update a member by ID
// @Description Update a member by ID
// @Tags members
// @Param id path string true "Member ID"
// @Accept json
// @Produce json
// @Param member body models.Member true "Updated member object"
// @Success 200 {object} models.Member
// @Router /members/{id} [put]
func PutMemberHandler(c *gin.Context) {
	id := c.Param("id")
	var member models.Member

	if err := c.ShouldBindJSON(&member); err != nil {
		Return400(c, err)
		return
	}

	if err := Validate.Struct(member); err != nil {
		Return400(c, err)
		return
	}

	member, err := MemberService.UpdateMember(id, member)
	if err != nil {
		Return500(c, err)
		return
	}

	Return200(c, member)
}
