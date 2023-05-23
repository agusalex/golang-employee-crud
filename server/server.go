package server

import (
	"github.com/agusalex/golang-employee-crud/config"
	"github.com/agusalex/golang-employee-crud/db"
	repositories2 "github.com/agusalex/golang-employee-crud/repositories"
	services2 "github.com/agusalex/golang-employee-crud/services"
)

func InitServer() {
	router := InitRouter()
	_ = db.Connect()

	initServices(repositories2.NewMemberRepository(db.DB), repositories2.NewTagRepository(db.DB), repositories2.NewSearchRepository(db.DB))
	err := router.Run(":" + config.Get().Server.Port)

	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}

}
func initServices(memberRepository *repositories2.MemberRepository, tagRepository *repositories2.TagRepository, searchRepository *repositories2.SearchRepository) {
	memberService := services2.NewMemberService(memberRepository, tagRepository)
	tagService := services2.NewTagService(tagRepository)
	searchService := services2.NewSearchService(searchRepository)

	services2.MemberService = memberService
	services2.TagService = tagService
	services2.SearchService = searchService
}
